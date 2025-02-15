package unreachable

import (
	"context"
	"errors"

	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/dao/types/status"
	optypes "github.com/seal-io/walrus/pkg/operator/types"
)

const OperatorType = "UnReachable"

// New returns types.Operator with the given options.
func New(ctx context.Context, opts optypes.CreateOptions) (optypes.Operator, error) {
	return Operator{}, nil
}

type Operator struct{}

func (Operator) Type() optypes.Type {
	return OperatorType
}

func (Operator) IsConnected(ctx context.Context) error {
	return nil
}

func (op Operator) Burst() int {
	return 100
}

func (op Operator) ID() string {
	return ""
}

func (Operator) GetKeys(
	ctx context.Context,
	resource *model.ResourceComponent,
) (*types.ResourceComponentOperationKeys, error) {
	return nil, nil
}

func (Operator) GetStatus(ctx context.Context, resource *model.ResourceComponent) (*status.Status, error) {
	return &status.Status{
		Summary: status.Summary{
			SummaryStatus:        "Unknown",
			SummaryStatusMessage: "Connector Error: unreachable",
			Error:                true,
		},
	}, nil
}

func (Operator) GetEndpoints(
	ctx context.Context,
	resource *model.ResourceComponent,
) ([]types.ResourceComponentEndpoint, error) {
	return nil, nil
}

func (Operator) GetComponents(
	ctx context.Context,
	resource *model.ResourceComponent,
) ([]*model.ResourceComponent, error) {
	return nil, nil
}

func (Operator) Log(ctx context.Context, s string, options optypes.LogOptions) error {
	return errors.New("cannot log")
}

func (Operator) Exec(ctx context.Context, s string, options optypes.ExecOptions) error {
	return errors.New("cannot execute")
}

func (Operator) Label(ctx context.Context, resource *model.ResourceComponent, m map[string]string) error {
	return nil
}
