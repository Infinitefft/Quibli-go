package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("phone").MaxLen(11).Unique(),
		field.String("nickname").MaxLen(50).Optional(),
		field.String("password").MaxLen(255),
		field.String("avatar").MaxLen(255).Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Post.Type),
		edge.To("questions", Question.Type),
		edge.To("comments", Comment.Type),
		edge.To("like_posts", UserLikePost.Type),
		edge.To("like_questions", UserLikeQuestion.Type),
		edge.To("favorite_posts", UserFavoritePost.Type),
		edge.To("favorite_questions", UserFavoriteQuestion.Type),
		edge.To("following", Follow.Type),
		edge.To("followed_by", Follow.Type),
	}
}

// Mixin of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Indexes of the User.
func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("phone"),
	}
}
