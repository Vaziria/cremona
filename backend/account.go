package backend

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

type ChatData struct {
	ChatUnread int32 `json:"chat_unread"`
}

type ACookie struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Account struct {
	Name     string    `json:"name"`
	ChatData ChatData  `json:"chat_data"`
	Cookies  []ACookie `json:"cookies"`
	Username string    `json:"username"`
	DeviceId string    `json:"device_id"`
	Token    string    `json:"tokenid"`
}

func NewFromJson(pathname string) Account {

	var akun Account

	file, _ := os.ReadFile(pathname)

	err := json.Unmarshal(file, &akun)

	if err != nil {
		panic(err)
	}

	return akun
}

func (ac *Account) Save() {
	data, err := json.MarshalIndent(ac, "", "    ")

	if err != nil {
		log.Fatal(err)
	}

	fname := BaseAccountDir + ac.Name + ".json"
	err = os.WriteFile(fname, data, 0644)

	if err != nil {
		log.Fatal(err)
	}

}

func (ac *Account) Delete() {
	fname := BaseAccountDir + ac.Name + ".json"
	e := os.Remove(fname)
	if e != nil {
		log.Fatal(e)
	}
}

func (ac *Account) GetCookies(browser *Browser) {
	profile := locProfile + ac.Name
	driver := browser.CreateDriver(profile)
	defer driver.Close()

	driver.Get("https://seller-id.tiktok.com/chat")
	time.Sleep(5 * time.Second)

	token := <-TokenChan
	log.Println(token)

	ac.Token = token.Token
	ac.DeviceId = token.PigeonCid

	cookies := getHttpCookies(driver)
	ac.Cookies = make([]ACookie, len(cookies))

	for key, value := range cookies {
		ac.Cookies[key] = ACookie{
			Name:  value.Name,
			Value: value.Value,
		}
	}
}
