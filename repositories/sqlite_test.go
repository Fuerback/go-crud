package repositories_test

import (
	"database/sql"
	"testing"
	"time"

	"github.com/Fuerback/go-crud/business/domain"
	"github.com/Fuerback/go-crud/repositories"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

const (
	UUID string = "81b71d14-bea8-41a0-a24f-5ee255e843c1"
	NAME string = "name"
)

func TestSaveAndGet(t *testing.T) {
	u := createUser(UUID)

	db, err := sql.Open("sqlite3", "../db/mydb_test.db")
	if err != nil {
		t.Fatalf("Erro conectando ao banco de dados %s", err.Error())
	}
	err = clearDB(db)
	if err != nil {
		t.Fatalf("Erro limpando o banco de dados: %s", err.Error())
	}
	defer db.Close()
	repository := repositories.NewService(db)
	err = repository.Save(u)
	if err != nil {
		t.Fatalf("Erro salvando no banco de dados: %s", err.Error())
	}
	saved, err := repository.Get(UUID)
	if err != nil {
		t.Fatalf("Erro buscando do banco de dados: %s", err.Error())
	}
	if saved.ID != UUID {
		t.Fatalf("Dados inválidos. Esperado %d, recebido %s", 1, saved.ID)
	}
}

func TestSaveAndDelete(t *testing.T) {
	u := createUser(UUID)

	db, err := sql.Open("sqlite3", "../db/mydb_test.db")
	if err != nil {
		t.Fatalf("Erro conectando ao banco de dados %s", err.Error())
	}
	err = clearDB(db)
	if err != nil {
		t.Fatalf("Erro limpando o banco de dados: %s", err.Error())
	}
	defer db.Close()
	repository := repositories.NewService(db)
	err = repository.Save(u)
	if err != nil {
		t.Fatalf("Erro salvando no banco de dados: %s", err.Error())
	}
	saved, err := repository.Get(UUID)
	if err != nil {
		t.Fatalf("Erro buscando do banco de dados: %s", err.Error())
	}
	if saved.ID != UUID {
		t.Fatalf("Dados inválidos. Esperado %d, recebido %s", 1, saved.ID)
	}
}

func TestSaveAndUpdate(t *testing.T) {
	u := createUser(UUID)

	db, err := sql.Open("sqlite3", "../db/mydb_test.db")
	if err != nil {
		t.Fatalf("Erro conectando ao banco de dados %s", err.Error())
	}
	err = clearDB(db)
	if err != nil {
		t.Fatalf("Erro limpando o banco de dados: %s", err.Error())
	}
	defer db.Close()

	repository := repositories.NewService(db)

	err = repository.Save(u)
	if err != nil {
		t.Fatalf("Erro salvando no banco de dados: %s", err.Error())
	}
	user, err := repository.Get(UUID)
	if err != nil {
		t.Fatalf("Erro buscando do banco de dados: %s", err.Error())
	}
	if user == nil {
		t.Fatalf("User not found")
	}

	err = repository.Delete(UUID)
	if err != nil {
		t.Fatalf("Erro salvando no banco de dados: %s", err.Error())
	}
	user, err = repository.Get(UUID)
	if err == nil {
		t.Fatalf("Erro buscando do banco de dados: %s", err.Error())
	}
}

func TestGetAll(t *testing.T) {
	u1 := createUser(uuid.NewString())
	u2 := createUser(uuid.NewString())
	u3 := createUser(uuid.NewString())

	db, err := sql.Open("sqlite3", "../db/mydb_test.db")
	if err != nil {
		t.Fatalf("Erro conectando ao banco de dados %s", err.Error())
	}
	err = clearDB(db)
	if err != nil {
		t.Fatalf("Erro limpando o banco de dados: %s", err.Error())
	}
	defer db.Close()
	repository := repositories.NewService(db)

	err = repository.Save(u1)
	if err != nil {
		t.Fatalf("Erro salvando no banco de dados: %s", err.Error())
	}
	err = repository.Save(u2)
	if err != nil {
		t.Fatalf("Erro salvando no banco de dados: %s", err.Error())
	}
	err = repository.Save(u3)
	if err != nil {
		t.Fatalf("Erro salvando no banco de dados: %s", err.Error())
	}

	paigantor := &domain.PaginatorDTO{
		Limit:  5,
		Offset: 0,
	}
	users, err := repository.GetAll(*paigantor)
	if err != nil {
		t.Fatalf("Erro buscando do banco de dados: %s", err.Error())
	}
	if len(users) != 3 {
		t.Fatalf("Dados inválidos. Esperado lista com %d itens, recebido %d", 3, len(users))
	}
}

func clearDB(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("delete from user_account")
	tx.Commit()
	return err
}

func createUser(ID string) *domain.UserAccount {
	return &domain.UserAccount{
		ID:            ID,
		Name:          NAME,
		Email:         "email@gmail.com",
		Document:      "07654376612",
		Bank:          "Caixa",
		Agency:        2354,
		AgencyDigit:   1,
		AccountNumber: 35345,
		AccountDigit:  4,
		AccountType:   "corrente",
		Status:        "rascunho",
		CreatedAt:     time.Now().Format(time.RFC3339),
	}
}
