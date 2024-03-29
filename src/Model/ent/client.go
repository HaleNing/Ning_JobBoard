// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/HaleNing/Ning_JobBoard/src/Model/ent/migrate"

	"github.com/HaleNing/Ning_JobBoard/src/Model/ent/job"
	"github.com/HaleNing/Ning_JobBoard/src/Model/ent/user"
	"github.com/HaleNing/Ning_JobBoard/src/Model/ent/user_info"
	"github.com/HaleNing/Ning_JobBoard/src/Model/ent/user_job"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Job is the client for interacting with the Job builders.
	Job *JobClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// User_info is the client for interacting with the User_info builders.
	User_info *User_infoClient
	// User_job is the client for interacting with the User_job builders.
	User_job *User_jobClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Job = NewJobClient(c.config)
	c.User = NewUserClient(c.config)
	c.User_info = NewUser_infoClient(c.config)
	c.User_job = NewUser_jobClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:       ctx,
		config:    cfg,
		Job:       NewJobClient(cfg),
		User:      NewUserClient(cfg),
		User_info: NewUser_infoClient(cfg),
		User_job:  NewUser_jobClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:       ctx,
		config:    cfg,
		Job:       NewJobClient(cfg),
		User:      NewUserClient(cfg),
		User_info: NewUser_infoClient(cfg),
		User_job:  NewUser_jobClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Job.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Job.Use(hooks...)
	c.User.Use(hooks...)
	c.User_info.Use(hooks...)
	c.User_job.Use(hooks...)
}

// JobClient is a client for the Job schema.
type JobClient struct {
	config
}

// NewJobClient returns a client for the Job from the given config.
func NewJobClient(c config) *JobClient {
	return &JobClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `job.Hooks(f(g(h())))`.
func (c *JobClient) Use(hooks ...Hook) {
	c.hooks.Job = append(c.hooks.Job, hooks...)
}

// Create returns a builder for creating a Job entity.
func (c *JobClient) Create() *JobCreate {
	mutation := newJobMutation(c.config, OpCreate)
	return &JobCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Job entities.
func (c *JobClient) CreateBulk(builders ...*JobCreate) *JobCreateBulk {
	return &JobCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Job.
func (c *JobClient) Update() *JobUpdate {
	mutation := newJobMutation(c.config, OpUpdate)
	return &JobUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *JobClient) UpdateOne(j *Job) *JobUpdateOne {
	mutation := newJobMutation(c.config, OpUpdateOne, withJob(j))
	return &JobUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *JobClient) UpdateOneID(id int) *JobUpdateOne {
	mutation := newJobMutation(c.config, OpUpdateOne, withJobID(id))
	return &JobUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Job.
func (c *JobClient) Delete() *JobDelete {
	mutation := newJobMutation(c.config, OpDelete)
	return &JobDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *JobClient) DeleteOne(j *Job) *JobDeleteOne {
	return c.DeleteOneID(j.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *JobClient) DeleteOneID(id int) *JobDeleteOne {
	builder := c.Delete().Where(job.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &JobDeleteOne{builder}
}

// Query returns a query builder for Job.
func (c *JobClient) Query() *JobQuery {
	return &JobQuery{
		config: c.config,
	}
}

// Get returns a Job entity by its id.
func (c *JobClient) Get(ctx context.Context, id int) (*Job, error) {
	return c.Query().Where(job.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *JobClient) GetX(ctx context.Context, id int) *Job {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *JobClient) Hooks() []Hook {
	return c.hooks.Job
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// User_infoClient is a client for the User_info schema.
type User_infoClient struct {
	config
}

// NewUser_infoClient returns a client for the User_info from the given config.
func NewUser_infoClient(c config) *User_infoClient {
	return &User_infoClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user_info.Hooks(f(g(h())))`.
func (c *User_infoClient) Use(hooks ...Hook) {
	c.hooks.User_info = append(c.hooks.User_info, hooks...)
}

// Create returns a builder for creating a User_info entity.
func (c *User_infoClient) Create() *UserInfoCreate {
	mutation := newUserInfoMutation(c.config, OpCreate)
	return &UserInfoCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User_info entities.
func (c *User_infoClient) CreateBulk(builders ...*UserInfoCreate) *UserInfoCreateBulk {
	return &UserInfoCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User_info.
func (c *User_infoClient) Update() *UserInfoUpdate {
	mutation := newUserInfoMutation(c.config, OpUpdate)
	return &UserInfoUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *User_infoClient) UpdateOne(ui *User_info) *UserInfoUpdateOne {
	mutation := newUserInfoMutation(c.config, OpUpdateOne, withUser_info(ui))
	return &UserInfoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *User_infoClient) UpdateOneID(id int) *UserInfoUpdateOne {
	mutation := newUserInfoMutation(c.config, OpUpdateOne, withUser_infoID(id))
	return &UserInfoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User_info.
func (c *User_infoClient) Delete() *UserInfoDelete {
	mutation := newUserInfoMutation(c.config, OpDelete)
	return &UserInfoDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *User_infoClient) DeleteOne(ui *User_info) *UserInfoDeleteOne {
	return c.DeleteOneID(ui.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *User_infoClient) DeleteOneID(id int) *UserInfoDeleteOne {
	builder := c.Delete().Where(user_info.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserInfoDeleteOne{builder}
}

// Query returns a query builder for User_info.
func (c *User_infoClient) Query() *UserInfoQuery {
	return &UserInfoQuery{
		config: c.config,
	}
}

// Get returns a User_info entity by its id.
func (c *User_infoClient) Get(ctx context.Context, id int) (*User_info, error) {
	return c.Query().Where(user_info.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *User_infoClient) GetX(ctx context.Context, id int) *User_info {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *User_infoClient) Hooks() []Hook {
	return c.hooks.User_info
}

// User_jobClient is a client for the User_job schema.
type User_jobClient struct {
	config
}

// NewUser_jobClient returns a client for the User_job from the given config.
func NewUser_jobClient(c config) *User_jobClient {
	return &User_jobClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user_job.Hooks(f(g(h())))`.
func (c *User_jobClient) Use(hooks ...Hook) {
	c.hooks.User_job = append(c.hooks.User_job, hooks...)
}

// Create returns a builder for creating a User_job entity.
func (c *User_jobClient) Create() *UserJobCreate {
	mutation := newUserJobMutation(c.config, OpCreate)
	return &UserJobCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User_job entities.
func (c *User_jobClient) CreateBulk(builders ...*UserJobCreate) *UserJobCreateBulk {
	return &UserJobCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User_job.
func (c *User_jobClient) Update() *UserJobUpdate {
	mutation := newUserJobMutation(c.config, OpUpdate)
	return &UserJobUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *User_jobClient) UpdateOne(uj *User_job) *UserJobUpdateOne {
	mutation := newUserJobMutation(c.config, OpUpdateOne, withUser_job(uj))
	return &UserJobUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *User_jobClient) UpdateOneID(id int) *UserJobUpdateOne {
	mutation := newUserJobMutation(c.config, OpUpdateOne, withUser_jobID(id))
	return &UserJobUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User_job.
func (c *User_jobClient) Delete() *UserJobDelete {
	mutation := newUserJobMutation(c.config, OpDelete)
	return &UserJobDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *User_jobClient) DeleteOne(uj *User_job) *UserJobDeleteOne {
	return c.DeleteOneID(uj.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *User_jobClient) DeleteOneID(id int) *UserJobDeleteOne {
	builder := c.Delete().Where(user_job.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserJobDeleteOne{builder}
}

// Query returns a query builder for User_job.
func (c *User_jobClient) Query() *UserJobQuery {
	return &UserJobQuery{
		config: c.config,
	}
}

// Get returns a User_job entity by its id.
func (c *User_jobClient) Get(ctx context.Context, id int) (*User_job, error) {
	return c.Query().Where(user_job.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *User_jobClient) GetX(ctx context.Context, id int) *User_job {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *User_jobClient) Hooks() []Hook {
	return c.hooks.User_job
}
