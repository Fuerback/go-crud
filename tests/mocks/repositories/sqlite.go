package repositories

import (
	"github.com/Fuerback/go-crud/business/domain"
	"github.com/stretchr/testify/mock"
)

type SqliteMock struct {
	mock.Mock
}

func (ref *SqliteMock) GetAll(paginator domain.PaginatorDTO) ([]*domain.UserAccount, error) {
	args := ref.Called(paginator)
	return args.Get(0).([]*domain.UserAccount), args.Error(1)
}

func (ref *SqliteMock) Get(ID string) (*domain.UserAccount, error) {
	args := ref.Called(ID)
	return args.Get(0).(*domain.UserAccount), args.Error(1)
}

func (ref *SqliteMock) Save(u *domain.UserAccount) error {
	args := ref.Called(u)
	return args.Error(0)
}

func (ref *SqliteMock) Update(u *domain.UserAccount) error {
	args := ref.Called(u)
	return args.Error(0)
}

func (ref *SqliteMock) Delete(ID string) error {
	args := ref.Called(ID)
	return args.Error(0)
}
