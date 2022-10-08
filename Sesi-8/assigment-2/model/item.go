package model

type Item struct {
	ID         uint   `gorm:"primaryKey" json:"item_id"`
	Item_code  string `gorm:"not null;type:varchar(50)" json:"item_code"`
	Decription string `gorm:"not null;type:varchar(50)" json:"description"`
	Quantity   int64  `gorm:"not null;type:int64" json:"quantity"`
	Order_id   uint 	`gorm:"not null;type:uint" json:"order_id
}

// produk == item
// user == order
