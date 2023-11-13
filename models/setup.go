package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=ep-holy-queen-80404705.ap-southeast-1.aws.neon.fl0.io user=fl0user password=5B1omyHNsriR dbname=dbramadhanapis port=5432 sslmode=require"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Pasien{}) // Pastikan Pasien sudah didefinisikan di paket yang sama

	DB = database
}
