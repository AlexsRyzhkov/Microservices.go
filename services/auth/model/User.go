package model

type User struct {
	ID   uint   `gorm:"primary_key"`
	Name string `gorm:"uniqueIndex;column:name"`
	Psw  string `gorm:"column:psw"`
}
