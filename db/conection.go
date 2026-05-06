package db

import (
	"log"

	"proyect/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "postgresql://neondb_owner:npg_HjzWGey2s9AJ@ep-aged-mud-am6iv1pn-pooler.c-5.us-east-1.aws.neon.tech/neondb?sslmode=require&channel_binding=require"

var DB *gorm.DB

func DBConection() {
	var err error
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Print("DB conectada ✅")
	}
	DB.AutoMigrate(&models.Product{}, &models.Category{})
}
