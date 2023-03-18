package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zoelabbb/go-restpi-gin/controllers/productcontroller"
	"github.com/zoelabbb/go-restpi-gin/models"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/pasien", productcontroller.Index)
	r.GET("/api/pasien/:id", productcontroller.Show)
	r.POST("/api/pasien", productcontroller.Create)
	r.PUT("/api/pasien/:id", productcontroller.Update)
	r.DELETE("/api/pasien", productcontroller.Delete)

	r.Run()
}
