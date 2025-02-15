// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/seal-io/walrus/pkg/dao/model/resourcerevision"
	"github.com/seal-io/walrus/pkg/dao/schema/intercept"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/dao/types/crypto"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/pkg/dao/types/property"
	"github.com/seal-io/walrus/pkg/dao/types/status"
	"github.com/seal-io/walrus/utils/json"
)

// ResourceRevisionCreateInput holds the creation input of the ResourceRevision entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type ResourceRevisionCreateInput struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to create ResourceRevision entity MUST under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`
	// Environment indicates to create ResourceRevision entity MUST under the Environment route.
	Environment *EnvironmentQueryInput `path:",inline" query:"-" json:"-"`
	// Resource indicates to create ResourceRevision entity MUST under the Resource route.
	Resource *ResourceQueryInput `path:",inline" query:"-" json:"-"`

	// Output of the revision.
	Output string `path:"-" query:"-" json:"output"`
	// Input plan of the revision.
	InputPlan string `path:"-" query:"-" json:"inputPlan"`
	// ID of the template.
	TemplateID object.ID `path:"-" query:"-" json:"templateID"`
	// Version of the template.
	TemplateVersion string `path:"-" query:"-" json:"templateVersion"`
	// Name of the template.
	TemplateName string `path:"-" query:"-" json:"templateName"`
	// Attributes to configure the template.
	Attributes property.Values `path:"-" query:"-" json:"attributes,omitempty"`
	// Variables of the revision.
	Variables crypto.Map[string, string] `path:"-" query:"-" json:"variables,omitempty"`
	// Type of deployer.
	DeployerType string `path:"-" query:"-" json:"deployerType,omitempty"`
	// Duration in seconds of the revision deploying.
	Duration int `path:"-" query:"-" json:"duration,omitempty"`
	// Previous provider requirement of the revision.
	PreviousRequiredProviders []types.ProviderRequirement `path:"-" query:"-" json:"previousRequiredProviders,omitempty"`
	// Record of the revision.
	Record string `path:"-" query:"-" json:"record,omitempty"`
	// Change comment of the revision.
	ChangeComment string `path:"-" query:"-" json:"changeComment,omitempty"`
}

// Model returns the ResourceRevision entity for creating,
// after validating.
func (rrci *ResourceRevisionCreateInput) Model() *ResourceRevision {
	if rrci == nil {
		return nil
	}

	_rr := &ResourceRevision{
		Output:                    rrci.Output,
		InputPlan:                 rrci.InputPlan,
		TemplateID:                rrci.TemplateID,
		TemplateVersion:           rrci.TemplateVersion,
		TemplateName:              rrci.TemplateName,
		Attributes:                rrci.Attributes,
		Variables:                 rrci.Variables,
		DeployerType:              rrci.DeployerType,
		Duration:                  rrci.Duration,
		PreviousRequiredProviders: rrci.PreviousRequiredProviders,
		Record:                    rrci.Record,
		ChangeComment:             rrci.ChangeComment,
	}

	if rrci.Project != nil {
		_rr.ProjectID = rrci.Project.ID
	}
	if rrci.Environment != nil {
		_rr.EnvironmentID = rrci.Environment.ID
	}
	if rrci.Resource != nil {
		_rr.ResourceID = rrci.Resource.ID
	}

	return _rr
}

// Validate checks the ResourceRevisionCreateInput entity.
func (rrci *ResourceRevisionCreateInput) Validate() error {
	if rrci == nil {
		return errors.New("nil receiver")
	}

	return rrci.ValidateWith(rrci.inputConfig.Context, rrci.inputConfig.Client, nil)
}

// ValidateWith checks the ResourceRevisionCreateInput entity with the given context and client set.
func (rrci *ResourceRevisionCreateInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if rrci == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	// Validate when creating under the Project route.
	if rrci.Project != nil {
		if err := rrci.Project.ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}
	// Validate when creating under the Environment route.
	if rrci.Environment != nil {
		if err := rrci.Environment.ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}
	// Validate when creating under the Resource route.
	if rrci.Resource != nil {
		if err := rrci.Resource.ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	return nil
}

// ResourceRevisionCreateInputs holds the creation input item of the ResourceRevision entities.
type ResourceRevisionCreateInputsItem struct {
	// Output of the revision.
	Output string `path:"-" query:"-" json:"output"`
	// Input plan of the revision.
	InputPlan string `path:"-" query:"-" json:"inputPlan"`
	// ID of the template.
	TemplateID object.ID `path:"-" query:"-" json:"templateID"`
	// Version of the template.
	TemplateVersion string `path:"-" query:"-" json:"templateVersion"`
	// Name of the template.
	TemplateName string `path:"-" query:"-" json:"templateName"`
	// Attributes to configure the template.
	Attributes property.Values `path:"-" query:"-" json:"attributes,omitempty"`
	// Variables of the revision.
	Variables crypto.Map[string, string] `path:"-" query:"-" json:"variables,omitempty"`
	// Type of deployer.
	DeployerType string `path:"-" query:"-" json:"deployerType,omitempty"`
	// Duration in seconds of the revision deploying.
	Duration int `path:"-" query:"-" json:"duration,omitempty"`
	// Previous provider requirement of the revision.
	PreviousRequiredProviders []types.ProviderRequirement `path:"-" query:"-" json:"previousRequiredProviders,omitempty"`
	// Record of the revision.
	Record string `path:"-" query:"-" json:"record,omitempty"`
	// Change comment of the revision.
	ChangeComment string `path:"-" query:"-" json:"changeComment,omitempty"`
}

// ValidateWith checks the ResourceRevisionCreateInputsItem entity with the given context and client set.
func (rrci *ResourceRevisionCreateInputsItem) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if rrci == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	return nil
}

// ResourceRevisionCreateInputs holds the creation input of the ResourceRevision entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type ResourceRevisionCreateInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to create ResourceRevision entity MUST under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`
	// Environment indicates to create ResourceRevision entity MUST under the Environment route.
	Environment *EnvironmentQueryInput `path:",inline" query:"-" json:"-"`
	// Resource indicates to create ResourceRevision entity MUST under the Resource route.
	Resource *ResourceQueryInput `path:",inline" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*ResourceRevisionCreateInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the ResourceRevision entities for creating,
// after validating.
func (rrci *ResourceRevisionCreateInputs) Model() []*ResourceRevision {
	if rrci == nil || len(rrci.Items) == 0 {
		return nil
	}

	_rrs := make([]*ResourceRevision, len(rrci.Items))

	for i := range rrci.Items {
		_rr := &ResourceRevision{
			Output:                    rrci.Items[i].Output,
			InputPlan:                 rrci.Items[i].InputPlan,
			TemplateID:                rrci.Items[i].TemplateID,
			TemplateVersion:           rrci.Items[i].TemplateVersion,
			TemplateName:              rrci.Items[i].TemplateName,
			Attributes:                rrci.Items[i].Attributes,
			Variables:                 rrci.Items[i].Variables,
			DeployerType:              rrci.Items[i].DeployerType,
			Duration:                  rrci.Items[i].Duration,
			PreviousRequiredProviders: rrci.Items[i].PreviousRequiredProviders,
			Record:                    rrci.Items[i].Record,
			ChangeComment:             rrci.Items[i].ChangeComment,
		}

		if rrci.Project != nil {
			_rr.ProjectID = rrci.Project.ID
		}
		if rrci.Environment != nil {
			_rr.EnvironmentID = rrci.Environment.ID
		}
		if rrci.Resource != nil {
			_rr.ResourceID = rrci.Resource.ID
		}

		_rrs[i] = _rr
	}

	return _rrs
}

// Validate checks the ResourceRevisionCreateInputs entity .
func (rrci *ResourceRevisionCreateInputs) Validate() error {
	if rrci == nil {
		return errors.New("nil receiver")
	}

	return rrci.ValidateWith(rrci.inputConfig.Context, rrci.inputConfig.Client, nil)
}

// ValidateWith checks the ResourceRevisionCreateInputs entity with the given context and client set.
func (rrci *ResourceRevisionCreateInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if rrci == nil {
		return errors.New("nil receiver")
	}

	if len(rrci.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	// Validate when creating under the Project route.
	if rrci.Project != nil {
		if err := rrci.Project.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				rrci.Project = nil
			}
		}
	}
	// Validate when creating under the Environment route.
	if rrci.Environment != nil {
		if err := rrci.Environment.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				rrci.Environment = nil
			}
		}
	}
	// Validate when creating under the Resource route.
	if rrci.Resource != nil {
		if err := rrci.Resource.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				rrci.Resource = nil
			}
		}
	}

	for i := range rrci.Items {
		if rrci.Items[i] == nil {
			continue
		}

		if err := rrci.Items[i].ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	return nil
}

// ResourceRevisionDeleteInput holds the deletion input of the ResourceRevision entity,
// please tags with `path:",inline"` if embedding.
type ResourceRevisionDeleteInput struct {
	ResourceRevisionQueryInput `path:",inline"`
}

// ResourceRevisionDeleteInputs holds the deletion input item of the ResourceRevision entities.
type ResourceRevisionDeleteInputsItem struct {
	// ID of the ResourceRevision entity.
	ID object.ID `path:"-" query:"-" json:"id"`
}

// ResourceRevisionDeleteInputs holds the deletion input of the ResourceRevision entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type ResourceRevisionDeleteInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to delete ResourceRevision entity MUST under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`
	// Environment indicates to delete ResourceRevision entity MUST under the Environment route.
	Environment *EnvironmentQueryInput `path:",inline" query:"-" json:"-"`
	// Resource indicates to delete ResourceRevision entity MUST under the Resource route.
	Resource *ResourceQueryInput `path:",inline" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*ResourceRevisionDeleteInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the ResourceRevision entities for deleting,
// after validating.
func (rrdi *ResourceRevisionDeleteInputs) Model() []*ResourceRevision {
	if rrdi == nil || len(rrdi.Items) == 0 {
		return nil
	}

	_rrs := make([]*ResourceRevision, len(rrdi.Items))
	for i := range rrdi.Items {
		_rrs[i] = &ResourceRevision{
			ID: rrdi.Items[i].ID,
		}
	}
	return _rrs
}

// IDs returns the ID list of the ResourceRevision entities for deleting,
// after validating.
func (rrdi *ResourceRevisionDeleteInputs) IDs() []object.ID {
	if rrdi == nil || len(rrdi.Items) == 0 {
		return nil
	}

	ids := make([]object.ID, len(rrdi.Items))
	for i := range rrdi.Items {
		ids[i] = rrdi.Items[i].ID
	}
	return ids
}

// Validate checks the ResourceRevisionDeleteInputs entity.
func (rrdi *ResourceRevisionDeleteInputs) Validate() error {
	if rrdi == nil {
		return errors.New("nil receiver")
	}

	return rrdi.ValidateWith(rrdi.inputConfig.Context, rrdi.inputConfig.Client, nil)
}

// ValidateWith checks the ResourceRevisionDeleteInputs entity with the given context and client set.
func (rrdi *ResourceRevisionDeleteInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if rrdi == nil {
		return errors.New("nil receiver")
	}

	if len(rrdi.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.ResourceRevisions().Query()

	// Validate when deleting under the Project route.
	if rrdi.Project != nil {
		if err := rrdi.Project.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			ctx = valueContext(ctx, intercept.WithProjectInterceptor)
			q.Where(
				resourcerevision.ProjectID(rrdi.Project.ID))
		}
	}

	// Validate when deleting under the Environment route.
	if rrdi.Environment != nil {
		if err := rrdi.Environment.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			q.Where(
				resourcerevision.EnvironmentID(rrdi.Environment.ID))
		}
	}

	// Validate when deleting under the Resource route.
	if rrdi.Resource != nil {
		if err := rrdi.Resource.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			q.Where(
				resourcerevision.ResourceID(rrdi.Resource.ID))
		}
	}

	ids := make([]object.ID, 0, len(rrdi.Items))

	for i := range rrdi.Items {
		if rrdi.Items[i] == nil {
			return errors.New("nil item")
		}

		if rrdi.Items[i].ID != "" {
			ids = append(ids, rrdi.Items[i].ID)
		} else {
			return errors.New("found item hasn't identify")
		}
	}

	if len(ids) != cap(ids) {
		return errors.New("found unrecognized item")
	}

	idsCnt, err := q.Where(resourcerevision.IDIn(ids...)).
		Count(ctx)
	if err != nil {
		return err
	}

	if idsCnt != cap(ids) {
		return errors.New("found unrecognized item")
	}

	return nil
}

// ResourceRevisionPatchInput holds the patch input of the ResourceRevision entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type ResourceRevisionPatchInput struct {
	ResourceRevisionUpdateInput `path:",inline" query:"-" json:",inline"`

	patchedEntity *ResourceRevision `path:"-" query:"-" json:"-"`
}

// Model returns the ResourceRevision patched entity,
// after validating.
func (rrpi *ResourceRevisionPatchInput) Model() *ResourceRevision {
	if rrpi == nil {
		return nil
	}

	return rrpi.patchedEntity
}

// Validate checks the ResourceRevisionPatchInput entity.
func (rrpi *ResourceRevisionPatchInput) Validate() error {
	if rrpi == nil {
		return errors.New("nil receiver")
	}

	return rrpi.ValidateWith(rrpi.inputConfig.Context, rrpi.inputConfig.Client, nil)
}

// ValidateWith checks the ResourceRevisionPatchInput entity with the given context and client set.
func (rrpi *ResourceRevisionPatchInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if cache == nil {
		cache = map[string]any{}
	}

	if err := rrpi.ResourceRevisionUpdateInput.ValidateWith(ctx, cs, cache); err != nil {
		return err
	}

	q := cs.ResourceRevisions().Query()

	// Validate when querying under the Project route.
	if rrpi.Project != nil {
		if err := rrpi.Project.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			ctx = valueContext(ctx, intercept.WithProjectInterceptor)
			q.Where(
				resourcerevision.ProjectID(rrpi.Project.ID))
		}
	}

	// Validate when querying under the Environment route.
	if rrpi.Environment != nil {
		if err := rrpi.Environment.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			q.Where(
				resourcerevision.EnvironmentID(rrpi.Environment.ID))
		}
	}

	// Validate when querying under the Resource route.
	if rrpi.Resource != nil {
		if err := rrpi.Resource.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			q.Where(
				resourcerevision.ResourceID(rrpi.Resource.ID))
		}
	}

	if rrpi.Refer != nil {
		if rrpi.Refer.IsID() {
			q.Where(
				resourcerevision.ID(rrpi.Refer.ID()))
		} else {
			return errors.New("invalid identify refer of resourcerevision")
		}
	} else if rrpi.ID != "" {
		q.Where(
			resourcerevision.ID(rrpi.ID))
	} else {
		return errors.New("invalid identify of resourcerevision")
	}

	q.Select(
		resourcerevision.WithoutFields(
			resourcerevision.FieldCreateTime,
			resourcerevision.FieldStatus,
			resourcerevision.FieldCreatedBy,
		)...,
	)

	var e *ResourceRevision
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
			e = cv.(*ResourceRevision)
		}
	}

	_rr := rrpi.ResourceRevisionUpdateInput.Model()

	_obj, err := json.PatchObject(e, _rr)
	if err != nil {
		return err
	}

	rrpi.patchedEntity = _obj.(*ResourceRevision)
	return nil
}

// ResourceRevisionQueryInput holds the query input of the ResourceRevision entity,
// please tags with `path:",inline"` if embedding.
type ResourceRevisionQueryInput struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to query ResourceRevision entity MUST under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"project"`
	// Environment indicates to query ResourceRevision entity MUST under the Environment route.
	Environment *EnvironmentQueryInput `path:",inline" query:"-" json:"environment"`
	// Resource indicates to query ResourceRevision entity MUST under the Resource route.
	Resource *ResourceQueryInput `path:",inline" query:"-" json:"resource"`

	// Refer holds the route path reference of the ResourceRevision entity.
	Refer *object.Refer `path:"resourcerevision,default=" query:"-" json:"-"`
	// ID of the ResourceRevision entity.
	ID object.ID `path:"-" query:"-" json:"id"`
}

// Model returns the ResourceRevision entity for querying,
// after validating.
func (rrqi *ResourceRevisionQueryInput) Model() *ResourceRevision {
	if rrqi == nil {
		return nil
	}

	return &ResourceRevision{
		ID: rrqi.ID,
	}
}

// Validate checks the ResourceRevisionQueryInput entity.
func (rrqi *ResourceRevisionQueryInput) Validate() error {
	if rrqi == nil {
		return errors.New("nil receiver")
	}

	return rrqi.ValidateWith(rrqi.inputConfig.Context, rrqi.inputConfig.Client, nil)
}

// ValidateWith checks the ResourceRevisionQueryInput entity with the given context and client set.
func (rrqi *ResourceRevisionQueryInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if rrqi == nil {
		return errors.New("nil receiver")
	}

	if rrqi.Refer != nil && *rrqi.Refer == "" {
		return fmt.Errorf("model: %s : %w", resourcerevision.Label, ErrBlankResourceRefer)
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.ResourceRevisions().Query()

	// Validate when querying under the Project route.
	if rrqi.Project != nil {
		if err := rrqi.Project.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			ctx = valueContext(ctx, intercept.WithProjectInterceptor)
			q.Where(
				resourcerevision.ProjectID(rrqi.Project.ID))
		}
	}

	// Validate when querying under the Environment route.
	if rrqi.Environment != nil {
		if err := rrqi.Environment.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			q.Where(
				resourcerevision.EnvironmentID(rrqi.Environment.ID))
		}
	}

	// Validate when querying under the Resource route.
	if rrqi.Resource != nil {
		if err := rrqi.Resource.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			q.Where(
				resourcerevision.ResourceID(rrqi.Resource.ID))
		}
	}

	if rrqi.Refer != nil {
		if rrqi.Refer.IsID() {
			q.Where(
				resourcerevision.ID(rrqi.Refer.ID()))
		} else {
			return errors.New("invalid identify refer of resourcerevision")
		}
	} else if rrqi.ID != "" {
		q.Where(
			resourcerevision.ID(rrqi.ID))
	} else {
		return errors.New("invalid identify of resourcerevision")
	}

	q.Select(
		resourcerevision.FieldID,
	)

	var e *ResourceRevision
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
			e = cv.(*ResourceRevision)
		}
	}

	rrqi.ID = e.ID
	return nil
}

// ResourceRevisionQueryInputs holds the query input of the ResourceRevision entities,
// please tags with `path:",inline" query:",inline"` if embedding.
type ResourceRevisionQueryInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to query ResourceRevision entity MUST under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`
	// Environment indicates to query ResourceRevision entity MUST under the Environment route.
	Environment *EnvironmentQueryInput `path:",inline" query:"-" json:"-"`
	// Resource indicates to query ResourceRevision entity MUST under the Resource route.
	Resource *ResourceQueryInput `path:",inline" query:"-" json:"-"`
}

// Validate checks the ResourceRevisionQueryInputs entity.
func (rrqi *ResourceRevisionQueryInputs) Validate() error {
	if rrqi == nil {
		return errors.New("nil receiver")
	}

	return rrqi.ValidateWith(rrqi.inputConfig.Context, rrqi.inputConfig.Client, nil)
}

// ValidateWith checks the ResourceRevisionQueryInputs entity with the given context and client set.
func (rrqi *ResourceRevisionQueryInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if rrqi == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	// Validate when querying under the Project route.
	if rrqi.Project != nil {
		if err := rrqi.Project.ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	// Validate when querying under the Environment route.
	if rrqi.Environment != nil {
		if err := rrqi.Environment.ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	// Validate when querying under the Resource route.
	if rrqi.Resource != nil {
		if err := rrqi.Resource.ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	return nil
}

// ResourceRevisionUpdateInput holds the modification input of the ResourceRevision entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type ResourceRevisionUpdateInput struct {
	ResourceRevisionQueryInput `path:",inline" query:"-" json:"-"`

	// Version of the template.
	TemplateVersion string `path:"-" query:"-" json:"templateVersion,omitempty"`
	// Attributes to configure the template.
	Attributes property.Values `path:"-" query:"-" json:"attributes,omitempty"`
	// Variables of the revision.
	Variables crypto.Map[string, string] `path:"-" query:"-" json:"variables,omitempty"`
	// Input plan of the revision.
	InputPlan string `path:"-" query:"-" json:"inputPlan,omitempty"`
	// Output of the revision.
	Output string `path:"-" query:"-" json:"output,omitempty"`
	// Type of deployer.
	DeployerType string `path:"-" query:"-" json:"deployerType,omitempty"`
	// Duration in seconds of the revision deploying.
	Duration int `path:"-" query:"-" json:"duration,omitempty"`
	// Previous provider requirement of the revision.
	PreviousRequiredProviders []types.ProviderRequirement `path:"-" query:"-" json:"previousRequiredProviders,omitempty"`
	// Record of the revision.
	Record string `path:"-" query:"-" json:"record,omitempty"`
	// Change comment of the revision.
	ChangeComment string `path:"-" query:"-" json:"changeComment,omitempty"`
}

// Model returns the ResourceRevision entity for modifying,
// after validating.
func (rrui *ResourceRevisionUpdateInput) Model() *ResourceRevision {
	if rrui == nil {
		return nil
	}

	_rr := &ResourceRevision{
		ID:                        rrui.ID,
		TemplateVersion:           rrui.TemplateVersion,
		Attributes:                rrui.Attributes,
		Variables:                 rrui.Variables,
		InputPlan:                 rrui.InputPlan,
		Output:                    rrui.Output,
		DeployerType:              rrui.DeployerType,
		Duration:                  rrui.Duration,
		PreviousRequiredProviders: rrui.PreviousRequiredProviders,
		Record:                    rrui.Record,
		ChangeComment:             rrui.ChangeComment,
	}

	return _rr
}

// Validate checks the ResourceRevisionUpdateInput entity.
func (rrui *ResourceRevisionUpdateInput) Validate() error {
	if rrui == nil {
		return errors.New("nil receiver")
	}

	return rrui.ValidateWith(rrui.inputConfig.Context, rrui.inputConfig.Client, nil)
}

// ValidateWith checks the ResourceRevisionUpdateInput entity with the given context and client set.
func (rrui *ResourceRevisionUpdateInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if cache == nil {
		cache = map[string]any{}
	}

	if err := rrui.ResourceRevisionQueryInput.ValidateWith(ctx, cs, cache); err != nil {
		return err
	}

	return nil
}

// ResourceRevisionUpdateInputs holds the modification input item of the ResourceRevision entities.
type ResourceRevisionUpdateInputsItem struct {
	// ID of the ResourceRevision entity.
	ID object.ID `path:"-" query:"-" json:"id"`

	// Version of the template.
	TemplateVersion string `path:"-" query:"-" json:"templateVersion"`
	// Attributes to configure the template.
	Attributes property.Values `path:"-" query:"-" json:"attributes,omitempty"`
	// Variables of the revision.
	Variables crypto.Map[string, string] `path:"-" query:"-" json:"variables"`
	// Input plan of the revision.
	InputPlan string `path:"-" query:"-" json:"inputPlan"`
	// Output of the revision.
	Output string `path:"-" query:"-" json:"output"`
	// Type of deployer.
	DeployerType string `path:"-" query:"-" json:"deployerType"`
	// Duration in seconds of the revision deploying.
	Duration int `path:"-" query:"-" json:"duration"`
	// Previous provider requirement of the revision.
	PreviousRequiredProviders []types.ProviderRequirement `path:"-" query:"-" json:"previousRequiredProviders"`
	// Record of the revision.
	Record string `path:"-" query:"-" json:"record,omitempty"`
	// Change comment of the revision.
	ChangeComment string `path:"-" query:"-" json:"changeComment,omitempty"`
}

// ValidateWith checks the ResourceRevisionUpdateInputsItem entity with the given context and client set.
func (rrui *ResourceRevisionUpdateInputsItem) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if rrui == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	return nil
}

// ResourceRevisionUpdateInputs holds the modification input of the ResourceRevision entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type ResourceRevisionUpdateInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to update ResourceRevision entity MUST under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`
	// Environment indicates to update ResourceRevision entity MUST under the Environment route.
	Environment *EnvironmentQueryInput `path:",inline" query:"-" json:"-"`
	// Resource indicates to update ResourceRevision entity MUST under the Resource route.
	Resource *ResourceQueryInput `path:",inline" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*ResourceRevisionUpdateInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the ResourceRevision entities for modifying,
// after validating.
func (rrui *ResourceRevisionUpdateInputs) Model() []*ResourceRevision {
	if rrui == nil || len(rrui.Items) == 0 {
		return nil
	}

	_rrs := make([]*ResourceRevision, len(rrui.Items))

	for i := range rrui.Items {
		_rr := &ResourceRevision{
			ID:                        rrui.Items[i].ID,
			TemplateVersion:           rrui.Items[i].TemplateVersion,
			Attributes:                rrui.Items[i].Attributes,
			Variables:                 rrui.Items[i].Variables,
			InputPlan:                 rrui.Items[i].InputPlan,
			Output:                    rrui.Items[i].Output,
			DeployerType:              rrui.Items[i].DeployerType,
			Duration:                  rrui.Items[i].Duration,
			PreviousRequiredProviders: rrui.Items[i].PreviousRequiredProviders,
			Record:                    rrui.Items[i].Record,
			ChangeComment:             rrui.Items[i].ChangeComment,
		}

		_rrs[i] = _rr
	}

	return _rrs
}

// IDs returns the ID list of the ResourceRevision entities for modifying,
// after validating.
func (rrui *ResourceRevisionUpdateInputs) IDs() []object.ID {
	if rrui == nil || len(rrui.Items) == 0 {
		return nil
	}

	ids := make([]object.ID, len(rrui.Items))
	for i := range rrui.Items {
		ids[i] = rrui.Items[i].ID
	}
	return ids
}

// Validate checks the ResourceRevisionUpdateInputs entity.
func (rrui *ResourceRevisionUpdateInputs) Validate() error {
	if rrui == nil {
		return errors.New("nil receiver")
	}

	return rrui.ValidateWith(rrui.inputConfig.Context, rrui.inputConfig.Client, nil)
}

// ValidateWith checks the ResourceRevisionUpdateInputs entity with the given context and client set.
func (rrui *ResourceRevisionUpdateInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if rrui == nil {
		return errors.New("nil receiver")
	}

	if len(rrui.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.ResourceRevisions().Query()

	// Validate when updating under the Project route.
	if rrui.Project != nil {
		if err := rrui.Project.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			ctx = valueContext(ctx, intercept.WithProjectInterceptor)
			q.Where(
				resourcerevision.ProjectID(rrui.Project.ID))
		}
	}

	// Validate when updating under the Environment route.
	if rrui.Environment != nil {
		if err := rrui.Environment.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			q.Where(
				resourcerevision.EnvironmentID(rrui.Environment.ID))
		}
	}

	// Validate when updating under the Resource route.
	if rrui.Resource != nil {
		if err := rrui.Resource.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			q.Where(
				resourcerevision.ResourceID(rrui.Resource.ID))
		}
	}

	ids := make([]object.ID, 0, len(rrui.Items))

	for i := range rrui.Items {
		if rrui.Items[i] == nil {
			return errors.New("nil item")
		}

		if rrui.Items[i].ID != "" {
			ids = append(ids, rrui.Items[i].ID)
		} else {
			return errors.New("found item hasn't identify")
		}
	}

	if len(ids) != cap(ids) {
		return errors.New("found unrecognized item")
	}

	idsCnt, err := q.Where(resourcerevision.IDIn(ids...)).
		Count(ctx)
	if err != nil {
		return err
	}

	if idsCnt != cap(ids) {
		return errors.New("found unrecognized item")
	}

	for i := range rrui.Items {
		if err := rrui.Items[i].ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	return nil
}

// ResourceRevisionOutput holds the output of the ResourceRevision entity.
type ResourceRevisionOutput struct {
	ID                        object.ID                   `json:"id,omitempty"`
	CreateTime                *time.Time                  `json:"createTime,omitempty"`
	Status                    status.Status               `json:"status,omitempty"`
	TemplateName              string                      `json:"templateName,omitempty"`
	TemplateVersion           string                      `json:"templateVersion,omitempty"`
	TemplateID                object.ID                   `json:"templateID,omitempty"`
	Attributes                property.Values             `json:"attributes,omitempty"`
	Variables                 crypto.Map[string, string]  `json:"variables,omitempty"`
	DeployerType              string                      `json:"deployerType,omitempty"`
	Duration                  int                         `json:"duration,omitempty"`
	PreviousRequiredProviders []types.ProviderRequirement `json:"previousRequiredProviders,omitempty"`
	Record                    string                      `json:"record,omitempty"`
	ChangeComment             string                      `json:"changeComment,omitempty"`
	CreatedBy                 string                      `json:"createdBy,omitempty"`

	Project     *ProjectOutput     `json:"project,omitempty"`
	Environment *EnvironmentOutput `json:"environment,omitempty"`
	Resource    *ResourceOutput    `json:"resource,omitempty"`
}

// View returns the output of ResourceRevision entity.
func (_rr *ResourceRevision) View() *ResourceRevisionOutput {
	return ExposeResourceRevision(_rr)
}

// View returns the output of ResourceRevision entities.
func (_rrs ResourceRevisions) View() []*ResourceRevisionOutput {
	return ExposeResourceRevisions(_rrs)
}

// ExposeResourceRevision converts the ResourceRevision to ResourceRevisionOutput.
func ExposeResourceRevision(_rr *ResourceRevision) *ResourceRevisionOutput {
	if _rr == nil {
		return nil
	}

	rro := &ResourceRevisionOutput{
		ID:                        _rr.ID,
		CreateTime:                _rr.CreateTime,
		Status:                    _rr.Status,
		TemplateName:              _rr.TemplateName,
		TemplateVersion:           _rr.TemplateVersion,
		TemplateID:                _rr.TemplateID,
		Attributes:                _rr.Attributes,
		Variables:                 _rr.Variables,
		DeployerType:              _rr.DeployerType,
		Duration:                  _rr.Duration,
		PreviousRequiredProviders: _rr.PreviousRequiredProviders,
		Record:                    _rr.Record,
		ChangeComment:             _rr.ChangeComment,
		CreatedBy:                 _rr.CreatedBy,
	}

	if _rr.Edges.Project != nil {
		rro.Project = ExposeProject(_rr.Edges.Project)
	} else if _rr.ProjectID != "" {
		rro.Project = &ProjectOutput{
			ID: _rr.ProjectID,
		}
	}
	if _rr.Edges.Environment != nil {
		rro.Environment = ExposeEnvironment(_rr.Edges.Environment)
	} else if _rr.EnvironmentID != "" {
		rro.Environment = &EnvironmentOutput{
			ID: _rr.EnvironmentID,
		}
	}
	if _rr.Edges.Resource != nil {
		rro.Resource = ExposeResource(_rr.Edges.Resource)
	} else if _rr.ResourceID != "" {
		rro.Resource = &ResourceOutput{
			ID: _rr.ResourceID,
		}
	}
	return rro
}

// ExposeResourceRevisions converts the ResourceRevision slice to ResourceRevisionOutput pointer slice.
func ExposeResourceRevisions(_rrs []*ResourceRevision) []*ResourceRevisionOutput {
	if len(_rrs) == 0 {
		return nil
	}

	rros := make([]*ResourceRevisionOutput, len(_rrs))
	for i := range _rrs {
		rros[i] = ExposeResourceRevision(_rrs[i])
	}
	return rros
}
