package repositories

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Fuerback/go-crud/business/domain"
)

type sqlite struct {
	DB *sql.DB
}

func (s *sqlite) GetAll(paginator domain.PaginatorDTO) ([]*domain.UserAccount, error) {
	result := []*domain.UserAccount{}

	query := fmt.Sprintf("select * from user_account where deleted_at is null limit %d offset %d", paginator.Limit, paginator.Offset)
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var u domain.UserAccount
		err = rows.Scan(&u.ID, &u.Name, &u.Email, &u.Document, &u.DocumentType, &u.Bank, &u.Agency, &u.AgencyDigit, &u.AccountNumber, &u.AccountDigit, &u.AccountType, &u.Status, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt)
		if err != nil {
			return nil, err
		}

		result = append(result, &u)
	}
	return result, nil
}

func (s *sqlite) Get(ID string) (*domain.UserAccount, error) {
	var u domain.UserAccount

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
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("insert into user_account(id, name, email, document, document_type, bank, agency, agency_digit, account_number, account_digit, account_type, status, created_at) values (?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(u.ID, u.Name, u.Email, u.Document, u.DocumentType, u.Bank, u.Agency, u.AgencyDigit, u.AccountNumber, u.AccountDigit, u.AccountType, u.Status, u.CreatedAt)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (s *sqlite) Update(u *domain.UserAccount) error {
	if u.ID == "" {
		return fmt.Errorf("invalid ID")
	}

	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("update user_account set name=?, email=?, document=?, document_type=?, bank=?, agency=?, agency_digit=?, account_number=?, account_digit=?, account_type=?, status=?, updated_at=? where id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.Name, u.Email, u.Document, u.DocumentType, u.Bank, u.Agency, u.AgencyDigit, u.AccountNumber, u.AccountDigit, u.AccountType, u.Status, u.UpdatedAt.String, u.ID)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (s *sqlite) Delete(ID string) error {
	if ID == "" {
		return fmt.Errorf("invalid ID")
	}

	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("update user_account set deleted_at=? where id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(time.Now().Format(time.RFC3339), ID)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
