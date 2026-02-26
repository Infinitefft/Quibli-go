package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// PostTag holds the schema definition for the PostTag entity.
type PostTag struct {
	ent.Schema
}

func (PostTag) Fields() []ent.Field {
	return []ent.Field{
		field.Int("post_id"),
		field.Int("tag_id"),
	}
}

func (PostTag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("post", Post.Type),
		edge.To("tag", Tag.Type),
	}
}

func (PostTag) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("tag_id"),
		index.Fields("post_id"),
	}
}
