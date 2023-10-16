package online

import (
	"log"
	"sync"
	"time"

	user_domain "server/domain"

	"github.com/gorilla/websocket"
)

type Client struct {
	User       *user_domain.Users
	Connection *websocket.Conn
	mu         sync.Mutex
}

func (u *Client) OnLine(w *Hub) {

	pongWait := time.Duration(w.Env.PongWait) * time.Second
	pingPeriod := (pongWait * 9) / 15

	u.Connection.SetReadLimit(int64(w.Env.MaxMessageSize))
	u.Connection.SetReadDeadline(time.Now().Add(pongWait))
	u.Connection.SetPongHandler(func(string) error { ; u.Connection.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	u.Connection.EnableWriteCompression(false)

	pingTicker := time.NewTicker(pingPeriod)

	//go u.aliveConection(pingTicker)

	for {
		_, _, err := u.Connection.ReadMessage()

		if err != nil {
			log.Println("Error on read message: ", err.Error())
			pingTicker.Stop()
			//u.mu.Lock()
			u.Connection.Close()
			//u.mu.Unlock()

			break
		} else {

			u.mu.Lock()
			u.Connection.SetReadDeadline(time.Now().Add(pongWait))
			u.mu.Unlock()

			err = u.sendMsgMutex([]byte("holaaa bebeeee"), 1)

			if err != nil {
				break
			}
		}

	}

	w.chanel.LeaveChanel <- u
}
