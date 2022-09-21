package book

type BookRequest struct {
	Title       string  `json:"title" binding:"required"`
	Price       float64 `json:"price" binding:"required,number"`
	Description string  `json:"description" binding:"required"`
	Rating      float64 `json:"rating" binding:"required,number"`
	Discount    float64 `json:"discount" binding:"required,number"`
}
