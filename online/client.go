package online

import (
	"fmt"
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
	pingPeriod := (pongWait * 9) / 20

	u.Connection.SetReadLimit(int64(w.Env.MaxMessageSize))
	u.Connection.SetReadDeadline(time.Now().Add(pongWait))
	u.Connection.SetPongHandler(func(string) error { ; u.Connection.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	u.Connection.EnableWriteCompression(false)

	pingTicker := time.NewTicker(pingPeriod)

	go u.aliveConection(pingTicker)

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

func (u *Client) aliveConection(pingTicker *time.Ticker) {

	for {

		<-pingTicker.C
		if err := u.sendMsgMutex([]byte{}, 0); err != nil {
			fmt.Println("ping: ", err)
		}

	}

}

func (u *Client) sendMsgMutex(msg []byte, pingmsg int) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	if pingmsg == 0 {

		err := u.Connection.WriteMessage(websocket.PingMessage, []byte{})
		if err != nil {
			return err
		}

	} else {

		err := u.Connection.WriteMessage(websocket.BinaryMessage, msg)
		if err != nil {
			return err
		}

	}

	return nil
}
