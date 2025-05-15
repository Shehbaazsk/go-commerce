package models

type Role struct {
	BaseModel
	Name  string `gorm:"unique;not null" json:"name"`
	Users []User `gorm:"foreignKey:RoleID" json:"users"`
}
