// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/seal-io/walrus/pkg/dao/model/role"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/utils/json"
)

// RoleCreateInput holds the creation input of the Role entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type RoleCreateInput struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// The kind of the role.
	Kind string `path:"-" query:"-" json:"kind,omitempty"`
	// The detail of the role.
	Description string `path:"-" query:"-" json:"description,omitempty"`
	// The policy list of the role.
	Policies types.RolePolicies `path:"-" query:"-" json:"policies,omitempty"`
	// The environment type list of the role to apply, only for system kind role.
	ApplicableEnvironmentTypes []string `path:"-" query:"-" json:"applicableEnvironmentTypes,omitempty"`
}

// Model returns the Role entity for creating,
// after validating.
func (rci *RoleCreateInput) Model() *Role {
	if rci == nil {
		return nil
	}

	_r := &Role{
		Kind:                       rci.Kind,
		Description:                rci.Description,
		Policies:                   rci.Policies,
		ApplicableEnvironmentTypes: rci.ApplicableEnvironmentTypes,
	}

	return _r
}

// Validate checks the RoleCreateInput entity.
func (rci *RoleCreateInput) Validate() error {
	if rci == nil {
		return errors.New("nil receiver")
	}

	return rci.ValidateWith(rci.inputConfig.Context, rci.inputConfig.Client, nil)
}

// ValidateWith checks the RoleCreateInput entity with the given context and client set.
func (rci *RoleCreateInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if rci == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	return nil
}

// RoleCreateInputs holds the creation input item of the Role entities.
type RoleCreateInputsItem struct {
	// The kind of the role.
	Kind string `path:"-" query:"-" json:"kind,omitempty"`
	// The detail of the role.
	Description string `path:"-" query:"-" json:"description,omitempty"`
	// The policy list of the role.
	Policies types.RolePolicies `path:"-" query:"-" json:"policies,omitempty"`
	// The environment type list of the role to apply, only for system kind role.
	ApplicableEnvironmentTypes []string `path:"-" query:"-" json:"applicableEnvironmentTypes,omitempty"`
}

// ValidateWith checks the RoleCreateInputsItem entity with the given context and client set.
func (rci *RoleCreateInputsItem) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if rci == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	return nil
}

// RoleCreateInputs holds the creation input of the Role entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type RoleCreateInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*RoleCreateInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the Role entities for creating,
// after validating.
func (rci *RoleCreateInputs) Model() []*Role {
	if rci == nil || len(rci.Items) == 0 {
		return nil
	}

	_rs := make([]*Role, len(rci.Items))

	for i := range rci.Items {
		_r := &Role{
			Kind:                       rci.Items[i].Kind,
			Description:                rci.Items[i].Description,
			Policies:                   rci.Items[i].Policies,
			ApplicableEnvironmentTypes: rci.Items[i].ApplicableEnvironmentTypes,
		}

		_rs[i] = _r
	}

	return _rs
}

// Validate checks the RoleCreateInputs entity .
func (rci *RoleCreateInputs) Validate() error {
	if rci == nil {
		return errors.New("nil receiver")
	}

	return rci.ValidateWith(rci.inputConfig.Context, rci.inputConfig.Client, nil)
}

// ValidateWith checks the RoleCreateInputs entity with the given context and client set.
func (rci *RoleCreateInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if rci == nil {
		return errors.New("nil receiver")
	}

	if len(rci.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	for i := range rci.Items {
		if rci.Items[i] == nil {
			continue
		}

		if err := rci.Items[i].ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	return nil
}

// RoleDeleteInput holds the deletion input of the Role entity,
// please tags with `path:",inline"` if embedding.
type RoleDeleteInput struct {
	RoleQueryInput `path:",inline"`
}

// RoleDeleteInputs holds the deletion input item of the Role entities.
type RoleDeleteInputsItem struct {
	// ID of the Role entity.
	ID string `path:"-" query:"-" json:"id"`
}

// RoleDeleteInputs holds the deletion input of the Role entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type RoleDeleteInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*RoleDeleteInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the Role entities for deleting,
// after validating.
func (rdi *RoleDeleteInputs) Model() []*Role {
	if rdi == nil || len(rdi.Items) == 0 {
		return nil
	}

	_rs := make([]*Role, len(rdi.Items))
	for i := range rdi.Items {
		_rs[i] = &Role{
			ID: rdi.Items[i].ID,
		}
	}
	return _rs
}

// IDs returns the ID list of the Role entities for deleting,
// after validating.
func (rdi *RoleDeleteInputs) IDs() []string {
	if rdi == nil || len(rdi.Items) == 0 {
		return nil
	}

	ids := make([]string, len(rdi.Items))
	for i := range rdi.Items {
		ids[i] = rdi.Items[i].ID
	}
	return ids
}

// Validate checks the RoleDeleteInputs entity.
func (rdi *RoleDeleteInputs) Validate() error {
	if rdi == nil {
		return errors.New("nil receiver")
	}

	return rdi.ValidateWith(rdi.inputConfig.Context, rdi.inputConfig.Client, nil)
}

// ValidateWith checks the RoleDeleteInputs entity with the given context and client set.
func (rdi *RoleDeleteInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if rdi == nil {
		return errors.New("nil receiver")
	}

	if len(rdi.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.Roles().Query()

	ids := make([]string, 0, len(rdi.Items))

	for i := range rdi.Items {
		if rdi.Items[i] == nil {
			return errors.New("nil item")
		}

		if rdi.Items[i].ID != "" {
			ids = append(ids, rdi.Items[i].ID)
		} else {
			return errors.New("found item hasn't identify")
		}
	}

	if len(ids) != cap(ids) {
		return errors.New("found unrecognized item")
	}

	idsCnt, err := q.Where(role.IDIn(ids...)).
		Count(ctx)
	if err != nil {
		return err
	}

	if idsCnt != cap(ids) {
		return errors.New("found unrecognized item")
	}

	return nil
}

// RolePatchInput holds the patch input of the Role entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type RolePatchInput struct {
	RoleUpdateInput `path:",inline" query:"-" json:",inline"`

	patchedEntity *Role `path:"-" query:"-" json:"-"`
}

// Model returns the Role patched entity,
// after validating.
func (rpi *RolePatchInput) Model() *Role {
	if rpi == nil {
		return nil
	}

	return rpi.patchedEntity
}

// Validate checks the RolePatchInput entity.
func (rpi *RolePatchInput) Validate() error {
	if rpi == nil {
		return errors.New("nil receiver")
	}

	return rpi.ValidateWith(rpi.inputConfig.Context, rpi.inputConfig.Client, nil)
}

// ValidateWith checks the RolePatchInput entity with the given context and client set.
func (rpi *RolePatchInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if cache == nil {
		cache = map[string]any{}
	}

	if err := rpi.RoleUpdateInput.ValidateWith(ctx, cs, cache); err != nil {
		return err
	}

	q := cs.Roles().Query()

	if rpi.Refer != nil {
		if rpi.Refer.IsString() {
			q.Where(
				role.ID(rpi.Refer.String()))
		} else {
			return errors.New("invalid identify refer of role")
		}
	} else if rpi.ID != "" {
		q.Where(
			role.ID(rpi.ID))
	} else {
		return errors.New("invalid identify of role")
	}

	q.Select(
		role.WithoutFields(
			role.FieldCreateTime,
			role.FieldUpdateTime,
			role.FieldSession,
			role.FieldBuiltin,
		)...,
	)

	var e *Role
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
			e = cv.(*Role)
		}
	}

	_r := rpi.RoleUpdateInput.Model()

	_obj, err := json.PatchObject(e, _r)
	if err != nil {
		return err
	}

	rpi.patchedEntity = _obj.(*Role)
	return nil
}

// RoleQueryInput holds the query input of the Role entity,
// please tags with `path:",inline"` if embedding.
type RoleQueryInput struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Refer holds the route path reference of the Role entity.
	Refer *object.Refer `path:"role,default=" query:"-" json:"-"`
	// ID of the Role entity.
	ID string `path:"-" query:"-" json:"id"`
}

// Model returns the Role entity for querying,
// after validating.
func (rqi *RoleQueryInput) Model() *Role {
	if rqi == nil {
		return nil
	}

	return &Role{
		ID: rqi.ID,
	}
}

// Validate checks the RoleQueryInput entity.
func (rqi *RoleQueryInput) Validate() error {
	if rqi == nil {
		return errors.New("nil receiver")
	}

	return rqi.ValidateWith(rqi.inputConfig.Context, rqi.inputConfig.Client, nil)
}

// ValidateWith checks the RoleQueryInput entity with the given context and client set.
func (rqi *RoleQueryInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if rqi == nil {
		return errors.New("nil receiver")
	}

	if rqi.Refer != nil && *rqi.Refer == "" {
		return fmt.Errorf("model: %s : %w", role.Label, ErrBlankResourceRefer)
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.Roles().Query()

	if rqi.Refer != nil {
		if rqi.Refer.IsString() {
			q.Where(
				role.ID(rqi.Refer.String()))
		} else {
			return errors.New("invalid identify refer of role")
		}
	} else if rqi.ID != "" {
		q.Where(
			role.ID(rqi.ID))
	} else {
		return errors.New("invalid identify of role")
	}

	q.Select(
		role.FieldID,
	)

	var e *Role
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
			e = cv.(*Role)
		}
	}

	rqi.ID = e.ID
	return nil
}

// RoleQueryInputs holds the query input of the Role entities,
// please tags with `path:",inline" query:",inline"` if embedding.
type RoleQueryInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`
}

// Validate checks the RoleQueryInputs entity.
func (rqi *RoleQueryInputs) Validate() error {
	if rqi == nil {
		return errors.New("nil receiver")
	}

	return rqi.ValidateWith(rqi.inputConfig.Context, rqi.inputConfig.Client, nil)
}

// ValidateWith checks the RoleQueryInputs entity with the given context and client set.
func (rqi *RoleQueryInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if rqi == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	return nil
}

// RoleUpdateInput holds the modification input of the Role entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type RoleUpdateInput struct {
	RoleQueryInput `path:",inline" query:"-" json:"-"`

	// The detail of the role.
	Description string `path:"-" query:"-" json:"description,omitempty"`
	// The policy list of the role.
	Policies types.RolePolicies `path:"-" query:"-" json:"policies,omitempty"`
	// The environment type list of the role to apply, only for system kind role.
	ApplicableEnvironmentTypes []string `path:"-" query:"-" json:"applicableEnvironmentTypes,omitempty"`
}

// Model returns the Role entity for modifying,
// after validating.
func (rui *RoleUpdateInput) Model() *Role {
	if rui == nil {
		return nil
	}

	_r := &Role{
		ID:                         rui.ID,
		Description:                rui.Description,
		Policies:                   rui.Policies,
		ApplicableEnvironmentTypes: rui.ApplicableEnvironmentTypes,
	}

	return _r
}

// Validate checks the RoleUpdateInput entity.
func (rui *RoleUpdateInput) Validate() error {
	if rui == nil {
		return errors.New("nil receiver")
	}

	return rui.ValidateWith(rui.inputConfig.Context, rui.inputConfig.Client, nil)
}

// ValidateWith checks the RoleUpdateInput entity with the given context and client set.
func (rui *RoleUpdateInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if cache == nil {
		cache = map[string]any{}
	}

	if err := rui.RoleQueryInput.ValidateWith(ctx, cs, cache); err != nil {
		return err
	}

	return nil
}

// RoleUpdateInputs holds the modification input item of the Role entities.
type RoleUpdateInputsItem struct {
	// ID of the Role entity.
	ID string `path:"-" query:"-" json:"id"`

	// The detail of the role.
	Description string `path:"-" query:"-" json:"description,omitempty"`
	// The policy list of the role.
	Policies types.RolePolicies `path:"-" query:"-" json:"policies"`
	// The environment type list of the role to apply, only for system kind role.
	ApplicableEnvironmentTypes []string `path:"-" query:"-" json:"applicableEnvironmentTypes,omitempty"`
}

// ValidateWith checks the RoleUpdateInputsItem entity with the given context and client set.
func (rui *RoleUpdateInputsItem) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if rui == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	return nil
}

// RoleUpdateInputs holds the modification input of the Role entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type RoleUpdateInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*RoleUpdateInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the Role entities for modifying,
// after validating.
func (rui *RoleUpdateInputs) Model() []*Role {
	if rui == nil || len(rui.Items) == 0 {
		return nil
	}

	_rs := make([]*Role, len(rui.Items))

	for i := range rui.Items {
		_r := &Role{
			ID:                         rui.Items[i].ID,
			Description:                rui.Items[i].Description,
			Policies:                   rui.Items[i].Policies,
			ApplicableEnvironmentTypes: rui.Items[i].ApplicableEnvironmentTypes,
		}

		_rs[i] = _r
	}

	return _rs
}

// IDs returns the ID list of the Role entities for modifying,
// after validating.
func (rui *RoleUpdateInputs) IDs() []string {
	if rui == nil || len(rui.Items) == 0 {
		return nil
	}

	ids := make([]string, len(rui.Items))
	for i := range rui.Items {
		ids[i] = rui.Items[i].ID
	}
	return ids
}

// Validate checks the RoleUpdateInputs entity.
func (rui *RoleUpdateInputs) Validate() error {
	if rui == nil {
		return errors.New("nil receiver")
	}

	return rui.ValidateWith(rui.inputConfig.Context, rui.inputConfig.Client, nil)
}

// ValidateWith checks the RoleUpdateInputs entity with the given context and client set.
func (rui *RoleUpdateInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if rui == nil {
		return errors.New("nil receiver")
	}

	if len(rui.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.Roles().Query()

	ids := make([]string, 0, len(rui.Items))

	for i := range rui.Items {
		if rui.Items[i] == nil {
			return errors.New("nil item")
		}

		if rui.Items[i].ID != "" {
			ids = append(ids, rui.Items[i].ID)
		} else {
			return errors.New("found item hasn't identify")
		}
	}

	if len(ids) != cap(ids) {
		return errors.New("found unrecognized item")
	}

	idsCnt, err := q.Where(role.IDIn(ids...)).
		Count(ctx)
	if err != nil {
		return err
	}

	if idsCnt != cap(ids) {
		return errors.New("found unrecognized item")
	}

	for i := range rui.Items {
		if err := rui.Items[i].ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	return nil
}

// RoleOutput holds the output of the Role entity.
type RoleOutput struct {
	ID                         string             `json:"id,omitempty"`
	CreateTime                 *time.Time         `json:"createTime,omitempty"`
	UpdateTime                 *time.Time         `json:"updateTime,omitempty"`
	Kind                       string             `json:"kind,omitempty"`
	Description                string             `json:"description,omitempty"`
	Policies                   types.RolePolicies `json:"policies,omitempty"`
	ApplicableEnvironmentTypes []string           `json:"applicableEnvironmentTypes,omitempty"`
	Builtin                    bool               `json:"builtin,omitempty"`
}

// View returns the output of Role entity.
func (_r *Role) View() *RoleOutput {
	return ExposeRole(_r)
}

// View returns the output of Role entities.
func (_rs Roles) View() []*RoleOutput {
	return ExposeRoles(_rs)
}

// ExposeRole converts the Role to RoleOutput.
func ExposeRole(_r *Role) *RoleOutput {
	if _r == nil {
		return nil
	}

	ro := &RoleOutput{
		ID:                         _r.ID,
		CreateTime:                 _r.CreateTime,
		UpdateTime:                 _r.UpdateTime,
		Kind:                       _r.Kind,
		Description:                _r.Description,
		Policies:                   _r.Policies,
		ApplicableEnvironmentTypes: _r.ApplicableEnvironmentTypes,
		Builtin:                    _r.Builtin,
	}

	return ro
}

// ExposeRoles converts the Role slice to RoleOutput pointer slice.
func ExposeRoles(_rs []*Role) []*RoleOutput {
	if len(_rs) == 0 {
		return nil
	}

	ros := make([]*RoleOutput, len(_rs))
	for i := range _rs {
		ros[i] = ExposeRole(_rs[i])
	}
	return ros
}
