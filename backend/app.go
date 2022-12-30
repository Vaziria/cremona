package backend

import (
	"context"
	"os"
)

// App struct
type App struct {
	ctx      context.Context
	akunRepo *AccountRepo
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
	a.ctx = ctx

}

func init() {
	os.MkdirAll(BaseAccountDir, os.ModePerm)
}
