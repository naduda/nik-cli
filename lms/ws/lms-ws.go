package ws

import (
	"github.com/sacOO7/gowebsocket"
	"log"
	"net/http"
	"os"
	"os/signal"
)

type LmsWebsocket struct {
	Cookies      []*http.Cookie
	CookieString string
}

func NewInstance(cookies []*http.Cookie) *LmsWebsocket {
	var cookieString string
	for _, c := range cookies {
		cookieString += c.String() + ";"
	}
	return &LmsWebsocket{
		Cookies:      cookies,
		CookieString: cookieString,
	}
}

func (r *LmsWebsocket) Connect() error {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	socket := gowebsocket.New("wss://lms.ua.energy/ws/subscribe/")
	socket.ConnectionOptions = gowebsocket.ConnectionOptions{
		UseSSL:         true,
		UseCompression: false,
	}
	socket.RequestHeader.Set("Accept-Encoding", "gzip, deflate, br")
	socket.RequestHeader.Set("Accept-Language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7,uk;q=0.6")
	socket.RequestHeader.Set("Pragma", "no-cache")
	socket.RequestHeader.Set("Cookie", r.CookieString)

	socket.OnConnected = func(socket gowebsocket.Socket) {
		log.Println("Connected to server")
	}

	socket.OnConnectError = func(err error, socket gowebsocket.Socket) {
		log.Println("Recieved connect error ", err)
	}

	socket.OnTextMessage = func(message string, socket gowebsocket.Socket) {
		log.Println("Recieved message " + message)
	}

	socket.OnBinaryMessage = func(data []byte, socket gowebsocket.Socket) {
		log.Println("Recieved binary data ", data)
	}

	socket.OnPingReceived = func(data string, socket gowebsocket.Socket) {
		log.Println("Recieved ping " + data)
	}

	socket.OnPongReceived = func(data string, socket gowebsocket.Socket) {
		log.Println("Recieved pong " + data)
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
