package domain

import "database/sql"

// TODO: Usar composição e não retornar os sql.NullString
type UserAccount struct {
	ID            string         `json:"id"`
	Name          string         `json:"name"`
	Email         string         `json:"email"`
	Document      string         `json:"document"`
	DocumentType  string         `json:"document_type"`
	Bank          string         `json:"bank"`
	Agency        int32          `json:"agency"`
	AgencyDigit   int8           `json:"agency_digit"`
	AccountNumber int64          `json:"account_number"`
	AccountDigit  int8           `json:"account_digit"`
	AccountType   string         `json:"account_type"`
	Status        string         `json:"status"`
	CreatedAt     string         `json:"created_at"`
	UpdatedAt     sql.NullString `json:"updated_at,omitempty"`
	DeletedAt     sql.NullString `json:"deleted_at,omitempty"`
}

/*
CREATE TABLE user_account (
    id             STRING   PRIMARY KEY
                            UNIQUE
                            NOT NULL,
    name           STRING   NOT NULL,
    email          STRING   NOT NULL,
    document       STRING   NOT NULL,
    document_type  STRING   NOT NULL,
    bank           STRING   NOT NULL,
    agency         INTEGER  NOT NULL,
    agency_digit   INT      NOT NULL,
    account_number INTEGER  NOT NULL,
    account_digit  INT      NOT NULL,
    account_type   STRING   NOT NULL,
    status         STRING   NOT NULL,
    created_at     DATETIME NOT NULL,
    updated_at     DATETIME,
    deleted_at     DATETIME
)
*/
