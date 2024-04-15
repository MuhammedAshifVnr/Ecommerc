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

type UserData struct {
	Name   string `json:"username"`
	Mobile int    `json:"mobile"`
	Gender string `json:"gender"`
}

type AddressData struct {
	Type    string `jsonz:"type"`
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode uint   `json:"zip"`
}

type ReviewData struct {
	Rating  float64 `json:"rating"`
	Comment string  `json:"review"`
}
