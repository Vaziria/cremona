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

func (akun *Account) createPing() []byte {

	data := &Request{
		Headers:        map[string]string{},
		AuthType:       2,
		DevicePlatform: "web",
		InboxType:      0,
		BuildNumber:    "12c929a:master",
		SdkVersion:     "0.3.8",
		Cmd:            203,
		Body: &RequestBody{
			MessagesPerUserInitV2Body: &MessagesPerUserInitV2RequestBody{
				Cursor: 0,
			},
		},
		Refer:      3,
		Token:      akun.Token,
		DeviceId:   akun.DeviceId,
		SequenceId: 10001,
	}

	payload, err := proto.Marshal(data)

	frame := &Frame{
		Seqid:       10001,
		Logid:       10001,
		Service:     10002,
		PayloadType: "pb",
		Method:      1,
		Payload:     payload,
	}

	if err != nil {
		panic("gagal create request websocket")
	}

	hasil, _ := proto.Marshal(frame)

	log.Println("creating ping", fmt.Sprintf("%x", hasil))

	return hasil
}

func (akun *Account) CreateWebsocket() {
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

			log.Println("message socket", dataframe.Payload)
		}
	}()

	ticker := time.NewTicker(10 * time.Second)
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

func StartWebsocketAll(akuns []*Account) {
	for _, value := range akuns {
		go value.CreateWebsocket()
	}
}
