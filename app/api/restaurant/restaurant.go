package restaurant

import (
	"github.com/gin-gonic/gin"
	"main/app/internal/model"
	"main/app/internal/service"
	"net/http"
)

type Api struct{}

var insRestaurant = Api{}

func (a *Api) Search(c *gin.Context) {
	latitude := c.Query("latitude")
	longitude := c.Query("longitude")
	term := c.Query("term")

	res, err := service.Restaurant().Info().SearchRestaurantFromYelp(latitude, longitude, term)
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

	restaurants := new([]*model.Restaurant)
	err = service.Restaurant().Info().UnmarshalRestaurantsData(res, restaurants)
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

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "get restaurant successfully",
		"ok":   true,
		"data": restaurants,
	})
}
