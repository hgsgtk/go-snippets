package repository

import "github.com/hgsgtk/go-snippets/layerd-port-arch/domain/model"

type UserRepository interface {
	GetByID(db DBHandler, id int) (model.User, error)
	Save(db DBHandler, user model.User) error
}
