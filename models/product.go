package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string   `json:"nombre"`
	Description string   `json:"descripcion"`
	SKU         string   `json:"sku"`
	Barcode     string   `gorm:"column:codigo_barras" json:"codigo_barras"`
	CostPrice   float64  `gorm:"column:precio_costo" json:"precio_costo"`
	SalePrice   float64  `gorm:"column:precio_venta" json:"precio_venta"`
	Stock       int      `json:"stock"`
	MinStock    int      `gorm:"column:stock_minimo" json:"stock_minimo"`
	Unit        string   `gorm:"column:unidad" json:"unidad"`
	CategoryID  uint     `gorm:"column:categoria_id" json:"categoria_id"`
	Category    Category `gorm:"foreignKey:CategoryID;references:ID" json:"categoria,omitempty"`
	Supplier    string   `json:"proveedor"`
	Brand       string   `json:"marca"`
	IsActive    bool     `gorm:"column:es_activo" json:"es_activo"`
	ImageURL    string   `gorm:"column:url_imagen" json:"url_imagen"`
}
