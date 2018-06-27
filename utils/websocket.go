package utils

import "github.com/gorilla/websocket"

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var Clients = make(map[*websocket.Conn]bool)

var Broadcast = make(chan map[string]interface{})           // broadcast channel