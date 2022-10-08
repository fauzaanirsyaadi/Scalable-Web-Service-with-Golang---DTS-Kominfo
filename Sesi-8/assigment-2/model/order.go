package model

import "time"

type Order struct {
	ID            int       `gorm:"primaryKey" json:"orderId"`
	Customer_name string    `gorm:"not null;type:varchar(50)" json:"customerName"`
	Order_at      time.Time `gorm:"not null;type:timestamp" json:"orderAt"`
	Items         []Item    `gorm:"foreignKey:OrderId"`
}
