package controller

import (
	"assigment-2/database"
	"assigment-2/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	db database.Database
}

func New(db database.Database) Controller {
	return Controller{
		db: db,
	}
}

// get item
func (c Controller) GetItems(ctx *gin.Context) {
	items, err := c.db.GetItems()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    "500",
			"message": "error get data",
		})
	}
	ctx.JSON(http.StatusOK, items)
}

func (c Controller) CreateItem(ctx *gin.Context) {
	var newItem model.Item
	err := ctx.BindJSON(&newItem)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    "500",
			"message": "error bind json request",
		})
	}

	itemResult, err := c.db.CreateItem(newItem)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    "500",
			"message": "error create person",
		})
	}

	ctx.JSON(http.StatusOK, itemResult)
}

func (c Controller) GetOrder(ctx *gin.Context) {
	orders, err := c.db.GetOrders()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    "500",
			"message": "error get data",
		})
	}
	ctx.JSON(http.StatusOK, orders)
}

func (c Controller) CreateOrder(ctx *gin.Context) {
	var newOrder model.Order
	err := ctx.BindJSON(&newOrder)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    "500",
			"message": "error bind json request",
		})
	}

	orderResult, err := c.db.CreateOrder(newOrder)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    "500",
			"message": "error create order",
		})
	}

	ctx.JSON(http.StatusOK, orderResult)
}
