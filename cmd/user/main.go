package main

import (
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"github.com/linzijie1998/mini-tiktok/cmd/user/global"
	"github.com/linzijie1998/mini-tiktok/cmd/user/initialize"
	user "github.com/linzijie1998/mini-tiktok/kitex_gen/douyin/user/userservice"
	"log"
	"net"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func LoadConfigsAndInit() {
	var err error
	if global.Viper, err = initialize.Viper("/home/nahida/devgo/src/mini-tiktok/cmd/user/config.yaml"); err != nil {
		panic(err)
	}
	if global.GormDB, err = initialize.GormMySQL(); err != nil {
		panic(err)
	}
	if global.RedisClient, err = initialize.Redis(); err != nil {
		panic(err)
	}
}

func main() {
	LoadConfigsAndInit()

	r, err := etcd.NewEtcdRegistry([]string{global.Configs.ETCD.Addr()})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", global.Configs.RPCServer.Addr())
	if err != nil {
		panic(err)
	}

	svr := user.NewServer(new(UserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: global.Configs.RPCServer.ServiceName,
		}),
		server.WithServiceAddr(addr),
		server.WithLimit(&limit.Option{
			MaxConnections: 1000,
			MaxQPS:         100,
		}),
		server.WithMuxTransport(),
		server.WithSuite(trace.NewDefaultServerSuite()),
		server.WithRegistry(r),
	)

	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
