package main

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/JcsnP/postgres-docker/controllers"
	"github.com/JcsnP/postgres-docker/models"
)

func main() {
  gin.DisableConsoleColor()

  f, _ := os.Create("gin.log")
  gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

  // create router
  r := gin.Default()
  
  // connect to database
  models.ConnectDatabase()

  r.GET("/products", controllers.GetAllProducts)
  r.POST("/products", controllers.CreateProduct)
  r.GET("/products/:id", controllers.FindProduct)
  r.PATCH("/products/:id", controllers.UpdateProduct)
  r.DELETE("/products/:id", controllers.DeleteProduct)

  // default port is 0.0.0.0:8080
  r.Run()

  fmt.Println("starting at 0.0.0.0:8080")
}
