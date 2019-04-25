package service

import (
	"net/http"

	"github.com/hgsgtk/go-snippets/layerd-port-arch/domain/model"
	"github.com/hgsgtk/go-snippets/layerd-port-arch/domain/repository"
	"github.com/hgsgtk/go-snippets/layerd-port-arch/infrastructure/datastore"
)

type UserViewService interface {
	Run(input UserViewInput) (UserViewOutput, *Error)
}

type UserViewInput struct {
	ID int
}

type UserViewOutput struct {
	User model.User
}

type Error struct {
	Code    string
	Message string
	Status  int
	Detail  map[string]string
}

type UserViewServiceImpl struct {
	Conn repository.DBConnector
	User repository.UserRepository
}

func NewUserViewService(conn repository.DBConnector) *UserViewServiceImpl {
	return &UserViewServiceImpl{
		Conn: conn,
		User: &datastore.UserStore{},
	}
}

func (s *UserViewServiceImpl) Run(input UserViewInput) (UserViewOutput, *Error) {
	user, err := s.User.GetByID(s.Conn, input.ID)
	if err != nil {
		return UserViewOutput{}, &Error{
			Code:    "internal_server_error",
			Message: "Internal Server Error",
			Status:  http.StatusInternalServerError,
		}
	}
	return UserViewOutput{User: user}, nil
}
