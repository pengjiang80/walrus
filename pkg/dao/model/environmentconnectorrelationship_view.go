// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/seal-io/walrus/pkg/dao/model/environmentconnectorrelationship"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/utils/json"
)

// EnvironmentConnectorRelationshipCreateInput holds the creation input of the EnvironmentConnectorRelationship entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type EnvironmentConnectorRelationshipCreateInput struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Connector specifies full inserting the new Connector entity of the EnvironmentConnectorRelationship entity.
	Connector *ConnectorQueryInput `uri:"-" query:"-" json:"connector"`
}

// Model returns the EnvironmentConnectorRelationship entity for creating,
// after validating.
func (ecrci *EnvironmentConnectorRelationshipCreateInput) Model() *EnvironmentConnectorRelationship {
	if ecrci == nil {
		return nil
	}

	_ecr := &EnvironmentConnectorRelationship{}

	if ecrci.Connector != nil {
		_ecr.ConnectorID = ecrci.Connector.ID
	}
	return _ecr
}

// Validate checks the EnvironmentConnectorRelationshipCreateInput entity.
func (ecrci *EnvironmentConnectorRelationshipCreateInput) Validate() error {
	if ecrci == nil {
		return errors.New("nil receiver")
	}

	return ecrci.ValidateWith(ecrci.inputConfig.Context, ecrci.inputConfig.Client, nil)
}

// ValidateWith checks the EnvironmentConnectorRelationshipCreateInput entity with the given context and client set.
func (ecrci *EnvironmentConnectorRelationshipCreateInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if ecrci == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	if ecrci.Connector != nil {
		if err := ecrci.Connector.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				ecrci.Connector = nil
			}
		}
	}

	return nil
}

// EnvironmentConnectorRelationshipCreateInputs holds the creation input item of the EnvironmentConnectorRelationship entities.
type EnvironmentConnectorRelationshipCreateInputsItem struct {

	// Connector specifies full inserting the new Connector entity.
	Connector *ConnectorQueryInput `uri:"-" query:"-" json:"connector"`
}

// ValidateWith checks the EnvironmentConnectorRelationshipCreateInputsItem entity with the given context and client set.
func (ecrci *EnvironmentConnectorRelationshipCreateInputsItem) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if ecrci == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	if ecrci.Connector != nil {
		if err := ecrci.Connector.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				ecrci.Connector = nil
			}
		}
	}

	return nil
}

// EnvironmentConnectorRelationshipCreateInputs holds the creation input of the EnvironmentConnectorRelationship entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type EnvironmentConnectorRelationshipCreateInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*EnvironmentConnectorRelationshipCreateInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the EnvironmentConnectorRelationship entities for creating,
// after validating.
func (ecrci *EnvironmentConnectorRelationshipCreateInputs) Model() []*EnvironmentConnectorRelationship {
	if ecrci == nil || len(ecrci.Items) == 0 {
		return nil
	}

	_ecrs := make([]*EnvironmentConnectorRelationship, len(ecrci.Items))

	for i := range ecrci.Items {
		_ecr := &EnvironmentConnectorRelationship{}

		if ecrci.Items[i].Connector != nil {
			_ecr.ConnectorID = ecrci.Items[i].Connector.ID
		}

		_ecrs[i] = _ecr
	}

	return _ecrs
}

// Validate checks the EnvironmentConnectorRelationshipCreateInputs entity .
func (ecrci *EnvironmentConnectorRelationshipCreateInputs) Validate() error {
	if ecrci == nil {
		return errors.New("nil receiver")
	}

	return ecrci.ValidateWith(ecrci.inputConfig.Context, ecrci.inputConfig.Client, nil)
}

// ValidateWith checks the EnvironmentConnectorRelationshipCreateInputs entity with the given context and client set.
func (ecrci *EnvironmentConnectorRelationshipCreateInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if ecrci == nil {
		return errors.New("nil receiver")
	}

	if len(ecrci.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	for i := range ecrci.Items {
		if ecrci.Items[i] == nil {
			continue
		}

		if err := ecrci.Items[i].ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	return nil
}

// EnvironmentConnectorRelationshipDeleteInput holds the deletion input of the EnvironmentConnectorRelationship entity,
// please tags with `path:",inline"` if embedding.
type EnvironmentConnectorRelationshipDeleteInput struct {
	EnvironmentConnectorRelationshipQueryInput `path:",inline"`
}

// EnvironmentConnectorRelationshipDeleteInputs holds the deletion input item of the EnvironmentConnectorRelationship entities.
type EnvironmentConnectorRelationshipDeleteInputsItem struct {
	// ID of the EnvironmentConnectorRelationship entity.
	ID object.ID `path:"-" query:"-" json:"id"`
}

// EnvironmentConnectorRelationshipDeleteInputs holds the deletion input of the EnvironmentConnectorRelationship entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type EnvironmentConnectorRelationshipDeleteInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*EnvironmentConnectorRelationshipDeleteInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the EnvironmentConnectorRelationship entities for deleting,
// after validating.
func (ecrdi *EnvironmentConnectorRelationshipDeleteInputs) Model() []*EnvironmentConnectorRelationship {
	if ecrdi == nil || len(ecrdi.Items) == 0 {
		return nil
	}

	_ecrs := make([]*EnvironmentConnectorRelationship, len(ecrdi.Items))
	for i := range ecrdi.Items {
		_ecrs[i] = &EnvironmentConnectorRelationship{
			ID: ecrdi.Items[i].ID,
		}
	}
	return _ecrs
}

// IDs returns the ID list of the EnvironmentConnectorRelationship entities for deleting,
// after validating.
func (ecrdi *EnvironmentConnectorRelationshipDeleteInputs) IDs() []object.ID {
	if ecrdi == nil || len(ecrdi.Items) == 0 {
		return nil
	}

	ids := make([]object.ID, len(ecrdi.Items))
	for i := range ecrdi.Items {
		ids[i] = ecrdi.Items[i].ID
	}
	return ids
}

// Validate checks the EnvironmentConnectorRelationshipDeleteInputs entity.
func (ecrdi *EnvironmentConnectorRelationshipDeleteInputs) Validate() error {
	if ecrdi == nil {
		return errors.New("nil receiver")
	}

	return ecrdi.ValidateWith(ecrdi.inputConfig.Context, ecrdi.inputConfig.Client, nil)
}

// ValidateWith checks the EnvironmentConnectorRelationshipDeleteInputs entity with the given context and client set.
func (ecrdi *EnvironmentConnectorRelationshipDeleteInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if ecrdi == nil {
		return errors.New("nil receiver")
	}

	if len(ecrdi.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.EnvironmentConnectorRelationships().Query()

	ids := make([]object.ID, 0, len(ecrdi.Items))

	for i := range ecrdi.Items {
		if ecrdi.Items[i] == nil {
			return errors.New("nil item")
		}

		if ecrdi.Items[i].ID != "" {
			ids = append(ids, ecrdi.Items[i].ID)
		} else {
			return errors.New("found item hasn't identify")
		}
	}

	if len(ids) != cap(ids) {
		return errors.New("found unrecognized item")
	}

	idsCnt, err := q.Where(environmentconnectorrelationship.IDIn(ids...)).
		Count(ctx)
	if err != nil {
		return err
	}

	if idsCnt != cap(ids) {
		return errors.New("found unrecognized item")
	}

	return nil
}

// EnvironmentConnectorRelationshipPatchInput holds the patch input of the EnvironmentConnectorRelationship entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type EnvironmentConnectorRelationshipPatchInput struct {
	EnvironmentConnectorRelationshipUpdateInput `path:",inline" query:"-" json:",inline"`

	patchedEntity *EnvironmentConnectorRelationship `path:"-" query:"-" json:"-"`
}

// Model returns the EnvironmentConnectorRelationship patched entity,
// after validating.
func (ecrpi *EnvironmentConnectorRelationshipPatchInput) Model() *EnvironmentConnectorRelationship {
	if ecrpi == nil {
		return nil
	}

	return ecrpi.patchedEntity
}

// Validate checks the EnvironmentConnectorRelationshipPatchInput entity.
func (ecrpi *EnvironmentConnectorRelationshipPatchInput) Validate() error {
	if ecrpi == nil {
		return errors.New("nil receiver")
	}

	return ecrpi.ValidateWith(ecrpi.inputConfig.Context, ecrpi.inputConfig.Client, nil)
}

// ValidateWith checks the EnvironmentConnectorRelationshipPatchInput entity with the given context and client set.
func (ecrpi *EnvironmentConnectorRelationshipPatchInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if cache == nil {
		cache = map[string]any{}
	}

	if err := ecrpi.EnvironmentConnectorRelationshipUpdateInput.ValidateWith(ctx, cs, cache); err != nil {
		return err
	}

	q := cs.EnvironmentConnectorRelationships().Query()

	if ecrpi.Refer != nil {
		if ecrpi.Refer.IsID() {
			q.Where(
				environmentconnectorrelationship.ID(ecrpi.Refer.ID()))
		} else {
			return errors.New("invalid identify refer of environmentconnectorrelationship")
		}
	} else if ecrpi.ID != "" {
		q.Where(
			environmentconnectorrelationship.ID(ecrpi.ID))
	} else {
		return errors.New("invalid identify of environmentconnectorrelationship")
	}

	q.Select(
		environmentconnectorrelationship.WithoutFields(
			environmentconnectorrelationship.FieldCreateTime,
		)...,
	)

	var e *EnvironmentConnectorRelationship
	{
		// Get cache from previous validation.
		queryStmt, queryArgs := q.sqlQuery(setContextOp(ctx, q.ctx, "cache")).Query()
		ck := fmt.Sprintf("stmt=%v, args=%v", queryStmt, queryArgs)
		if cv, existed := cache[ck]; !existed {
			var err error
			e, err = q.Only(ctx)
			if err != nil {
				return err
			}

			// Set cache for other validation.
			cache[ck] = e
		} else {
			e = cv.(*EnvironmentConnectorRelationship)
		}
	}

	_ecr := ecrpi.EnvironmentConnectorRelationshipUpdateInput.Model()

	_obj, err := json.PatchObject(e, _ecr)
	if err != nil {
		return err
	}

	ecrpi.patchedEntity = _obj.(*EnvironmentConnectorRelationship)
	return nil
}

// EnvironmentConnectorRelationshipQueryInput holds the query input of the EnvironmentConnectorRelationship entity,
// please tags with `path:",inline"` if embedding.
type EnvironmentConnectorRelationshipQueryInput struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Refer holds the route path reference of the EnvironmentConnectorRelationship entity.
	Refer *object.Refer `path:"environmentconnectorrelationship,default=" query:"-" json:"-"`
	// ID of the EnvironmentConnectorRelationship entity.
	ID object.ID `path:"-" query:"-" json:"id"`
}

// Model returns the EnvironmentConnectorRelationship entity for querying,
// after validating.
func (ecrqi *EnvironmentConnectorRelationshipQueryInput) Model() *EnvironmentConnectorRelationship {
	if ecrqi == nil {
		return nil
	}

	return &EnvironmentConnectorRelationship{
		ID: ecrqi.ID,
	}
}

// Validate checks the EnvironmentConnectorRelationshipQueryInput entity.
func (ecrqi *EnvironmentConnectorRelationshipQueryInput) Validate() error {
	if ecrqi == nil {
		return errors.New("nil receiver")
	}

	return ecrqi.ValidateWith(ecrqi.inputConfig.Context, ecrqi.inputConfig.Client, nil)
}

// ValidateWith checks the EnvironmentConnectorRelationshipQueryInput entity with the given context and client set.
func (ecrqi *EnvironmentConnectorRelationshipQueryInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if ecrqi == nil {
		return errors.New("nil receiver")
	}

	if ecrqi.Refer != nil && *ecrqi.Refer == "" {
		return fmt.Errorf("model: %s : %w", environmentconnectorrelationship.Label, ErrBlankResourceRefer)
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.EnvironmentConnectorRelationships().Query()

	if ecrqi.Refer != nil {
		if ecrqi.Refer.IsID() {
			q.Where(
				environmentconnectorrelationship.ID(ecrqi.Refer.ID()))
		} else {
			return errors.New("invalid identify refer of environmentconnectorrelationship")
		}
	} else if ecrqi.ID != "" {
		q.Where(
			environmentconnectorrelationship.ID(ecrqi.ID))
	} else {
		return errors.New("invalid identify of environmentconnectorrelationship")
	}

	q.Select(
		environmentconnectorrelationship.FieldID,
	)

	var e *EnvironmentConnectorRelationship
	{
		// Get cache from previous validation.
		queryStmt, queryArgs := q.sqlQuery(setContextOp(ctx, q.ctx, "cache")).Query()
		ck := fmt.Sprintf("stmt=%v, args=%v", queryStmt, queryArgs)
		if cv, existed := cache[ck]; !existed {
			var err error
			e, err = q.Only(ctx)
			if err != nil {
				return err
			}

			// Set cache for other validation.
			cache[ck] = e
		} else {
			e = cv.(*EnvironmentConnectorRelationship)
		}
	}

	ecrqi.ID = e.ID
	return nil
}

// EnvironmentConnectorRelationshipQueryInputs holds the query input of the EnvironmentConnectorRelationship entities,
// please tags with `path:",inline" query:",inline"` if embedding.
type EnvironmentConnectorRelationshipQueryInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`
}

// Validate checks the EnvironmentConnectorRelationshipQueryInputs entity.
func (ecrqi *EnvironmentConnectorRelationshipQueryInputs) Validate() error {
	if ecrqi == nil {
		return errors.New("nil receiver")
	}

	return ecrqi.ValidateWith(ecrqi.inputConfig.Context, ecrqi.inputConfig.Client, nil)
}

// ValidateWith checks the EnvironmentConnectorRelationshipQueryInputs entity with the given context and client set.
func (ecrqi *EnvironmentConnectorRelationshipQueryInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if ecrqi == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	return nil
}

// EnvironmentConnectorRelationshipUpdateInput holds the modification input of the EnvironmentConnectorRelationship entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type EnvironmentConnectorRelationshipUpdateInput struct {
	EnvironmentConnectorRelationshipQueryInput `path:",inline" query:"-" json:"-"`

	// Connector indicates replacing the stale Connector entity.
	Connector *ConnectorQueryInput `uri:"-" query:"-" json:"connector"`
}

// Model returns the EnvironmentConnectorRelationship entity for modifying,
// after validating.
func (ecrui *EnvironmentConnectorRelationshipUpdateInput) Model() *EnvironmentConnectorRelationship {
	if ecrui == nil {
		return nil
	}

	_ecr := &EnvironmentConnectorRelationship{
		ID: ecrui.ID,
	}

	if ecrui.Connector != nil {
		_ecr.ConnectorID = ecrui.Connector.ID
	}
	return _ecr
}

// Validate checks the EnvironmentConnectorRelationshipUpdateInput entity.
func (ecrui *EnvironmentConnectorRelationshipUpdateInput) Validate() error {
	if ecrui == nil {
		return errors.New("nil receiver")
	}

	return ecrui.ValidateWith(ecrui.inputConfig.Context, ecrui.inputConfig.Client, nil)
}

// ValidateWith checks the EnvironmentConnectorRelationshipUpdateInput entity with the given context and client set.
func (ecrui *EnvironmentConnectorRelationshipUpdateInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if cache == nil {
		cache = map[string]any{}
	}

	if err := ecrui.EnvironmentConnectorRelationshipQueryInput.ValidateWith(ctx, cs, cache); err != nil {
		return err
	}

	if ecrui.Connector != nil {
		if err := ecrui.Connector.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				ecrui.Connector = nil
			}
		}
	}

	return nil
}

// EnvironmentConnectorRelationshipUpdateInputs holds the modification input item of the EnvironmentConnectorRelationship entities.
type EnvironmentConnectorRelationshipUpdateInputsItem struct {
	// ID of the EnvironmentConnectorRelationship entity.
	ID object.ID `path:"-" query:"-" json:"id"`

	// Connector indicates replacing the stale Connector entity.
	Connector *ConnectorQueryInput `uri:"-" query:"-" json:"connector"`
}

// ValidateWith checks the EnvironmentConnectorRelationshipUpdateInputsItem entity with the given context and client set.
func (ecrui *EnvironmentConnectorRelationshipUpdateInputsItem) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if ecrui == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	if ecrui.Connector != nil {
		if err := ecrui.Connector.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				ecrui.Connector = nil
			}
		}
	}

	return nil
}

// EnvironmentConnectorRelationshipUpdateInputs holds the modification input of the EnvironmentConnectorRelationship entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type EnvironmentConnectorRelationshipUpdateInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*EnvironmentConnectorRelationshipUpdateInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the EnvironmentConnectorRelationship entities for modifying,
// after validating.
func (ecrui *EnvironmentConnectorRelationshipUpdateInputs) Model() []*EnvironmentConnectorRelationship {
	if ecrui == nil || len(ecrui.Items) == 0 {
		return nil
	}

	_ecrs := make([]*EnvironmentConnectorRelationship, len(ecrui.Items))

	for i := range ecrui.Items {
		_ecr := &EnvironmentConnectorRelationship{
			ID: ecrui.Items[i].ID,
		}

		if ecrui.Items[i].Connector != nil {
			_ecr.ConnectorID = ecrui.Items[i].Connector.ID
		}

		_ecrs[i] = _ecr
	}

	return _ecrs
}

// IDs returns the ID list of the EnvironmentConnectorRelationship entities for modifying,
// after validating.
func (ecrui *EnvironmentConnectorRelationshipUpdateInputs) IDs() []object.ID {
	if ecrui == nil || len(ecrui.Items) == 0 {
		return nil
	}

	ids := make([]object.ID, len(ecrui.Items))
	for i := range ecrui.Items {
		ids[i] = ecrui.Items[i].ID
	}
	return ids
}

// Validate checks the EnvironmentConnectorRelationshipUpdateInputs entity.
func (ecrui *EnvironmentConnectorRelationshipUpdateInputs) Validate() error {
	if ecrui == nil {
		return errors.New("nil receiver")
	}

	return ecrui.ValidateWith(ecrui.inputConfig.Context, ecrui.inputConfig.Client, nil)
}

// ValidateWith checks the EnvironmentConnectorRelationshipUpdateInputs entity with the given context and client set.
func (ecrui *EnvironmentConnectorRelationshipUpdateInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if ecrui == nil {
		return errors.New("nil receiver")
	}

	if len(ecrui.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.EnvironmentConnectorRelationships().Query()

	ids := make([]object.ID, 0, len(ecrui.Items))

	for i := range ecrui.Items {
		if ecrui.Items[i] == nil {
			return errors.New("nil item")
		}

		if ecrui.Items[i].ID != "" {
			ids = append(ids, ecrui.Items[i].ID)
		} else {
			return errors.New("found item hasn't identify")
		}
	}

	if len(ids) != cap(ids) {
		return errors.New("found unrecognized item")
	}

	idsCnt, err := q.Where(environmentconnectorrelationship.IDIn(ids...)).
		Count(ctx)
	if err != nil {
		return err
	}

	if idsCnt != cap(ids) {
		return errors.New("found unrecognized item")
	}

	for i := range ecrui.Items {
		if err := ecrui.Items[i].ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	return nil
}

// EnvironmentConnectorRelationshipOutput holds the output of the EnvironmentConnectorRelationship entity.
type EnvironmentConnectorRelationshipOutput struct {
	ID         object.ID  `json:"id,omitempty"`
	CreateTime *time.Time `json:"createTime,omitempty"`

	Connector *ConnectorOutput `json:"connector,omitempty"`
}

// View returns the output of EnvironmentConnectorRelationship entity.
func (_ecr *EnvironmentConnectorRelationship) View() *EnvironmentConnectorRelationshipOutput {
	return ExposeEnvironmentConnectorRelationship(_ecr)
}

// View returns the output of EnvironmentConnectorRelationship entities.
func (_ecrs EnvironmentConnectorRelationships) View() []*EnvironmentConnectorRelationshipOutput {
	return ExposeEnvironmentConnectorRelationships(_ecrs)
}

// ExposeEnvironmentConnectorRelationship converts the EnvironmentConnectorRelationship to EnvironmentConnectorRelationshipOutput.
func ExposeEnvironmentConnectorRelationship(_ecr *EnvironmentConnectorRelationship) *EnvironmentConnectorRelationshipOutput {
	if _ecr == nil {
		return nil
	}

	ecro := &EnvironmentConnectorRelationshipOutput{
		ID:         _ecr.ID,
		CreateTime: _ecr.CreateTime,
	}

	if _ecr.Edges.Connector != nil {
		ecro.Connector = ExposeConnector(_ecr.Edges.Connector)
	} else if _ecr.ConnectorID != "" {
		ecro.Connector = &ConnectorOutput{
			ID: _ecr.ConnectorID,
		}
	}
	return ecro
}

// ExposeEnvironmentConnectorRelationships converts the EnvironmentConnectorRelationship slice to EnvironmentConnectorRelationshipOutput pointer slice.
func ExposeEnvironmentConnectorRelationships(_ecrs []*EnvironmentConnectorRelationship) []*EnvironmentConnectorRelationshipOutput {
	if len(_ecrs) == 0 {
		return nil
	}

	ecros := make([]*EnvironmentConnectorRelationshipOutput, len(_ecrs))
	for i := range _ecrs {
		ecros[i] = ExposeEnvironmentConnectorRelationship(_ecrs[i])
	}
	return ecros
}
