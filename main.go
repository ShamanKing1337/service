package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/websocket"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type Message struct {
	Messagetype int
	Message  []byte
}

type ViewData struct{

	User string
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func homePage(w http.ResponseWriter, r * http.Request){


	data := ViewData{
		User: "",
	}

	tmpl := template.Must(template.ParseFiles("index.html"))

	tmpl.Execute(w, data)

}


func reader(conn * websocket.Conn){
	for{
		messageType, p, err := conn.ReadMessage()
		if err != nil{
			log.Println(err)
			return
		}
		log.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil{
			log.Println(err)
			return
		}

	}
}

func wsEndpoint(w http.ResponseWriter, r * http.Request){


	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil{
		log.Print(err)
	}

	log.Println("Client connected")

	defer ws.Close()
	clients[ws] = true
	for {
		var msg Message

		messageType, p, err := ws.ReadMessage()
		msg.Messagetype = messageType
		msg.Message = p

		connStr := "user=postgres password=zopa123 dbname=postgres sslmode=disable host=localhost"
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			panic(err)
		}
		defer db.Close()
		fmt.Println(p)
		fmt.Println(string(p))
		str := strings.Split(string(p), " ")
		name := str[0]
		surname:= strings.Split(str[1], ":")[0]
		message := str[2]

		rows,err := db.Exec("INSERT INTO allmessages (id_sender, message) VALUES((SELECT id from patient where first_name = $1 and last_name = $2), $3) ", name, surname, message)
		if err != nil{
			fmt.Println(rows)
			panic(err)
		}

		if err != nil {
			log.Printf("errorpopa: %v", err)
			delete(clients, ws)
			break
		}

		broadcast <- msg
	}

}

func handleMessages() {
	for {

		msg := <-broadcast

		for client := range clients {
			err := client.WriteMessage(msg.Messagetype, msg.Message)
			if err != nil {
				log.Printf("errorjopa: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func setupRoutes(){
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}

func main(){
	fmt.Println("Go WebSockets")
	setupRoutes()
	go handleMessages()
	log.Fatal(http.ListenAndServe(":8080", nil))
}