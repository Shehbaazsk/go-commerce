package models

type Address struct {
	BaseModel
	UserID  uint   `gorm:"index" json:"user_id"`
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
	Pincode string `json:"pincode"`
	Type    string `json:"type"`
}
