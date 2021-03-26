package handler_test

import (
	"context"
	"testing"

	handler "github.com/Fuerback/go-crud/api/handlers"
	"github.com/Fuerback/go-crud/tests/mocks/services"
	"github.com/stretchr/testify/suite"
)

type handlerServerSuite struct {
	suite.Suite
	ctx     context.Context
	service *services.UserAccountMock
	handler *handler.Handler
}

func TestHanlderServer(t *testing.T) {
	suite.Run(t, &handlerServerSuite{
		ctx: context.Background(),
	})
}

func (ref *handlerServerSuite) SetupTest() {
	ref.service = new(services.UserAccountMock)
	ref.handler = handler.New(ref.service)
}
