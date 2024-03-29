// Code generated by ent, DO NOT EDIT.

package generate

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ZNotify/server/app/db/ent/generate/device"
	"github.com/ZNotify/server/app/db/ent/generate/message"
	"github.com/ZNotify/server/app/db/ent/generate/user"
	"github.com/google/uuid"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (uc *UserCreate) SetCreatedAt(t time.Time) *UserCreate {
	uc.mutation.SetCreatedAt(t)
	return uc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableCreatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetCreatedAt(*t)
	}
	return uc
}

// SetUpdatedAt sets the "updated_at" field.
func (uc *UserCreate) SetUpdatedAt(t time.Time) *UserCreate {
	uc.mutation.SetUpdatedAt(t)
	return uc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableUpdatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetUpdatedAt(*t)
	}
	return uc
}

// SetSecret sets the "secret" field.
func (uc *UserCreate) SetSecret(s string) *UserCreate {
	uc.mutation.SetSecret(s)
	return uc
}

// SetGithubID sets the "githubID" field.
func (uc *UserCreate) SetGithubID(i int64) *UserCreate {
	uc.mutation.SetGithubID(i)
	return uc
}

// SetGithubName sets the "githubName" field.
func (uc *UserCreate) SetGithubName(s string) *UserCreate {
	uc.mutation.SetGithubName(s)
	return uc
}

// SetGithubLogin sets the "githubLogin" field.
func (uc *UserCreate) SetGithubLogin(s string) *UserCreate {
	uc.mutation.SetGithubLogin(s)
	return uc
}

// SetGithubOauthToken sets the "githubOauthToken" field.
func (uc *UserCreate) SetGithubOauthToken(s string) *UserCreate {
	uc.mutation.SetGithubOauthToken(s)
	return uc
}

// AddDeviceIDs adds the "devices" edge to the Device entity by IDs.
func (uc *UserCreate) AddDeviceIDs(ids ...int) *UserCreate {
	uc.mutation.AddDeviceIDs(ids...)
	return uc
}

// AddDevices adds the "devices" edges to the Device entity.
func (uc *UserCreate) AddDevices(d ...*Device) *UserCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return uc.AddDeviceIDs(ids...)
}

// AddMessageIDs adds the "messages" edge to the Message entity by IDs.
func (uc *UserCreate) AddMessageIDs(ids ...uuid.UUID) *UserCreate {
	uc.mutation.AddMessageIDs(ids...)
	return uc
}

// AddMessages adds the "messages" edges to the Message entity.
func (uc *UserCreate) AddMessages(m ...*Message) *UserCreate {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return uc.AddMessageIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	uc.defaults()
	return withHooks[*User, UserMutation](ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.CreatedAt(); !ok {
		v := user.DefaultCreatedAt()
		uc.mutation.SetCreatedAt(v)
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		v := user.DefaultUpdatedAt()
		uc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`generate: missing required field "User.created_at"`)}
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`generate: missing required field "User.updated_at"`)}
	}
	if _, ok := uc.mutation.Secret(); !ok {
		return &ValidationError{Name: "secret", err: errors.New(`generate: missing required field "User.secret"`)}
	}
	if _, ok := uc.mutation.GithubID(); !ok {
		return &ValidationError{Name: "githubID", err: errors.New(`generate: missing required field "User.githubID"`)}
	}
	if _, ok := uc.mutation.GithubName(); !ok {
		return &ValidationError{Name: "githubName", err: errors.New(`generate: missing required field "User.githubName"`)}
	}
	if _, ok := uc.mutation.GithubLogin(); !ok {
		return &ValidationError{Name: "githubLogin", err: errors.New(`generate: missing required field "User.githubLogin"`)}
	}
	if v, ok := uc.mutation.GithubLogin(); ok {
		if err := user.GithubLoginValidator(v); err != nil {
			return &ValidationError{Name: "githubLogin", err: fmt.Errorf(`generate: validator failed for field "User.githubLogin": %w`, err)}
		}
	}
	if _, ok := uc.mutation.GithubOauthToken(); !ok {
		return &ValidationError{Name: "githubOauthToken", err: errors.New(`generate: missing required field "User.githubOauthToken"`)}
	}
	if v, ok := uc.mutation.GithubOauthToken(); ok {
		if err := user.GithubOauthTokenValidator(v); err != nil {
			return &ValidationError{Name: "githubOauthToken", err: fmt.Errorf(`generate: validator failed for field "User.githubOauthToken": %w`, err)}
		}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	)
	_spec.OnConflict = uc.conflict
	if value, ok := uc.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := uc.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := uc.mutation.Secret(); ok {
		_spec.SetField(user.FieldSecret, field.TypeString, value)
		_node.Secret = value
	}
	if value, ok := uc.mutation.GithubID(); ok {
		_spec.SetField(user.FieldGithubID, field.TypeInt64, value)
		_node.GithubID = value
	}
	if value, ok := uc.mutation.GithubName(); ok {
		_spec.SetField(user.FieldGithubName, field.TypeString, value)
		_node.GithubName = value
	}
	if value, ok := uc.mutation.GithubLogin(); ok {
		_spec.SetField(user.FieldGithubLogin, field.TypeString, value)
		_node.GithubLogin = value
	}
	if value, ok := uc.mutation.GithubOauthToken(); ok {
		_spec.SetField(user.FieldGithubOauthToken, field.TypeString, value)
		_node.GithubOauthToken = value
	}
	if nodes := uc.mutation.DevicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.DevicesTable,
			Columns: []string{user.DevicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: device.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.MessagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.MessagesTable,
			Columns: []string{user.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: message.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.User.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (uc *UserCreate) OnConflict(opts ...sql.ConflictOption) *UserUpsertOne {
	uc.conflict = opts
	return &UserUpsertOne{
		create: uc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (uc *UserCreate) OnConflictColumns(columns ...string) *UserUpsertOne {
	uc.conflict = append(uc.conflict, sql.ConflictColumns(columns...))
	return &UserUpsertOne{
		create: uc,
	}
}

type (
	// UserUpsertOne is the builder for "upsert"-ing
	//  one User node.
	UserUpsertOne struct {
		create *UserCreate
	}

	// UserUpsert is the "OnConflict" setter.
	UserUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *UserUpsert) SetUpdatedAt(v time.Time) *UserUpsert {
	u.Set(user.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *UserUpsert) UpdateUpdatedAt() *UserUpsert {
	u.SetExcluded(user.FieldUpdatedAt)
	return u
}

// SetSecret sets the "secret" field.
func (u *UserUpsert) SetSecret(v string) *UserUpsert {
	u.Set(user.FieldSecret, v)
	return u
}

// UpdateSecret sets the "secret" field to the value that was provided on create.
func (u *UserUpsert) UpdateSecret() *UserUpsert {
	u.SetExcluded(user.FieldSecret)
	return u
}

// SetGithubID sets the "githubID" field.
func (u *UserUpsert) SetGithubID(v int64) *UserUpsert {
	u.Set(user.FieldGithubID, v)
	return u
}

// UpdateGithubID sets the "githubID" field to the value that was provided on create.
func (u *UserUpsert) UpdateGithubID() *UserUpsert {
	u.SetExcluded(user.FieldGithubID)
	return u
}

// AddGithubID adds v to the "githubID" field.
func (u *UserUpsert) AddGithubID(v int64) *UserUpsert {
	u.Add(user.FieldGithubID, v)
	return u
}

// SetGithubName sets the "githubName" field.
func (u *UserUpsert) SetGithubName(v string) *UserUpsert {
	u.Set(user.FieldGithubName, v)
	return u
}

// UpdateGithubName sets the "githubName" field to the value that was provided on create.
func (u *UserUpsert) UpdateGithubName() *UserUpsert {
	u.SetExcluded(user.FieldGithubName)
	return u
}

// SetGithubLogin sets the "githubLogin" field.
func (u *UserUpsert) SetGithubLogin(v string) *UserUpsert {
	u.Set(user.FieldGithubLogin, v)
	return u
}

// UpdateGithubLogin sets the "githubLogin" field to the value that was provided on create.
func (u *UserUpsert) UpdateGithubLogin() *UserUpsert {
	u.SetExcluded(user.FieldGithubLogin)
	return u
}

// SetGithubOauthToken sets the "githubOauthToken" field.
func (u *UserUpsert) SetGithubOauthToken(v string) *UserUpsert {
	u.Set(user.FieldGithubOauthToken, v)
	return u
}

// UpdateGithubOauthToken sets the "githubOauthToken" field to the value that was provided on create.
func (u *UserUpsert) UpdateGithubOauthToken() *UserUpsert {
	u.SetExcluded(user.FieldGithubOauthToken)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *UserUpsertOne) UpdateNewValues() *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(user.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.User.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *UserUpsertOne) Ignore() *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserUpsertOne) DoNothing() *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserCreate.OnConflict
// documentation for more info.
func (u *UserUpsertOne) Update(set func(*UserUpsert)) *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *UserUpsertOne) SetUpdatedAt(v time.Time) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateUpdatedAt() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetSecret sets the "secret" field.
func (u *UserUpsertOne) SetSecret(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetSecret(v)
	})
}

// UpdateSecret sets the "secret" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateSecret() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateSecret()
	})
}

// SetGithubID sets the "githubID" field.
func (u *UserUpsertOne) SetGithubID(v int64) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetGithubID(v)
	})
}

// AddGithubID adds v to the "githubID" field.
func (u *UserUpsertOne) AddGithubID(v int64) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.AddGithubID(v)
	})
}

// UpdateGithubID sets the "githubID" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateGithubID() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateGithubID()
	})
}

// SetGithubName sets the "githubName" field.
func (u *UserUpsertOne) SetGithubName(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetGithubName(v)
	})
}

// UpdateGithubName sets the "githubName" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateGithubName() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateGithubName()
	})
}

// SetGithubLogin sets the "githubLogin" field.
func (u *UserUpsertOne) SetGithubLogin(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetGithubLogin(v)
	})
}

// UpdateGithubLogin sets the "githubLogin" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateGithubLogin() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateGithubLogin()
	})
}

// SetGithubOauthToken sets the "githubOauthToken" field.
func (u *UserUpsertOne) SetGithubOauthToken(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetGithubOauthToken(v)
	})
}

// UpdateGithubOauthToken sets the "githubOauthToken" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateGithubOauthToken() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateGithubOauthToken()
	})
}

// Exec executes the query.
func (u *UserUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("generate: missing options for UserCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *UserUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *UserUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	builders []*UserCreate
	conflict []sql.ConflictOption
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
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
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ucb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.User.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (ucb *UserCreateBulk) OnConflict(opts ...sql.ConflictOption) *UserUpsertBulk {
	ucb.conflict = opts
	return &UserUpsertBulk{
		create: ucb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ucb *UserCreateBulk) OnConflictColumns(columns ...string) *UserUpsertBulk {
	ucb.conflict = append(ucb.conflict, sql.ConflictColumns(columns...))
	return &UserUpsertBulk{
		create: ucb,
	}
}

// UserUpsertBulk is the builder for "upsert"-ing
// a bulk of User nodes.
type UserUpsertBulk struct {
	create *UserCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *UserUpsertBulk) UpdateNewValues() *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(user.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *UserUpsertBulk) Ignore() *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserUpsertBulk) DoNothing() *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserCreateBulk.OnConflict
// documentation for more info.
func (u *UserUpsertBulk) Update(set func(*UserUpsert)) *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *UserUpsertBulk) SetUpdatedAt(v time.Time) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateUpdatedAt() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetSecret sets the "secret" field.
func (u *UserUpsertBulk) SetSecret(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetSecret(v)
	})
}

// UpdateSecret sets the "secret" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateSecret() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateSecret()
	})
}

// SetGithubID sets the "githubID" field.
func (u *UserUpsertBulk) SetGithubID(v int64) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetGithubID(v)
	})
}

// AddGithubID adds v to the "githubID" field.
func (u *UserUpsertBulk) AddGithubID(v int64) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.AddGithubID(v)
	})
}

// UpdateGithubID sets the "githubID" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateGithubID() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateGithubID()
	})
}

// SetGithubName sets the "githubName" field.
func (u *UserUpsertBulk) SetGithubName(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetGithubName(v)
	})
}

// UpdateGithubName sets the "githubName" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateGithubName() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateGithubName()
	})
}

// SetGithubLogin sets the "githubLogin" field.
func (u *UserUpsertBulk) SetGithubLogin(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetGithubLogin(v)
	})
}

// UpdateGithubLogin sets the "githubLogin" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateGithubLogin() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateGithubLogin()
	})
}

// SetGithubOauthToken sets the "githubOauthToken" field.
func (u *UserUpsertBulk) SetGithubOauthToken(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetGithubOauthToken(v)
	})
}

// UpdateGithubOauthToken sets the "githubOauthToken" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateGithubOauthToken() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateGithubOauthToken()
	})
}

// Exec executes the query.
func (u *UserUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("generate: OnConflict was set for builder %d. Set it on the UserCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("generate: missing options for UserCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
