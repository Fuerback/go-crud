package services_test

import (
	"context"
	"testing"
	"time"

	"github.com/Fuerback/go-crud/business/domain"
	"github.com/Fuerback/go-crud/services"
	"github.com/Fuerback/go-crud/tests/mocks/repositories"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type userAccountSuite struct {
	suite.Suite
	ctx        context.Context
	repository *repositories.SqliteMock
	services   services.UserAccountService
}

func TestHanlderServer(t *testing.T) {
	suite.Run(t, &userAccountSuite{
		ctx: context.Background(),
	})
}

func (ref *userAccountSuite) SetupTest() {
	ref.repository = new(repositories.SqliteMock)
	ref.services = services.NewService(ref.repository)
}

func (ref *userAccountSuite) TestSaveUser_Success() {
	u := getUserDTO()

	ref.repository.On("Save", mock.Anything).Return(nil).Once()

	err := ref.services.Save(u)
	ref.NoError(err)
}

func (ref *userAccountSuite) TestGetAll_Success() {
	p := domain.PaginatorDTO{
		Limit:  5,
		Offset: 0,
	}

	u := getUserEntity()
	users := []*domain.UserAccount{&u}

	ref.repository.On("GetAll", p).Return(users, nil).Once()

	list, err := ref.services.GetAll(p)
	ref.NoError(err)
	ref.Equal(len(list), 1)
}

func (ref *userAccountSuite) TestGet_Success() {
	id := uuid.NewString()

	u := getUserEntity()

	ref.repository.On("Get", id).Return(&u, nil).Once()

	user, err := ref.services.Get(id)
	ref.NoError(err)
	ref.NotNil(user)
}

func (ref *userAccountSuite) TestUpdate_Success() {
	id := uuid.NewString()

	dto := getUserDTO()

	ref.repository.On("Update", mock.Anything).Return(nil).Once()

	err := ref.services.Update(id, dto)
	ref.NoError(err)
}

func (ref *userAccountSuite) TestDelete_Success() {
	id := uuid.NewString()

	ref.repository.On("Delete", id).Return(nil).Once()

	err := ref.services.Delete(id)
	ref.NoError(err)
}

func getUserEntity() domain.UserAccount {
	return domain.UserAccount{
		ID:            uuid.NewString(),
		Name:          "felipe",
		Email:         "felipe@gmail.com",
		Document:      "40843908744",
		Bank:          "Caixa",
		Agency:        43665,
		AgencyDigit:   4,
		AccountNumber: 2463456,
		AccountDigit:  4,
		AccountType:   "poupança",
		Status:        "rascunho",
		CreatedAt:     time.Now().Format(time.RFC3339),
	}
}

func getUserDTO() *domain.UserAccountRequestDto {
	return &domain.UserAccountRequestDto{
		Name:          "felipe",
		Email:         "felipe@gmail.com",
		Document:      "40843908744",
		Bank:          "Caixa",
		Agency:        43665,
		AgencyDigit:   4,
		AccountNumber: 2463456,
		AccountDigit:  4,
		AccountType:   "poupança",
		Status:        "rascunho",
	}
}
