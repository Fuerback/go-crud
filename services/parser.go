package services

import (
	"database/sql"
	"time"

	"github.com/Fuerback/go-crud/business/domain"
	"github.com/google/uuid"
)

type (
	SaveUserParser   struct{}
	UpdateUserParser struct{}
)

func getDocumentType(document string) (string, error) {
	switch len(document) {
	case 11:
		return "NATURAL", nil
	case 14:
		return "LEGAL", nil
	default:
		return "", domain.ErrInvalidDocument
	}
}

func (ref UpdateUserParser) ParseMessageToDomain(ID string, u *domain.UserAccountRequestDto) (*domain.UserAccount, error) {
	documentType, err := getDocumentType(u.Document)
	if err != nil {
		return &domain.UserAccount{}, err
	}

	user := &domain.UserAccount{
		ID:            ID,
		Name:          u.Name,
		Email:         u.Email,
		Document:      u.Document,
		DocumentType:  documentType,
		Bank:          u.Bank,
		Agency:        u.Agency,
		AgencyDigit:   u.AgencyDigit,
		AccountNumber: u.AccountNumber,
		AccountDigit:  u.AccountDigit,
		AccountType:   u.AccountType,
		Status:        u.Status,
		UpdatedAt:     sql.NullString{String: time.Now().Format(time.RFC3339)},
	}

	return user, nil
}

func (ref SaveUserParser) ParseMessageToDomain(u *domain.UserAccountRequestDto) (*domain.UserAccount, error) {
	documentType, err := getDocumentType(u.Document)
	if err != nil {
		return &domain.UserAccount{}, err
	}

	user := &domain.UserAccount{
		ID:            uuid.New().String(),
		Name:          u.Name,
		Email:         u.Email,
		Document:      u.Document,
		DocumentType:  documentType,
		Bank:          u.Bank,
		Agency:        u.Agency,
		AgencyDigit:   u.AgencyDigit,
		AccountNumber: u.AccountNumber,
		AccountDigit:  u.AccountDigit,
		AccountType:   u.AccountType,
		Status:        u.Status,
		CreatedAt:     time.Now().Format(time.RFC3339),
	}

	return user, nil
}
