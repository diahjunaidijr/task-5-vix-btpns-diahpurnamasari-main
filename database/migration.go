package database

import "task-5-vix-btpns-Diah_Purnama_Sari/models"

func MigrateDb() {
	db := ConnDb()
	db.AutoMigrate(&models.User{}, &models.Photo{})
}
