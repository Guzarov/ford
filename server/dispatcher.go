package main

import (
	"log"
)

type DispatchTo int

const (
	ToClient DispatchTo = iota
	ToRoom
)

type Dispatcher struct {
	model      *Model
	clients    map[string]*ClientSession
	connect    chan *ClientSession
	fromClient chan *Message
}

func NewDispatcher(model *Model) (*Dispatcher, error) {
	return &Dispatcher{
		model:      model,
		clients:    make(map[string]*ClientSession),
		connect:    make(chan *ClientSession, 10),
		fromClient: make(chan *Message, 10),
	}, nil
}

func (d *Dispatcher) Start() {
	go func() {
		for {
			select {
			case client := <-d.connect:
				d.clients[client.clientId] = client
				log.Println("ClientSession connected", client)
			case msg := <-d.fromClient:
				client, founded := d.clients[msg.clientId] //todo
				if !founded {
					log.Printf("Cannot find user by id %s", client.clientId)
					break
				}
				target, message, err := d.model.HandleMessage(client.roomId, msg.payload)
				if err != nil {
					log.Println("Error while process message", err)
					break
				}
				if target == ToClient {
					client.mailbox <- message
				} else if target == ToRoom {
					for _, c := range d.clients {
						if c.roomId == client.roomId {
							c.mailbox <- message
						}
					}
				} else {
					// to all
				}
			}
		}
	}()
}
