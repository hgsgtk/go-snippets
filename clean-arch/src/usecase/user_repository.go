package usecase

import "github.com/Khigashiguchi/go-snippets/clean-arch/src/domain"

type UserRepository interface {
	Store(domain.User) (int, error)
	FindById(int) (domain.User, error)
	FindAll() (domain.Users, error)
}
