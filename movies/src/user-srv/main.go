package main

import (
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"go.uber.org/zap"
	"hongbao/movies/src/share/config"
	"hongbao/movies/src/share/pb"
	"hongbao/movies/src/share/utils/log"
	"hongbao/movies/src/user-srv/db"
	"hongbao/movies/src/user-srv/handler"
)
func main()  {

	log.Init("user")
	logger := log.Instance()
	//var op = micro.Option()
	// 我这里用的etcd 做为服务发现，如果使用consul可以去掉
	//reg := etcdv3.NewRegistry(func(options *registry.Options) {
	//	options.Addrs = []string{
	//		"http://127.0.0.1:2379",
	//	}
	//
	//})


	service := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
		//micro.Registry(reg),
		)
	// 定义Service动作操作
	fmt.Println("kaiqi")
	service.Init(
		micro.Action(func(context *cli.Context) {
			logger.Info("Info", zap.Any("user-srv", "user-srv is start ..."))
			db.Init(config.MysqlDSN)
			pb.RegisterUserServiceExtHandler(service.Server(),handler.NewUserServiceExtHandler(),server.InternalHandler(true))

		}),
		micro.AfterStop(func() error {
			logger.Info("Info", zap.Any("user-srv", "user-srv is stop ..."))
			return nil
		}),
		micro.AfterStart(func() error {
			return nil
		}),

		)
	//启动service
	if err := service.Run(); err != nil {
		logger.Panic("user-srv服务启动失败 ...")
	}





}