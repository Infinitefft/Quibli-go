package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/index"
)

// UserLikeQuestion holds the schema definition for the UserLikeQuestion entity.
type UserLikeQuestion struct {
	ent.Schema
}

func (UserLikeQuestion) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.Int("question_id"),
	}
}

func (UserLikeQuestion) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type),
		edge.To("question", Question.Type),
	}
}

func (UserLikeQuestion) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("question_id"),
		index.Fields("user_id"),
	}
}
