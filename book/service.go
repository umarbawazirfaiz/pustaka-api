package book

type Service interface {
	FindAll() ([]Book, error)
	FindById(ID int) (Book, error)
	Create(book BookRequest) (Book, error)
	Update(ID int, book BookRequest) (Book, error)
	Delete(ID int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	books, err := s.repository.FindAll()
	return books, err
}

func (s *service) FindById(ID int) (Book, error) {
	book, err := s.repository.FindById(ID)
	return book, err
}

func (s *service) Create(bookRequest BookRequest) (Book, error) {
	book := Book{
		Title:       bookRequest.Title,
		Price:       int(bookRequest.Price),
		Rating:      int(bookRequest.Rating),
		Discount:    int(bookRequest.Discount),
		Description: bookRequest.Description,
	}
	books, err := s.repository.Create(book)
	return books, err
}

func (s *service) Update(ID int, bookRequest BookRequest) (Book, error) {
	book, err := s.repository.FindById(ID)
	if err != nil {
		return book, err
	}

	book.Title = bookRequest.Title
	book.Description = bookRequest.Description
	book.Rating = int(bookRequest.Rating)
	book.Price = int(bookRequest.Price)
	book.Discount = int(bookRequest.Discount)
	book, err = s.repository.Update(book)
	return book, err
}

func (s *service) Delete(ID int) error {
	book, err := s.repository.FindById(ID)
	err = s.repository.Delete(book)

	return err
}
