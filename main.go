package main

import (
	"proyect/db"
	"proyect/models"

	"github.com/gin-gonic/gin"
)

func main() {
	db.DBConection()
	r := gin.Default()

	// rutas para productos
	r.GET("/products", func(c *gin.Context) {
		var productos []models.Product
		db.DB.Preload("category").Find(&productos)
		c.JSON(200, productos)
	})
	r.GET("/products/:id", func(c *gin.Context) {
		var producto models.Product
		id := c.Param("id")
		db.DB.Preload("category").First(&producto, id)
		c.JSON(200, producto)
	})

	// ruta post para crear un producto
	r.POST("/products", func(c *gin.Context) {
		var producto models.Product
		if err := c.ShouldBindJSON(&producto); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		db.DB.Create(&producto)
		db.DB.Save(&producto)
		c.JSON(201, producto)
	})

	// rutas para eliminar un producto
	r.DELETE("/products/:id", func(c *gin.Context) {
		var producto models.Product
		id := c.Param("id")
		db.DB.Delete(&models.Product{}, id)
		db.DB.Save(&producto)
		c.JSON(200, gin.H{"message": "Producto eliminado"})
	})

	// rutas para categorias
	r.GET("/categories", func(c *gin.Context) {
		var categorias []models.Category
		db.DB.Preload("products").Find(&categorias)
		c.JSON(200, categorias)
	})
	r.GET("/categories/:id", func(c *gin.Context) {
		var categoria models.Category
		id := c.Param("id")
		db.DB.Preload("products").First(&categoria, id)
		c.JSON(200, categoria)
	})

	r.Run(":8080")
}
