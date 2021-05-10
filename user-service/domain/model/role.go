package model

type Role struct {
	Id string `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"unique" json:"name"`

}