package model

import "time"

type User struct {
	ID         int       `json:"id" db:"id"`
	FamilyName string    `json:"family_name" db:"family_name" validate:"require,max=50"`
	GivenName  string    `json:"given_name" db:"given_name" validate:"require,max=50"`
	Created    time.Time `json:"created" db:"created"`
	Modified   time.Time `json:"modified" db:"modified"`
}

func (u *User) GetFullName() string {
	return u.FamilyName + " " + u.GivenName
}
