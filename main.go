package main

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"realtime_chat/pkg/httpserver"
	"realtime_chat/pkg/ws"
)

func init() {
	// load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("unable to load .env", err)
	}
}

// TODO: implement client
func main() {
	server := flag.String("server", "", "http,websocket")
	flag.Parse()

	if *server == "http" {
		fmt.Println("http server is starting on :8080")
		httpserver.StartHTTPServer()
	} else if *server == "websocket" {
		fmt.Println("websocket server is starting on :8081")
		ws.StartWebSocketServer()
	} else {
		fmt.Println("invalid server")
	}
}
