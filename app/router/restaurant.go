package router

import (
	"github.com/gin-gonic/gin"
	"main/app/api"
)

type RestaurantRouter struct{}

func (r *RestaurantRouter) InitRestaurantRouter(router *gin.RouterGroup) (R gin.IRoutes) {
	restaurantRouter := router.Group("/restaurant")
	restaurantApi := api.Restaurant()
	{
		restaurantRouter.GET("", restaurantApi.Restaurant().Search)
	}

	return restaurantRouter
}
