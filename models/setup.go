package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// keys := "DSN"
	// dsn := "29iju1fem3r93rb4k8ci:pscale_pw_IncoiqiF88s9HNwnyNYkfGxIYbT5Rk3SDDSA2yGhea1@tcp(ap-southeast.connect.psdb.cloud)/go_restpi_gin?tls=true"
	database, err := gorm.Open(mysql.Open("yn6xhx6hndet7wkh36js:pscale_pw_DhYO0DMBR700TMhsfNINbEhUfN2SRdeikETNaOP0ytw@tcp(ap-southeast.connect.psdb.cloud)/go_restpi_gin?tls=true"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Pasien{})

	DB = database
}
