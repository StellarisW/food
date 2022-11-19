package recipe

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	g "main/app/global"
	"main/app/internal/model"
	"strings"
	"time"
)

type sInfo struct{}

var insInfo = sInfo{}

func (s *sInfo) GetRecipeById(ctx context.Context, recipeId int64) *model.Recipe {
	filter := bson.D{
		{
			"recipe_id",
			recipeId,
		},
	}

	var elem model.Recipe

	cur := g.MongoDB.Database("food").Collection("recipe").
		FindOne(ctx, filter)

	err := cur.Decode(&elem)
	if err != nil {
		return nil
	}

	return &elem
}

func (s *sInfo) GetTimeDuration(timeStr string) (time.Duration, time.Duration) {
	if timeStr == "" {
		return 0, 0
	}
	output := strings.Split(timeStr, "-")
	if len(output) != 2 {
		return 0, 0
	}
	from, _ := time.ParseDuration(output[0])
	to, _ := time.ParseDuration(output[1])
	return from, to
}
