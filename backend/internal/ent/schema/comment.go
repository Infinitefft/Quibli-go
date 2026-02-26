package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.String("content").Optional(),
		field.Int("user_id"),
		field.Int("post_id").Optional(),
		field.Int("question_id").Optional(),
		field.Int("parent_id").Optional(),
	}
}

func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Required(),
		edge.To("post", Post.Type),
		edge.To("question", Question.Type),
		edge.To("parent", Comment.Type),
		edge.To("replies", Comment.Type),
	}
}

func (Comment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (Comment) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("post_id"),
		index.Fields("question_id"),
		index.Fields("user_id"),
		index.Fields("parent_id"),
	}
}
