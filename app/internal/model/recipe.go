package model

type Recipe struct {
	Id           string   `bson:"_id"`
	RecipeId     int64    `bson:"recipe_id"`
	Images       []string `bson:"images"`
	Name         string   `bson:"name"`
	Category     string   `bson:"category"`
	Dietary      []string `bson:"dietary"`
	Description  string   `bson:"description"`
	Keywords     []string `bson:"keywords"`
	Instruction  []string `bson:"instruction"`
	Ingredients  []string `bson:"ingredients"`
	CookTime     int64    `bson:"cook_time"`
	PerpTime     int64    `bson:"perp_time"`
	TotalTime    int64    `bson:"total_time"`
	Calories     float64  `bson:"calories"`
	Fat          float64  `bson:"fat"`
	SaturatedFat float64  `bson:"saturated_fat"`
	Sodium       float64  `bson:"sodium"`
	Carbohydrate float64  `bson:"carbohydrate"`
	Fiber        float64  `bson:"fiber"`
	Sugar        float64  `bson:"sugar"`
	Protein      float64  `bson:"protein"`
}
