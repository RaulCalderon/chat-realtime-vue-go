package main

import (
	"log"
	"net/http"
	"github.com/googollee/go-socket.io"
)

func main() {

	// Initialize a new Server
	server := socketio.NewServer(nil)

	// If there is an error, show it
	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalf("SocketIO server error: %v\n", err)
		}
	}()
	defer server.Close()

	// Web Sockets
	// If there is a connection, it is saved on 'so' variable
	server.OnConnect("/", func(so socketio.Conn) error {
		so.SetContext("")
		log.Println("A new user connected:", so.ID())
		return nil
	})

	// Everytime a new user has enter in the Room
	server.OnEvent("/", "join room", func(so socketio.Conn, room string) {
		so.Join(room)
		log.Printf("User %s joined the room: %s", so.ID(), room)
	})


	// Listen to 'chat message' event from client on sendMessage()
	server.OnEvent("/", "chat message", func(so socketio.Conn, msg string) {
		
		// Get messages form server
		log.Printf("Message from %s in room %s: %s", so.ID(), "chat_room", msg)

		// Send message to all user in the room (except the emissor)
		server.BroadcastToRoom("/", "chat_room", "broad message", so.ID()+": "+msg)
	})

	

	// HTTP

	http.Handle("/socket.io/", server)
	// Send all resources to the explorer (html, images, etc.)
	http.Handle("/", http.FileServer(http.Dir("./public")))
	// Show custom message
	log.Println("Server on port 3000")
	// Make the server listen to a specific port
	log.Fatal(http.ListenAndServe(":3000", nil))

}

