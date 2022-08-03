package nacos

import (
	"context"
	"github.com/asim/go-micro/v3/registry"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
)

type AddressKey struct{}
type ConfigKey struct{}

// WithAddress sets the nacos address
func WithAddress(addrs []string) registry.Option {
	return func(o *registry.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, AddressKey{}, addrs)
	}
}

// WithClientConfig sets the nacos config
func WithClientConfig(cc constant.ClientConfig) registry.Option {
	return func(o *registry.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, ConfigKey{}, cc)
	}
}
