package models

import "errors"

//Account errors
var (
	ErrNoAccountFound = errors.New("no account found")
	ErrAddAccount     = errors.New("add account failed")
	ErrDeleteAccount  = errors.New("delete account failed")
	ErrGetAccount     = errors.New("get account failed")
)

//Account table
type Account struct {
	Email     string
	FirstName string
	LastName  string
	Password  string
}

//AccountStore interface to maintain config records
type AccountStore interface {
	Save(account *Account) error
	Delete(email string) error
	Get(email string) (*Account, error)
	GetAll() ([]Account, error)
}
