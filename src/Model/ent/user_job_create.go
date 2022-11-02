// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/HaleNing/Ning_JobBoard/src/Model/ent/user_job"
)

// UserJobCreate is the builder for creating a User_job entity.
type UserJobCreate struct {
	config
	mutation *UserJobMutation
	hooks    []Hook
}

// SetJobID sets the "job_id" field.
func (ujc *UserJobCreate) SetJobID(i int64) *UserJobCreate {
	ujc.mutation.SetJobID(i)
	return ujc
}

// SetUserID sets the "user_id" field.
func (ujc *UserJobCreate) SetUserID(i int64) *UserJobCreate {
	ujc.mutation.SetUserID(i)
	return ujc
}

// SetCreateTime sets the "create_time" field.
func (ujc *UserJobCreate) SetCreateTime(t time.Time) *UserJobCreate {
	ujc.mutation.SetCreateTime(t)
	return ujc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (ujc *UserJobCreate) SetNillableCreateTime(t *time.Time) *UserJobCreate {
	if t != nil {
		ujc.SetCreateTime(*t)
	}
	return ujc
}

// SetUpdateTime sets the "update_time" field.
func (ujc *UserJobCreate) SetUpdateTime(t time.Time) *UserJobCreate {
	ujc.mutation.SetUpdateTime(t)
	return ujc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (ujc *UserJobCreate) SetNillableUpdateTime(t *time.Time) *UserJobCreate {
	if t != nil {
		ujc.SetUpdateTime(*t)
	}
	return ujc
}

// Mutation returns the UserJobMutation object of the builder.
func (ujc *UserJobCreate) Mutation() *UserJobMutation {
	return ujc.mutation
}

// Save creates the User_job in the database.
func (ujc *UserJobCreate) Save(ctx context.Context) (*User_job, error) {
	var (
		err  error
		node *User_job
	)
	ujc.defaults()
	if len(ujc.hooks) == 0 {
		if err = ujc.check(); err != nil {
			return nil, err
		}
		node, err = ujc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserJobMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ujc.check(); err != nil {
				return nil, err
			}
			ujc.mutation = mutation
			if node, err = ujc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ujc.hooks) - 1; i >= 0; i-- {
			if ujc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ujc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ujc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*User_job)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from UserJobMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ujc *UserJobCreate) SaveX(ctx context.Context) *User_job {
	v, err := ujc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ujc *UserJobCreate) Exec(ctx context.Context) error {
	_, err := ujc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ujc *UserJobCreate) ExecX(ctx context.Context) {
	if err := ujc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ujc *UserJobCreate) defaults() {
	if _, ok := ujc.mutation.CreateTime(); !ok {
		v := user_job.DefaultCreateTime
		ujc.mutation.SetCreateTime(v)
	}
	if _, ok := ujc.mutation.UpdateTime(); !ok {
		v := user_job.DefaultUpdateTime()
		ujc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ujc *UserJobCreate) check() error {
	if _, ok := ujc.mutation.JobID(); !ok {
		return &ValidationError{Name: "job_id", err: errors.New(`ent: missing required field "User_job.job_id"`)}
	}
	if v, ok := ujc.mutation.JobID(); ok {
		if err := user_job.JobIDValidator(v); err != nil {
			return &ValidationError{Name: "job_id", err: fmt.Errorf(`ent: validator failed for field "User_job.job_id": %w`, err)}
		}
	}
	if _, ok := ujc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "User_job.user_id"`)}
	}
	if v, ok := ujc.mutation.UserID(); ok {
		if err := user_job.UserIDValidator(v); err != nil {
			return &ValidationError{Name: "user_id", err: fmt.Errorf(`ent: validator failed for field "User_job.user_id": %w`, err)}
		}
	}
	if _, ok := ujc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "User_job.create_time"`)}
	}
	if _, ok := ujc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "User_job.update_time"`)}
	}
	return nil
}

func (ujc *UserJobCreate) sqlSave(ctx context.Context) (*User_job, error) {
	_node, _spec := ujc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ujc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ujc *UserJobCreate) createSpec() (*User_job, *sqlgraph.CreateSpec) {
	var (
		_node = &User_job{config: ujc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: user_job.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user_job.FieldID,
			},
		}
	)
	if value, ok := ujc.mutation.JobID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: user_job.FieldJobID,
		})
		_node.JobID = value
	}
	if value, ok := ujc.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: user_job.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := ujc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user_job.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := ujc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user_job.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	return _node, _spec
}

// UserJobCreateBulk is the builder for creating many User_job entities in bulk.
type UserJobCreateBulk struct {
	config
	builders []*UserJobCreate
}

// Save creates the User_job entities in the database.
func (ujcb *UserJobCreateBulk) Save(ctx context.Context) ([]*User_job, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ujcb.builders))
	nodes := make([]*User_job, len(ujcb.builders))
	mutators := make([]Mutator, len(ujcb.builders))
	for i := range ujcb.builders {
		func(i int, root context.Context) {
			builder := ujcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserJobMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ujcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ujcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ujcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ujcb *UserJobCreateBulk) SaveX(ctx context.Context) []*User_job {
	v, err := ujcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ujcb *UserJobCreateBulk) Exec(ctx context.Context) error {
	_, err := ujcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ujcb *UserJobCreateBulk) ExecX(ctx context.Context) {
	if err := ujcb.Exec(ctx); err != nil {
		panic(err)
	}
}
