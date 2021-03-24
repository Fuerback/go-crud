package services

import (
	"database/sql"

	"github.com/Fuerback/go-crud/business/domain"
)

type UserAccountService interface {
	GetAll(paginator domain.PaginatorDTO) ([]*domain.UserAccount, error)
	Get(ID string) (*domain.UserAccount, error)
	Save(b *domain.UserAccountRequestDto) error
	Update(b *domain.UserAccountRequestDto) error
	Delete(ID string) error
}

type Service struct {
	DB *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{
		DB: db,
	}
}

func (s *Service) GetAll(paginator domain.PaginatorDTO) ([]*domain.UserAccount, error) {
	return nil, nil
}
func (s *Service) Get(ID string) (*domain.UserAccount, error) {
	//b é um tipo Beer
	var b domain.UserAccount

	//o comando Prepare verifica se a consulta está válida
	stmt, err := s.DB.Prepare("select id, name, email from user_account where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(ID).Scan(&b.ID, &b.Name, &b.Email)
	if err != nil {
		return nil, err
	}
	//deve retornar a posição da memória de b
	return &b, nil
}
func (s *Service) Save(b *domain.UserAccount) error {
	return nil
}
func (s *Service) Update(b *domain.UserAccount) error {
	return nil
}
func (s *Service) Delete(ID string) error {
	return nil
}
