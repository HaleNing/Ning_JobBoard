package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int8("role").Default(0).Comment("role of user 0-user，1-administer"),
		field.String("user_name").NotEmpty().MaxLen(16).Unique().Comment("user name,unique and not empty"),
		field.Bool("is_exist").Default(true).Comment("user delete or not"),
		field.Time("create_time").Default(time.Now()).Comment("create_time"),
		field.Time("update_time").Default(time.Now).
			UpdateDefault(time.Now).Comment("update_time"),
		field.String("user_title").Comment("user title,like software engineer"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
