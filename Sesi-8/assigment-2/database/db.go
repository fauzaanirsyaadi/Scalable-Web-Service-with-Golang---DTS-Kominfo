package database

import (

	// import model
	"assigment-2/model"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func Start() (Database, error) {
	var host = "localhost"
	var port = 5432
	var username = "postgres"
	var password = "admin"
	var dbName = "db-go-sql"

	var conn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbName)

	db, err := gorm.Open(postgres.Open(conn))
	if err != nil {
		fmt.Println("error open connection to db", err)
		return Database{}, err
	}

	err = db.Debug().AutoMigrate(model.Item{})
	if err != nil {
		fmt.Println("error on migration", err)
		return Database{}, err
	}

	err = db.Debug().AutoMigrate(model.Order{})
	if err != nil {
		fmt.Println("error on migration", err)
		return Database{}, err
	}

	return Database{
		db: db,
	}, nil
}

// make func get item
func (d Database) GetItems() ([]model.Item, error) {
	dbg := d.db.Find(&[]model.Item{})

	rows, err := dbg.Rows()
	if err != nil {
		return nil, err
	}

	items := make([]model.Item, 0)

	for rows.Next() {
		var item model.Item
		err := rows.Scan(&item.ID, &item.Item_code, &item.Decryption, &item.Quantity, &item.Order_id)
		if err != nil {
			continue
		}
		items = append(items, item)
	}
	return items, nil
}

// mae func get item by id
func (d Database) GetItemById(id uint) (model.Item, error) {
	dbg := d.db.Find(&model.Item{}, id)

	row := dbg.Row()

	var itemResult model.Item

	err := row.Scan(&itemResult.ID, &itemResult.Item_code, &itemResult.Decryption, &itemResult.Quantity, &itemResult.Order_id)

	return itemResult, err
}

// make func create item
func (d Database) CreateItem(item model.Item) (model.Item, error) {
	dbg := d.db.Create(&item)

	row := dbg.Row()

	var itemResult model.Item

	err := row.Scan(&itemResult.ID, &itemResult.Item_code, &itemResult.Decryption, &itemResult.Quantity, &itemResult.Order_id)

	return itemResult, err
}

// make func update item
func (d Database) UpdateItem(item model.Item) (model.Item, error) {
	dbg := d.db.Save(&item)

	row := dbg.Row()

	var itemResult model.Item

	err := row.Scan(&itemResult.ID, &itemResult.Item_code, &itemResult.Decryption, &itemResult.Quantity, &itemResult.Order_id)

	return itemResult, err
}

// make func delete item
func (d Database) DeleteItem(item model.Item) (model.Item, error) {
	dbg := d.db.Delete(&item)

	row := dbg.Row()

	var itemResult model.Item

	err := row.Scan(&itemResult.ID, &itemResult.Item_code, &itemResult.Decryption, &itemResult.Quantity, &itemResult.Order_id)

	return itemResult, err
}

// make func get order
func (d Database) GetOrders() ([]model.Order, error) {
	dbg := d.db.Find(&[]model.Order{})

	rows, err := dbg.Rows()
	if err != nil {
		return nil, err
	}

	orders := make([]model.Order, 0)

	for rows.Next() {
		var order model.Order
		err := rows.Scan(&order.ID, &order.Customer_name, &order.Order_at)
		if err != nil {
			continue
		}
		orders = append(orders, order)
	}
	return orders, nil
}

// make func get order by id
func (d Database) GetOrderById(id uint) (model.Order, error) {
	dbg := d.db.Find(&model.Order{}, id)

	row := dbg.Row()

	var orderResult model.Order

	err := row.Scan(&orderResult.ID, &orderResult.Customer_name, &orderResult.Order_at)

	return orderResult, err
}

// make func create order
func (d Database) CreateOrder(order model.Order) (model.Order, error) {
	dbg := d.db.Create(&order)

	row := dbg.Row()

	var orderResult model.Order

	err := row.Scan(&orderResult.ID, &orderResult.Customer_name, &orderResult.Order_at)

	return orderResult, err
}

// make func update order
func (d Database) UpdateOrder(order model.Order) (model.Order, error) {
	dbg := d.db.Save(&order)

	row := dbg.Row()

	var orderResult model.Order

	err := row.Scan(&orderResult.ID, &orderResult.Customer_name, &orderResult.Order_at)

	return orderResult, err
}

// make func delete order
func (d Database) DeleteOrder(order model.Order) (model.Order, error) {
	dbg := d.db.Delete(&order)

	row := dbg.Row()

	var orderResult model.Order

	err := row.Scan(&orderResult.ID, &orderResult.Customer_name, &orderResult.Order_at)

	return orderResult, err
}
