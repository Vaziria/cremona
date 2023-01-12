package backend

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/gorilla/schema"
	"github.com/gorilla/websocket"
	"github.com/pierrec/lz4/v4"
	"google.golang.org/protobuf/proto"
)

var encoder = schema.NewEncoder()
var appkey = "b42d99769353ce6304e74fb597e36e90"

func cookieString(cookies []ACookie) string {
	var hasil string

	for _, value := range cookies {
		hasil += value.Name + "=" + value.Value + ";"
	}

	return hasil

}

type SocketQuery struct {
	Token                 string `schema:"token"`
	Aid                   int32  `schema:"aid"`
	Fpid                  int32  `schema:"fpid"`
	DeviceId              string `schema:"device_id"`
	AccessKey             string `schema:"access_key"`
	DevicePlatform        string `schema:"device_platform"`
	VersionCode           int32  `schema:"version_code"`
	WebsocketSwitchRegion string `schema:"websocket_switch_region"`
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func NewSocketQuery() SocketQuery {
	return SocketQuery{
		WebsocketSwitchRegion: "ID",
		Aid:                   5341,
		Fpid:                  92,
		DevicePlatform:        "web",
		VersionCode:           10000,
	}
}

func (akun *Account) createQuery() *SocketQuery {

	query := NewSocketQuery()

	// shopInfo := akun.GetShopAndCsInfo()
	// tokeninfo := akun.GetTokenInfo(shopInfo.CustomerServiceInfo.OuterCid)

	// query.Token = tokeninfo.Token
	// query.DeviceId = tokeninfo.PigeonCid

	query.Token = akun.Token
	query.DeviceId = akun.DeviceId

	contentHash := strconv.Itoa(int(query.Fpid)) + appkey + akun.DeviceId + "f8a69f1719916z"
	query.AccessKey = GetMD5Hash(contentHash)

	return &query

}

func (akun *Account) createUrl() *url.URL {
	query := akun.createQuery()
	q := url.Values{}

	err := encoder.Encode(query, q)

	if err != nil {
		fmt.Printf("error = %s \n", err)
	}

	u := url.URL{
		Scheme: "wss",
		Host:   "oec-im-frontier-va.tiktokglobalshop.com",
		Path:   "/ws/v2",
	}

	u.RawQuery = q.Encode()

	return &u
}

func (akun *Account) CreateWebsocket(reschan chan<- *Response) {
	u := akun.createUrl()

	requester := http.Header{
		"User-Agent": {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36"},
		"Origin":     {"https://seller-id.tiktok.com"},
		"Cookie":     {cookieString(akun.Cookies)},
	}

	c, res, err := websocket.DefaultDialer.Dial(u.String(), requester)

	log.Println(res.StatusCode)

	if err != nil {
		log.Fatal("dial:", err)
	}

	log.Println("create connection websocket")

	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		log.Println("running reader")
		for {
			_, message, err := c.ReadMessage()

			if err != nil {
				log.Println("read socket error:", err)
				return
			}

			var dataframe Frame
			err = proto.Unmarshal(message, &dataframe)

			if err != nil {
				log.Println("parse dataframe error:", err)
				return
			}

			truePayload := make([]byte, len(dataframe.Payload)*10)

			size, err := lz4.UncompressBlock(dataframe.Payload, truePayload)
			if err != nil {

				for _, b := range truePayload {
					fmt.Printf("%02x ", b)
				}
				fmt.Println()

				panic(err)
			}

			// fmt.Println("\ndecompressed Data:", truePayload)

			// for _, b := range dataframe.Payload {
			// 	fmt.Printf("%02x ", b)
			// }
			// fmt.Println()

			// log.Println("payload", dataframe.PayloadType, dataframe.PayloadEncoding)

			var response Response

			err = proto.Unmarshal(truePayload[:size], &response)

			if err != nil {
				log.Println("parse response payload error:", err)
				return
			}

			reschan <- &response
			akun.handleRes(&response)
			log.Println("message socket", &response)
		}
	}()

	akun.CreateV2Init()

	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return

		case <-ticker.C:
			log.Println("sending ping")

			ping := akun.createPing()

			err = c.WriteMessage(websocket.BinaryMessage, ping)
			if err != nil {
				log.Println("write:", err)
				return
			}

		}
	}
}

func (akun *Account) handleRes(res *Response) {
	switch res.Cmd {
	case 203:
		akun.CursorChat = res.Body.MessagesPerUserInitV2Body.PerUserCursor
		// log.Println(res.Body.MessagesPerUserInitV2Body.PerUserCursor)
	}
}

func StartWebsocketAll(akuns []*Account, reschan chan<- *Response) {
	for _, value := range akuns {
		go value.CreateWebsocket(reschan)
	}
}
