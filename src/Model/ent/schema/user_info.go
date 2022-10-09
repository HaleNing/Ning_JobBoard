package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User_info holds the schema definition for the User_info entity.
type User_info struct {
	ent.Schema
}

// Fields of the User_info.
func (User_info) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_name").NotEmpty().Unique().Comment("user_name"),
		field.String("salt").NotEmpty().Comment("salt will with passwd hash to result"),
		field.String("passwd_hash").NotEmpty().Comment("passwd hash res"),
	}
}

// Edges of the User_info.
func (User_info) Edges() []ent.Edge {
	return nil
}
