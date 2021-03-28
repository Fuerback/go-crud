package services

import (
	"github.com/Fuerback/go-crud/business/domain"
	"github.com/Fuerback/go-crud/repositories"
)

var (
	saveParser   SaveUserParser
	updateParser UpdateUserParser
	getParser    GetUserParser
)

type userAccountService struct {
	repository repositories.Repository
}

func (s *userAccountService) GetAll(paginator domain.PaginatorDTO) ([]*domain.UserAccountRequestDto, error) {
	users, err := s.repository.GetAll(paginator)
	if err != nil {
		return []*domain.UserAccountRequestDto{}, err
	}

	usersDTO := []*domain.UserAccountRequestDto{}
	for _, u := range users {
		userDTO, _ := getParser.ParseMessageToDomain(u)
		usersDTO = append(usersDTO, userDTO)
	}

	return usersDTO, nil
}

func (s *userAccountService) Get(ID string) (*domain.UserAccountRequestDto, error) {
	user, err := s.repository.Get(ID)
	if err != nil {
		return &domain.UserAccountRequestDto{}, err
	}

	userDTO, _ := getParser.ParseMessageToDomain(user)
	return userDTO, nil
}

func (s *userAccountService) Save(u *domain.UserAccountRequestDto) error {
	user, err := saveParser.ParseMessageToDomain(u)
	if err != nil {
		return err
	}
	return s.repository.Save(user)
}

func (s *userAccountService) Update(ID string, u *domain.UserAccountRequestDto) error {
	user, err := updateParser.ParseMessageToDomain(ID, u)
	if err != nil {
		return err
	}
	return s.repository.Update(user)
}

func (s *userAccountService) Delete(ID string) error {
	return s.repository.Delete(ID)
}
