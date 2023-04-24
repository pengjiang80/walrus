package platformk8s

import (
	"context"
	"fmt"

	kerrors "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	dynamicclient "k8s.io/client-go/dynamic"

	"github.com/seal-io/seal/pkg/dao/model"
	"github.com/seal-io/seal/pkg/dao/types"
	"github.com/seal-io/seal/pkg/dao/types/status"
	"github.com/seal-io/seal/pkg/platformk8s/helm"
	"github.com/seal-io/seal/pkg/platformk8s/intercept"
	"github.com/seal-io/seal/pkg/platformk8s/kube"
	"github.com/seal-io/seal/pkg/platformk8s/kubestatus"
)

// GetStatus implements operator.Operator.
func (op Operator) GetStatus(ctx context.Context, res *model.ApplicationResource) (*status.Status, error) {
	if res == nil {
		return kubestatus.StatusError(""), nil
	}

	if res.DeployerType != types.DeployerTypeTF {
		op.Logger.Warn("error resource stating: unknown deployer type: " + res.DeployerType)
		return kubestatus.StatusError("unknown deployer type"), nil
	}

	if res.Type == "helm_release" {
		var opts = helm.GetReleaseOptions{
			RESTClientGetter: helm.IncompleteRestClientGetter(*op.RestConfig),
			Log:              op.Logger.Debugf,
		}

		return helm.GetReleaseStatus(ctx, res, opts)
	}

	var gvr, ok = intercept.Terraform().GetGVR(res.Type)
	if !ok {
		// mark ready if it's unresolved type.
		return &kubestatus.GeneralStatusReady, nil
	}
	var ns, n = kube.ParseNamespacedName(res.Name)

	// fetch label selector with dynamic client.
	dynamicCli, err := dynamicclient.NewForConfig(op.RestConfig)
	if err != nil {
		err = fmt.Errorf("error creating kubernetes dynamic client: %w", err)
		return kubestatus.StatusError(err.Error()), err
	}
	o, err := dynamicCli.Resource(gvr).Namespace(ns).
		Get(ctx, n, meta.GetOptions{ResourceVersion: "0"}) // non quorum read
	if err != nil {
		if !kerrors.IsNotFound(err) {
			err = fmt.Errorf("error getting kubernetes %s %s/%s: %w", gvr.Resource, ns, n, err)
			return kubestatus.StatusError(err.Error()), err
		}
		// mark unknown if not found.
		return kubestatus.StatusError("resource not found"), nil
	}

	os, err := kubestatus.Get(ctx, dynamicCli, o)
	if err != nil {
		err = fmt.Errorf("error stating status of kubernetes %s %s/%s: %w", gvr.Resource, ns, n, err)
		return kubestatus.StatusError(err.Error()), err

	}
	return os, nil
}
