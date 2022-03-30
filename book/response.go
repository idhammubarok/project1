package book

type BookResponse struct {
	ID                int                 `json:"id"`
	Title             string              `json:"title"`
	Description       string              `json:"description"`
	Price             int                 `json:"price"`
	Rating            int                 `json:"rating"`
	BookResponseChild []BookResponseChild `json:"book_response_child"`
}

type BookResponseChild struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
