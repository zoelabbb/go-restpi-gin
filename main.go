package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zoelabbb/go-restpi-gin/controllers/productcontroller"
	"github.com/zoelabbb/go-restpi-gin/models"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/products", productcontroller.Index)
	r.GET("/api/product/:Id", productcontroller.Show)
	r.POST("/api/product", productcontroller.Create)
	r.PUT("/api/product/:Id", productcontroller.Update)
	r.DELETE("/api/product", productcontroller.Delete)

	r.Run()
}
