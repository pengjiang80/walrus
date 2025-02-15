// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package model

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"

	"github.com/seal-io/walrus/pkg/dao/model/environment"
	"github.com/seal-io/walrus/pkg/dao/model/project"
	"github.com/seal-io/walrus/pkg/dao/model/resource"
	"github.com/seal-io/walrus/pkg/dao/model/resourcerevision"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/dao/types/crypto"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/pkg/dao/types/property"
	"github.com/seal-io/walrus/pkg/dao/types/status"
	"github.com/seal-io/walrus/utils/json"
)

// ResourceRevision is the model entity for the ResourceRevision schema.
type ResourceRevision struct {
	config `json:"-"`
	// ID of the ent.
	ID object.ID `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// Status holds the value of the "status" field.
	Status status.Status `json:"status,omitempty"`
	// ID of the project to belong.
	ProjectID object.ID `json:"project_id,omitempty"`
	// ID of the environment to which the revision belongs.
	EnvironmentID object.ID `json:"environment_id,omitempty"`
	// ID of the resource to which the revision belongs.
	ResourceID object.ID `json:"resource_id,omitempty"`
	// Name of the template.
	TemplateName string `json:"template_name,omitempty"`
	// Version of the template.
	TemplateVersion string `json:"template_version,omitempty"`
	// ID of the template.
	TemplateID object.ID `json:"template_id,omitempty"`
	// Attributes to configure the template.
	Attributes property.Values `json:"attributes,omitempty"`
	// Variables of the revision.
	Variables crypto.Map[string, string] `json:"variables,omitempty"`
	// Input plan of the revision.
	InputPlan string `json:"-"`
	// Output of the revision.
	Output string `json:"-"`
	// Type of deployer.
	DeployerType string `json:"deployer_type,omitempty"`
	// Duration in seconds of the revision deploying.
	Duration int `json:"duration,omitempty"`
	// Previous provider requirement of the revision.
	PreviousRequiredProviders []types.ProviderRequirement `json:"previous_required_providers,omitempty"`
	// Record of the revision.
	Record string `json:"record,omitempty"`
	// Change comment of the revision.
	ChangeComment string `json:"change_comment,omitempty"`
	// User who created the revision.
	CreatedBy string `json:"created_by,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ResourceRevisionQuery when eager-loading is set.
	Edges        ResourceRevisionEdges `json:"edges,omitempty"`
	selectValues sql.SelectValues
}

// ResourceRevisionEdges holds the relations/edges for other nodes in the graph.
type ResourceRevisionEdges struct {
	// Project to which the revision belongs.
	Project *Project `json:"project,omitempty"`
	// Environment to which the revision deploys.
	Environment *Environment `json:"environment,omitempty"`
	// Resource to which the revision belongs.
	Resource *Resource `json:"resource,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ProjectOrErr returns the Project value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ResourceRevisionEdges) ProjectOrErr() (*Project, error) {
	if e.loadedTypes[0] {
		if e.Project == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: project.Label}
		}
		return e.Project, nil
	}
	return nil, &NotLoadedError{edge: "project"}
}

// EnvironmentOrErr returns the Environment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ResourceRevisionEdges) EnvironmentOrErr() (*Environment, error) {
	if e.loadedTypes[1] {
		if e.Environment == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: environment.Label}
		}
		return e.Environment, nil
	}
	return nil, &NotLoadedError{edge: "environment"}
}

// ResourceOrErr returns the Resource value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ResourceRevisionEdges) ResourceOrErr() (*Resource, error) {
	if e.loadedTypes[2] {
		if e.Resource == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: resource.Label}
		}
		return e.Resource, nil
	}
	return nil, &NotLoadedError{edge: "resource"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ResourceRevision) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case resourcerevision.FieldStatus, resourcerevision.FieldPreviousRequiredProviders:
			values[i] = new([]byte)
		case resourcerevision.FieldVariables:
			values[i] = new(crypto.Map[string, string])
		case resourcerevision.FieldID, resourcerevision.FieldProjectID, resourcerevision.FieldEnvironmentID, resourcerevision.FieldResourceID, resourcerevision.FieldTemplateID:
			values[i] = new(object.ID)
		case resourcerevision.FieldAttributes:
			values[i] = new(property.Values)
		case resourcerevision.FieldDuration:
			values[i] = new(sql.NullInt64)
		case resourcerevision.FieldTemplateName, resourcerevision.FieldTemplateVersion, resourcerevision.FieldInputPlan, resourcerevision.FieldOutput, resourcerevision.FieldDeployerType, resourcerevision.FieldRecord, resourcerevision.FieldChangeComment, resourcerevision.FieldCreatedBy:
			values[i] = new(sql.NullString)
		case resourcerevision.FieldCreateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ResourceRevision fields.
func (rr *ResourceRevision) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case resourcerevision.FieldID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				rr.ID = *value
			}
		case resourcerevision.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				rr.CreateTime = new(time.Time)
				*rr.CreateTime = value.Time
			}
		case resourcerevision.FieldStatus:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &rr.Status); err != nil {
					return fmt.Errorf("unmarshal field status: %w", err)
				}
			}
		case resourcerevision.FieldProjectID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field project_id", values[i])
			} else if value != nil {
				rr.ProjectID = *value
			}
		case resourcerevision.FieldEnvironmentID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field environment_id", values[i])
			} else if value != nil {
				rr.EnvironmentID = *value
			}
		case resourcerevision.FieldResourceID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field resource_id", values[i])
			} else if value != nil {
				rr.ResourceID = *value
			}
		case resourcerevision.FieldTemplateName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field template_name", values[i])
			} else if value.Valid {
				rr.TemplateName = value.String
			}
		case resourcerevision.FieldTemplateVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field template_version", values[i])
			} else if value.Valid {
				rr.TemplateVersion = value.String
			}
		case resourcerevision.FieldTemplateID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field template_id", values[i])
			} else if value != nil {
				rr.TemplateID = *value
			}
		case resourcerevision.FieldAttributes:
			if value, ok := values[i].(*property.Values); !ok {
				return fmt.Errorf("unexpected type %T for field attributes", values[i])
			} else if value != nil {
				rr.Attributes = *value
			}
		case resourcerevision.FieldVariables:
			if value, ok := values[i].(*crypto.Map[string, string]); !ok {
				return fmt.Errorf("unexpected type %T for field variables", values[i])
			} else if value != nil {
				rr.Variables = *value
			}
		case resourcerevision.FieldInputPlan:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field input_plan", values[i])
			} else if value.Valid {
				rr.InputPlan = value.String
			}
		case resourcerevision.FieldOutput:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field output", values[i])
			} else if value.Valid {
				rr.Output = value.String
			}
		case resourcerevision.FieldDeployerType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deployer_type", values[i])
			} else if value.Valid {
				rr.DeployerType = value.String
			}
		case resourcerevision.FieldDuration:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field duration", values[i])
			} else if value.Valid {
				rr.Duration = int(value.Int64)
			}
		case resourcerevision.FieldPreviousRequiredProviders:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field previous_required_providers", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &rr.PreviousRequiredProviders); err != nil {
					return fmt.Errorf("unmarshal field previous_required_providers: %w", err)
				}
			}
		case resourcerevision.FieldRecord:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field record", values[i])
			} else if value.Valid {
				rr.Record = value.String
			}
		case resourcerevision.FieldChangeComment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field change_comment", values[i])
			} else if value.Valid {
				rr.ChangeComment = value.String
			}
		case resourcerevision.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				rr.CreatedBy = value.String
			}
		default:
			rr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ResourceRevision.
// This includes values selected through modifiers, order, etc.
func (rr *ResourceRevision) Value(name string) (ent.Value, error) {
	return rr.selectValues.Get(name)
}

// QueryProject queries the "project" edge of the ResourceRevision entity.
func (rr *ResourceRevision) QueryProject() *ProjectQuery {
	return NewResourceRevisionClient(rr.config).QueryProject(rr)
}

// QueryEnvironment queries the "environment" edge of the ResourceRevision entity.
func (rr *ResourceRevision) QueryEnvironment() *EnvironmentQuery {
	return NewResourceRevisionClient(rr.config).QueryEnvironment(rr)
}

// QueryResource queries the "resource" edge of the ResourceRevision entity.
func (rr *ResourceRevision) QueryResource() *ResourceQuery {
	return NewResourceRevisionClient(rr.config).QueryResource(rr)
}

// Update returns a builder for updating this ResourceRevision.
// Note that you need to call ResourceRevision.Unwrap() before calling this method if this ResourceRevision
// was returned from a transaction, and the transaction was committed or rolled back.
func (rr *ResourceRevision) Update() *ResourceRevisionUpdateOne {
	return NewResourceRevisionClient(rr.config).UpdateOne(rr)
}

// Unwrap unwraps the ResourceRevision entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rr *ResourceRevision) Unwrap() *ResourceRevision {
	_tx, ok := rr.config.driver.(*txDriver)
	if !ok {
		panic("model: ResourceRevision is not a transactional entity")
	}
	rr.config.driver = _tx.drv
	return rr
}

// String implements the fmt.Stringer.
func (rr *ResourceRevision) String() string {
	var builder strings.Builder
	builder.WriteString("ResourceRevision(")
	builder.WriteString(fmt.Sprintf("id=%v, ", rr.ID))
	if v := rr.CreateTime; v != nil {
		builder.WriteString("create_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", rr.Status))
	builder.WriteString(", ")
	builder.WriteString("project_id=")
	builder.WriteString(fmt.Sprintf("%v", rr.ProjectID))
	builder.WriteString(", ")
	builder.WriteString("environment_id=")
	builder.WriteString(fmt.Sprintf("%v", rr.EnvironmentID))
	builder.WriteString(", ")
	builder.WriteString("resource_id=")
	builder.WriteString(fmt.Sprintf("%v", rr.ResourceID))
	builder.WriteString(", ")
	builder.WriteString("template_name=")
	builder.WriteString(rr.TemplateName)
	builder.WriteString(", ")
	builder.WriteString("template_version=")
	builder.WriteString(rr.TemplateVersion)
	builder.WriteString(", ")
	builder.WriteString("template_id=")
	builder.WriteString(fmt.Sprintf("%v", rr.TemplateID))
	builder.WriteString(", ")
	builder.WriteString("attributes=")
	builder.WriteString(fmt.Sprintf("%v", rr.Attributes))
	builder.WriteString(", ")
	builder.WriteString("variables=")
	builder.WriteString(fmt.Sprintf("%v", rr.Variables))
	builder.WriteString(", ")
	builder.WriteString("input_plan=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("output=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("deployer_type=")
	builder.WriteString(rr.DeployerType)
	builder.WriteString(", ")
	builder.WriteString("duration=")
	builder.WriteString(fmt.Sprintf("%v", rr.Duration))
	builder.WriteString(", ")
	builder.WriteString("previous_required_providers=")
	builder.WriteString(fmt.Sprintf("%v", rr.PreviousRequiredProviders))
	builder.WriteString(", ")
	builder.WriteString("record=")
	builder.WriteString(rr.Record)
	builder.WriteString(", ")
	builder.WriteString("change_comment=")
	builder.WriteString(rr.ChangeComment)
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(rr.CreatedBy)
	builder.WriteByte(')')
	return builder.String()
}

// ResourceRevisions is a parsable slice of ResourceRevision.
type ResourceRevisions []*ResourceRevision
