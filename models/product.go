package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string   `json:"name"`
	Description string   `json:"description"`
	SKU         string   `json:"sku"`
	Barcode     string   `json:"barcode"`
	CostPrice   float64  `json:"cost_price"`
	SalePrice   float64  `json:"sale_price"`
	Stock       int      `json:"stock"`
	MinStock    int      `json:"min_stock"`
	Unit        string   `json:"unit"`
	CategoryID  uint     `json:"category_id"`
	Category    Category `json:"category,omitempty"`
	Supplier    string   `json:"supplier"`
	Brand       string   `json:"brand"`
	IsActive    bool     `json:"is_active"`
	ImageURL    string   `json:"image_url"`
}
