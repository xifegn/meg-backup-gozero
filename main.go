package main

import (
	"flag"
	"fmt"
	middleware "github.com/muhfajar/go-zero-cors-middleware"
	"meg-backup-gozero/internal/config"
	"meg-backup-gozero/internal/handler"
	"meg-backup-gozero/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/main.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	// Register go-zero-cors-middleware handler to handle preflight request
	cors := middleware.NewCORSMiddleware(&middleware.Options{})

	// Add run option WithNotAllowedHandler and register `.Handler()` to handle `OPTIONS` request (preflight)
	server := rest.MustNewServer(c.RestConf,
		rest.WithNotAllowedHandler(cors.Handler()),
	)
	defer server.Stop()
	handler.RegisterHandlers(server, ctx)
	// Register go-zero-cors-middleware
	server.Use(cors.Handle)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
