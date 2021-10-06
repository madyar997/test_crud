package models

type Author struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type AuthorRepository interface {
	Get() (a []Author, err error)
	Create(a *Author) error
	Delete(id int64) error
	Update(a *Author) error
}

type AuthorUsecase interface {
	Get() ([]Author, error)
	Create(a *Author) error
	Update(a *Author) error
	Delete(id int64) error
}
