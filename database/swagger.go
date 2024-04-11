package database

type Logging struct {
	Username string `json:"useremail"`
	Password string `json:"password"`
}

type CategoryData struct {
	Name        string `gorm:"unique;not null" json:"catagory"`
	Description string `gorm:"not null" json:"description"`
}
