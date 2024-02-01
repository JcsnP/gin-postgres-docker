package models

type Product struct {
  ID            uint    `json:"id" gorm:"primary_key"`
  Title         string  `json:"title"`
  Description   string  `json:"description"`
}

type CreateProduct struct {
  Title         string  `json:"title" binding:"required"`
  Description   string  `json:"description" binding:"required"`
}

type UpdateProduct struct {
  Title         string  `json:"title"`
  Description   string  `json:"description"`
}
