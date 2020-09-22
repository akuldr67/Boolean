package control

import (
	"github.com/akuldr67/Boolean/config"
	"github.com/akuldr67/Boolean/models"
)

func getAllBooleansHelper(bools *[]models.Boolean) (err error) {
	if err = config.DB.Find(bools).Error; err != nil {
		return err
	}
	return nil
}

func createBooleanHelper(newBool *models.Boolean) (err error) {
	if err = config.DB.Create(newBool).Error; err != nil {
		return err
	}
	return nil
}

func getBooleanByIDHelper(boolean *models.Boolean, id string) error {
	if err := config.DB.Where("id = ?", id).First(boolean).Error; err != nil {
		return err
	}
	return nil
}

func updateBooleanHelper(oldBool *models.Boolean, newBool models.Boolean) error {
	config.DB.Model(oldBool).Update(newBool)
	return nil
}

func deleteBooleanHelper(boolean *models.Boolean) error {
	config.DB.Delete(boolean)
	return nil
}
