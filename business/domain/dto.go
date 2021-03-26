package domain

// TODO: Permitir digito 0 e validar valor maximo e minimo

type UserAccountRequestDto struct {
	Name          string `json:"name" validate:"required"`
	Email         string `json:"email" validate:"required,email"`
	Document      string `json:"document" validate:"required"`
	Bank          string `json:"bank" validate:"required"`
	Agency        int32  `json:"agency" validate:"required"`
	AgencyDigit   int8   `json:"agency_digit" validate:"required"`
	AccountNumber int64  `json:"account_number" validate:"required"`
	AccountDigit  int8   `json:"account_digit" validate:"required"`
	AccountType   string `json:"account_type" validate:"required"`
	Status        string `json:"status" validate:"required"`
}

type PaginatorDTO struct {
	Limit  int `json:"limit" validate:"required"`
	Offset int `json:"offset"`
}

// usar composição aqui
