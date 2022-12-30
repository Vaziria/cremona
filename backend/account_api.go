package backend

func (app *App) AddAccount(name string) *Account {
	akun := &Account{
		Name: name,
	}

	akun.Save()
	return app.akunRepo.Add(akun)
}

func (app *App) ListAccount() []*Account {
	return app.akunRepo.accounts
}

func (app *App) DeleteAccount(name string) {

	akun, _ := app.akunRepo.Find(name)

	akun.Delete()
}

func (app *App) RefreshAccount() {}
