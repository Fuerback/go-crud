package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	handler "github.com/Fuerback/go-crud/api/handlers"
	"github.com/Fuerback/go-crud/business/domain"
	"github.com/Fuerback/go-crud/tests/mocks/services"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

const UUID string = "81b71d14-bea8-41a0-a24f-5ee255e843c1"

type handlerSuite struct {
	suite.Suite
	ctx     context.Context
	service *services.UserAccountMock
	handler *handler.Handler
	router  *mux.Router
}

func TestHanlderServer(t *testing.T) {
	suite.Run(t, &handlerSuite{
		ctx: context.Background(),
	})
}

func (ref *handlerSuite) SetupTest() {
	ref.service = new(services.UserAccountMock)
	ref.handler = handler.New(ref.service)
	ref.router = mux.NewRouter()
	ref.handler.Handlers(ref.router)
}

func (ref *handlerSuite) TestGetAllUsers_Success() {
	u := getUserDTO()
	users := []*domain.UserAccountRequestDto{u}

	ref.service.On("GetAll", mock.Anything).Return(users, nil).Once()

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/api/useraccount", nil)
	ref.router.ServeHTTP(w, r)

	ref.Equal(http.StatusOK, w.Code)
}

func (ref *handlerSuite) TestDelete_Success() {
	ref.service.On("Delete", UUID).Return(nil).Once()

	path := fmt.Sprintf("/api/useraccount/%s", UUID)
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("DELETE", path, nil)
	ref.router.ServeHTTP(w, r)

	ref.Equal(http.StatusOK, w.Code)
}

func (ref *handlerSuite) TestUpdate_Success() {
	userDTO := getUserDTO()
	ref.service.On("Update", UUID, mock.Anything).Return(nil).Once()

	path := fmt.Sprintf("/api/useraccount/%s", UUID)
	body, _ := json.Marshal(userDTO)
	requestReader := bytes.NewReader(body)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("PUT", path, requestReader)
	ref.router.ServeHTTP(w, r)

	ref.Equal(http.StatusOK, w.Code)
}

func (ref *handlerSuite) TestGetUser_Success() {
	user := getUserDTO()
	ref.service.On("Get", UUID).Return(user, nil).Once()

	path := fmt.Sprintf("/api/useraccount/%s", UUID)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", path, nil)
	ref.router.ServeHTTP(w, r)

	ref.Equal(http.StatusOK, w.Code)
}

func (ref *handlerSuite) TestSaveUser_Success() {
	userDTO := getUserDTO()
	ref.service.On("Save", userDTO).Return(nil).Once()

	body, _ := json.Marshal(userDTO)
	requestReader := bytes.NewReader(body)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/api/useraccount", requestReader)
	ref.router.ServeHTTP(w, r)

	ref.Equal(http.StatusCreated, w.Code)
}

func getUserEntity() domain.UserAccount {
	return domain.UserAccount{
		ID:            uuid.NewString(),
		Name:          "felipe",
		Email:         "felipe@gmail.com",
		Document:      "40843908744",
		Bank:          "Caixa",
		Agency:        43665,
		AgencyDigit:   4,
		AccountNumber: 2463456,
		AccountDigit:  4,
		AccountType:   "poupança",
		Status:        "rascunho",
		CreatedAt:     time.Now().Format(time.RFC3339),
	}
}

func getUserDTO() *domain.UserAccountRequestDto {
	return &domain.UserAccountRequestDto{
		Name:          "felipe",
		Email:         "felipe@gmail.com",
		Document:      "40843908744",
		Bank:          "Caixa",
		Agency:        43665,
		AgencyDigit:   4,
		AccountNumber: 2463456,
		AccountDigit:  4,
		AccountType:   "poupança",
		Status:        "rascunho",
	}
}
