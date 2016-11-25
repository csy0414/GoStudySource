package RESTServer

import (
	"errors"
	"fmt"
)

type AccountAccess interface {
	GetAccount(id string) (Account, error)
	ModifyAccount(id string, account Account) error
	CreateAccount(account Account) (string, error)
	RemoveAccount(id string) error
}

type MemoryDataAccess struct {
	account map[string]Account
}

func NewMemoryDataAccess() AccountAccess {
	return &MemoryDataAccess{
		account: map[string]Account{},
	}
}

var ErrAccountNotExist = errors.New("Account does not exist.")

func (m *MemoryDataAccess) GetAccount(id string) (Account, error) {
	account, exists := m.account[id]
	if !exists {
		return Account{}, ErrAccountNotExist
	}
	return account, nil
}

func (m *MemoryDataAccess) ModifyAccount(id string, account Account) error {
	if _, exists := m.account[id]; !exists {
		return ErrAccountNotExist
	}

	m.account[id] = account
	return nil
}

func (m *MemoryDataAccess) CreateAccount(account Account) (string, error) {
	if account.id == "" {
		fmt.Printf("Check input ID : %s\n", account.id)
		return "", ErrAccountNotExist
	}

	if account.pw == "" {
		fmt.Printf("Check input password : %s\n", account.pw)
		return "", ErrAccountNotExist
	}

	if account.name == "" {
		fmt.Printf("Check input name : %s\n", account.name)
		return "", ErrAccountNotExist
	}
	m.account[account.id] = account
	return account.id, nil
}

func (m *MemoryDataAccess) RemoveAccount(id string) error {
	if _, exists := m.account[id]; !exists {
		return ErrAccountNotExist
	}
	delete(m.account, id)
	return nil
}
