package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	"newbtc-ws/utils"
	"github.com/gorilla/websocket"
	"fmt"
)

type WebsocketController struct {
}


func (this *WebsocketController)Router(r *gin.Engine){
		r.GET("/websocket", this.Websocket)
		//go this.broadcast()

}

func (this *WebsocketController)Websocket(ctx *gin.Context){

	this.websocketHander(ctx.Writer,ctx.Request)
}

func (this *WebsocketController)websocketHander(w http.ResponseWriter, r *http.Request){

	ws, err := utils.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf(err.Error())
		return
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

		err:=ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("onMessageerror: %v", err)
			break
		}
		event,ok:=msg["event"]
		fmt.Println(event.(string) )
		if !ok {
			continue
		}

		switch event.(string) {

		case "login":
			this.login(ws,msg["params"].(map[string]interface{}))
		case "addChannel":
			this.addChannel(ws,msg["params"].(map[string]interface{}))
		case "removeChannel":
			this.removeChannel(ws,msg["params"].(map[string]interface{}))
		default:
			//utils.Broadcast <- msg
		}

	}
}

func (this *WebsocketController)login(ws *websocket.Conn,params map[string]interface{}){

}
func (this *WebsocketController)addChannel(ws *websocket.Conn,params map[string]interface{}){
	chname,ok :=params["channel"]
	if !ok{
		return
	}
	Channel,has:=utils.Channelist[chname.(string)]
	if !has{
		return
	}
	Channel.NewClient(ws)
}
func (this *WebsocketController)removeChannel(ws *websocket.Conn,params map[string]interface{}){
	chname,ok :=params["channel"]
	if !ok{
		return
	}
	Channel,has:=utils.Channelist[chname.(string)]
	if !has{
		return
	}
	Channel.RemoveClient(ws)
}
func (this *WebsocketController)broadcast(){
	for {
		// Grab the next message from the broadcast channel
		msg := <-utils.Broadcast
			// Send it out to every client that is currently connected
		for client := range utils.Clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("broadcasterror: %v", err)
			}
		}
	}
}