package main

import (
	"context"
	"fmt"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/registry"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"my-micro/registry/nacos"
)

const (
	ServerName = "go.micro.srv.GetArea" // server name
)

func main() {
	addrs := make([]string, 1)
	addrs[0] = "nacos.topwhere.cn:8848"

	reg := nacos.NewRegistry(func(options *registry.Options) {
		if options.Context == nil {
			options.Context = context.Background()
		}
		//// 支持 namespace
		options.Context = context.WithValue(options.Context, nacos.AddressKey{}, addrs)
		options.Context = context.WithValue(options.Context, nacos.ConfigKey{}, GetNacosClientConfig())
	})

	// Create service
	service := micro.NewService(
		micro.Name(ServerName),
		micro.Version("latest"),
		micro.Registry(reg),
	)
	// Initialise service
	service.Init()

	// Run service
	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}

func GetNacosClientConfig() constant.ClientConfig {
	clientConfig := constant.ClientConfig{
		NamespaceId:         "f509d58f-275b-4d60-b9f7-37369a7a608c",
		RegionId:            "go-oss",
		TimeoutMs:           3000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}
	fmt.Println("定义客户端配置", clientConfig)
	return clientConfig
}
