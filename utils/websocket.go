package utils

import (
	"github.com/gorilla/websocket"
	"net/http"
	)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin :func(r *http.Request) bool {
		// allow all connections by default
		return true
	},
}

var Clients = make(map[*websocket.Conn]bool)  //all clients

var Broadcast = make(chan map[string]interface{})           // broadcast channel

