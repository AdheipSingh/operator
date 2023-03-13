package parseabletenantcontroller

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Builder struct {
	ConfigMaps              []BuilderConfigMap
	DeploymentOrStatefulset []BuilderDeploymentStatefulSet
}

type CommonBuilder struct {
	ObjectMeta   metav1.ObjectMeta
	Client       client.Client
	OwnerRef     metav1.OwnerReference
	CrObject     client.Object
	DesiredState client.Object
	CurrentState client.Object
}

type ToBuilder func(opts *Builder)

func NewBuilder(opts ...ToBuilder) *Builder {
	builder := &Builder{}
	for _, opt := range opts {
		opt(builder)
	}
	return builder
}
