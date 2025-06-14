package model

type User struct {
	BaseModel
	Email       string `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	Password    string `gorm:"type:varchar(255);not null" json:"-"`
	FirstName   string `gorm:"type:varchar(100)" json:"first_name,omitempty"`
	LastName    string `gorm:"type:varchar(100)" json:"last_name,omitempty"`
	MobilePhone string `gorm:"type:varchar(20);index" json:"mobile_phone,omitempty"`
}
