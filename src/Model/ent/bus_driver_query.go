// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/HaleNing/bustrack/src/Model/ent/bus_driver"
	"github.com/HaleNing/bustrack/src/Model/ent/predicate"
)

// BusDriverQuery is the builder for querying Bus_driver entities.
type BusDriverQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Bus_driver
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BusDriverQuery builder.
func (bdq *BusDriverQuery) Where(ps ...predicate.Bus_driver) *BusDriverQuery {
	bdq.predicates = append(bdq.predicates, ps...)
	return bdq
}

// Limit adds a limit step to the query.
func (bdq *BusDriverQuery) Limit(limit int) *BusDriverQuery {
	bdq.limit = &limit
	return bdq
}

// Offset adds an offset step to the query.
func (bdq *BusDriverQuery) Offset(offset int) *BusDriverQuery {
	bdq.offset = &offset
	return bdq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bdq *BusDriverQuery) Unique(unique bool) *BusDriverQuery {
	bdq.unique = &unique
	return bdq
}

// Order adds an order step to the query.
func (bdq *BusDriverQuery) Order(o ...OrderFunc) *BusDriverQuery {
	bdq.order = append(bdq.order, o...)
	return bdq
}

// First returns the first Bus_driver entity from the query.
// Returns a *NotFoundError when no Bus_driver was found.
func (bdq *BusDriverQuery) First(ctx context.Context) (*Bus_driver, error) {
	nodes, err := bdq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{bus_driver.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bdq *BusDriverQuery) FirstX(ctx context.Context) *Bus_driver {
	node, err := bdq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Bus_driver ID from the query.
// Returns a *NotFoundError when no Bus_driver ID was found.
func (bdq *BusDriverQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = bdq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{bus_driver.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bdq *BusDriverQuery) FirstIDX(ctx context.Context) int64 {
	id, err := bdq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Bus_driver entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Bus_driver entity is found.
// Returns a *NotFoundError when no Bus_driver entities are found.
func (bdq *BusDriverQuery) Only(ctx context.Context) (*Bus_driver, error) {
	nodes, err := bdq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{bus_driver.Label}
	default:
		return nil, &NotSingularError{bus_driver.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bdq *BusDriverQuery) OnlyX(ctx context.Context) *Bus_driver {
	node, err := bdq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Bus_driver ID in the query.
// Returns a *NotSingularError when more than one Bus_driver ID is found.
// Returns a *NotFoundError when no entities are found.
func (bdq *BusDriverQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = bdq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{bus_driver.Label}
	default:
		err = &NotSingularError{bus_driver.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bdq *BusDriverQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := bdq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Bus_drivers.
func (bdq *BusDriverQuery) All(ctx context.Context) ([]*Bus_driver, error) {
	if err := bdq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return bdq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (bdq *BusDriverQuery) AllX(ctx context.Context) []*Bus_driver {
	nodes, err := bdq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Bus_driver IDs.
func (bdq *BusDriverQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := bdq.Select(bus_driver.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bdq *BusDriverQuery) IDsX(ctx context.Context) []int64 {
	ids, err := bdq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bdq *BusDriverQuery) Count(ctx context.Context) (int, error) {
	if err := bdq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return bdq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (bdq *BusDriverQuery) CountX(ctx context.Context) int {
	count, err := bdq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bdq *BusDriverQuery) Exist(ctx context.Context) (bool, error) {
	if err := bdq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return bdq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (bdq *BusDriverQuery) ExistX(ctx context.Context) bool {
	exist, err := bdq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BusDriverQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bdq *BusDriverQuery) Clone() *BusDriverQuery {
	if bdq == nil {
		return nil
	}
	return &BusDriverQuery{
		config:     bdq.config,
		limit:      bdq.limit,
		offset:     bdq.offset,
		order:      append([]OrderFunc{}, bdq.order...),
		predicates: append([]predicate.Bus_driver{}, bdq.predicates...),
		// clone intermediate query.
		sql:    bdq.sql.Clone(),
		path:   bdq.path,
		unique: bdq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserName string `json:"user_name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.BusDriver.Query().
//		GroupBy(bus_driver.FieldUserName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (bdq *BusDriverQuery) GroupBy(field string, fields ...string) *BusDriverGroupBy {
	grbuild := &BusDriverGroupBy{config: bdq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := bdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return bdq.sqlQuery(ctx), nil
	}
	grbuild.label = bus_driver.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UserName string `json:"user_name,omitempty"`
//	}
//
//	client.BusDriver.Query().
//		Select(bus_driver.FieldUserName).
//		Scan(ctx, &v)
//
func (bdq *BusDriverQuery) Select(fields ...string) *BusDriverSelect {
	bdq.fields = append(bdq.fields, fields...)
	selbuild := &BusDriverSelect{BusDriverQuery: bdq}
	selbuild.label = bus_driver.Label
	selbuild.flds, selbuild.scan = &bdq.fields, selbuild.Scan
	return selbuild
}

func (bdq *BusDriverQuery) prepareQuery(ctx context.Context) error {
	for _, f := range bdq.fields {
		if !bus_driver.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if bdq.path != nil {
		prev, err := bdq.path(ctx)
		if err != nil {
			return err
		}
		bdq.sql = prev
	}
	return nil
}

func (bdq *BusDriverQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Bus_driver, error) {
	var (
		nodes = []*Bus_driver{}
		_spec = bdq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*Bus_driver).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &Bus_driver{config: bdq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bdq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (bdq *BusDriverQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bdq.querySpec()
	_spec.Node.Columns = bdq.fields
	if len(bdq.fields) > 0 {
		_spec.Unique = bdq.unique != nil && *bdq.unique
	}
	return sqlgraph.CountNodes(ctx, bdq.driver, _spec)
}

func (bdq *BusDriverQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := bdq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (bdq *BusDriverQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   bus_driver.Table,
			Columns: bus_driver.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: bus_driver.FieldID,
			},
		},
		From:   bdq.sql,
		Unique: true,
	}
	if unique := bdq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := bdq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, bus_driver.FieldID)
		for i := range fields {
			if fields[i] != bus_driver.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bdq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bdq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bdq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bdq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bdq *BusDriverQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bdq.driver.Dialect())
	t1 := builder.Table(bus_driver.Table)
	columns := bdq.fields
	if len(columns) == 0 {
		columns = bus_driver.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bdq.sql != nil {
		selector = bdq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bdq.unique != nil && *bdq.unique {
		selector.Distinct()
	}
	for _, p := range bdq.predicates {
		p(selector)
	}
	for _, p := range bdq.order {
		p(selector)
	}
	if offset := bdq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bdq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BusDriverGroupBy is the group-by builder for Bus_driver entities.
type BusDriverGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bdgb *BusDriverGroupBy) Aggregate(fns ...AggregateFunc) *BusDriverGroupBy {
	bdgb.fns = append(bdgb.fns, fns...)
	return bdgb
}

// Scan applies the group-by query and scans the result into the given value.
func (bdgb *BusDriverGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := bdgb.path(ctx)
	if err != nil {
		return err
	}
	bdgb.sql = query
	return bdgb.sqlScan(ctx, v)
}

func (bdgb *BusDriverGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range bdgb.fields {
		if !bus_driver.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := bdgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bdgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (bdgb *BusDriverGroupBy) sqlQuery() *sql.Selector {
	selector := bdgb.sql.Select()
	aggregation := make([]string, 0, len(bdgb.fns))
	for _, fn := range bdgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(bdgb.fields)+len(bdgb.fns))
		for _, f := range bdgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(bdgb.fields...)...)
}

// BusDriverSelect is the builder for selecting fields of BusDriver entities.
type BusDriverSelect struct {
	*BusDriverQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (bds *BusDriverSelect) Scan(ctx context.Context, v interface{}) error {
	if err := bds.prepareQuery(ctx); err != nil {
		return err
	}
	bds.sql = bds.BusDriverQuery.sqlQuery(ctx)
	return bds.sqlScan(ctx, v)
}

func (bds *BusDriverSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := bds.sql.Query()
	if err := bds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
