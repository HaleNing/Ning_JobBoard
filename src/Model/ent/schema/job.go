package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Job holds the schema definition for the Job entity.
type Job struct {
	ent.Schema
}

// Fields of the Job.
func (Job) Fields() []ent.Field {
	return []ent.Field{
		field.String("job_name").NotEmpty().Comment("job name"),
		field.String("company_name").NotEmpty().Comment("company name"),
		field.Bool("is_exist").Default(true).Comment("delete or not"),
		field.Text("description").NotEmpty().Comment("job desc"),
		field.Bool("is_remote").Comment("job is remote or not"),
		field.Int8("exp").Positive().Max(100).Comment("job exp"),
		field.String("area").NotEmpty().Comment("job area"),
		field.Time("create_time").Default(time.Now()).Comment("create_time"),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now).Comment("update_time"),
	}
}

// Edges of the Job.
func (Job) Edges() []ent.Edge {
	return nil
}
