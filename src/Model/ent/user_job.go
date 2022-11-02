// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/HaleNing/Ning_JobBoard/src/Model/ent/user_job"
)

// User_job is the model entity for the User_job schema.
type User_job struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// job id,unique for same user id
	JobID int64 `json:"job_id,omitempty"`
	// user id
	UserID int64 `json:"user_id,omitempty"`
	// create_time
	CreateTime time.Time `json:"create_time,omitempty"`
	// update_time
	UpdateTime time.Time `json:"update_time,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User_job) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user_job.FieldID, user_job.FieldJobID, user_job.FieldUserID:
			values[i] = new(sql.NullInt64)
		case user_job.FieldCreateTime, user_job.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User_job", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User_job fields.
func (uj *User_job) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user_job.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			uj.ID = int(value.Int64)
		case user_job.FieldJobID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field job_id", values[i])
			} else if value.Valid {
				uj.JobID = value.Int64
			}
		case user_job.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				uj.UserID = value.Int64
			}
		case user_job.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				uj.CreateTime = value.Time
			}
		case user_job.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				uj.UpdateTime = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this User_job.
// Note that you need to call User_job.Unwrap() before calling this method if this User_job
// was returned from a transaction, and the transaction was committed or rolled back.
func (uj *User_job) Update() *UserJobUpdateOne {
	return (&User_jobClient{config: uj.config}).UpdateOne(uj)
}

// Unwrap unwraps the User_job entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (uj *User_job) Unwrap() *User_job {
	_tx, ok := uj.config.driver.(*txDriver)
	if !ok {
		panic("ent: User_job is not a transactional entity")
	}
	uj.config.driver = _tx.drv
	return uj
}

// String implements the fmt.Stringer.
func (uj *User_job) String() string {
	var builder strings.Builder
	builder.WriteString("User_job(")
	builder.WriteString(fmt.Sprintf("id=%v, ", uj.ID))
	builder.WriteString("job_id=")
	builder.WriteString(fmt.Sprintf("%v", uj.JobID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", uj.UserID))
	builder.WriteString(", ")
	builder.WriteString("create_time=")
	builder.WriteString(uj.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(uj.UpdateTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// User_jobs is a parsable slice of User_job.
type User_jobs []*User_job

func (uj User_jobs) config(cfg config) {
	for _i := range uj {
		uj[_i].config = cfg
	}
}