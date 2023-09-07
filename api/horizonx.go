package main

import (
	"HorizonX/common/result"
	"flag"
	"fmt"
	"net/http"

	"HorizonX/api/internal/config"
	"HorizonX/api/internal/handler"
	"HorizonX/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/horizonx-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf,
		rest.WithUnauthorizedCallback(
			func(w http.ResponseWriter, r *http.Request, err error) {
				result.AuthHttpResult(r, w, nil, err)
			}))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
