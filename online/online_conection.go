package online

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"server/config"
	user_domain "server/domain"
	"server/jwt"
	user_services "server/services"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type LeaveChan chan *Client
type JoinChan chan *Client

type Chanel struct {
	LeaveChanel chan *Client
	Join        chan *Client
}

type Hub struct {
	clients map[string]*Client
	chanel  *Chanel
	userDB  user_services.UserPort
	mu      sync.Mutex
	Env     *config.Env
}

func NewHub(Env *config.Env, user_DB user_services.UserPort) *Hub {
	return &Hub{
		clients: make(map[string]*Client),
		chanel: &Chanel{
			LeaveChanel: make(LeaveChan, 1000),
			Join:        make(JoinChan, 1000),
		},
		Env:    Env,
		userDB: user_DB,
	}
}

func check(r *http.Request) bool {
	log.Printf("%s %s%s %v", r.Method, r.Host, r.RequestURI, r.Proto)
	return r.Method == http.MethodGet
}

var upGradeWebSocket = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     check,
}

func (w *Hub) HandlerWebSocket(c *gin.Context) {

	TK := c.Query("token")

	conn, err := upGradeWebSocket.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	token, is := jwt.ExtractClaims(TK, w.Env)

	if !is {
		conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Error decoding token"))
	}

	strName := fmt.Sprint(token["name"])

	UserLooked, err := w.userDB.GetUser(&user_domain.Users{Username: strName})

	if err != nil {
		conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Error buscando usuario"))
		return
	}

	ClientUser := &Client{User: UserLooked, Connection: conn}

	w.chanel.Join <- ClientUser

	ClientUser.OnLine(w)
}

func (w *Hub) UsersManager() {
	for {
		select {
		case userClient := <-w.chanel.Join:
			if _, exists := w.clients[userClient.User.Username]; exists {
				userClient.Connection.Close()
			} else {
				w.mu.Lock()
				w.clients[userClient.User.Username] = userClient
				fmt.Println(w.clients)
				w.mu.Unlock()
			}
		case user := <-w.chanel.LeaveChanel:
			w.DisconnectUser(user.User.Username)
		}
	}
}

func (w *Hub) DisconnectUser(username string) {
	w.mu.Lock()
	if user, ok := w.clients[username]; ok {
		defer user.Connection.Close()
		delete(w.clients, username)
	}
	w.mu.Unlock()
}
