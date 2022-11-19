package router

import (
	"github.com/gin-gonic/gin"
	"main/app/api"
)

type RecipeRouter struct{}

func (r *RecipeRouter) InitRecipeRouter(router *gin.RouterGroup) (R gin.IRoutes) {
	recipeRouter := router.Group("/recipe")
	recipeApi := api.Recipe()
	{
		recipeRouter.GET("", recipeApi.Recipe().Search)
	}

	return recipeRouter
}
