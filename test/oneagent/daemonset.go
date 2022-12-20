package oneagent

import (
	"context"
	"testing"

	"github.com/Dynatrace/dynatrace-operator/test/kubeobjects/daemonset"
	appsv1 "k8s.io/api/apps/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/e2e-framework/klient/k8s/resources"
	"sigs.k8s.io/e2e-framework/pkg/envconf"
	"sigs.k8s.io/e2e-framework/pkg/features"
)

const (
	name      = "dynakube-oneagent"
	namespace = "dynatrace"
)

func WaitForDaemonset() features.Func {
	return daemonset.WaitFor(name, namespace)
}

func WaitForDaemonSetPodsDeletion() func(ctx context.Context, environmentConfig *envconf.Config, t *testing.T) (context.Context, error) {
	return daemonset.WaitForPodsDeletion(name, namespace)
}

func Get(ctx context.Context, resource *resources.Resources) (appsv1.DaemonSet, error) {
	return daemonset.NewQuery(ctx, resource, client.ObjectKey{
		Name:      name,
		Namespace: namespace,
	}).Get()
}

func ForEachPod(ctx context.Context, resource *resources.Resources, consumer daemonset.PodConsumer) error {
	return daemonset.NewQuery(ctx, resource, client.ObjectKey{
		Name:      name,
		Namespace: namespace,
	}).ForEachPod(consumer)
}
