package dto

type CreateProductRequest struct {
	ProductCode string  `json:"product_code" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	Category    string  `json:"category" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required,gte=0"`
}

type UpdateProductRequest struct {
	ProductCode string  `json:"product_code" binding:"required"`
	Name        string  `json:"name" binding:"omitempty"`
	Category    string  `json:"category" binding:"omitempty"`
	Description string  `json:"description" binding:"omitempty"`
	Price       float64 `json:"price" binding:"omitempty,gte=0"`
}

type ProductResponse struct {
	ProductCode string  `json:"product_code"`
	Name        string  `json:"name"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
