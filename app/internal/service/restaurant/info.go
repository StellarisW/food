package restaurant

import (
	"encoding/json"
	"fmt"
	"github.com/imroc/req/v3"
	"github.com/tidwall/gjson"
	g "main/app/global"
	"main/app/internal/model"
	"net/http"
)

type sInfo struct{}

var insInfo = sInfo{}

const (
	yelpSearchApi = "https://api.yelp.com/v3/businesses/search"
	yelpIdApi     = "https://api.yelp.com/v3/businesses/"
)

func (s *sInfo) SearchRestaurantFromYelp(latitude, longitude, term string) (string, error) {
	res, err := req.SetHeader("Authorization", fmt.Sprintf("Bearer %s", g.Config.YelpApiKey)).
		SetQueryParams(map[string]string{
			"latitude":  latitude,
			"longitude": longitude,
			"term":      term,
		}).Get(yelpSearchApi)
	if err != nil {
		g.Logger.Errorf("query yelp api failed, err: %v", err)
		return "", fmt.Errorf("internal err")
	}
	if res.StatusCode != http.StatusOK {
		g.Logger.Errorf("query yelp api failed, err: %v", res)
		return "", fmt.Errorf("internal err")
	}

	return res.String(), nil
}

func (s *sInfo) UnmarshalRestaurantsData(src string, dst *[]*model.Restaurant) error {
	resJson := gjson.Parse(src)
	resArray := resJson.Get("businesses").Array()

	var dstData []*model.Restaurant

	for _, restaurantRes := range resArray {
		restaurant := &model.Restaurant{}
		err := json.Unmarshal([]byte(restaurantRes.Raw), restaurant)
		if err != nil {
			g.Logger.Errorf("unmarshal restaurant json failed, err: %v", err)
			return fmt.Errorf("internal err")
		}
		dstData = append(dstData, restaurant)
	}

	*dst = dstData

	return nil
}

func (s *sInfo) GetRestaurantByID(id string) (string, error) {
	res, err := req.SetHeader("Authorization", fmt.Sprintf("Bearer %s", g.Config.YelpApiKey)).
		Get(fmt.Sprintf("%s%s", yelpIdApi, id))
	if err != nil {
		g.Logger.Errorf("query yelp api failed, err: %v", err)
		return "", fmt.Errorf("internal err")
	}
	if res.StatusCode != http.StatusOK {
		g.Logger.Errorf("query yelp api failed, err: %v", res)
		return "", fmt.Errorf("internal err")
	}

	return res.String(), nil
}
