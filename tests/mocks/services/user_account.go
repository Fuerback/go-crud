package services

import (
	"github.com/Fuerback/go-crud/business/domain"
	"github.com/stretchr/testify/mock"
)

type UserAccountMock struct {
	mock.Mock
}

func (ref *UserAccountMock) GetAll(paginator domain.PaginatorDTO) ([]*domain.UserAccountRequestDto, error) {
	args := ref.Called(paginator)
	return args.Get(0).([]*domain.UserAccountRequestDto), args.Error(1)
}

func (ref *UserAccountMock) Get(ID string) (*domain.UserAccountRequestDto, error) {
	args := ref.Called(ID)
	return args.Get(0).(*domain.UserAccountRequestDto), args.Error(1)
}

func (ref *UserAccountMock) Save(u *domain.UserAccountRequestDto) error {
	args := ref.Called(u)
	return args.Error(0)
}

func (ref *UserAccountMock) Update(ID string, u *domain.UserAccountRequestDto) error {
	args := ref.Called(ID)
	return args.Error(0)
}

func (ref *UserAccountMock) Delete(ID string) error {
	args := ref.Called(ID)
	return args.Error(0)
}
