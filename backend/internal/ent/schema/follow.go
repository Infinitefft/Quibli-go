package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// Follow holds the schema definition for the Follow entity.
type Follow struct {
	ent.Schema
}

func (Follow) Fields() []ent.Field {
	return []ent.Field{
		field.Int("follower_id"),
		field.Int("following_id"),
	}
}

func (Follow) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("follower", User.Type),
		edge.To("following", User.Type),
	}
}

func (Follow) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (Follow) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("follower_id"),
		index.Fields("following_id"),
	}
}
