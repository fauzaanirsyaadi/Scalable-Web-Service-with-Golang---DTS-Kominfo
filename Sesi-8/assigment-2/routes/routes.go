package routes

import (
	"assigment-2/controller"

	"github.com/gin-gonic/gin"
)

func StartServer(ctl controller.Controller) error {
	r := gin.Default()

	r.GET("/items", ctl.GetItems)
	r.GET("/items/:id")
	r.POST("/item", ctl.CreateItem)
	r.PUT("/item/:id")
	r.DELETE("/item/:id")

	return r.Run("localhost:8080")
}
