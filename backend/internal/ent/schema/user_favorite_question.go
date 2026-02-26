package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/index"
)

// UserFavoriteQuestion holds the schema definition for the UserFavoriteQuestion entity.
type UserFavoriteQuestion struct {
	ent.Schema
}

func (UserFavoriteQuestion) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.Int("question_id"),
	}
}

func (UserFavoriteQuestion) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type),
		edge.To("question", Question.Type),
	}
}

func (UserFavoriteQuestion) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("question_id"),
		index.Fields("user_id"),
	}
}
