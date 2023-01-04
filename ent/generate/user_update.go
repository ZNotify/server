// Code generated by ent, DO NOT EDIT.

package generate

import (
	"context"
	"errors"
	"fmt"
	"notify-api/ent/generate/device"
	"notify-api/ent/generate/message"
	"notify-api/ent/generate/predicate"
	"notify-api/ent/generate/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUpdatedAt sets the "updated_at" field.
func (uu *UserUpdate) SetUpdatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetUpdatedAt(t)
	return uu
}

// SetSecret sets the "secret" field.
func (uu *UserUpdate) SetSecret(s string) *UserUpdate {
	uu.mutation.SetSecret(s)
	return uu
}

// SetGithubID sets the "githubID" field.
func (uu *UserUpdate) SetGithubID(s string) *UserUpdate {
	uu.mutation.SetGithubID(s)
	return uu
}

// SetGithubName sets the "githubName" field.
func (uu *UserUpdate) SetGithubName(s string) *UserUpdate {
	uu.mutation.SetGithubName(s)
	return uu
}

// SetGithubLogin sets the "githubLogin" field.
func (uu *UserUpdate) SetGithubLogin(s string) *UserUpdate {
	uu.mutation.SetGithubLogin(s)
	return uu
}

// SetGithubOauthToken sets the "githubOauthToken" field.
func (uu *UserUpdate) SetGithubOauthToken(s string) *UserUpdate {
	uu.mutation.SetGithubOauthToken(s)
	return uu
}

// AddDeviceIDs adds the "devices" edge to the Device entity by IDs.
func (uu *UserUpdate) AddDeviceIDs(ids ...int) *UserUpdate {
	uu.mutation.AddDeviceIDs(ids...)
	return uu
}

// AddDevices adds the "devices" edges to the Device entity.
func (uu *UserUpdate) AddDevices(d ...*Device) *UserUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return uu.AddDeviceIDs(ids...)
}

// AddMessageIDs adds the "messages" edge to the Message entity by IDs.
func (uu *UserUpdate) AddMessageIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddMessageIDs(ids...)
	return uu
}

// AddMessages adds the "messages" edges to the Message entity.
func (uu *UserUpdate) AddMessages(m ...*Message) *UserUpdate {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return uu.AddMessageIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearDevices clears all "devices" edges to the Device entity.
func (uu *UserUpdate) ClearDevices() *UserUpdate {
	uu.mutation.ClearDevices()
	return uu
}

// RemoveDeviceIDs removes the "devices" edge to Device entities by IDs.
func (uu *UserUpdate) RemoveDeviceIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveDeviceIDs(ids...)
	return uu
}

// RemoveDevices removes "devices" edges to Device entities.
func (uu *UserUpdate) RemoveDevices(d ...*Device) *UserUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return uu.RemoveDeviceIDs(ids...)
}

// ClearMessages clears all "messages" edges to the Message entity.
func (uu *UserUpdate) ClearMessages() *UserUpdate {
	uu.mutation.ClearMessages()
	return uu
}

// RemoveMessageIDs removes the "messages" edge to Message entities by IDs.
func (uu *UserUpdate) RemoveMessageIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveMessageIDs(ids...)
	return uu
}

// RemoveMessages removes "messages" edges to Message entities.
func (uu *UserUpdate) RemoveMessages(m ...*Message) *UserUpdate {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return uu.RemoveMessageIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	uu.defaults()
	if len(uu.hooks) == 0 {
		if err = uu.check(); err != nil {
			return 0, err
		}
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uu.check(); err != nil {
				return 0, err
			}
			uu.mutation = mutation
			affected, err = uu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uu.hooks) - 1; i >= 0; i-- {
			if uu.hooks[i] == nil {
				return 0, fmt.Errorf("generate: uninitialized hook (forgotten import generate/runtime?)")
			}
			mut = uu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uu *UserUpdate) defaults() {
	if _, ok := uu.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.GithubLogin(); ok {
		if err := user.GithubLoginValidator(v); err != nil {
			return &ValidationError{Name: "githubLogin", err: fmt.Errorf(`generate: validator failed for field "User.githubLogin": %w`, err)}
		}
	}
	if v, ok := uu.mutation.GithubOauthToken(); ok {
		if err := user.GithubOauthTokenValidator(v); err != nil {
			return &ValidationError{Name: "githubOauthToken", err: fmt.Errorf(`generate: validator failed for field "User.githubOauthToken": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uu.mutation.Secret(); ok {
		_spec.SetField(user.FieldSecret, field.TypeString, value)
	}
	if value, ok := uu.mutation.GithubID(); ok {
		_spec.SetField(user.FieldGithubID, field.TypeString, value)
	}
	if value, ok := uu.mutation.GithubName(); ok {
		_spec.SetField(user.FieldGithubName, field.TypeString, value)
	}
	if value, ok := uu.mutation.GithubLogin(); ok {
		_spec.SetField(user.FieldGithubLogin, field.TypeString, value)
	}
	if value, ok := uu.mutation.GithubOauthToken(); ok {
		_spec.SetField(user.FieldGithubOauthToken, field.TypeString, value)
	}
	if uu.mutation.DevicesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedDevicesIDs(); len(nodes) > 0 && !uu.mutation.DevicesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.DevicesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.MessagesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedMessagesIDs(); len(nodes) > 0 && !uu.mutation.MessagesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.MessagesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (uuo *UserUpdateOne) SetUpdatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdatedAt(t)
	return uuo
}

// SetSecret sets the "secret" field.
func (uuo *UserUpdateOne) SetSecret(s string) *UserUpdateOne {
	uuo.mutation.SetSecret(s)
	return uuo
}

// SetGithubID sets the "githubID" field.
func (uuo *UserUpdateOne) SetGithubID(s string) *UserUpdateOne {
	uuo.mutation.SetGithubID(s)
	return uuo
}

// SetGithubName sets the "githubName" field.
func (uuo *UserUpdateOne) SetGithubName(s string) *UserUpdateOne {
	uuo.mutation.SetGithubName(s)
	return uuo
}

// SetGithubLogin sets the "githubLogin" field.
func (uuo *UserUpdateOne) SetGithubLogin(s string) *UserUpdateOne {
	uuo.mutation.SetGithubLogin(s)
	return uuo
}

// SetGithubOauthToken sets the "githubOauthToken" field.
func (uuo *UserUpdateOne) SetGithubOauthToken(s string) *UserUpdateOne {
	uuo.mutation.SetGithubOauthToken(s)
	return uuo
}

// AddDeviceIDs adds the "devices" edge to the Device entity by IDs.
func (uuo *UserUpdateOne) AddDeviceIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddDeviceIDs(ids...)
	return uuo
}

// AddDevices adds the "devices" edges to the Device entity.
func (uuo *UserUpdateOne) AddDevices(d ...*Device) *UserUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return uuo.AddDeviceIDs(ids...)
}

// AddMessageIDs adds the "messages" edge to the Message entity by IDs.
func (uuo *UserUpdateOne) AddMessageIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddMessageIDs(ids...)
	return uuo
}

// AddMessages adds the "messages" edges to the Message entity.
func (uuo *UserUpdateOne) AddMessages(m ...*Message) *UserUpdateOne {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return uuo.AddMessageIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearDevices clears all "devices" edges to the Device entity.
func (uuo *UserUpdateOne) ClearDevices() *UserUpdateOne {
	uuo.mutation.ClearDevices()
	return uuo
}

// RemoveDeviceIDs removes the "devices" edge to Device entities by IDs.
func (uuo *UserUpdateOne) RemoveDeviceIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveDeviceIDs(ids...)
	return uuo
}

// RemoveDevices removes "devices" edges to Device entities.
func (uuo *UserUpdateOne) RemoveDevices(d ...*Device) *UserUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return uuo.RemoveDeviceIDs(ids...)
}

// ClearMessages clears all "messages" edges to the Message entity.
func (uuo *UserUpdateOne) ClearMessages() *UserUpdateOne {
	uuo.mutation.ClearMessages()
	return uuo
}

// RemoveMessageIDs removes the "messages" edge to Message entities by IDs.
func (uuo *UserUpdateOne) RemoveMessageIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveMessageIDs(ids...)
	return uuo
}

// RemoveMessages removes "messages" edges to Message entities.
func (uuo *UserUpdateOne) RemoveMessages(m ...*Message) *UserUpdateOne {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return uuo.RemoveMessageIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	uuo.defaults()
	if len(uuo.hooks) == 0 {
		if err = uuo.check(); err != nil {
			return nil, err
		}
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uuo.check(); err != nil {
				return nil, err
			}
			uuo.mutation = mutation
			node, err = uuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuo.hooks) - 1; i >= 0; i-- {
			if uuo.hooks[i] == nil {
				return nil, fmt.Errorf("generate: uninitialized hook (forgotten import generate/runtime?)")
			}
			mut = uuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, uuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*User)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from UserMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uuo *UserUpdateOne) defaults() {
	if _, ok := uuo.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.GithubLogin(); ok {
		if err := user.GithubLoginValidator(v); err != nil {
			return &ValidationError{Name: "githubLogin", err: fmt.Errorf(`generate: validator failed for field "User.githubLogin": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.GithubOauthToken(); ok {
		if err := user.GithubOauthTokenValidator(v); err != nil {
			return &ValidationError{Name: "githubOauthToken", err: fmt.Errorf(`generate: validator failed for field "User.githubOauthToken": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generate: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generate: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.Secret(); ok {
		_spec.SetField(user.FieldSecret, field.TypeString, value)
	}
	if value, ok := uuo.mutation.GithubID(); ok {
		_spec.SetField(user.FieldGithubID, field.TypeString, value)
	}
	if value, ok := uuo.mutation.GithubName(); ok {
		_spec.SetField(user.FieldGithubName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.GithubLogin(); ok {
		_spec.SetField(user.FieldGithubLogin, field.TypeString, value)
	}
	if value, ok := uuo.mutation.GithubOauthToken(); ok {
		_spec.SetField(user.FieldGithubOauthToken, field.TypeString, value)
	}
	if uuo.mutation.DevicesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedDevicesIDs(); len(nodes) > 0 && !uuo.mutation.DevicesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.DevicesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.MessagesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedMessagesIDs(); len(nodes) > 0 && !uuo.mutation.MessagesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.MessagesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
