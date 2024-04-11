package database

import "time"

type Logging struct {
	Username string `json:"useremail"`
	Password string `json:"password"`
}

type CategoryData struct {
	Name        string `gorm:"unique;not null" json:"catagory"`
	Description string `gorm:"not null" json:"description"`
}

type OfferProductData struct {
	Percentage float64   `json:"percentage"`
	Expirey    time.Time `json:"expirey"`
}
