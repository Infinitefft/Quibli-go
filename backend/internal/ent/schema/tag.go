package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/index"
)

// Tag holds the schema definition for the Tag entity.
type Tag struct {
	ent.Schema
}

func (Tag) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(255).Unique(),
	}
}

func (Tag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", PostTag.Type),
		edge.To("questions", QuestionTag.Type),
	}
}

func (Tag) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name"),
	}
}
