package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"

	"notify-api/ent/mixin"
	"notify-api/push/item"
)

type Message struct {
	ent.Schema
}

func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()),
		field.String("title"),
		field.String("content").NotEmpty(),
		field.String("long"),
		field.String("priority").NotEmpty().GoType(item.Priority("")),
		field.Int("sequenceID").Immutable().Unique(),
	}
}

func (Message) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeCreateMixin{},
	}
}

func (Message) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("messages").Unique().Required(),
	}
}
