package main

import (
	"HorizonX/rpc/mqueue/worker/internal/config"
	"HorizonX/rpc/mqueue/worker/internal/logic"
	"HorizonX/rpc/mqueue/worker/internal/svc"
	"context"
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
)

var configFile = flag.String("f", "etc/mqueue.yaml", "Specify the config file")

func main() {
	flag.Parse()
	var c config.Config

	conf.MustLoad(*configFile, &c, conf.UseEnv())

	// log、prometheus、trace、metricsUrl
	if err := c.SetUp(); err != nil {
		panic(err)
	}

	//logx.DisableStat()

	svcContext := svc.NewServiceContext(c)
	ctx := context.Background()
	mqWorker := logic.NewMqWorker(ctx, svcContext)
	mux := mqWorker.Register()

	if err := svcContext.AsynqServer.Run(mux); err != nil {
		logx.WithContext(ctx).Errorf("!!!MqWorkerErr!!! run err:%+v", err)
		os.Exit(1)
	}
}
