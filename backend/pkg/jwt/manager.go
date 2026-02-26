package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID string `json:"sub"`
	Phone  string `json:"name"`
	jwt.RegisteredClaims
}

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Manager struct {
	secretKey string
}

func NewManager(secretKey string) *Manager {
	return &Manager{secretKey: secretKey}
}

func (m *Manager) GenerateTokens(userID, phone string) (*TokenPair, error) {
	now := time.Now()
	nowNumeric := jwt.NewNumericDate(now)
	expireNumeric := jwt.NewNumericDate(now.Add(15 * time.Minute))

	claims := Claims{
		UserID: userID,
		Phone:  phone,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expireNumeric,
			IssuedAt:  nowNumeric,
		},
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(m.secretKey))
	if err != nil {
		return nil, err
	}

	refreshExpireNumeric := jwt.NewNumericDate(now.Add(7 * 24 * time.Hour))
	refreshClaims := Claims{
		UserID: userID,
		Phone:  phone,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: refreshExpireNumeric,
			IssuedAt:  nowNumeric,
		},
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(m.secretKey))
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (m *Manager) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(m.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
