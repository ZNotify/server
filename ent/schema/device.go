package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"notify-api/ent/mixin"
)

type Device struct {
	ent.Schema
}

func (Device) Fields() []ent.Field {
	return []ent.Field{
		field.String("identifier").Unique(),
		field.String("channel"),
		field.String("channelMeta"),
		field.String("channelToken"),
		field.String("deviceName"),
		field.String("deviceMeta"),
	}
}

func (Device) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeCreateMixin{},
		mixin.TimeUpdateMixin{},
	}
}

func (Device) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("devices").Unique().Required(),
	}
}
