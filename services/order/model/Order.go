package model

type Order struct {
	ID     uint64 `gorm:"primary_key;AUTO_INCREMENT"`
	status string `gorm:"column:status"`

	CarsID uint64 `gorm:"column:cars_id"`
}
