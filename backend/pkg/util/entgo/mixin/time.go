package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"time"
)

var _ ent.Mixin = (*CreateTime)(nil)

type CreateTime struct{ mixin.Schema }

func (CreateTime) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("create_time").
			Comment("create time").
			Immutable().
			Nillable().
			DefaultFunc(time.Now().UnixMilli),
	}
}
