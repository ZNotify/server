// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ZNotify/server/app/db/ent/generate/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdatedAt, v))
}

// Secret applies equality check predicate on the "secret" field. It's identical to SecretEQ.
func Secret(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldSecret, v))
}

// GithubID applies equality check predicate on the "githubID" field. It's identical to GithubIDEQ.
func GithubID(v int64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldGithubID, v))
}

// GithubName applies equality check predicate on the "githubName" field. It's identical to GithubNameEQ.
func GithubName(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldGithubName, v))
}

// GithubLogin applies equality check predicate on the "githubLogin" field. It's identical to GithubLoginEQ.
func GithubLogin(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldGithubLogin, v))
}

// GithubOauthToken applies equality check predicate on the "githubOauthToken" field. It's identical to GithubOauthTokenEQ.
func GithubOauthToken(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldGithubOauthToken, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUpdatedAt, v))
}

// SecretEQ applies the EQ predicate on the "secret" field.
func SecretEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldSecret, v))
}

// SecretNEQ applies the NEQ predicate on the "secret" field.
func SecretNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldSecret, v))
}

// SecretIn applies the In predicate on the "secret" field.
func SecretIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldSecret, vs...))
}

// SecretNotIn applies the NotIn predicate on the "secret" field.
func SecretNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldSecret, vs...))
}

// SecretGT applies the GT predicate on the "secret" field.
func SecretGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldSecret, v))
}

// SecretGTE applies the GTE predicate on the "secret" field.
func SecretGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldSecret, v))
}

// SecretLT applies the LT predicate on the "secret" field.
func SecretLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldSecret, v))
}

// SecretLTE applies the LTE predicate on the "secret" field.
func SecretLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldSecret, v))
}

// SecretContains applies the Contains predicate on the "secret" field.
func SecretContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldSecret, v))
}

// SecretHasPrefix applies the HasPrefix predicate on the "secret" field.
func SecretHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldSecret, v))
}

// SecretHasSuffix applies the HasSuffix predicate on the "secret" field.
func SecretHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldSecret, v))
}

// SecretEqualFold applies the EqualFold predicate on the "secret" field.
func SecretEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldSecret, v))
}

// SecretContainsFold applies the ContainsFold predicate on the "secret" field.
func SecretContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldSecret, v))
}

// GithubIDEQ applies the EQ predicate on the "githubID" field.
func GithubIDEQ(v int64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldGithubID, v))
}

// GithubIDNEQ applies the NEQ predicate on the "githubID" field.
func GithubIDNEQ(v int64) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldGithubID, v))
}

// GithubIDIn applies the In predicate on the "githubID" field.
func GithubIDIn(vs ...int64) predicate.User {
	return predicate.User(sql.FieldIn(FieldGithubID, vs...))
}

// GithubIDNotIn applies the NotIn predicate on the "githubID" field.
func GithubIDNotIn(vs ...int64) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldGithubID, vs...))
}

// GithubIDGT applies the GT predicate on the "githubID" field.
func GithubIDGT(v int64) predicate.User {
	return predicate.User(sql.FieldGT(FieldGithubID, v))
}

// GithubIDGTE applies the GTE predicate on the "githubID" field.
func GithubIDGTE(v int64) predicate.User {
	return predicate.User(sql.FieldGTE(FieldGithubID, v))
}

// GithubIDLT applies the LT predicate on the "githubID" field.
func GithubIDLT(v int64) predicate.User {
	return predicate.User(sql.FieldLT(FieldGithubID, v))
}

// GithubIDLTE applies the LTE predicate on the "githubID" field.
func GithubIDLTE(v int64) predicate.User {
	return predicate.User(sql.FieldLTE(FieldGithubID, v))
}

// GithubNameEQ applies the EQ predicate on the "githubName" field.
func GithubNameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldGithubName, v))
}

// GithubNameNEQ applies the NEQ predicate on the "githubName" field.
func GithubNameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldGithubName, v))
}

// GithubNameIn applies the In predicate on the "githubName" field.
func GithubNameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldGithubName, vs...))
}

// GithubNameNotIn applies the NotIn predicate on the "githubName" field.
func GithubNameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldGithubName, vs...))
}

// GithubNameGT applies the GT predicate on the "githubName" field.
func GithubNameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldGithubName, v))
}

// GithubNameGTE applies the GTE predicate on the "githubName" field.
func GithubNameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldGithubName, v))
}

// GithubNameLT applies the LT predicate on the "githubName" field.
func GithubNameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldGithubName, v))
}

// GithubNameLTE applies the LTE predicate on the "githubName" field.
func GithubNameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldGithubName, v))
}

// GithubNameContains applies the Contains predicate on the "githubName" field.
func GithubNameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldGithubName, v))
}

// GithubNameHasPrefix applies the HasPrefix predicate on the "githubName" field.
func GithubNameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldGithubName, v))
}

// GithubNameHasSuffix applies the HasSuffix predicate on the "githubName" field.
func GithubNameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldGithubName, v))
}

// GithubNameEqualFold applies the EqualFold predicate on the "githubName" field.
func GithubNameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldGithubName, v))
}

// GithubNameContainsFold applies the ContainsFold predicate on the "githubName" field.
func GithubNameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldGithubName, v))
}

// GithubLoginEQ applies the EQ predicate on the "githubLogin" field.
func GithubLoginEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldGithubLogin, v))
}

// GithubLoginNEQ applies the NEQ predicate on the "githubLogin" field.
func GithubLoginNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldGithubLogin, v))
}

// GithubLoginIn applies the In predicate on the "githubLogin" field.
func GithubLoginIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldGithubLogin, vs...))
}

// GithubLoginNotIn applies the NotIn predicate on the "githubLogin" field.
func GithubLoginNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldGithubLogin, vs...))
}

// GithubLoginGT applies the GT predicate on the "githubLogin" field.
func GithubLoginGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldGithubLogin, v))
}

// GithubLoginGTE applies the GTE predicate on the "githubLogin" field.
func GithubLoginGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldGithubLogin, v))
}

// GithubLoginLT applies the LT predicate on the "githubLogin" field.
func GithubLoginLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldGithubLogin, v))
}

// GithubLoginLTE applies the LTE predicate on the "githubLogin" field.
func GithubLoginLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldGithubLogin, v))
}

// GithubLoginContains applies the Contains predicate on the "githubLogin" field.
func GithubLoginContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldGithubLogin, v))
}

// GithubLoginHasPrefix applies the HasPrefix predicate on the "githubLogin" field.
func GithubLoginHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldGithubLogin, v))
}

// GithubLoginHasSuffix applies the HasSuffix predicate on the "githubLogin" field.
func GithubLoginHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldGithubLogin, v))
}

// GithubLoginEqualFold applies the EqualFold predicate on the "githubLogin" field.
func GithubLoginEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldGithubLogin, v))
}

// GithubLoginContainsFold applies the ContainsFold predicate on the "githubLogin" field.
func GithubLoginContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldGithubLogin, v))
}

// GithubOauthTokenEQ applies the EQ predicate on the "githubOauthToken" field.
func GithubOauthTokenEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldGithubOauthToken, v))
}

// GithubOauthTokenNEQ applies the NEQ predicate on the "githubOauthToken" field.
func GithubOauthTokenNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldGithubOauthToken, v))
}

// GithubOauthTokenIn applies the In predicate on the "githubOauthToken" field.
func GithubOauthTokenIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldGithubOauthToken, vs...))
}

// GithubOauthTokenNotIn applies the NotIn predicate on the "githubOauthToken" field.
func GithubOauthTokenNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldGithubOauthToken, vs...))
}

// GithubOauthTokenGT applies the GT predicate on the "githubOauthToken" field.
func GithubOauthTokenGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldGithubOauthToken, v))
}

// GithubOauthTokenGTE applies the GTE predicate on the "githubOauthToken" field.
func GithubOauthTokenGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldGithubOauthToken, v))
}

// GithubOauthTokenLT applies the LT predicate on the "githubOauthToken" field.
func GithubOauthTokenLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldGithubOauthToken, v))
}

// GithubOauthTokenLTE applies the LTE predicate on the "githubOauthToken" field.
func GithubOauthTokenLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldGithubOauthToken, v))
}

// GithubOauthTokenContains applies the Contains predicate on the "githubOauthToken" field.
func GithubOauthTokenContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldGithubOauthToken, v))
}

// GithubOauthTokenHasPrefix applies the HasPrefix predicate on the "githubOauthToken" field.
func GithubOauthTokenHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldGithubOauthToken, v))
}

// GithubOauthTokenHasSuffix applies the HasSuffix predicate on the "githubOauthToken" field.
func GithubOauthTokenHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldGithubOauthToken, v))
}

// GithubOauthTokenEqualFold applies the EqualFold predicate on the "githubOauthToken" field.
func GithubOauthTokenEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldGithubOauthToken, v))
}

// GithubOauthTokenContainsFold applies the ContainsFold predicate on the "githubOauthToken" field.
func GithubOauthTokenContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldGithubOauthToken, v))
}

// HasDevices applies the HasEdge predicate on the "devices" edge.
func HasDevices() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DevicesTable, DevicesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDevicesWith applies the HasEdge predicate on the "devices" edge with a given conditions (other predicates).
func HasDevicesWith(preds ...predicate.Device) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DevicesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DevicesTable, DevicesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMessages applies the HasEdge predicate on the "messages" edge.
func HasMessages() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MessagesTable, MessagesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMessagesWith applies the HasEdge predicate on the "messages" edge with a given conditions (other predicates).
func HasMessagesWith(preds ...predicate.Message) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MessagesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MessagesTable, MessagesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		p(s.Not())
	})
}
