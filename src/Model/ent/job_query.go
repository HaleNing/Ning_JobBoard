// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/HaleNing/bustrack/src/Model/ent/job"
	"github.com/HaleNing/bustrack/src/Model/ent/predicate"
)

// JobQuery is the builder for querying Job entities.
type JobQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Job
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the JobQuery builder.
func (jq *JobQuery) Where(ps ...predicate.Job) *JobQuery {
	jq.predicates = append(jq.predicates, ps...)
	return jq
}

// Limit adds a limit step to the query.
func (jq *JobQuery) Limit(limit int) *JobQuery {
	jq.limit = &limit
	return jq
}

// Offset adds an offset step to the query.
func (jq *JobQuery) Offset(offset int) *JobQuery {
	jq.offset = &offset
	return jq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (jq *JobQuery) Unique(unique bool) *JobQuery {
	jq.unique = &unique
	return jq
}

// Order adds an order step to the query.
func (jq *JobQuery) Order(o ...OrderFunc) *JobQuery {
	jq.order = append(jq.order, o...)
	return jq
}

// First returns the first Job entity from the query.
// Returns a *NotFoundError when no Job was found.
func (jq *JobQuery) First(ctx context.Context) (*Job, error) {
	nodes, err := jq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{job.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (jq *JobQuery) FirstX(ctx context.Context) *Job {
	node, err := jq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Job ID from the query.
// Returns a *NotFoundError when no Job ID was found.
func (jq *JobQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = jq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{job.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (jq *JobQuery) FirstIDX(ctx context.Context) int {
	id, err := jq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Job entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Job entity is found.
// Returns a *NotFoundError when no Job entities are found.
func (jq *JobQuery) Only(ctx context.Context) (*Job, error) {
	nodes, err := jq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{job.Label}
	default:
		return nil, &NotSingularError{job.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (jq *JobQuery) OnlyX(ctx context.Context) *Job {
	node, err := jq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Job ID in the query.
// Returns a *NotSingularError when more than one Job ID is found.
// Returns a *NotFoundError when no entities are found.
func (jq *JobQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = jq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{job.Label}
	default:
		err = &NotSingularError{job.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (jq *JobQuery) OnlyIDX(ctx context.Context) int {
	id, err := jq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Jobs.
func (jq *JobQuery) All(ctx context.Context) ([]*Job, error) {
	if err := jq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return jq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (jq *JobQuery) AllX(ctx context.Context) []*Job {
	nodes, err := jq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Job IDs.
func (jq *JobQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := jq.Select(job.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (jq *JobQuery) IDsX(ctx context.Context) []int {
	ids, err := jq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (jq *JobQuery) Count(ctx context.Context) (int, error) {
	if err := jq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return jq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (jq *JobQuery) CountX(ctx context.Context) int {
	count, err := jq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (jq *JobQuery) Exist(ctx context.Context) (bool, error) {
	if err := jq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return jq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (jq *JobQuery) ExistX(ctx context.Context) bool {
	exist, err := jq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the JobQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (jq *JobQuery) Clone() *JobQuery {
	if jq == nil {
		return nil
	}
	return &JobQuery{
		config:     jq.config,
		limit:      jq.limit,
		offset:     jq.offset,
		order:      append([]OrderFunc{}, jq.order...),
		predicates: append([]predicate.Job{}, jq.predicates...),
		// clone intermediate query.
		sql:    jq.sql.Clone(),
		path:   jq.path,
		unique: jq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		JobName string `json:"job_name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Job.Query().
//		GroupBy(job.FieldJobName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (jq *JobQuery) GroupBy(field string, fields ...string) *JobGroupBy {
	grbuild := &JobGroupBy{config: jq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := jq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return jq.sqlQuery(ctx), nil
	}
	grbuild.label = job.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		JobName string `json:"job_name,omitempty"`
//	}
//
//	client.Job.Query().
//		Select(job.FieldJobName).
//		Scan(ctx, &v)
func (jq *JobQuery) Select(fields ...string) *JobSelect {
	jq.fields = append(jq.fields, fields...)
	selbuild := &JobSelect{JobQuery: jq}
	selbuild.label = job.Label
	selbuild.flds, selbuild.scan = &jq.fields, selbuild.Scan
	return selbuild
}

func (jq *JobQuery) prepareQuery(ctx context.Context) error {
	for _, f := range jq.fields {
		if !job.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if jq.path != nil {
		prev, err := jq.path(ctx)
		if err != nil {
			return err
		}
		jq.sql = prev
	}
	return nil
}

func (jq *JobQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Job, error) {
	var (
		nodes = []*Job{}
		_spec = jq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Job).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Job{config: jq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, jq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (jq *JobQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := jq.querySpec()
	_spec.Node.Columns = jq.fields
	if len(jq.fields) > 0 {
		_spec.Unique = jq.unique != nil && *jq.unique
	}
	return sqlgraph.CountNodes(ctx, jq.driver, _spec)
}

func (jq *JobQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := jq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (jq *JobQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   job.Table,
			Columns: job.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: job.FieldID,
			},
		},
		From:   jq.sql,
		Unique: true,
	}
	if unique := jq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := jq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, job.FieldID)
		for i := range fields {
			if fields[i] != job.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := jq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := jq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := jq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := jq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (jq *JobQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(jq.driver.Dialect())
	t1 := builder.Table(job.Table)
	columns := jq.fields
	if len(columns) == 0 {
		columns = job.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if jq.sql != nil {
		selector = jq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if jq.unique != nil && *jq.unique {
		selector.Distinct()
	}
	for _, p := range jq.predicates {
		p(selector)
	}
	for _, p := range jq.order {
		p(selector)
	}
	if offset := jq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := jq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// JobGroupBy is the group-by builder for Job entities.
type JobGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (jgb *JobGroupBy) Aggregate(fns ...AggregateFunc) *JobGroupBy {
	jgb.fns = append(jgb.fns, fns...)
	return jgb
}

// Scan applies the group-by query and scans the result into the given value.
func (jgb *JobGroupBy) Scan(ctx context.Context, v any) error {
	query, err := jgb.path(ctx)
	if err != nil {
		return err
	}
	jgb.sql = query
	return jgb.sqlScan(ctx, v)
}

func (jgb *JobGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range jgb.fields {
		if !job.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := jgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := jgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (jgb *JobGroupBy) sqlQuery() *sql.Selector {
	selector := jgb.sql.Select()
	aggregation := make([]string, 0, len(jgb.fns))
	for _, fn := range jgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(jgb.fields)+len(jgb.fns))
		for _, f := range jgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(jgb.fields...)...)
}

// JobSelect is the builder for selecting fields of Job entities.
type JobSelect struct {
	*JobQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (js *JobSelect) Scan(ctx context.Context, v any) error {
	if err := js.prepareQuery(ctx); err != nil {
		return err
	}
	js.sql = js.JobQuery.sqlQuery(ctx)
	return js.sqlScan(ctx, v)
}

func (js *JobSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := js.sql.Query()
	if err := js.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}