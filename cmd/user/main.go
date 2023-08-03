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
	"github.com/linzijie1998/mini-tiktok/pkg/bound"
	"github.com/linzijie1998/mini-tiktok/pkg/middleware"
	"github.com/linzijie1998/mini-tiktok/pkg/path"
	"log"
	"net"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func LoadConfigsAndInit() {
	var err error
	if exist, err := path.FileExist("config.yaml"); err != nil || !exist {
		log.Fatal("未找到配置文件，无法启动服务")
	}
	if global.Viper, err = initialize.Viper("config.yaml"); err != nil {
		log.Fatal(err)
	}
	if global.GormDB, err = initialize.GormMySQL(); err != nil {
		log.Fatal(err)
	}
	if global.RedisClient, err = initialize.Redis(); err != nil {
		log.Fatal(err)
	}
	if global.MongoClient, err = initialize.Mongo(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// 加载配置文件并且初始化
	LoadConfigsAndInit()
	// etcd 服务注册
	r, err := etcd.NewEtcdRegistry([]string{global.Configs.ETCD.Addr()})
	if err != nil {
		log.Fatal(err)
	}
	// 微服务地址
	addr, err := net.ResolveTCPAddr("tcp", global.Configs.RPCServer.Addr())
	if err != nil {
		log.Fatal(err)
	}
	// 微服务定义
	svr := user.NewServer(new(UserServiceImpl),
		// 服务名称
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: global.Configs.RPCServer.ServiceName,
		}),
		// 服务地址
		server.WithServiceAddr(addr),
		// 中间件
		server.WithMiddleware(middleware.CommonMiddleware),
		server.WithMiddleware(middleware.ServerMiddleware),
		// 连接限制
		server.WithLimit(&limit.Option{
			MaxConnections: 1000,
			MaxQPS:         100,
		}),
		// BoundHandler
		server.WithBoundHandler(bound.NewCpuLimitHandler()),
		// 多路复用?
		server.WithMuxTransport(),
		// 链路追踪
		server.WithSuite(trace.NewDefaultServerSuite()),
		// 服务注册
		server.WithRegistry(r),
	)
	// 启动服务
	if err = svr.Run(); err != nil {
		log.Fatal(err)
	}
}
