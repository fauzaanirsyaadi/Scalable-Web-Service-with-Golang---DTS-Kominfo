package model

type Item struct {
	ID         uint   `gorm:"primaryKey" json:"itemId"`
	Item_code  string `gorm:"not null;type:varchar(50)" json:"itemCode"`
	Decryption string `gorm:"not null;type:varchar(50)" json:"descryption"`
	Quantity   int    `gorm:"not null;type:int" json:"quantity"`
	Order_id   uint   `gorm:"not null;type:uint" json:"orderId"`
}
