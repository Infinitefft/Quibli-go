package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/index"
)

// UserLikePost holds the schema definition for the UserLikePost entity.
type UserLikePost struct {
	ent.Schema
}

func (UserLikePost) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.Int("post_id"),
	}
}

func (UserLikePost) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type),
		edge.To("post", Post.Type),
	}
}

func (UserLikePost) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("post_id"),
		index.Fields("user_id"),
	}
}
