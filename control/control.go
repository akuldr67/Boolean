package control

import (
	"fmt"
	"net/http"

	"github.com/akuldr67/Boolean/config"
	"github.com/akuldr67/Boolean/models"
	uuid "github.com/satori/go.uuid"

	"github.com/gin-gonic/gin"
)

func getAllBooleans(c *gin.Context) {
	var bools []models.Boolean
	// err := models.GetAllBooleans(&bools)
	err := config.DB.Find(&bools).Error
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, bools)
	}
}

func createBoolean(c *gin.Context) {
	var newBool models.Boolean
	err := c.ShouldBindJSON(&newBool)

	// Jugaad for pre create hook!!
	uuid, err := uuid.NewV4()
	newBool = models.Boolean{ID: uuid, Key: newBool.Key, Value: newBool.Value}

	fmt.Println(newBool)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = config.DB.Create(&newBool).Error
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, newBool)
	}
}
