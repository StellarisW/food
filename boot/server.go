package boot

import (
	"github.com/gin-gonic/gin"
	g "main/app/global"
	"main/app/router"
	"net/http"
)

func ServerSetup() {
	config := g.Config.Server

	gin.SetMode(config.Mode)
	routers := router.InitRouter()
	server := &http.Server{
		Addr:              config.Addr(),
		Handler:           routers,
		TLSConfig:         nil,
		ReadTimeout:       config.GetReadTimeout(),
		ReadHeaderTimeout: 0,
		WriteTimeout:      config.GetWriteTimeout(),
		IdleTimeout:       0,
		MaxHeaderBytes:    1 << 20, // 16mb
	}

	g.Logger.Infof("server running on %s ...", config.Addr())
	g.Logger.Errorf(server.ListenAndServe().Error())
}
