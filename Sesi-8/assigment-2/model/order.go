package model

import "time"

type Order struct {
	ID            int       `gorm:"primaryKey" json:"order_id"`
	Customer_name string    `gorm:"not null;type:varchar(50)" json:"customer_name"`
	Order_at      time.Time `gorm:"not null;type:timestamp" json:"order_at"`
	Items         []Items   `gorm:"foreignKey:Order_id"`
}
