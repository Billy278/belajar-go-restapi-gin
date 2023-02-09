package main

import (
	produkcontroller "github.com/Billy278/belajar-go-restapi-gin/controller/produkController"
	"github.com/Billy278/belajar-go-restapi-gin/model"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	model.NewDB()
	router.GET("/api/produk", produkcontroller.Index)
	router.GET("/api/produk/:id", produkcontroller.Show)
	router.POST("/api/produk", produkcontroller.Create)
	router.PUT("/api/produk/:id", produkcontroller.Update)
	router.DELETE("/api/produk", produkcontroller.Delete)
	//router.Run(":4000")
	router.Run()

}
