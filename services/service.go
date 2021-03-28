package services

import (
	"github.com/Fuerback/go-crud/business/domain"
	"github.com/Fuerback/go-crud/repositories"
)

type UserAccountService interface {
	GetAll(paginator domain.PaginatorDTO) ([]*domain.UserAccountRequestDto, error)
	Get(ID string) (*domain.UserAccountRequestDto, error)
	Save(u *domain.UserAccountRequestDto) error
	Update(ID string, u *domain.UserAccountRequestDto) error
	Delete(ID string) error
}

func NewService(r repositories.Repository) UserAccountService {
	return &userAccountService{
		repository: r,
	}
}
