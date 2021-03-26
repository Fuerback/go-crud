package domain

type UserAccountRequestDto struct {
	Name          string `json:"name" validate:"required,min=2,max=100"`
	Email         string `json:"email" validate:"required,email"`
	Document      string `json:"document" validate:"required,min=11,max=14"`
	Bank          string `json:"bank" validate:"required"`
	Agency        int32  `json:"agency" validate:"required"`
	AgencyDigit   int8   `json:"agency_digit" validate:"required,min=1,max=1"`
	AccountNumber int64  `json:"account_number" validate:"required"`
	AccountDigit  int8   `json:"account_digit" validate:"required,min=1,max=1"`
	AccountType   string `json:"account_type" validate:"required"`
	Status        string `json:"status" validate:"required"`
}

type PaginatorDTO struct {
	Limit  int `json:"limit" validate:"number:min=1&max=100"`
	Offset int `json:"offset"`
}

// usar composição aqui
