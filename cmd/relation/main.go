package main

import (
	"fmt"
	"github.com/linzijie1998/mini-tiktok/pkg/path"
	"log"
	"net"
	"os"

	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"github.com/linzijie1998/mini-tiktok/cmd/relation/global"
	"github.com/linzijie1998/mini-tiktok/cmd/relation/initialize"
	relation "github.com/linzijie1998/mini-tiktok/kitex_gen/douyin/relation/relationservice"
)

func LoadConfigsAndInit() {
	var err error
	if exist, err := path.FileExist("/home/nahida/devgo/src/mini-tiktok/cmd/relation/config.yaml"); err != nil || !exist {
		fmt.Println("未找到配置文件，无法启动服务")
		os.Exit(0)
	}
	if global.Viper, err = initialize.Viper("/home/nahida/devgo/src/mini-tiktok/cmd/relation/config.yaml"); err != nil {
		panic(err)
	}
	if global.GormDB, err = initialize.GormMySQL(); err != nil {
		panic(err)
	}
	if global.RedisClient, err = initialize.Redis(); err != nil {
		panic(err)
	}
	if global.MongoClient, err = initialize.Mongo(); err != nil {
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

	svr := relation.NewServer(new(RelationServiceImpl),
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
