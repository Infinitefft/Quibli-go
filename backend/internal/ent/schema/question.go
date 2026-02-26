package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// Question holds the schema definition for the Question entity.
type Question struct {
	ent.Schema
}

func (Question) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").MaxLen(255),
		field.Bytes("embedding").Optional(),
	}
}

func (Question) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type),
		edge.To("comments", Comment.Type),
		edge.To("tags", QuestionTag.Type),
		edge.To("likes", UserLikeQuestion.Type),
		edge.To("favorites", UserFavoriteQuestion.Type),
	}
}

func (Question) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (Question) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("title"),
	}
}
