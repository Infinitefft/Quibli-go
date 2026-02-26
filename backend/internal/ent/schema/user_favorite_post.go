package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/index"
)

// UserFavoritePost holds the schema definition for the UserFavoritePost entity.
type UserFavoritePost struct {
	ent.Schema
}

func (UserFavoritePost) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.Int("post_id"),
	}
}

func (UserFavoritePost) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type),
		edge.To("post", Post.Type),
	}
}

func (UserFavoritePost) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("post_id"),
		index.Fields("user_id"),
	}
}
