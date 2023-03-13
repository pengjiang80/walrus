package platformk8s

import (
	"context"
	"time"

	"k8s.io/client-go/rest"

	"github.com/seal-io/seal/pkg/dao/types"
	"github.com/seal-io/seal/pkg/k8s"
	"github.com/seal-io/seal/pkg/platform/operator"
	"github.com/seal-io/seal/utils/log"
)

const OperatorType = types.ConnectorTypeK8s

// NewOperator returns operator.Operator with the given options.
func NewOperator(ctx context.Context, opts operator.CreateOptions) (operator.Operator, error) {
	var restConfig, err = GetConfig(opts.Connector)
	if err != nil {
		return nil, err
	}
	var op = Operator{
		Logger:     log.WithName("operator").WithName("k8s"),
		RestConfig: restConfig,
	}
	return op, nil
}

type Operator struct {
	Logger     log.Logger
	RestConfig *rest.Config
}

// Type implements operator.Operator.
func (Operator) Type() operator.Type {
	return OperatorType
}

// IsConnected implements operator.Operator.
func (op Operator) IsConnected(ctx context.Context) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	var err = k8s.Wait(ctx, op.RestConfig)
	return err == nil, err
}
