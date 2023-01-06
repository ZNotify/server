// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"
	"notify-api/ent/generate"
)

// The DeviceFunc type is an adapter to allow the use of ordinary
// function as Device mutator.
type DeviceFunc func(context.Context, *generate.DeviceMutation) (generate.Value, error)

// Mutate calls f(ctx, m).
func (f DeviceFunc) Mutate(ctx context.Context, m generate.Mutation) (generate.Value, error) {
	mv, ok := m.(*generate.DeviceMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *generate.DeviceMutation", m)
	}
	return f(ctx, mv)
}

// The MessageFunc type is an adapter to allow the use of ordinary
// function as Message mutator.
type MessageFunc func(context.Context, *generate.MessageMutation) (generate.Value, error)

// Mutate calls f(ctx, m).
func (f MessageFunc) Mutate(ctx context.Context, m generate.Mutation) (generate.Value, error) {
	mv, ok := m.(*generate.MessageMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *generate.MessageMutation", m)
	}
	return f(ctx, mv)
}

// The UserFunc type is an adapter to allow the use of ordinary
// function as User mutator.
type UserFunc func(context.Context, *generate.UserMutation) (generate.Value, error)

// Mutate calls f(ctx, m).
func (f UserFunc) Mutate(ctx context.Context, m generate.Mutation) (generate.Value, error) {
	mv, ok := m.(*generate.UserMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *generate.UserMutation", m)
	}
	return f(ctx, mv)
}

// Condition is a hook condition function.
type Condition func(context.Context, generate.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m generate.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m generate.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m generate.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op generate.Op) Condition {
	return func(_ context.Context, m generate.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m generate.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m generate.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m generate.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
func If(hk generate.Hook, cond Condition) generate.Hook {
	return func(next generate.Mutator) generate.Mutator {
		return generate.MutateFunc(func(ctx context.Context, m generate.Mutation) (generate.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, generate.Delete|generate.Create)
func On(hk generate.Hook, op generate.Op) generate.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, generate.Update|generate.UpdateOne)
func Unless(hk generate.Hook, op generate.Op) generate.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) generate.Hook {
	return func(generate.Mutator) generate.Mutator {
		return generate.MutateFunc(func(context.Context, generate.Mutation) (generate.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []generate.Hook {
//		return []generate.Hook{
//			Reject(generate.Delete|generate.Update),
//		}
//	}
func Reject(op generate.Op) generate.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []generate.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...generate.Hook) Chain {
	return Chain{append([]generate.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() generate.Hook {
	return func(mutator generate.Mutator) generate.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...generate.Hook) Chain {
	newHooks := make([]generate.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}