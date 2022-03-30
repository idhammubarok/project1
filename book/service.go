package book

type Service interface {
	FindAll() ([]Book, error)                             // method() (callback)
	FindByID(ID int) (Book, error)                        // method(param) (callback)
	Create(bookRequest BookRequest) (Book, error)         // method(param) (callback)
	Update(ID int, bookRequest BookRequest) (Book, error) // method(param) (callback)
	Delete(ID int) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	book, err := s.repository.FindAll()

	return book, err
}

func (s *service) FindByID(ID int) (Book, error) {
	book, err := s.repository.FindByID(ID)

	return book, err
}

func (s *service) Create(bookRequest BookRequest) (Book, error) {

	price, _ := bookRequest.Price.Int64()

	book := Book{
		Title:       bookRequest.Title,
		Description: bookRequest.Description,
		Price:       int(price),
		Rating:      bookRequest.Rating,
	}

	newBook, err := s.repository.Create(book)

	return newBook, err
}

func (s *service) Update(ID int, bookRequest BookRequest) (Book, error) {

	/*
		pertama cari data berdasarkan id FindByID()
		setelah itu update data yang di inginkan
	*/
	book, err := s.repository.FindByID(ID)
	price, _ := bookRequest.Price.Int64()

	book.Title = bookRequest.Title
	book.Description = bookRequest.Description
	book.Price = int(price)
	book.Rating = bookRequest.Rating

	newBook, err := s.repository.Update(book)

	return newBook, err
}

func (s *service) Delete(ID int) (Book, error) {
	book, err := s.repository.FindByID(ID)
	newBook, err := s.repository.Delete(book)

	return newBook, err
}
