package users

import (
	"crypto/md5"
	"errors"
	"fmt"
	"time"
)

var (
	ErrorNameRequired     = errors.New("Name is required")
	ErrorLoginRequired    = errors.New("Login is required")
	ErrorPasswordRequired = errors.New("Password is required")
	ErrorPasswordLength   = errors.New("Password must be have at 6 characters")
	ErrorPasswordEmpty    = errors.New("Password can't be empty")
)

func New(name, login, password string) (*User, error) {
	now := time.Now()
	u := User{Name: name, Login: login, CreatedAt: now, ModifiedAt: now}
	err := u.SetPassword(password)
	if err != nil {
		return nil, err
	}

	err = u.Validate()
	if err != nil {
		return nil, err
	}

	return &u, nil
}

type User struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Login      string    `json:"login"`
	Password   string    `json:"password"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	Deleted    bool      `json:"-"`
	LastLogin  time.Time `json:"last_login"`
}

func (u *User) SetPassword(password string) error {
	if password == "" {
		return ErrorPasswordRequired
	}

	if len(password) < 6 {
		return ErrorPasswordLength
	}

	u.Password = fmt.Sprintf("%x", md5.Sum([]byte(password)))
	return nil
}

func (u *User) Validate() error {
	if u.Name == "" {
		return ErrorNameRequired
	}

	if u.Login == "" {
		return ErrorLoginRequired
	}

	var emptyPassword = fmt.Sprintf("%x", md5.Sum([]byte("")))
	if u.Password == emptyPassword {
		return ErrorPasswordEmpty
	}

	return nil
}
