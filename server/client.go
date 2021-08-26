package main

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var (
	upgrader = websocket.Upgrader{
		//ReadBufferSize:  1024,
		//WriteBufferSize: 1024,
	}
)

type Message struct {
	clientId string
	payload  []byte
}

type ClientSession struct {
	roomId     string
	clientId   string
	ws         *websocket.Conn
	dispatcher *Dispatcher
	mailbox    chan []byte
}

func NewSession(writer http.ResponseWriter, request *http.Request, dispatcher *Dispatcher) (*ClientSession, error) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			log.Println(err)
		}
		return nil, err
	}
	roomId := request.RequestURI //todo
	log.Println("New client in", roomId)
	return &ClientSession{
		roomId:     roomId,
		clientId:   uuid.NewString(),
		ws:         ws,
		dispatcher: dispatcher,
		mailbox:    make(chan []byte, 10),
	}, nil
}

func (c *ClientSession) Start() {
	c.dispatcher.connect <- c
	go c.read()
	go c.write()
}

func (c *ClientSession) read() {
	defer func() {
		_ = c.ws.Close()
	}()
	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			log.Println("Cannot read from client", err)
			break
		} else {
			c.dispatcher.fromClient <- &Message{
				clientId: c.clientId,
				payload:  message,
			}
		}
	}
}

func (c *ClientSession) write() {
	defer func() {
		_ = c.ws.Close()
	}()
	for {
		select {
		case message := <-c.mailbox:
			w, err := c.ws.NextWriter(websocket.TextMessage)
			if err != nil {
				log.Println("Cannot get writer", err)
				return
			}
			_, err = w.Write(message)
			if err != nil {
				log.Print("Cannot write message", err)
			}
			_ = w.Close()
		}
	}

}
