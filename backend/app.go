package backend

import (
	"context"
	"os"
)

// App struct
type App struct {
	ctx      context.Context
	akunRepo *AccountRepo
	Browser  *Browser
}

// NewApp creates a new App application struct
func NewApp() *App {

	akunRepo := AccountRepo{
		accounts: []*Account{},
	}
	akunRepo.Initialize()

	return &App{
		akunRepo: &akunRepo,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	config := NewConfigProxy()
	browser := NewBrowser(config.Addr)

	go StartProxy(config)
	StartWebsocketAll(a.akunRepo.accounts)

	a.ctx = ctx
	a.Browser = browser
}

func init() {
	os.MkdirAll(BaseAccountDir, os.ModePerm)
}
