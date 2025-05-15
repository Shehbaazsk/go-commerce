package models

type User struct {
	BaseModel
	Email       string      `gorm:"unique;not null" json:"email"`
	Password    string      `gorm:"not null" json:"-"`
	RoleID      uint        `json:"role_id"`
	Role        Role        `gorm:"foreignKey:RoleID" json:"role"`
	UserProfile UserProfile `gorm:"foreignKey:UserID" json:"profile"`
}
