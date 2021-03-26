package services

import (
	"github.com/Fuerback/go-crud/business/domain"
	"github.com/Fuerback/go-crud/repositories"
)

type UserAccountService interface {
	GetAll(paginator domain.PaginatorDTO) ([]*domain.UserAccount, error)
	Get(ID string) (*domain.UserAccount, error)
	Save(u *domain.UserAccountRequestDto) error
	Update(ID string, u *domain.UserAccountRequestDto) error
	Delete(ID string) error
}

func NewService(r repositories.Repository) UserAccountService {
	return &userAccountService{
		repository: r,
	}
}
