package router

import (
	"github.com/gin-gonic/gin"
	g "main/app/global"
	"main/app/internal/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.ZapLogger(g.Logger), middleware.ZapRecovery(g.Logger, true))
	r.Use(middleware.CorsByRules())

	routerGroup := new(Group)

	PublicGroup := r.Group("/api")
	{
		routerGroup.InitUserSignRouter(PublicGroup)
	}

	PrivateGroup := r.Group("/api")
	PrivateGroup.Use(middleware.JWTAuthMiddleware())
	{
		routerGroup.InitRecipeRouter(PrivateGroup)
		routerGroup.InitRestaurantRouter(PrivateGroup)
		routerGroup.InitUserInfoRouter(PrivateGroup)
	}

	g.Logger.Infof("initialize routers successfully")
	return r
}
