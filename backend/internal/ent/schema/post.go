package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").MaxLen(255),
		field.String("content").Optional(),
		field.Bytes("embedding").Optional(),
	}
}

func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type),
		edge.To("comments", Comment.Type),
		edge.To("tags", PostTag.Type),
		edge.To("likes", UserLikePost.Type),
		edge.To("favorites", UserFavoritePost.Type),
	}
}

func (Post) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (Post) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("title"),
	}
}
