package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	"newbtc-ws/utils"
)

type WebsocketController struct {
}


func (this *WebsocketController)Router(r *gin.Engine){
	r.GET("/websocket", this.Websocket)
	go this.broadcast()
}

func (this *WebsocketController)Websocket(ctx *gin.Context){
	this.websocketHander(ctx.Writer,ctx.Request)
}
func (this *WebsocketController)websocketHander(w http.ResponseWriter, r *http.Request){

	ws, err := utils.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer this.onClose(ws)

	this.onConnect(ws)
	this.onMessage(ws)
}
func (this *WebsocketController)onConnect(ws *websocket.Conn){
	utils.Clients[ws] = true
}

func (this *WebsocketController)onClose(ws *websocket.Conn){
	delete(utils.Clients, ws)
	ws.Close()
}

func (this *WebsocketController)onMessage(ws *websocket.Conn){
	for {
		var msg=make(map[string]interface{})          // Read in a new message as JSON and map it to a Message object

		ws.ReadJSON(&msg)
		//if err != nil {
		//	log.Printf("error: %v", err)
		//	continue
		//
		//}
		event,_:=msg["event"]
		//if !ok {
		//	log.Printf("error: event empty")
		//	continue
		//}
		switch event.(string) {
		case "login":
			this.login(msg["params"].(map[string]interface{}))
		case "addChannel":
			this.addChannel(msg["params"].(map[string]interface{}))
		case "removeChannel":
			this.removeChannel(msg["params"].(map[string]interface{}))
		default:
			utils.Broadcast <- msg
		}
	}
}

func (this *WebsocketController)login(params map[string]interface{}){

}
func (this *WebsocketController)addChannel(params map[string]interface{}){

}
func (this *WebsocketController)removeChannel(params map[string]interface{}){

}
func (this *WebsocketController)broadcast(){
	for {
		// Grab the next message from the broadcast channel
		msg := <-utils.Broadcast
			// Send it out to every client that is currently connected
		for client := range utils.Clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				this.onClose(client)
			}
		}
	}
}