// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/HaleNing/Ning_JobBoard/src/Model/ent/predicate"
	"github.com/HaleNing/Ning_JobBoard/src/Model/ent/user_info"
)

// UserInfoUpdate is the builder for updating User_info entities.
type UserInfoUpdate struct {
	config
	hooks    []Hook
	mutation *UserInfoMutation
}

// Where appends a list predicates to the UserInfoUpdate builder.
func (uiu *UserInfoUpdate) Where(ps ...predicate.User_info) *UserInfoUpdate {
	uiu.mutation.Where(ps...)
	return uiu
}

// SetUserName sets the "user_name" field.
func (uiu *UserInfoUpdate) SetUserName(s string) *UserInfoUpdate {
	uiu.mutation.SetUserName(s)
	return uiu
}

// SetSalt sets the "salt" field.
func (uiu *UserInfoUpdate) SetSalt(s string) *UserInfoUpdate {
	uiu.mutation.SetSalt(s)
	return uiu
}

// SetPasswdHash sets the "passwd_hash" field.
func (uiu *UserInfoUpdate) SetPasswdHash(s string) *UserInfoUpdate {
	uiu.mutation.SetPasswdHash(s)
	return uiu
}

// Mutation returns the UserInfoMutation object of the builder.
func (uiu *UserInfoUpdate) Mutation() *UserInfoMutation {
	return uiu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uiu *UserInfoUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uiu.hooks) == 0 {
		if err = uiu.check(); err != nil {
			return 0, err
		}
		affected, err = uiu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uiu.check(); err != nil {
				return 0, err
			}
			uiu.mutation = mutation
			affected, err = uiu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uiu.hooks) - 1; i >= 0; i-- {
			if uiu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uiu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uiu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uiu *UserInfoUpdate) SaveX(ctx context.Context) int {
	affected, err := uiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uiu *UserInfoUpdate) Exec(ctx context.Context) error {
	_, err := uiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uiu *UserInfoUpdate) ExecX(ctx context.Context) {
	if err := uiu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uiu *UserInfoUpdate) check() error {
	if v, ok := uiu.mutation.UserName(); ok {
		if err := user_info.UserNameValidator(v); err != nil {
			return &ValidationError{Name: "user_name", err: fmt.Errorf(`ent: validator failed for field "User_info.user_name": %w`, err)}
		}
	}
	if v, ok := uiu.mutation.Salt(); ok {
		if err := user_info.SaltValidator(v); err != nil {
			return &ValidationError{Name: "salt", err: fmt.Errorf(`ent: validator failed for field "User_info.salt": %w`, err)}
		}
	}
	if v, ok := uiu.mutation.PasswdHash(); ok {
		if err := user_info.PasswdHashValidator(v); err != nil {
			return &ValidationError{Name: "passwd_hash", err: fmt.Errorf(`ent: validator failed for field "User_info.passwd_hash": %w`, err)}
		}
	}
	return nil
}

func (uiu *UserInfoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user_info.Table,
			Columns: user_info.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user_info.FieldID,
			},
		},
	}
	if ps := uiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uiu.mutation.UserName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user_info.FieldUserName,
		})
	}
	if value, ok := uiu.mutation.Salt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user_info.FieldSalt,
		})
	}
	if value, ok := uiu.mutation.PasswdHash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user_info.FieldPasswdHash,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user_info.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// UserInfoUpdateOne is the builder for updating a single User_info entity.
type UserInfoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserInfoMutation
}

// SetUserName sets the "user_name" field.
func (uiuo *UserInfoUpdateOne) SetUserName(s string) *UserInfoUpdateOne {
	uiuo.mutation.SetUserName(s)
	return uiuo
}

// SetSalt sets the "salt" field.
func (uiuo *UserInfoUpdateOne) SetSalt(s string) *UserInfoUpdateOne {
	uiuo.mutation.SetSalt(s)
	return uiuo
}

// SetPasswdHash sets the "passwd_hash" field.
func (uiuo *UserInfoUpdateOne) SetPasswdHash(s string) *UserInfoUpdateOne {
	uiuo.mutation.SetPasswdHash(s)
	return uiuo
}

// Mutation returns the UserInfoMutation object of the builder.
func (uiuo *UserInfoUpdateOne) Mutation() *UserInfoMutation {
	return uiuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uiuo *UserInfoUpdateOne) Select(field string, fields ...string) *UserInfoUpdateOne {
	uiuo.fields = append([]string{field}, fields...)
	return uiuo
}

// Save executes the query and returns the updated User_info entity.
func (uiuo *UserInfoUpdateOne) Save(ctx context.Context) (*User_info, error) {
	var (
		err  error
		node *User_info
	)
	if len(uiuo.hooks) == 0 {
		if err = uiuo.check(); err != nil {
			return nil, err
		}
		node, err = uiuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uiuo.check(); err != nil {
				return nil, err
			}
			uiuo.mutation = mutation
			node, err = uiuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uiuo.hooks) - 1; i >= 0; i-- {
			if uiuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uiuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, uiuo.mutation)
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

// SaveX is like Save, but panics if an error occurs.
func (uiuo *UserInfoUpdateOne) SaveX(ctx context.Context) *User_info {
	node, err := uiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uiuo *UserInfoUpdateOne) Exec(ctx context.Context) error {
	_, err := uiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uiuo *UserInfoUpdateOne) ExecX(ctx context.Context) {
	if err := uiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uiuo *UserInfoUpdateOne) check() error {
	if v, ok := uiuo.mutation.UserName(); ok {
		if err := user_info.UserNameValidator(v); err != nil {
			return &ValidationError{Name: "user_name", err: fmt.Errorf(`ent: validator failed for field "User_info.user_name": %w`, err)}
		}
	}
	if v, ok := uiuo.mutation.Salt(); ok {
		if err := user_info.SaltValidator(v); err != nil {
			return &ValidationError{Name: "salt", err: fmt.Errorf(`ent: validator failed for field "User_info.salt": %w`, err)}
		}
	}
	if v, ok := uiuo.mutation.PasswdHash(); ok {
		if err := user_info.PasswdHashValidator(v); err != nil {
			return &ValidationError{Name: "passwd_hash", err: fmt.Errorf(`ent: validator failed for field "User_info.passwd_hash": %w`, err)}
		}
	}
	return nil
}

func (uiuo *UserInfoUpdateOne) sqlSave(ctx context.Context) (_node *User_info, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user_info.Table,
			Columns: user_info.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user_info.FieldID,
			},
		},
	}
	id, ok := uiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User_info.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uiuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user_info.FieldID)
		for _, f := range fields {
			if !user_info.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user_info.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uiuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uiuo.mutation.UserName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user_info.FieldUserName,
		})
	}
	if value, ok := uiuo.mutation.Salt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user_info.FieldSalt,
		})
	}
	if value, ok := uiuo.mutation.PasswdHash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user_info.FieldPasswdHash,
		})
	}
	_node = &User_info{config: uiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user_info.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
