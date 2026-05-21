package middlewares

import (
	"errors"
	"proyect/db"
	"proyect/models"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MiddlewareProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var product models.Product
		result := c.Param("id")

		if err := db.DB.First(&product, "id = ?", result).Error; err != nil {
			c.AbortWithStatusJSON(400, gin.H{
				"error": "product not found",
			})
			return
		}

		c.Set("product", product)
		c.Next()
	}
}

func MiddlewareCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		var category models.Category
		result := c.Param("id")

		if err := db.DB.First(&category, "id = ?", result).Error; err != nil {
			c.JSON(400, gin.H{
				"error": "category not found",
			})
			c.Abort()
			return
		}

		c.Set("category", category)
		c.Next()
	}
}
func ValidateNewProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var producto models.Product
		if err := c.ShouldBindJSON(&producto); err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": "json inválido: " + err.Error()})
			return
		}

		if strings.TrimSpace(producto.Name) == "" || strings.TrimSpace(producto.SKU) == "" || strings.TrimSpace(producto.Barcode) == "" || strings.TrimSpace(producto.Unit) == "" {
			c.AbortWithStatusJSON(400, gin.H{"error": "Debe enviar name, sku, barcode y unit para diferenciar el producto"})
			return
		}

		if producto.CategoryID == 0 {
			c.AbortWithStatusJSON(400, gin.H{"error": "category_id es obligatorio y debe coincidir con una categoría existente"})
			return
		}

		var category models.Category
		if err := db.DB.First(&category, producto.CategoryID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.AbortWithStatusJSON(400, gin.H{"error": "category_id no coincide con ninguna categoría"})
				return
			}
			c.AbortWithStatusJSON(500, gin.H{"error": "error interno al validar categoría"})
			return
		}

		var existing models.Product
		if err := db.DB.Where("name = ? OR sku = ? OR barcode = ?", producto.Name, producto.SKU, producto.Barcode).First(&existing).Error; err == nil {
			var duplicated []string
			if existing.Name == producto.Name {
				duplicated = append(duplicated, "name")
			}
			if existing.SKU == producto.SKU {
				duplicated = append(duplicated, "sku")
			}
			if existing.Barcode == producto.Barcode {
				duplicated = append(duplicated, "barcode")
			}
			c.AbortWithStatusJSON(400, gin.H{
				"error":             "producto no está completamente diferenciado",
				"duplicated_fields": duplicated,
			})
			return
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(500, gin.H{"error": "error interno al validar producto existente"})
			return
		}

		c.Set("newProduct", producto)
		c.Next()
	}
}
