package control

import (
	"fmt"
	"net/http"

	"github.com/akuldr67/Boolean/models"
	uuid "github.com/satori/go.uuid"

	"github.com/gin-gonic/gin"
)

func getAllBooleans(c *gin.Context) {
	var bools []models.Boolean

	// err := config.DB.Find(&bools).Error
	err := getAllBooleansHelper(&bools)
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

	// err = config.DB.Create(&newBool).Error
	err = createBooleanHelper(&newBool)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, newBool)
	}
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
	err = c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}
	// config.DB.Model(&boolean).Updates(input)
	err = updateBooleanHelper(&boolean, input)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, boolean)
	}
}

func deleteBoolean(c *gin.Context) {
	var boolean models.Boolean
	id := c.Params.ByName("id")

	err := getBooleanByIDHelper(&boolean, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	// config.DB.Delete(&boolean)
	err = deleteBooleanHelper(&boolean)

	// c.JSON(http.StatusOK, gin.H{"id" + id: "deleted"})
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.Writer.WriteHeader(204)
	}
}
