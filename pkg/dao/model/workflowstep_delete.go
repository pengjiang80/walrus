// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package model

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/walrus/pkg/dao/model/internal"
	"github.com/seal-io/walrus/pkg/dao/model/predicate"
	"github.com/seal-io/walrus/pkg/dao/model/workflowstep"
)

// WorkflowStepDelete is the builder for deleting a WorkflowStep entity.
type WorkflowStepDelete struct {
	config
	hooks    []Hook
	mutation *WorkflowStepMutation
}

// Where appends a list predicates to the WorkflowStepDelete builder.
func (wsd *WorkflowStepDelete) Where(ps ...predicate.WorkflowStep) *WorkflowStepDelete {
	wsd.mutation.Where(ps...)
	return wsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (wsd *WorkflowStepDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, wsd.sqlExec, wsd.mutation, wsd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (wsd *WorkflowStepDelete) ExecX(ctx context.Context) int {
	n, err := wsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (wsd *WorkflowStepDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(workflowstep.Table, sqlgraph.NewFieldSpec(workflowstep.FieldID, field.TypeString))
	_spec.Node.Schema = wsd.schemaConfig.WorkflowStep
	ctx = internal.NewSchemaConfigContext(ctx, wsd.schemaConfig)
	if ps := wsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, wsd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	wsd.mutation.done = true
	return affected, err
}

// WorkflowStepDeleteOne is the builder for deleting a single WorkflowStep entity.
type WorkflowStepDeleteOne struct {
	wsd *WorkflowStepDelete
}

// Where appends a list predicates to the WorkflowStepDelete builder.
func (wsdo *WorkflowStepDeleteOne) Where(ps ...predicate.WorkflowStep) *WorkflowStepDeleteOne {
	wsdo.wsd.mutation.Where(ps...)
	return wsdo
}

// Exec executes the deletion query.
func (wsdo *WorkflowStepDeleteOne) Exec(ctx context.Context) error {
	n, err := wsdo.wsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{workflowstep.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (wsdo *WorkflowStepDeleteOne) ExecX(ctx context.Context) {
	if err := wsdo.Exec(ctx); err != nil {
		panic(err)
	}
}
