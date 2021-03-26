package repositories

import (
	"database/sql"

	"github.com/Fuerback/go-crud/business/domain"
)

type Repository interface {
	GetAll(paginator domain.PaginatorDTO) ([]*domain.UserAccount, error)
	Get(ID string) (*domain.UserAccount, error)
	Save(u *domain.UserAccount) error
	Update(u *domain.UserAccount) error
	Delete(ID string) error
}

func NewService(db *sql.DB) Repository {
	return &sqlite{
		DB: db,
	}
}
