package registry

import (
	v1 "github.com/solo-io/gloo/projects/controller/pkg/api/v1"
	"github.com/solo-io/gloo/projects/controller/pkg/bootstrap"
	"github.com/solo-io/gloo/projects/controller/pkg/plugins"
	"github.com/solo-io/gloo/projects/discovery/pkg/fds"
	"github.com/solo-io/gloo/projects/discovery/pkg/fds/discoveries/aws"
	"github.com/solo-io/gloo/projects/discovery/pkg/fds/discoveries/grpc"
	"github.com/solo-io/gloo/projects/discovery/pkg/fds/discoveries/swagger"
)

type FunctionDiscoveryPlugin func(u *v1.Upstream) fds.UpstreamFunctionDiscovery

type registry struct {
	plugins []fds.FunctionDiscoveryFactory
}

var globalRegistry = func(opts bootstrap.Opts, pluginExtensions ...func() plugins.Plugin) *registry {
	reg := &registry{}
	// plugins should be added here
	reg.plugins = append(reg.plugins,
		aws.NewFunctionDiscoveryFactory(),
		grpc.NewFunctionDiscoveryFactory(),
		swagger.NewFunctionDiscoveryFactory(),
	)

	return reg
}

func Plugins(opts bootstrap.Opts) []fds.FunctionDiscoveryFactory {
	return globalRegistry(opts).plugins
}
