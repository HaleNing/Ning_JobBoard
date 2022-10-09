// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/HaleNing/Ning_JobBoard/src/Model/ent/user_info"
)

// UserInfoCreate is the builder for creating a User_info entity.
type UserInfoCreate struct {
	config
	mutation *UserInfoMutation
	hooks    []Hook
}

// SetUserName sets the "user_name" field.
func (uic *UserInfoCreate) SetUserName(s string) *UserInfoCreate {
	uic.mutation.SetUserName(s)
	return uic
}

// SetSalt sets the "salt" field.
func (uic *UserInfoCreate) SetSalt(s string) *UserInfoCreate {
	uic.mutation.SetSalt(s)
	return uic
}

// SetPasswdHash sets the "passwd_hash" field.
func (uic *UserInfoCreate) SetPasswdHash(s string) *UserInfoCreate {
	uic.mutation.SetPasswdHash(s)
	return uic
}

// Mutation returns the UserInfoMutation object of the builder.
func (uic *UserInfoCreate) Mutation() *UserInfoMutation {
	return uic.mutation
}

// Save creates the User_info in the database.
func (uic *UserInfoCreate) Save(ctx context.Context) (*User_info, error) {
	var (
		err  error
		node *User_info
	)
	if len(uic.hooks) == 0 {
		if err = uic.check(); err != nil {
			return nil, err
		}
		node, err = uic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uic.check(); err != nil {
				return nil, err
			}
			uic.mutation = mutation
			if node, err = uic.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(uic.hooks) - 1; i >= 0; i-- {
			if uic.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uic.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, uic.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*User_info)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from UserInfoMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uic *UserInfoCreate) SaveX(ctx context.Context) *User_info {
	v, err := uic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uic *UserInfoCreate) Exec(ctx context.Context) error {
	_, err := uic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uic *UserInfoCreate) ExecX(ctx context.Context) {
	if err := uic.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uic *UserInfoCreate) check() error {
	if _, ok := uic.mutation.UserName(); !ok {
		return &ValidationError{Name: "user_name", err: errors.New(`ent: missing required field "User_info.user_name"`)}
	}
	if v, ok := uic.mutation.UserName(); ok {
		if err := user_info.UserNameValidator(v); err != nil {
			return &ValidationError{Name: "user_name", err: fmt.Errorf(`ent: validator failed for field "User_info.user_name": %w`, err)}
		}
	}
	if _, ok := uic.mutation.Salt(); !ok {
		return &ValidationError{Name: "salt", err: errors.New(`ent: missing required field "User_info.salt"`)}
	}
	if v, ok := uic.mutation.Salt(); ok {
		if err := user_info.SaltValidator(v); err != nil {
			return &ValidationError{Name: "salt", err: fmt.Errorf(`ent: validator failed for field "User_info.salt": %w`, err)}
		}
	}
	if _, ok := uic.mutation.PasswdHash(); !ok {
		return &ValidationError{Name: "passwd_hash", err: errors.New(`ent: missing required field "User_info.passwd_hash"`)}
	}
	if v, ok := uic.mutation.PasswdHash(); ok {
		if err := user_info.PasswdHashValidator(v); err != nil {
			return &ValidationError{Name: "passwd_hash", err: fmt.Errorf(`ent: validator failed for field "User_info.passwd_hash": %w`, err)}
		}
	}
	return nil
}

func (uic *UserInfoCreate) sqlSave(ctx context.Context) (*User_info, error) {
	_node, _spec := uic.createSpec()
	if err := sqlgraph.CreateNode(ctx, uic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (uic *UserInfoCreate) createSpec() (*User_info, *sqlgraph.CreateSpec) {
	var (
		_node = &User_info{config: uic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: user_info.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user_info.FieldID,
			},
		}
	)
	if value, ok := uic.mutation.UserName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user_info.FieldUserName,
		})
		_node.UserName = value
	}
	if value, ok := uic.mutation.Salt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user_info.FieldSalt,
		})
		_node.Salt = value
	}
	if value, ok := uic.mutation.PasswdHash(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user_info.FieldPasswdHash,
		})
		_node.PasswdHash = value
	}
	return _node, _spec
}

// UserInfoCreateBulk is the builder for creating many User_info entities in bulk.
type UserInfoCreateBulk struct {
	config
	builders []*UserInfoCreate
}

// Save creates the User_info entities in the database.
func (uicb *UserInfoCreateBulk) Save(ctx context.Context) ([]*User_info, error) {
	specs := make([]*sqlgraph.CreateSpec, len(uicb.builders))
	nodes := make([]*User_info, len(uicb.builders))
	mutators := make([]Mutator, len(uicb.builders))
	for i := range uicb.builders {
		func(i int, root context.Context) {
			builder := uicb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserInfoMutation)
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
					_, err = mutators[i+1].Mutate(root, uicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uicb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, uicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uicb *UserInfoCreateBulk) SaveX(ctx context.Context) []*User_info {
	v, err := uicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uicb *UserInfoCreateBulk) Exec(ctx context.Context) error {
	_, err := uicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uicb *UserInfoCreateBulk) ExecX(ctx context.Context) {
	if err := uicb.Exec(ctx); err != nil {
		panic(err)
	}
}