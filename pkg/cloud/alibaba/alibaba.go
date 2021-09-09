package alibaba

import (
	"embed"

	"github.com/openshift/cluster-cloud-controller-manager-operator/pkg/cloud/common"
	appsv1 "k8s.io/api/apps/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	//go:embed assets/*
	aliFS embed.FS

	aliResources []client.Object

	aliSources = []common.ObjectSource{
		{Object: &appsv1.Deployment{}, Path: "assets/cloud-controller-manager-deployment.yaml"},
	}
)

func init() {
	var err error
	aliResources, err = common.ReadResources(aliFS, aliSources)
	utilruntime.Must(err)
}

// GetResources returns a list of Alibaba Cloud resources for provisioning CCM in running cluster
func GetResources() []client.Object {
	resources := make([]client.Object, len(aliResources))
	for i := range aliResources {
		resources[i] = aliResources[i].DeepCopyObject().(client.Object)
	}

	return resources
}
