package backend

import (
	"encoding/json"
	"log"
	"os"
)

type ChatData struct {
	ChatUnread int32 `json:"chat_unread"`
}

type Account struct {
	Name     string   `json:"name"`
	ChatData ChatData `json:"chat_data"`
	// Cookies  []http.Cookie `json:"cookies"`
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
