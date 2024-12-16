package model

type Cars struct {
	ID    uint64 `gorm:"primary_key;AUTO_INCREMENT"`
	Name  string `gorm:"column:name"`
	Price int32  `gorm:"column:price"`
}
