package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Bus_driver holds the schema definition for the Bus_driver entity.
type Bus_driver struct {
	ent.Schema
}

// Fields of the Bus_driver.
func (Bus_driver) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("user_name").NotEmpty().MaxLen(25).Comment("user name"),
		field.Time("create_time").Default(time.Now()).Comment("创建时间"),
		field.Time("update_time").Default(time.Now()).Comment("更新时间"),
		field.Int8("is_delete").Default(1).Comment("0-无效，1-有效"),
		field.Int8("user_age").NonNegative().Max(100).Comment("年龄"),
		field.Bool("sex").Comment("性别"),
		field.Int8("career_age").NonNegative().Max(100).Min(3).Comment("驾龄"),
	}
}

// Edges of the Bus_driver.
func (Bus_driver) Edges() []ent.Edge {
	return nil
}
