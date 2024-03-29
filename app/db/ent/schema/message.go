package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"

	"github.com/ZNotify/server/app/db/ent/mixin"
	"github.com/ZNotify/server/app/manager/push/enum"
)

type Message struct {
	ent.Schema
}

func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).Immutable(),
		field.String("title").Immutable(),
		field.String("content").NotEmpty().Immutable(),
		field.String("long").Immutable(),
		field.String("priority").NotEmpty().GoType(enum.Priority("")).Immutable(),
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
