package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	//"github.com/gorilla/websocket"
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
	defer ws.Close()
	utils.Clients[ws] = true
	for {
		var msg=make(map[string]interface{})          // Read in a new message as JSON and map it to a Message object
		
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(utils.Clients, ws)
			break
		}
		// Send the newly received message to the broadcast channel
		utils.Broadcast <- msg
	}
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
					client.Close()
					delete(utils.Clients, client)
				}
			}
		}
}