// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "seal". DO NOT EDIT.

package distributelock

import (
	"entgo.io/ent/dialect/sql"

	"github.com/seal-io/seal/pkg/dao/model/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.DistributeLock {
	return predicate.DistributeLock(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.DistributeLock {
	return predicate.DistributeLock(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.DistributeLock {
	return predicate.DistributeLock(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.DistributeLock {
	return predicate.DistributeLock(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.DistributeLock {
	return predicate.DistributeLock(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.DistributeLock {
	return predicate.DistributeLock(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.DistributeLock {
	return predicate.DistributeLock(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.DistributeLock {
	return predicate.DistributeLock(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.DistributeLock {
	return predicate.DistributeLock(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.DistributeLock {
	return predicate.DistributeLock(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.DistributeLock {
	return predicate.DistributeLock(sql.FieldContainsFold(FieldID, id))
}

// ExpireAt applies equality check predicate on the "expireAt" field. It's identical to ExpireAtEQ.
func ExpireAt(v int64) predicate.DistributeLock {
	return predicate.DistributeLock(sql.FieldEQ(FieldExpireAt, v))
}

// ExpireAtEQ applies the EQ predicate on the "expireAt" field.
func ExpireAtEQ(v int64) predicate.DistributeLock {
	return predicate.DistributeLock(sql.FieldEQ(FieldExpireAt, v))
}

// ExpireAtNEQ applies the NEQ predicate on the "expireAt" field.
func ExpireAtNEQ(v int64) predicate.DistributeLock {
	return predicate.DistributeLock(sql.FieldNEQ(FieldExpireAt, v))
}

// ExpireAtIn applies the In predicate on the "expireAt" field.
func ExpireAtIn(vs ...int64) predicate.DistributeLock {
	return predicate.DistributeLock(sql.FieldIn(FieldExpireAt, vs...))
}

// ExpireAtNotIn applies the NotIn predicate on the "expireAt" field.
func ExpireAtNotIn(vs ...int64) predicate.DistributeLock {
	return predicate.DistributeLock(sql.FieldNotIn(FieldExpireAt, vs...))
}

// ExpireAtGT applies the GT predicate on the "expireAt" field.
func ExpireAtGT(v int64) predicate.DistributeLock {
	return predicate.DistributeLock(sql.FieldGT(FieldExpireAt, v))
}

// ExpireAtGTE applies the GTE predicate on the "expireAt" field.
func ExpireAtGTE(v int64) predicate.DistributeLock {
	return predicate.DistributeLock(sql.FieldGTE(FieldExpireAt, v))
}

// ExpireAtLT applies the LT predicate on the "expireAt" field.
func ExpireAtLT(v int64) predicate.DistributeLock {
	return predicate.DistributeLock(sql.FieldLT(FieldExpireAt, v))
}

// ExpireAtLTE applies the LTE predicate on the "expireAt" field.
func ExpireAtLTE(v int64) predicate.DistributeLock {
	return predicate.DistributeLock(sql.FieldLTE(FieldExpireAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.DistributeLock) predicate.DistributeLock {
	return predicate.DistributeLock(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.DistributeLock) predicate.DistributeLock {
	return predicate.DistributeLock(func(s *sql.Selector) {
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
func Not(p predicate.DistributeLock) predicate.DistributeLock {
	return predicate.DistributeLock(func(s *sql.Selector) {
		p(s.Not())
	})
}
