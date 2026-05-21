package main

import (
	"proyect/db"
	"proyect/middlewares"
	"proyect/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.DBConection()
	r := gin.Default()

	r.GET("/products", routes.ProductRoutes)

	r.GET("/products/:id", routes.ProductId)

	r.POST("/products", middlewares.ValidateNewProduct(), routes.NewProduct)

	r.DELETE("/products/:id", routes.DeleteProduct)

	r.GET("/categories", routes.CategoryRoutes)

	r.GET("/categories/:id", routes.CategoryId)

	r.Run(":8080")
}
