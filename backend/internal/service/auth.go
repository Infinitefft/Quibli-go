package service

import (
	"context"
	"fmt"

	"backend/internal/config"
	"backend/internal/ent"
	"backend/internal/ent/user"
	"backend/pkg/errors"
	"backend/pkg/jwt"
)

type AuthService struct {
	client *ent.Client
	jwtMgr *jwt.Manager
}

func NewAuthService(client *ent.Client) *AuthService {
	return &AuthService{
		client: client,
		jwtMgr: jwt.NewManager(config.AppConfig.JWTSecret),
	}
}

type LoginResponse struct {
	AccessToken  string      `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
	User         *UserDetail `json:"user"`
}

type UserDetail struct {
	ID                string `json:"id"`
	Nickname          string `json:"nickname"`
	Phone             string `json:"phone"`
	Following         []int  `json:"following"`
	LikePosts         []int  `json:"likePosts"`
	FavoritePosts     []int  `json:"favoritePosts"`
	LikeQuestions     []int  `json:"likeQuestions"`
	FavoriteQuestions []int  `json:"favoriteQuestions"`
	FollowingCount    int    `json:"followingCount"`
	FollowerCount     int    `json:"followerCount"`
}

func (s *AuthService) Login(ctx context.Context, phone, password string) (*LoginResponse, error) {
	// 1. 查询用户及所有关联关系
	u, err := s.client.User.
		Query().
		Where(user.PhoneEQ(phone)).
		WithFollowing().
		WithFollowedBy(). // 必须加上，否则计算 FollowerCount 会报错
		WithLikePosts(func(q *ent.UserLikePostQuery) {
			q.WithPost()
		}).
		WithFavoritePosts(func(q *ent.UserFavoritePostQuery) {
			q.WithPost()
		}).
		WithLikeQuestions(func(q *ent.UserLikeQuestionQuery) {
			q.WithQuestion()
		}).
		WithFavoriteQuestions(func(q *ent.UserFavoriteQuestionQuery) {
			q.WithQuestion()
		}).
		Only(ctx)

	if err != nil {
		return nil, errors.ErrInvalidCredentials
	}

	// 2. 密码校验 (暂时改为明文比对以匹配你数据库里的 123456)
	// 等之后写了注册接口，记得换回 bcrypt.CompareHashAndPassword
	if u.Password != password {
		return nil, errors.ErrInvalidCredentials
	}

	// 3. 生成 JWT Tokens
	tokens, err := s.jwtMgr.GenerateTokens(fmt.Sprintf("%d", u.ID), u.Phone)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	// 4. 构建返回的用户详情
	userDetail := s.buildUserDetail(u)

	return &LoginResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		User:         userDetail,
	}, nil
}

func (s *AuthService) RefreshToken(ctx context.Context, refreshToken string) (*LoginResponse, error) {
	claims, err := s.jwtMgr.ValidateToken(refreshToken)
	if err != nil {
		return nil, errors.ErrTokenExpired
	}

	tokens, err := s.jwtMgr.GenerateTokens(claims.UserID, claims.Phone)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	var userID int
	fmt.Sscanf(claims.UserID, "%d", &userID)

	// 为了 Refresh 之后也能拿到完整详情，这里建议也用 Query 补全 Edges
	u, err := s.client.User.Query().
		Where(user.IDEQ(userID)).
		WithFollowing().
		WithFollowedBy().
		Only(ctx)

	if err != nil {
		return nil, errors.ErrNotFound
	}

	userDetail := s.buildUserDetail(u)

	return &LoginResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		User:         userDetail,
	}, nil
}

func (s *AuthService) buildUserDetail(u *ent.User) *UserDetail {
	detail := &UserDetail{
		ID:                fmt.Sprintf("%d", u.ID),
		Nickname:          u.Nickname,
		Phone:             u.Phone,
		Following:         []int{},
		LikePosts:         []int{},
		FavoritePosts:     []int{},
		LikeQuestions:     []int{},
		FavoriteQuestions: []int{},
	}

	// 安全检查：处理 Following
	if u.Edges.Following != nil {
		for _, f := range u.Edges.Following {
			detail.Following = append(detail.Following, int(f.ID))
		}
		detail.FollowingCount = len(u.Edges.Following)
	}

	// 安全检查：处理 Follower
	if u.Edges.FollowedBy != nil {
		detail.FollowerCount = len(u.Edges.FollowedBy)
	}

	// 处理点赞文章 (Ent 的嵌套 Edge 逻辑)
	if u.Edges.LikePosts != nil {
		for _, lp := range u.Edges.LikePosts {
			if lp.Edges.Post != nil {
				detail.LikePosts = append(detail.LikePosts, int(lp.Edges.Post.ID))
			}
		}
	}

	// 处理收藏文章
	if u.Edges.FavoritePosts != nil {
		for _, fp := range u.Edges.FavoritePosts {
			if fp.Edges.Post != nil {
				detail.FavoritePosts = append(detail.FavoritePosts, int(fp.Edges.Post.ID))
			}
		}
	}

	// 处理点赞问题
	if u.Edges.LikeQuestions != nil {
		for _, lq := range u.Edges.LikeQuestions {
			if lq.Edges.Question != nil {
				detail.LikeQuestions = append(detail.LikeQuestions, int(lq.Edges.Question.id))
			}
		}
	}

	// 处理收藏问题
	if u.Edges.FavoriteQuestions != nil {
		for _, fq := range u.Edges.FavoriteQuestions {
			if fq.Edges.Question != nil {
				detail.FavoriteQuestions = append(detail.FavoriteQuestions, int(fq.Edges.Question.ID))
			}
		}
	}

	return detail
}
