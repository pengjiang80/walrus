package template

import (
	"entgo.io/ent/dialect/sql"

	modbus "github.com/seal-io/walrus/pkg/bus/template"
	"github.com/seal-io/walrus/pkg/dao"
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/model/template"
	"github.com/seal-io/walrus/pkg/dao/model/templateversion"
	"github.com/seal-io/walrus/pkg/dao/types/status"
)

func (h Handler) RouteRefresh(req RouteRefreshRequest) error {
	entity, err := h.modelClient.Templates().
		Get(req.Context, req.ID)
	if err != nil {
		return err
	}

	status.TemplateStatusInitialized.Unknown(entity, "Initializing template")
	entity.Status.SetSummary(status.WalkTemplate(&entity.Status))

	entity, err = h.modelClient.Templates().UpdateOne(entity).
		Set(entity).
		Save(req.Context)
	if err != nil {
		return err
	}

	return modbus.Notify(req.Context, entity)
}

var (
	queryVersionFields = []string{
		templateversion.FieldVersion,
	}
	getVersionFields = templateversion.WithoutFields()
)

func (h Handler) RouteGetVersions(
	req RouteGetVersionsRequest,
) (RouteGetVersionsResponse, int, error) {
	query := h.modelClient.TemplateVersions().Query().
		Where(templateversion.TemplateID(req.ID))

	if queries, ok := req.Querying(queryVersionFields); ok {
		query.Where(queries)
	}

	// Get count.
	cnt, err := query.Clone().Count(req.Context)
	if err != nil {
		return nil, 0, err
	}

	// Get entities.
	if limit, offset, ok := req.Paging(); ok {
		query.Limit(limit).Offset(offset)
	}

	if fields, ok := req.Extracting(getVersionFields, getVersionFields...); ok {
		query.Select(fields...)
	}

	// Borrow From https://github.com/Masterminds/semver/blob/2f39fdc11c33c38e8b8b15b1f04334ba84e751f2/version.go#L42.
	semverExpression := `^v?([0-9]+)(\.[0-9]+)?(\.[0-9]+)?` +
		`(-([0-9A-Za-z\-]+(\.[0-9A-Za-z\-]+)*))?` +
		`(\+([0-9A-Za-z\-]+(\.[0-9A-Za-z\-]+)*))?$`

	query.Order(model.Desc(templateversion.FieldCreateTime), func(s *sql.Selector) {
		s.OrderExprFunc(func(b *sql.Builder) {
			b.WriteString("CASE WHEN")
			b.Ident(templateversion.FieldVersion)
			b.WriteString(" ~ '" + semverExpression + "' THEN string_to_array(regexp_replace(")
			b.Ident(templateversion.FieldVersion)
			b.WriteString(", E'[^0-9\\.]+','', 'g'), '.', '')::int[] ELSE NULL END DESC")
		})
	})

	entities, err := query.
		// Must extract template.
		Select(templateversion.FieldTemplateID).
		WithTemplate(func(tq *model.TemplateQuery) {
			tq.Select(
				template.FieldID,
				template.FieldName)
		}).
		Unique(false).
		All(req.Context)
	if err != nil {
		return nil, 0, err
	}

	// Set expose schema.
	return dao.ExposeTemplateVersions(entities), cnt, nil
}
