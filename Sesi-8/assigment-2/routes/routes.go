package routes

import (
	"assigment-2/controller"

	"github.com/gin-gonic/gin"
)

func StartServer(ctl controller.Controller) error {
	r := gin.Default()

	// item order
	r.GET("/items", ctl.GetItems)
	r.GET("/items/:id")
	r.POST("/item", ctl.CreateItem)
	r.PUT("/item/:id")
	r.DELETE("/item/:id")

	// order controller
	r.GET("/orders", ctl.GetOrders)
	r.GET("/orders/:id")
	r.POST("/order", ctl.CreateOrder)
	r.PUT("/order/:id")
	r.DELETE("/order/:id")

	return r.Run("localhost:8080")
}
