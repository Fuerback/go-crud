package services

import (
	"github.com/Fuerback/go-crud/business/domain"
	"github.com/Fuerback/go-crud/repositories"
)

var (
	saveParser   SaveUserParser
	updateParser UpdateUserParser
)

type userAccountService struct {
	repository repositories.Repository
}

func (s *userAccountService) GetAll(paginator domain.PaginatorDTO) ([]*domain.UserAccount, error) {
	return s.repository.GetAll(paginator)
}
func (s *userAccountService) Get(ID string) (*domain.UserAccount, error) {
	return s.repository.Get(ID)
}
func (s *userAccountService) Save(u *domain.UserAccountRequestDto) error {
	user, err := saveParser.ParseMessageToDomain(u)
	if err != nil {
		return err
	}
	return s.repository.Save(user)
}
func (s *userAccountService) Update(u *domain.UserAccountRequestDto) error {
	user, err := updateParser.ParseMessageToDomain(u)
	if err != nil {
		return err
	}
	return s.repository.Update(user)
}
func (s *userAccountService) Delete(ID string) error {
	return s.repository.Delete(ID)
}
