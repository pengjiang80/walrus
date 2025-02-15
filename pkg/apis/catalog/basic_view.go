package catalog

import (
	"fmt"
	"net/url"

	"github.com/seal-io/walrus/pkg/apis/runtime"
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/model/catalog"
	"github.com/seal-io/walrus/pkg/dao/model/predicate"
	"github.com/seal-io/walrus/pkg/dao/types"
)

type (
	CreateRequest struct {
		model.CatalogCreateInput `path:",inline" json:",inline"`
	}

	CreateResponse = *model.CatalogOutput
)

func (r *CreateRequest) Validate() error {
	if err := r.CatalogCreateInput.Validate(); err != nil {
		return err
	}

	_, err := url.Parse(r.Source)
	if err != nil {
		return err
	}

	switch r.Type {
	case types.GitDriverGithub, types.GitDriverGitlab, types.GitDriverGitee:
	default:
		return fmt.Errorf("unsupported catalog type %q", r.Type)
	}

	return nil
}

type (
	GetRequest struct {
		model.CatalogQueryInput `path:",inline"`
	}

	GetResponse = *model.CatalogOutput
)

type UpdateRequest struct {
	model.CatalogUpdateInput `path:",inline" json:",inline"`
}

func (r *UpdateRequest) Validate() error {
	if err := r.CatalogUpdateInput.Validate(); err != nil {
		return err
	}

	return nil
}

type DeleteRequest = model.CatalogDeleteInput

type (
	CollectionGetRequest struct {
		model.CatalogQueryInputs `path:",inline" query:",inline"`

		runtime.RequestCollection[
			predicate.Catalog, catalog.OrderOption,
		] `query:",inline"`

		Stream *runtime.RequestUnidiStream
	}

	CollectionGetResponse = []*model.CatalogOutput
)

func (r *CollectionGetRequest) SetStream(stream runtime.RequestUnidiStream) {
	r.Stream = &stream
}

type CollectionDeleteRequest = model.CatalogDeleteInputs
