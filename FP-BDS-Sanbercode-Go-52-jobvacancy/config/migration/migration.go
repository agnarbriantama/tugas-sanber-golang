package migration

import (
	"fmt"
	"info-loker/config"
	"info-loker/models"
	"log"
)

func Migration() {
	err := config.DB.AutoMigrate(&models.User{}, &models.Jobvacancy{},  &models.ApplyJob{})
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Migrated successfully")
}