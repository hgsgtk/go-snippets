package service

import (
	"github.com/hgsgtk/go-snippets/layerd-port-arch/domain/model"
	"github.com/hgsgtk/go-snippets/layerd-port-arch/domain/repository"
)

type UserAddServiceImpl struct {
	Conn repository.DBConnector
	User repository.UserRepository
}

type UserAddServiceInput struct {
	FamilyName string
	GivenName  string
}

type UserAddServiceOutput struct {
}

func (s *UserAddServiceImpl) Run(input UserAddServiceInput) (UserAddServiceOutput, *Error) {
	tx, err := s.Conn.Beginx()
	if err != nil {
		// error handling
	}

	u := model.User{
		FamilyName: input.FamilyName,
		GivenName:  input.GivenName,
	}
	if err := s.User.Save(tx, u); err != nil {
		// error handling
		tx.Rollback()
	}

	tx.Commit()
	return UserAddServiceOutput{}, nil
}
