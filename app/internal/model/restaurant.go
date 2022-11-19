package model

type Restaurant struct {
	Id           string      `json:"id"`
	Name         string      `json:"name"`
	Alias        string      `json:"alias"`
	ImageUrl     string      `json:"image_url"`
	Url          string      `json:"url"`
	IsClose      bool        `json:"is_close"`
	ReviewCount  int32       `json:"reviewCount"`
	Categories   []Category  `json:"categories"`
	Rating       float64     `json:"rating"`
	Coordinates  Coordinates `json:"coordinates"`
	Transactions []string    `json:"transactions"`
	Price        string      `json:"price"`
	Location     Location    `json:"location"`
	Phone        string      `json:"phone"`
	DisplayPhone string      `json:"display_phone"`
	Distance     float64     `json:"distance"`
}

type Category struct {
	Title string `json:"title"`
	Alias string `json:"alias"`
}

type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Location struct {
	Address1       string   `json:"address1"`
	Address2       string   `json:"address2"`
	Address3       string   `json:"address3"`
	City           string   `json:"city"`
	ZipCode        string   `json:"zip_code"`
	Country        string   `json:"country"`
	State          string   `json:"state"`
	DisplayAddress []string `json:"display_address"`
}
