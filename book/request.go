package book

import "encoding/json"

type BookRequest struct {
	Title       string      `json:"title" binding:"required"`
	Description string      `json:"description" binding:"required"`
	Price       json.Number `json:"price" binding:"required,number"`  // json.Number = untuk menangkap nilai number/ string dari client
	Rating      int         `json:"rating" binding:"required,number"` // json.Number = untuk menangkap nilai number/ string dari client

	// SubTitle string `json:"sub_title"` // json sub_title jika parameter struck _ lebih dari 2 kata
}
