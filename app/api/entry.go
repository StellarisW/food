package api

import (
	"main/app/api/recipe"
	"main/app/api/restaurant"
	"main/app/api/user"
)

var insUser = user.Group{}

func User() *user.Group {
	return &insUser
}

var insRecipe = recipe.Group{}

func Recipe() *recipe.Group {
	return &insRecipe
}

var insRestaurant = restaurant.Group{}

func Restaurant() *restaurant.Group {
	return &insRestaurant
}
