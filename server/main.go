package main

import (
	"log"
	"net/http"
)

func main() {
	//storage, err := NewStorage(os.Getenv("DB_URL"), os.Getenv("DB_NAME"))
	storage, err := NewStorage("mongodb://localhost:27017", "ford")
	if err != nil {
		log.Fatal("Cannot init storage", err)
	}
	log.Println("Database found")
	model, err := NewModel(storage)
	if err != nil {
		log.Fatal("Cannot create model", err)
	}
	log.Println("Model created")

	dispatcher, err := NewDispatcher(model)
	if err != nil {
		log.Fatal("Cannot init dispatcher", err)
	}
	log.Println("Dispatcher created")

	dispatcher.Start()

	http.HandleFunc("/ws/v1/", func(writer http.ResponseWriter, request *http.Request) {
		client, err := NewSession(writer, request, dispatcher)
		if err != nil {
			log.Print("Cannot init client session", err)
		}
		client.Start()
	})

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
