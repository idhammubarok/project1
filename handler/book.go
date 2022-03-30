package handler

import (
	"fmt"
	"net/http"
	"project1/book"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) GetBooks(c *gin.Context) {
	books, err := h.bookService.FindAll()

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})

	}

	var booksResponse []book.BookResponse

	for _, b := range books {
		bookResponse := convertBookResponse(b)

		booksResponse = append(booksResponse, bookResponse) //append(slice, object)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
	})
}

func (h *bookHandler) GetBook(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	bookObj, err := h.bookService.FindByID(id)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})

		return

	}

	bookResponse := convertBookResponse(bookObj)

	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

func (h *bookHandler) CreateBook(c *gin.Context) {
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest) //cek error

	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage) // error message pada perulangan di tambahkan pakai append dan di tampung di slice
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})

		return //agar mengembalikan nilai kosong
	}

	book, err := h.bookService.Create(bookRequest)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})

		return //agar mengembalikan nilai kosong
	}

	bookResponse := convertBookResponse(book)

	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
		// "sub_title": bookInput.SubTitle,
	})
}

func (h *bookHandler) UpdateBook(c *gin.Context) {

	var bookRequest book.BookRequest
	err := c.ShouldBindJSON(&bookRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})

		return
	}

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	book, err := h.bookService.Update(id, bookRequest)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})

		return //agar mengembalikan nilai kosong
	}

	bookResponse := convertBookResponse(book)

	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

func (h *bookHandler) DeleteBook(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	book, err := h.bookService.Delete(id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"messaage":    "Data has been deleted!",
		"status_code": http.StatusOK,
		"data":        convertBookResponse(book),
	})
}

func convertBookResponse(bookOjt book.Book) book.BookResponse {
	return book.BookResponse{
		ID:          bookOjt.ID,
		Title:       bookOjt.Title,
		Description: bookOjt.Description,
		Price:       bookOjt.Price,
		Rating:      bookOjt.Rating,
	}
}
