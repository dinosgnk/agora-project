package dto

type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Category    string  `json:"category" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required,gte=0"`
}

type UpdateProductRequest struct {
	ProductId   int     `json:"id" binding:"required"`
	Name        *string `json:"name"`
	Category    *string `json:"category"`
	Description *string `json:"description"`
	Price       *float64 `json:"price" binding:"omitempty,gte=0"`
}

type ProductResponse struct {
	ProductId   int     `json:"id"`
	Name        string  `json:"name"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
