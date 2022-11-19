package user

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	g "main/app/global"
	"main/app/internal/model"
	"main/app/internal/service"
	"net/http"
)

type CollectApi struct{}

var insCollect = CollectApi{}

func (a *CollectApi) GetList(c *gin.Context) {
	userId := c.GetInt64("id")

	collectType := cast.ToInt32(c.Query("collect_type"))
	limit := cast.ToInt32(c.Query("limit"))
	page := cast.ToInt32(c.Query("page"))

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

	var userCollections []*model.UserCollection
	switch collectType {
	case 1, 2:
		cnt, err := service.User().Collect().GetUserCollectionCount(c, userId, collectType)
		if cnt == -1 || err != nil {
			switch err.Error() {
			case "internal err":
				c.JSON(http.StatusInternalServerError, gin.H{
					"code": http.StatusInternalServerError,
					"msg":  "internal err",
					"ok":   false,
				})
			}

			return
		}

		pageCount := int32(cnt) / limit
		if int32(cnt)%limit > 0 {
			pageCount = pageCount + 1
		}

		if page > pageCount {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  fmt.Sprintf("the maximum number of pages is %d", pageCount),
			})
			return
		}

		userCollections, err = service.User().Collect().GetUserCollectionsWithLimit(c, userId, collectType, int(limit), int(page))
		if err != nil {
			switch err.Error() {
			case "internal err":
				c.JSON(http.StatusInternalServerError, gin.H{
					"code": http.StatusInternalServerError,
					"msg":  "internal err",
					"ok":   false,
				})
			}

			return
		}

	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "invalid collect type",
			"ok":   false,
		})
		return
	}

	switch collectType {
	case 1:
		var collections []*model.Collection
		for _, userCollection := range userCollections {
			res, err := service.Restaurant().Info().GetRestaurantByID(userCollection.RestaurantId)
			if err != nil {
				switch err.Error() {
				case "internal err":
					c.JSON(http.StatusInternalServerError, gin.H{
						"code": http.StatusInternalServerError,
						"msg":  "internal err",
						"ok":   false,
					})
				}

				return
			}

			restaurant := &model.Restaurant{}
			err = json.Unmarshal([]byte(res), restaurant)
			if err != nil {
				g.Logger.Errorf("unmarshal restaurant json failed, err: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"code": http.StatusInternalServerError,
					"msg":  "internal err",
					"ok":   false,
				})
			}
			collection := &model.Collection{
				Id:             userCollection.Id,
				CollectionType: "restaurant",
				CollectionData: restaurant,
			}
			collections = append(collections, collection)
		}

		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "get collection successfully",
			"ok":   true,
			"data": collections,
		})
	case 2:
		var collections []*model.Collection
		for _, userCollection := range userCollections {
			recipe := service.Recipe().Info().GetRecipeById(c, userCollection.RecipeId)
			if recipe != nil {
				collection := &model.Collection{
					Id:             userCollection.Id,
					CollectionType: "recipe",
					CollectionData: recipe,
				}
				collections = append(collections, collection)
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "get collection successfully",
			"ok":   true,
			"data": collections,
		})
	}
}

func (a *CollectApi) Create(c *gin.Context) {
	userId := c.GetInt64("id")

	collectType := cast.ToInt32(c.PostForm("collect_type"))

	userCollection := &model.UserCollection{
		UserId:      userId,
		CollectType: collectType,
	}

	var id interface{}

	switch collectType {
	case 1:
		// collect restaurant
		restaurantId := c.PostForm("restaurant_id")
		if restaurantId == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  "restaurant_id cannot be null",
				"ok":   false,
			})
			return
		}

		id = restaurantId
		userCollection.RestaurantId = restaurantId

	case 2:
		// collect recipe
		recipeId := cast.ToInt64(c.PostForm("recipe_id"))
		if recipeId == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  "recipe_id cannot be null",
				"ok":   false,
			})
			return
		}

		id = recipeId
		userCollection.RecipeId = recipeId

	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "invalid collect type",
			"ok":   false,
		})
		return
	}

	err := service.User().Collect().CheckCollectionIsExist(c, collectType, userId, id)
	if err != nil {
		switch err.Error() {
		case "internal err":
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  "internal err",
				"ok":   false,
			})

		case "duplicate collect":
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  err.Error(),
				"ok":   false,
			})
		}

		return
	}

	service.User().Collect().CreateCollection(c, userCollection)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "create collection successfully",
		"ok":   true,
	})
}

func (a *CollectApi) Delete(c *gin.Context) {
	id := cast.ToInt64(c.Query("id"))
	userId := c.GetInt64("id")

	if id == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "id cannot be null",
			"ok":   false,
		})
		return
	}

	err := service.User().Collect().CheckCollectionIdIsExist(c, id, userId)
	if err != nil {
		switch err.Error() {
		case "internal err":
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  "internal err",
				"ok":   false,
			})
		case "collection not found":
			c.JSON(http.StatusNotFound, gin.H{
				"code": http.StatusNotFound,
				"msg":  "collection not found",
				"ok":   false,
			})
		}

		return
	}

	err = service.User().Collect().DeleteCollection(c, id)
	if err != nil {
		switch err.Error() {
		case "internal err":
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  "internal err",
				"ok":   false,
			})
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "delete collection successfully",
		"ok":   true,
	})
}
