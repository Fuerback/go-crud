package repositories

import (
	"database/sql"

	"github.com/Fuerback/go-crud/business/domain"
)

type sqlite struct {
	DB *sql.DB
}

func (s *sqlite) GetAll(paginator domain.PaginatorDTO) ([]*domain.UserAccount, error) {
	return nil, nil
}
func (s *sqlite) Get(ID string) (*domain.UserAccount, error) {
	var u domain.UserAccount

	//o comando Prepare verifica se a consulta está válida
	stmt, err := s.DB.Prepare("select * from user_account where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(ID).Scan(&u.ID, &u.Name, &u.Email, &u.Document, &u.DocumentType, &u.Bank, &u.Agency, &u.AgencyDigit, &u.AccountNumber, &u.AccountDigit, &u.AccountType, &u.Status, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
func (s *sqlite) Save(u *domain.UserAccount) error {
	//iniciamos uma transação
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("insert into user_account(id, name, email, document, document_type, bank, agency, agency_digit, account_number, account_digit, account_type, status, created_at) values (?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	//o comando Exec retorna um Result, mas não temos interesse nele, por isso podemos ignorá-lo com o _
	_, err = stmt.Exec(u.ID, u.Name, u.Email, u.Document, u.DocumentType, u.Bank, u.Agency, u.AgencyDigit, u.AccountNumber, u.AccountDigit, u.AccountType, u.Status, u.CreatedAt)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
func (s *sqlite) Update(u *domain.UserAccount) error {
	return nil
}
func (s *sqlite) Delete(ID string) error {
	return nil
}
