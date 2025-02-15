// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/seal-io/walrus/pkg/dao/model/predicate"
	"github.com/seal-io/walrus/pkg/dao/model/subject"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/utils/json"
)

// SubjectCreateInput holds the creation input of the Subject entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type SubjectCreateInput struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// The name of the subject.
	Name string `path:"-" query:"-" json:"name"`
	// The kind of the subject.
	Kind string `path:"-" query:"-" json:"kind,omitempty"`
	// The domain of the subject.
	Domain string `path:"-" query:"-" json:"domain,omitempty"`
	// The detail of the subject.
	Description string `path:"-" query:"-" json:"description,omitempty"`
	// Indicate whether the subject is builtin, decide when creating.
	Builtin bool `path:"-" query:"-" json:"builtin,omitempty"`

	// Roles specifies full inserting the new SubjectRoleRelationship entities of the Subject entity.
	Roles []*SubjectRoleRelationshipCreateInput `uri:"-" query:"-" json:"roles,omitempty"`
}

// Model returns the Subject entity for creating,
// after validating.
func (sci *SubjectCreateInput) Model() *Subject {
	if sci == nil {
		return nil
	}

	_s := &Subject{
		Name:        sci.Name,
		Kind:        sci.Kind,
		Domain:      sci.Domain,
		Description: sci.Description,
		Builtin:     sci.Builtin,
	}

	if sci.Roles != nil {
		// Empty slice is used for clearing the edge.
		_s.Edges.Roles = make([]*SubjectRoleRelationship, 0, len(sci.Roles))
	}
	for j := range sci.Roles {
		if sci.Roles[j] == nil {
			continue
		}
		_s.Edges.Roles = append(_s.Edges.Roles,
			sci.Roles[j].Model())
	}
	return _s
}

// Validate checks the SubjectCreateInput entity.
func (sci *SubjectCreateInput) Validate() error {
	if sci == nil {
		return errors.New("nil receiver")
	}

	return sci.ValidateWith(sci.inputConfig.Context, sci.inputConfig.Client, nil)
}

// ValidateWith checks the SubjectCreateInput entity with the given context and client set.
func (sci *SubjectCreateInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if sci == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	for i := range sci.Roles {
		if sci.Roles[i] == nil {
			continue
		}

		if err := sci.Roles[i].ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				sci.Roles[i] = nil
			}
		}
	}

	return nil
}

// SubjectCreateInputs holds the creation input item of the Subject entities.
type SubjectCreateInputsItem struct {
	// The name of the subject.
	Name string `path:"-" query:"-" json:"name"`
	// The kind of the subject.
	Kind string `path:"-" query:"-" json:"kind,omitempty"`
	// The domain of the subject.
	Domain string `path:"-" query:"-" json:"domain,omitempty"`
	// The detail of the subject.
	Description string `path:"-" query:"-" json:"description,omitempty"`
	// Indicate whether the subject is builtin, decide when creating.
	Builtin bool `path:"-" query:"-" json:"builtin,omitempty"`

	// Roles specifies full inserting the new SubjectRoleRelationship entities.
	Roles []*SubjectRoleRelationshipCreateInput `uri:"-" query:"-" json:"roles,omitempty"`
}

// ValidateWith checks the SubjectCreateInputsItem entity with the given context and client set.
func (sci *SubjectCreateInputsItem) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if sci == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	for i := range sci.Roles {
		if sci.Roles[i] == nil {
			continue
		}

		if err := sci.Roles[i].ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				sci.Roles[i] = nil
			}
		}
	}

	return nil
}

// SubjectCreateInputs holds the creation input of the Subject entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type SubjectCreateInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*SubjectCreateInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the Subject entities for creating,
// after validating.
func (sci *SubjectCreateInputs) Model() []*Subject {
	if sci == nil || len(sci.Items) == 0 {
		return nil
	}

	_ss := make([]*Subject, len(sci.Items))

	for i := range sci.Items {
		_s := &Subject{
			Name:        sci.Items[i].Name,
			Kind:        sci.Items[i].Kind,
			Domain:      sci.Items[i].Domain,
			Description: sci.Items[i].Description,
			Builtin:     sci.Items[i].Builtin,
		}

		if sci.Items[i].Roles != nil {
			// Empty slice is used for clearing the edge.
			_s.Edges.Roles = make([]*SubjectRoleRelationship, 0, len(sci.Items[i].Roles))
		}
		for j := range sci.Items[i].Roles {
			if sci.Items[i].Roles[j] == nil {
				continue
			}
			_s.Edges.Roles = append(_s.Edges.Roles,
				sci.Items[i].Roles[j].Model())
		}

		_ss[i] = _s
	}

	return _ss
}

// Validate checks the SubjectCreateInputs entity .
func (sci *SubjectCreateInputs) Validate() error {
	if sci == nil {
		return errors.New("nil receiver")
	}

	return sci.ValidateWith(sci.inputConfig.Context, sci.inputConfig.Client, nil)
}

// ValidateWith checks the SubjectCreateInputs entity with the given context and client set.
func (sci *SubjectCreateInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if sci == nil {
		return errors.New("nil receiver")
	}

	if len(sci.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	for i := range sci.Items {
		if sci.Items[i] == nil {
			continue
		}

		if err := sci.Items[i].ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	return nil
}

// SubjectDeleteInput holds the deletion input of the Subject entity,
// please tags with `path:",inline"` if embedding.
type SubjectDeleteInput struct {
	SubjectQueryInput `path:",inline"`
}

// SubjectDeleteInputs holds the deletion input item of the Subject entities.
type SubjectDeleteInputsItem struct {
	// ID of the Subject entity, tries to retrieve the entity with the following unique index parts if no ID provided.
	ID object.ID `path:"-" query:"-" json:"id,omitempty"`
	// Kind of the Subject entity, a part of the unique index.
	Kind string `path:"-" query:"-" json:"kind,omitempty"`
	// Domain of the Subject entity, a part of the unique index.
	Domain string `path:"-" query:"-" json:"domain,omitempty"`
	// Name of the Subject entity, a part of the unique index.
	Name string `path:"-" query:"-" json:"name,omitempty"`
}

// SubjectDeleteInputs holds the deletion input of the Subject entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type SubjectDeleteInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*SubjectDeleteInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the Subject entities for deleting,
// after validating.
func (sdi *SubjectDeleteInputs) Model() []*Subject {
	if sdi == nil || len(sdi.Items) == 0 {
		return nil
	}

	_ss := make([]*Subject, len(sdi.Items))
	for i := range sdi.Items {
		_ss[i] = &Subject{
			ID: sdi.Items[i].ID,
		}
	}
	return _ss
}

// IDs returns the ID list of the Subject entities for deleting,
// after validating.
func (sdi *SubjectDeleteInputs) IDs() []object.ID {
	if sdi == nil || len(sdi.Items) == 0 {
		return nil
	}

	ids := make([]object.ID, len(sdi.Items))
	for i := range sdi.Items {
		ids[i] = sdi.Items[i].ID
	}
	return ids
}

// Validate checks the SubjectDeleteInputs entity.
func (sdi *SubjectDeleteInputs) Validate() error {
	if sdi == nil {
		return errors.New("nil receiver")
	}

	return sdi.ValidateWith(sdi.inputConfig.Context, sdi.inputConfig.Client, nil)
}

// ValidateWith checks the SubjectDeleteInputs entity with the given context and client set.
func (sdi *SubjectDeleteInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if sdi == nil {
		return errors.New("nil receiver")
	}

	if len(sdi.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.Subjects().Query()

	ids := make([]object.ID, 0, len(sdi.Items))
	ors := make([]predicate.Subject, 0, len(sdi.Items))
	indexers := make(map[any][]int)

	for i := range sdi.Items {
		if sdi.Items[i] == nil {
			return errors.New("nil item")
		}

		if sdi.Items[i].ID != "" {
			ids = append(ids, sdi.Items[i].ID)
			ors = append(ors, subject.ID(sdi.Items[i].ID))
			indexers[sdi.Items[i].ID] = append(indexers[sdi.Items[i].ID], i)
		} else if (sdi.Items[i].Kind != "") && (sdi.Items[i].Domain != "") && (sdi.Items[i].Name != "") {
			ors = append(ors, subject.And(
				subject.Kind(sdi.Items[i].Kind),
				subject.Domain(sdi.Items[i].Domain),
				subject.Name(sdi.Items[i].Name)))
			indexerKey := fmt.Sprint("/", sdi.Items[i].Kind, "/", sdi.Items[i].Domain, "/", sdi.Items[i].Name)
			indexers[indexerKey] = append(indexers[indexerKey], i)
		} else {
			return errors.New("found item hasn't identify")
		}
	}

	p := subject.IDIn(ids...)
	if len(ids) != cap(ids) {
		p = subject.Or(ors...)
	}

	es, err := q.
		Where(p).
		Select(
			subject.FieldID,
			subject.FieldKind,
			subject.FieldDomain,
			subject.FieldName,
		).
		All(ctx)
	if err != nil {
		return err
	}

	if len(es) != cap(ids) {
		return errors.New("found unrecognized item")
	}

	for i := range es {
		indexer := indexers[es[i].ID]
		if indexer == nil {
			indexerKey := fmt.Sprint("/", es[i].Kind, "/", es[i].Domain, "/", es[i].Name)
			indexer = indexers[indexerKey]
		}
		for _, j := range indexer {
			sdi.Items[j].ID = es[i].ID
			sdi.Items[j].Kind = es[i].Kind
			sdi.Items[j].Domain = es[i].Domain
			sdi.Items[j].Name = es[i].Name
		}
	}

	return nil
}

// SubjectPatchInput holds the patch input of the Subject entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type SubjectPatchInput struct {
	SubjectUpdateInput `path:",inline" query:"-" json:",inline"`

	patchedEntity *Subject `path:"-" query:"-" json:"-"`
}

// Model returns the Subject patched entity,
// after validating.
func (spi *SubjectPatchInput) Model() *Subject {
	if spi == nil {
		return nil
	}

	return spi.patchedEntity
}

// Validate checks the SubjectPatchInput entity.
func (spi *SubjectPatchInput) Validate() error {
	if spi == nil {
		return errors.New("nil receiver")
	}

	return spi.ValidateWith(spi.inputConfig.Context, spi.inputConfig.Client, nil)
}

// ValidateWith checks the SubjectPatchInput entity with the given context and client set.
func (spi *SubjectPatchInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if cache == nil {
		cache = map[string]any{}
	}

	if err := spi.SubjectUpdateInput.ValidateWith(ctx, cs, cache); err != nil {
		return err
	}

	q := cs.Subjects().Query()

	if spi.Refer != nil {
		if spi.Refer.IsID() {
			q.Where(
				subject.ID(spi.Refer.ID()))
		} else if refers := spi.Refer.Split(3); len(refers) == 3 {
			q.Where(
				subject.Kind(refers[0].String()),
				subject.Domain(refers[1].String()),
				subject.Name(refers[2].String()))
		} else {
			return errors.New("invalid identify refer of subject")
		}
	} else if spi.ID != "" {
		q.Where(
			subject.ID(spi.ID))
	} else if (spi.Kind != "") && (spi.Domain != "") && (spi.Name != "") {
		q.Where(
			subject.Kind(spi.Kind),
			subject.Domain(spi.Domain),
			subject.Name(spi.Name))
	} else {
		return errors.New("invalid identify of subject")
	}

	q.Select(
		subject.WithoutFields(
			subject.FieldCreateTime,
			subject.FieldUpdateTime,
		)...,
	)

	var e *Subject
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
			e = cv.(*Subject)
		}
	}

	_s := spi.SubjectUpdateInput.Model()

	_obj, err := json.PatchObject(e, _s)
	if err != nil {
		return err
	}

	spi.patchedEntity = _obj.(*Subject)
	return nil
}

// SubjectQueryInput holds the query input of the Subject entity,
// please tags with `path:",inline"` if embedding.
type SubjectQueryInput struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Refer holds the route path reference of the Subject entity.
	Refer *object.Refer `path:"subject,default=" query:"-" json:"-"`
	// ID of the Subject entity, tries to retrieve the entity with the following unique index parts if no ID provided.
	ID object.ID `path:"-" query:"-" json:"id,omitempty"`
	// Kind of the Subject entity, a part of the unique index.
	Kind string `path:"-" query:"-" json:"kind,omitempty"`
	// Domain of the Subject entity, a part of the unique index.
	Domain string `path:"-" query:"-" json:"domain,omitempty"`
	// Name of the Subject entity, a part of the unique index.
	Name string `path:"-" query:"-" json:"name,omitempty"`
}

// Model returns the Subject entity for querying,
// after validating.
func (sqi *SubjectQueryInput) Model() *Subject {
	if sqi == nil {
		return nil
	}

	return &Subject{
		ID:     sqi.ID,
		Kind:   sqi.Kind,
		Domain: sqi.Domain,
		Name:   sqi.Name,
	}
}

// Validate checks the SubjectQueryInput entity.
func (sqi *SubjectQueryInput) Validate() error {
	if sqi == nil {
		return errors.New("nil receiver")
	}

	return sqi.ValidateWith(sqi.inputConfig.Context, sqi.inputConfig.Client, nil)
}

// ValidateWith checks the SubjectQueryInput entity with the given context and client set.
func (sqi *SubjectQueryInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if sqi == nil {
		return errors.New("nil receiver")
	}

	if sqi.Refer != nil && *sqi.Refer == "" {
		return fmt.Errorf("model: %s : %w", subject.Label, ErrBlankResourceRefer)
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.Subjects().Query()

	if sqi.Refer != nil {
		if sqi.Refer.IsID() {
			q.Where(
				subject.ID(sqi.Refer.ID()))
		} else if refers := sqi.Refer.Split(3); len(refers) == 3 {
			q.Where(
				subject.Kind(refers[0].String()),
				subject.Domain(refers[1].String()),
				subject.Name(refers[2].String()))
		} else {
			return errors.New("invalid identify refer of subject")
		}
	} else if sqi.ID != "" {
		q.Where(
			subject.ID(sqi.ID))
	} else if (sqi.Kind != "") && (sqi.Domain != "") && (sqi.Name != "") {
		q.Where(
			subject.Kind(sqi.Kind),
			subject.Domain(sqi.Domain),
			subject.Name(sqi.Name))
	} else {
		return errors.New("invalid identify of subject")
	}

	q.Select(
		subject.FieldID,
		subject.FieldKind,
		subject.FieldDomain,
		subject.FieldName,
	)

	var e *Subject
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
			e = cv.(*Subject)
		}
	}

	sqi.ID = e.ID
	sqi.Kind = e.Kind
	sqi.Domain = e.Domain
	sqi.Name = e.Name
	return nil
}

// SubjectQueryInputs holds the query input of the Subject entities,
// please tags with `path:",inline" query:",inline"` if embedding.
type SubjectQueryInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`
}

// Validate checks the SubjectQueryInputs entity.
func (sqi *SubjectQueryInputs) Validate() error {
	if sqi == nil {
		return errors.New("nil receiver")
	}

	return sqi.ValidateWith(sqi.inputConfig.Context, sqi.inputConfig.Client, nil)
}

// ValidateWith checks the SubjectQueryInputs entity with the given context and client set.
func (sqi *SubjectQueryInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if sqi == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	return nil
}

// SubjectUpdateInput holds the modification input of the Subject entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type SubjectUpdateInput struct {
	SubjectQueryInput `path:",inline" query:"-" json:"-"`

	// The kind of the subject.
	Kind string `path:"-" query:"-" json:"kind,omitempty"`
	// The domain of the subject.
	Domain string `path:"-" query:"-" json:"domain,omitempty"`
	// The name of the subject.
	Name string `path:"-" query:"-" json:"name,omitempty"`
	// The detail of the subject.
	Description string `path:"-" query:"-" json:"description,omitempty"`

	// Roles indicates replacing the stale SubjectRoleRelationship entities.
	Roles []*SubjectRoleRelationshipCreateInput `uri:"-" query:"-" json:"roles,omitempty"`
}

// Model returns the Subject entity for modifying,
// after validating.
func (sui *SubjectUpdateInput) Model() *Subject {
	if sui == nil {
		return nil
	}

	_s := &Subject{
		ID:          sui.ID,
		Kind:        sui.Kind,
		Domain:      sui.Domain,
		Name:        sui.Name,
		Description: sui.Description,
	}

	if sui.Roles != nil {
		// Empty slice is used for clearing the edge.
		_s.Edges.Roles = make([]*SubjectRoleRelationship, 0, len(sui.Roles))
	}
	for j := range sui.Roles {
		if sui.Roles[j] == nil {
			continue
		}
		_s.Edges.Roles = append(_s.Edges.Roles,
			sui.Roles[j].Model())
	}
	return _s
}

// Validate checks the SubjectUpdateInput entity.
func (sui *SubjectUpdateInput) Validate() error {
	if sui == nil {
		return errors.New("nil receiver")
	}

	return sui.ValidateWith(sui.inputConfig.Context, sui.inputConfig.Client, nil)
}

// ValidateWith checks the SubjectUpdateInput entity with the given context and client set.
func (sui *SubjectUpdateInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if cache == nil {
		cache = map[string]any{}
	}

	if err := sui.SubjectQueryInput.ValidateWith(ctx, cs, cache); err != nil {
		return err
	}

	for i := range sui.Roles {
		if sui.Roles[i] == nil {
			continue
		}

		if err := sui.Roles[i].ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				sui.Roles[i] = nil
			}
		}
	}

	return nil
}

// SubjectUpdateInputs holds the modification input item of the Subject entities.
type SubjectUpdateInputsItem struct {
	// ID of the Subject entity, tries to retrieve the entity with the following unique index parts if no ID provided.
	ID object.ID `path:"-" query:"-" json:"id,omitempty"`
	// Kind of the Subject entity, a part of the unique index.
	Kind string `path:"-" query:"-" json:"kind,omitempty"`
	// Domain of the Subject entity, a part of the unique index.
	Domain string `path:"-" query:"-" json:"domain,omitempty"`
	// Name of the Subject entity, a part of the unique index.
	Name string `path:"-" query:"-" json:"name,omitempty"`

	// The detail of the subject.
	Description string `path:"-" query:"-" json:"description,omitempty"`

	// Roles indicates replacing the stale SubjectRoleRelationship entities.
	Roles []*SubjectRoleRelationshipCreateInput `uri:"-" query:"-" json:"roles,omitempty"`
}

// ValidateWith checks the SubjectUpdateInputsItem entity with the given context and client set.
func (sui *SubjectUpdateInputsItem) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if sui == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	for i := range sui.Roles {
		if sui.Roles[i] == nil {
			continue
		}

		if err := sui.Roles[i].ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				sui.Roles[i] = nil
			}
		}
	}

	return nil
}

// SubjectUpdateInputs holds the modification input of the Subject entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type SubjectUpdateInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*SubjectUpdateInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the Subject entities for modifying,
// after validating.
func (sui *SubjectUpdateInputs) Model() []*Subject {
	if sui == nil || len(sui.Items) == 0 {
		return nil
	}

	_ss := make([]*Subject, len(sui.Items))

	for i := range sui.Items {
		_s := &Subject{
			ID:          sui.Items[i].ID,
			Kind:        sui.Items[i].Kind,
			Domain:      sui.Items[i].Domain,
			Name:        sui.Items[i].Name,
			Description: sui.Items[i].Description,
		}

		if sui.Items[i].Roles != nil {
			// Empty slice is used for clearing the edge.
			_s.Edges.Roles = make([]*SubjectRoleRelationship, 0, len(sui.Items[i].Roles))
		}
		for j := range sui.Items[i].Roles {
			if sui.Items[i].Roles[j] == nil {
				continue
			}
			_s.Edges.Roles = append(_s.Edges.Roles,
				sui.Items[i].Roles[j].Model())
		}

		_ss[i] = _s
	}

	return _ss
}

// IDs returns the ID list of the Subject entities for modifying,
// after validating.
func (sui *SubjectUpdateInputs) IDs() []object.ID {
	if sui == nil || len(sui.Items) == 0 {
		return nil
	}

	ids := make([]object.ID, len(sui.Items))
	for i := range sui.Items {
		ids[i] = sui.Items[i].ID
	}
	return ids
}

// Validate checks the SubjectUpdateInputs entity.
func (sui *SubjectUpdateInputs) Validate() error {
	if sui == nil {
		return errors.New("nil receiver")
	}

	return sui.ValidateWith(sui.inputConfig.Context, sui.inputConfig.Client, nil)
}

// ValidateWith checks the SubjectUpdateInputs entity with the given context and client set.
func (sui *SubjectUpdateInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if sui == nil {
		return errors.New("nil receiver")
	}

	if len(sui.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.Subjects().Query()

	ids := make([]object.ID, 0, len(sui.Items))
	ors := make([]predicate.Subject, 0, len(sui.Items))
	indexers := make(map[any][]int)

	for i := range sui.Items {
		if sui.Items[i] == nil {
			return errors.New("nil item")
		}

		if sui.Items[i].ID != "" {
			ids = append(ids, sui.Items[i].ID)
			ors = append(ors, subject.ID(sui.Items[i].ID))
			indexers[sui.Items[i].ID] = append(indexers[sui.Items[i].ID], i)
		} else if (sui.Items[i].Kind != "") && (sui.Items[i].Domain != "") && (sui.Items[i].Name != "") {
			ors = append(ors, subject.And(
				subject.Kind(sui.Items[i].Kind),
				subject.Domain(sui.Items[i].Domain),
				subject.Name(sui.Items[i].Name)))
			indexerKey := fmt.Sprint("/", sui.Items[i].Kind, "/", sui.Items[i].Domain, "/", sui.Items[i].Name)
			indexers[indexerKey] = append(indexers[indexerKey], i)
		} else {
			return errors.New("found item hasn't identify")
		}
	}

	p := subject.IDIn(ids...)
	if len(ids) != cap(ids) {
		p = subject.Or(ors...)
	}

	es, err := q.
		Where(p).
		Select(
			subject.FieldID,
			subject.FieldKind,
			subject.FieldDomain,
			subject.FieldName,
		).
		All(ctx)
	if err != nil {
		return err
	}

	if len(es) != cap(ids) {
		return errors.New("found unrecognized item")
	}

	for i := range es {
		indexer := indexers[es[i].ID]
		if indexer == nil {
			indexerKey := fmt.Sprint("/", es[i].Kind, "/", es[i].Domain, "/", es[i].Name)
			indexer = indexers[indexerKey]
		}
		for _, j := range indexer {
			sui.Items[j].ID = es[i].ID
			sui.Items[j].Kind = es[i].Kind
			sui.Items[j].Domain = es[i].Domain
			sui.Items[j].Name = es[i].Name
		}
	}

	for i := range sui.Items {
		if err := sui.Items[i].ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	return nil
}

// SubjectOutput holds the output of the Subject entity.
type SubjectOutput struct {
	ID          object.ID  `json:"id,omitempty"`
	CreateTime  *time.Time `json:"createTime,omitempty"`
	UpdateTime  *time.Time `json:"updateTime,omitempty"`
	Kind        string     `json:"kind,omitempty"`
	Domain      string     `json:"domain,omitempty"`
	Name        string     `json:"name,omitempty"`
	Description string     `json:"description,omitempty"`
	Builtin     bool       `json:"builtin,omitempty"`

	Roles []*SubjectRoleRelationshipOutput `json:"roles,omitempty"`
}

// View returns the output of Subject entity.
func (_s *Subject) View() *SubjectOutput {
	return ExposeSubject(_s)
}

// View returns the output of Subject entities.
func (_ss Subjects) View() []*SubjectOutput {
	return ExposeSubjects(_ss)
}

// ExposeSubject converts the Subject to SubjectOutput.
func ExposeSubject(_s *Subject) *SubjectOutput {
	if _s == nil {
		return nil
	}

	so := &SubjectOutput{
		ID:          _s.ID,
		CreateTime:  _s.CreateTime,
		UpdateTime:  _s.UpdateTime,
		Kind:        _s.Kind,
		Domain:      _s.Domain,
		Name:        _s.Name,
		Description: _s.Description,
		Builtin:     _s.Builtin,
	}

	if _s.Edges.Roles != nil {
		so.Roles = ExposeSubjectRoleRelationships(_s.Edges.Roles)
	}
	return so
}

// ExposeSubjects converts the Subject slice to SubjectOutput pointer slice.
func ExposeSubjects(_ss []*Subject) []*SubjectOutput {
	if len(_ss) == 0 {
		return nil
	}

	sos := make([]*SubjectOutput, len(_ss))
	for i := range _ss {
		sos[i] = ExposeSubject(_ss[i])
	}
	return sos
}
