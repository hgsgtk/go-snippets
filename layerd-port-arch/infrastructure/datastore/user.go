package datastore

import (
	"github.com/hgsgtk/go-snippets/layerd-port-arch/domain/model"
	"github.com/hgsgtk/go-snippets/layerd-port-arch/domain/repository"
	"github.com/pkg/errors"
)

type UserStore struct {
}

func (*UserStore) GetByID(db repository.DBHandler, id int) (user model.User, err error) {
	q := `
SELECT
	id, family_name, given_name, created, modified
FROM users
WHERE id = ?`

	st, err := db.Preparex(q)
	if err != nil {
		return user, errors.Wrap(err, "UserStore.GetByID got error")
	}
	if err := st.QueryRowx(id).StructScan(&user); err != nil {
		return user, errors.Wrap(err, "UserStore.GetByID got error")
	}
	return
}
