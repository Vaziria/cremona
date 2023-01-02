package backend

import (
	"net/http"
	"path/filepath"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

var locProfile = "data/profile/"

type Browser struct {
	ProxyAddr string
	Service   *selenium.Service
}

func (b *Browser) CreateDriver(profile string) selenium.WebDriver {

	proxy := "--proxy-server=" + b.ProxyAddr
	pathAbs, _ := filepath.Abs(profile)
	profileArg := "--user-data-dir=" + pathAbs

	caps := selenium.Capabilities{}
	caps.AddChrome(chrome.Capabilities{
		Args: []string{
			proxy,
			profileArg,
		},
	})

	driver, err := selenium.NewRemote(caps, "")

	if err != nil {
		panic(err)
	}

	return driver
}

func NewBrowser(proxy string) *Browser {
	service, err := selenium.NewChromeDriverService("./chromedriver", 4444)

	if err != nil {
		panic(err)
	}

	browser := &Browser{
		ProxyAddr: proxy,
		Service:   service,
	}

	return browser
}

func getHttpCookies(driver selenium.WebDriver) []http.Cookie {

	seleCookies, err := driver.GetCookies()

	if err != nil {
		panic("error getting cookies")
	}

	httpCookies := make([]http.Cookie, len(seleCookies))

	for key, cookie := range seleCookies {

		httpCookies[key] = http.Cookie{
			Name:  cookie.Name,
			Value: cookie.Value,
			Path:  cookie.Path,
		}
	}

	return httpCookies

}
