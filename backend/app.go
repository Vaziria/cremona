package backend

import (
	"context"
	"log"
	"os"

	"github.com/wailsapp/wails/v2/pkg/runtime"
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
	resSocket := make(chan *Response)
	// defer close(resSocket)

	config := NewConfigProxy()
	browser := NewBrowser(config.Addr)

	a.ctx = ctx
	a.Browser = browser

	go StartProxy(config)
	go a.sendToFrontend(resSocket)
	StartWebsocketAll(a.akunRepo.accounts, resSocket)

}

func init() {
	os.MkdirAll(BaseAccountDir, os.ModePerm)
}

func (a *App) sendToFrontend(resSocket <-chan *Response) {
	for res := range resSocket {
		runtime.EventsEmit(a.ctx, "message", &res)
		log.Println("send to frontend", &res)
	}
}
