package backend

import (
	"errors"
	"log"
	"path/filepath"
)

var BaseAccountDir = "data/account/"

type AccountRepo struct {
	accounts []*Account
}

func (repo *AccountRepo) Initialize() {
	pattern := BaseAccountDir + "*.json"
	matches, err := filepath.Glob(pattern)

	if err != nil {
		log.Println(err)
	}

	akuns := make([]*Account, len(matches))

	for key, value := range matches {
		akun := NewFromJson(value)
		akuns[key] = &akun
	}

	repo.accounts = akuns
}

func (repo *AccountRepo) List() []*Account {
	return repo.accounts
}

func (repo *AccountRepo) Add(akun *Account) *Account {
	repo.accounts = append(repo.accounts, akun)

	return akun
}

func (repo *AccountRepo) Find(name string) (*Account, error) {
	for _, value := range repo.accounts {
		if value.Name == name {
			return value, errors.New("akun tidak ditemukan")
		}
	}

	return nil, errors.New("akun tidak ditemukan")
}

func (repo *AccountRepo) Delete(name string) error {
	newAccounts := []*Account{}

	account, _ := repo.Find(name)
	account.Delete()

	for _, value := range repo.accounts {
		if value.Name != name {
			newAccounts = append(newAccounts, value)
		}
	}

	repo.accounts = newAccounts

	return nil
}
