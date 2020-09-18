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
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	newBool.ID, _ = uuid.NewV4()
	// newBool = models.Boolean{ID: uuid, Key: newBool.Key, Value: newBool.Value}

	err = config.DB.Create(&newBool).Error
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, newBool)
	}
}

func getBooleanByIDHelper(boolean *models.Boolean, id string) error {
	if err := config.DB.Where("id = ?", id).First(boolean).Error; err != nil {
		return err
	}
	return nil
}

func getBooleanByID(c *gin.Context) {
	var boolean models.Boolean
	fmt.Println(c.Params)
	id := c.Params.ByName("id")

	err := getBooleanByIDHelper(&boolean, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, boolean)
	}
}

func updateBoolean(c *gin.Context) {
	var boolean models.Boolean
	id := c.Params.ByName("id")

	err := getBooleanByIDHelper(&boolean, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	// c.BindJSON(&boolean)
	// config.DB.Save(&boolean)

	var input models.Boolean
	err2 := c.ShouldBindJSON(&input)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}
	config.DB.Model(&boolean).Updates(input)

	c.JSON(http.StatusOK, boolean)
}

func deleteBoolean(c *gin.Context) {
	var boolean models.Boolean
	id := c.Params.ByName("id")

	err := getBooleanByIDHelper(&boolean, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	config.DB.Delete(&boolean)

	// c.JSON(http.StatusOK, gin.H{"id" + id: "deleted"})
	c.Writer.WriteHeader(204)
}
