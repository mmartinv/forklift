package ocp

import (
	api "github.com/konveyor/forklift-controller/pkg/apis/forklift/v1beta1"
	"github.com/konveyor/forklift-controller/pkg/controller/plan/adapter/base"
	plancontext "github.com/konveyor/forklift-controller/pkg/controller/plan/context"
	k8sclient "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

// Openshift adapter.
type Adapter struct{}

// Constructs a openstack builder.
func (r *Adapter) Builder(ctx *plancontext.Context) (builder base.Builder, err error) {
	builder = &Builder{Context: ctx}
	return
}

// Constructs a openshift validator.
func (r *Adapter) Validator(plan *api.Plan) (validator base.Validator, err error) {
	conf, err := config.GetConfig()
	if err != nil {
		return
	}

	client, err := k8sclient.New(conf, k8sclient.Options{})
	if err != nil {
		return
	}

	v := &Validator{plan: plan, client: client}
	return v, nil
}

// Constructs an openshift client.
func (r *Adapter) Client(ctx *plancontext.Context) (client base.Client, err error) {
	return Client{Context: ctx}, nil
}

// Constucts a destination client.
func (r *Adapter) DestinationClient(ctx *plancontext.Context) (destinationClient base.DestinationClient, err error) {
	destinationClient = &DestinsationClient{Context: ctx}
	return
}