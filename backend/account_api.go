package backend

func (app *App) AddAccount(name string) *Account {
	akun := &Account{
		Name: name,
	}
	akun.GetCookies(app.Browser)
	akun.Save()

	return app.akunRepo.Add(akun)
}

func (app *App) ListAccount() []*Account {
	return app.akunRepo.accounts
}

func (app *App) DeleteAccount(name string) {
	app.akunRepo.Delete(name)
}

func (app *App) RefreshAccount() {}
