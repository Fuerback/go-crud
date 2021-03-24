package services

import "github.com/Fuerback/go-crud/business/domain"

type UserAccountService interface {
	GetAll(paginator domain.PaginatorDTO) ([]*domain.UserAccount, error)
	Get(ID string) (*domain.UserAccount, error)
	Save(b *domain.UserAccount) error
	Update(b *domain.UserAccount) error
	Delete(ID string) error
}
