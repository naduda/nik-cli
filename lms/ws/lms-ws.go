package ws

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/sacOO7/gowebsocket"
	"log"
	"net/http"
	"nik-cli/lms/ws/protoschema"
	"nik-cli/logger"
	"os"
	"os/signal"
	"strings"
	"time"
)

type LmsWebsocket struct {
	Cookies      []*http.Cookie
	CookieString string
	Logger       *log.Logger
}

func NewInstance(cookies []*http.Cookie) *LmsWebsocket {
	cookieString := "SL_GWPT_Show_Hide_tmp=1; SL_wptGlobTipTmp=1; "
	for _, c := range cookies {
		idx := strings.Index(c.String(), ";")
		t := c.String()[0:idx]
		cookieString += t + ";"
	}
	logfile := fmt.Sprintf("./ws_%s.log", time.Now().Format("2006_01_02"))
	lmsLogger, err := logger.InitLogger(logfile)
	if err != nil {
		log.Fatal(err)
	}
	return &LmsWebsocket{
		Cookies:      cookies,
		CookieString: cookieString,
		Logger:       lmsLogger,
	}
}

func (r *LmsWebsocket) Connect() error {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	//var u utils.NikUtils

	socket := gowebsocket.New("wss://lms.ua.energy/ws/lms/subscribe/")
	socket.ConnectionOptions = gowebsocket.ConnectionOptions{
		UseSSL:         true,
		UseCompression: false,
	}
	socket.RequestHeader.Set("Accept-Encoding", "gzip, deflate, br")
	socket.RequestHeader.Set("Accept-Language", "uk,ru;q=0.9,en-US;q=0.8,en;q=0.7,de;q=0.6")
	socket.RequestHeader.Set("Pragma", "no-cache")
	socket.RequestHeader.Set("Cookie", r.CookieString)

	socket.OnConnected = func(socket gowebsocket.Socket) {
		log.Println("Connected to server")
	}

	socket.OnConnectError = func(err error, socket gowebsocket.Socket) {
		log.Println("Recieved connect error ", err)
	}

	socket.OnBinaryMessage = func(data []byte, socket gowebsocket.Socket) {
		m := &protoschema.Message{}
		if err := proto.Unmarshal(data, m); err != nil {
			log.Println(err)
			return
		}
		r.Logger.Println(m.GetType())
		var p proto.Message = nil
		switch m.GetType() {
		case protoschema.Message_HEARTBEAT:
			p = &protoschema.Heartbeat{}
		case protoschema.Message_ACTIVATION,
			protoschema.Message_PROPORTIONAL_ACTIVATION:
			p = &protoschema.Activation{}
		case protoschema.Message_ACTIVATION_AUDIT:
			p = &protoschema.ActivationAudit{}
		case protoschema.Message_MARKET_AREA_SUMMARY:
			p = &protoschema.MarketAreaSummary{}
		default:
			r.Logger.Println("Not implement type:", m.GetType())
		}
		if p != nil {
			r.logMessage(m.GetPayload(), p)
		}
	}

	socket.OnDisconnected = func(err error, socket gowebsocket.Socket) {
		log.Println("Disconnected from server ")
		return
	}

	socket.Connect()

	for {
		select {
		case <-interrupt:
			socket.Close()
			return nil
		}
	}
}

func (r *LmsWebsocket) logMessage(payload []byte, m proto.Message) {
	if err := proto.Unmarshal(payload, m); err != nil {
		r.Logger.Println(err)
	}
	jsn, err := json.Marshal(m)
	if err != nil {
		r.Logger.Println(err)
	}
	r.Logger.Println(string(jsn))
}
