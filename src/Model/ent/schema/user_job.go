package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// User_job holds the schema definition for the User_job entity.
type User_job struct {
	ent.Schema
}

// Fields of the User_job.
func (User_job) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("job_id").NonNegative().Comment("job id,unique for same user id "),
		field.Int64("user_id").NonNegative().Comment("user id"),
		field.Time("create_time").Default(time.Now()).Comment("create_time"),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now).Comment("update_time"),
	}
}

// Edges of the User_job.
func (User_job) Edges() []ent.Edge {
	return nil
}
