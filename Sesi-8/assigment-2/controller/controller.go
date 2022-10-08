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

func (c Controller) GetPersons(ctx *gin.Context) {
	persons, err := c.db.GetPersons()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    "500",
			"message": "error get data",
		})
	}
	ctx.JSON(http.StatusOK, persons)
}

func (c Controller) CreatePerson(ctx *gin.Context) {
	var newPerson model.Person
	err := ctx.BindJSON(&newPerson)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    "500",
			"message": "error bind json request",
		})
	}

	personResult, err := c.db.CreatePerson(newPerson)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    "500",
			"message": "error create person",
		})
	}

	ctx.JSON(http.StatusOK, personResult)
}
