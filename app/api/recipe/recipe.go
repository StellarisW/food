package recipe

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	g "main/app/global"
	"main/app/internal/model"
	"main/app/internal/service"
	"net/http"
)

type Api struct{}

var insRecipe = Api{}

func (a *Api) Search(c *gin.Context) {
	dietary := c.Query("dietary")
	cookTimeString := c.Query("cook_time")
	perpTimeString := c.Query("perp_time")
	totalTimeString := c.Query("total_time")
	taste := c.QueryArray("taste")
	ingredients := c.QueryArray("ingredients")

	if dietary != "" {
		if dietary != "halal" && dietary != "vegan" && dietary != "vegetarian" {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  "invalid dietary",
				"ok":   false,
			})
			return
		}
	}

	cookBeginTime, cookEndTime := service.Recipe().Info().GetTimeDuration(cookTimeString)
	perpBeginTime, perpEndTime := service.Recipe().Info().GetTimeDuration(perpTimeString)
	totalBeginTime, totalEndTime := service.Recipe().Info().GetTimeDuration(totalTimeString)

	limit := cast.ToInt64(c.Query("limit"))
	page := cast.ToInt64(c.Query("page"))

	if limit <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  `invalid param "limit"`,
			"ok":   false,
		})
		return
	}
	if page <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  `invalid param "page"`,
			"ok":   false,
		})
		return
	}

	database := g.MongoDB.Database("food")

	collection := database.Collection("recipe")

	filter := bson.D{}

	if dietary != "" {
		if dietary == "halal" {
			filter = append(filter, bson.E{
				Key: "dietary",
				Value: bson.D{
					{
						"$not",
						bson.D{
							{
								"$in",
								[]string{"non-halal"},
							},
						},
					},
				},
			})
		} else if dietary == "vegetarian" {
			filter = append(filter, bson.E{
				Key: "dietary",
				Value: bson.D{
					{
						"$nin",
						[]string{"non-vegetarian"},
					},
				},
			})
		} else if dietary == "vegan" {
			filter = append(filter, bson.E{
				Key: "dietary",
				Value: bson.D{
					{
						"$nin",
						[]string{"non-vegan"},
					},
				},
			})
		}
	}

	if cookEndTime != 0 {
		filter = append(filter, bson.E{
			Key: "cook_time",
			Value: bson.D{
				{
					"$gte",
					cookBeginTime.Seconds(),
				},
				{
					"$lte",
					cookEndTime.Seconds(),
				},
			},
		})
	}

	if perpEndTime != 0 {
		filter = append(filter, bson.E{
			Key: "perp_time",
			Value: bson.D{
				{
					"$gte",
					perpBeginTime.Seconds(),
				},
				{
					"$lte",
					perpEndTime.Seconds(),
				},
			},
		})
	}

	if totalEndTime != 0 {
		filter = append(filter, bson.E{
			Key: "total_time",
			Value: bson.D{
				{
					"$gte",
					totalBeginTime.Seconds(),
				},
				{
					"$lte",
					totalEndTime.Seconds(),
				},
			},
		})
	}

	for _, ingredient := range ingredients {
		reg := bson.E{
			Key: "ingredients",
			Value: primitive.Regex{
				Pattern: ingredient,
				Options: "i",
			},
		}
		filter = append(filter, reg)
	}

	for _, tas := range taste {
		reg := bson.E{
			Key: "keywords",
			Value: primitive.Regex{
				Pattern: tas,
				Options: "i",
			},
		}
		filter = append(filter, reg)
	}

	g.Logger.Debugf("%v", filter)

	option := &options.FindOptions{}
	option.SetLimit(limit)
	option.SetSkip(limit * (page - 1))
	cur, err := collection.Find(c, filter, option)
	if err != nil {
		g.Logger.Errorf("find [recipe] document failed, err: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "internal err",
			"ok":   false,
		})
		return
	}
	defer cur.Close(c)

	var results []*model.Recipe

	for cur.Next(c) {
		var elem model.Recipe

		err := cur.Decode(&elem)
		if err != nil {
			g.Logger.Errorf("decode [recipe] document failed, err: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  "internal err",
				"ok":   false,
			})
			return
		}

		results = append(results, &elem)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "get recipe successfully",
		"ok":   true,
		"data": results,
	})
}
