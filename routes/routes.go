package routes

import (
	"proyect/db"
	"proyect/models"

	"github.com/gin-gonic/gin"
)

// rutas para productos
func ProductRoutes(c *gin.Context) {
	var productos []models.Product
	db.DB.Preload("category").Find(&productos)
	c.JSON(200, productos)
}

func ProductId(c *gin.Context) {
	var producto models.Product
	id := c.Param("id")
	db.DB.Preload("category").First(&producto, id)
	c.JSON(200, producto)
}

// ruta post para crear un producto
func NewProduct(c *gin.Context) {
	producto := c.MustGet("newProduct").(models.Product)
	if err := db.DB.Create(&producto).Error; err != nil {
		c.JSON(500, gin.H{"error": "no se pudo crear el producto"})
		return
	}
	c.JSON(201, producto)
}

// rutas para eliminar un producto
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := db.DB.Delete(&models.Product{}, id).Error; err != nil {
		c.JSON(500, gin.H{"error": "no se pudo eliminar el producto"})
		return
	}
	c.JSON(200, gin.H{"message": "Producto eliminado"})
}

// rutas para categorias
func CategoryRoutes(c *gin.Context) {
	var categorias []models.Category
	db.DB.Preload("products").Find(&categorias)
	c.JSON(200, categorias)
}

func CategoryId(c *gin.Context) {
	var categoria models.Category
	id := c.Param("id")
	db.DB.Preload("products").First(&categoria, id)
	c.JSON(200, categoria)
}
