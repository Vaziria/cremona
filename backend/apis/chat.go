package apis

import (
	"bytes"
	"changeme/backend"
	"io"
	"log"
	"net/http"
)

var defaultHeader = map[string][]string{
	"Accept":             {"*/*"},
	"Content-Type":       {"application/x-protobuf"},
	"Accept-Language":    {"en-US,en;q=0.9"},
	"Referer":            {"https://seller-id.tiktok.com/chat"},
	"Sec-ch-ua":          {"\"Not?A_Brand\";v=\"8\", \"Chromium\";v=\"108\", \"Google Chrome\";v=\"108\""},
	"Sec-ch-ua-mobile":   {"?0"},
	"Sec-ch-ua-platform": {"\"Windows\""},
	"Sec-Fetch-Dest":     {"empty"},
	"Sec-Fetch-Mode":     {"cors"},
	"Sec-Fetch-Site":     {"same-origin"},
	"User-Agent":         {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36"},
}

type ChatApi struct {
	akun *backend.Account
}

func NewChatApi(akun *backend.Account) *ChatApi {
	return &ChatApi{
		akun: akun,
	}
}

func (api *ChatApi) UserInit() {

	data := api.akun.CreateV2Init()

	url := "https://imapi-va-oth.isnssdk.com/v2/message/get_by_user_init"

	res, err := http.NewRequest("POST", url, bytes.NewReader(data))
	for _, value := range api.akun.Cookies {
		res.AddCookie(&http.Cookie{
			Name:  value.Name,
			Value: value.Value,
		})
	}

	if err != nil {
		log.Panicln(err, "gagal init chat")
	}

	b, _ := io.ReadAll(res.Body)

	backend.DecodeProto(b)

	log.Println(data, res)
}
