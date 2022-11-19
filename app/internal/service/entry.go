package service

import (
	"main/app/internal/service/recipe"
	"main/app/internal/service/restaurant"
	"main/app/internal/service/user"
)

var insUser = user.Group{}

func User() *user.Group {
	return &insUser
}

var insRestaurant = restaurant.Group{}

func Restaurant() *restaurant.Group {
	return &insRestaurant
}

var insRecipe = recipe.Group{}

func Recipe() *recipe.Group {
	return &insRecipe
}
