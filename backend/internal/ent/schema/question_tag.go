package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/index"
)

// QuestionTag holds the schema definition for the QuestionTag entity.
type QuestionTag struct {
	ent.Schema
}

func (QuestionTag) Fields() []ent.Field {
	return []ent.Field{
		field.Int("question_id"),
		field.Int("tag_id"),
	}
}

func (QuestionTag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("question", Question.Type),
		edge.To("tag", Tag.Type),
	}
}

func (QuestionTag) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("tag_id"),
		index.Fields("question_id"),
	}
}
