package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Bus holds the schema definition for the Bus entity.
type Bus struct {
	ent.Schema
}

// Fields of the Bus.
func (Bus) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("bus_name").NotEmpty().Comment("bus name"),
		field.Time("create_time").Default(time.Now()).Comment("创建时间"),
		field.Time("update_time").Default(time.Now()).Comment("更新时间"),
		field.Int64("user_id").Comment("用户id"),
		field.Int8("is_delete").Default(1).Comment("0-无效，1-有效"),
	}
}

// Edges of the Bus.
func (Bus) Edges() []ent.Edge {
	return nil
}
