package book

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Book, error)       // method() (callback)
	FindByID(ID int) (Book, error)  // method(param) (callback)
	Create(book Book) (Book, error) // method(param) (callback)
	Update(book Book) (Book, error) // method(param) (callback)
	Delete(book Book) (Book, error)
}

type repository struct { // membuat struck private
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Book, error) {
	var book []Book

	err := r.db.Find(&book).Error
	return book, err
}

func (r *repository) FindByID(ID int) (Book, error) {
	var book Book

	err := r.db.First(&book, ID).Error

	return book, err
}

func (r *repository) Create(book Book) (Book, error) {

	err := r.db.Create(&book).Error

	return book, err
}

func (r *repository) Update(book Book) (Book, error) {

	err := r.db.Save(&book).Error
	return book, err
}

func (r *repository) Delete(book Book) (Book, error) {
	err := r.db.Delete(&book).Error

	return book, err
}
