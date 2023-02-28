package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"notify-api/db/ent/mixin"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("secret").Unique(),
		field.Int64("githubID").Unique(),
		field.String("githubName"),
		field.String("githubLogin").NotEmpty(),
		field.String("githubOauthToken").NotEmpty(),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeCreateMixin{},
		mixin.TimeUpdateMixin{},
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("devices", Device.Type),
		edge.To("messages", Message.Type),
	}
}
