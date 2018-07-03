package utils

import (
	"github.com/gorilla/websocket"
	"fmt"
)
//
//频道列表
var Channelist = make(map[string]*Channel)

//频道后端处理接口
type channelbHander interface {
	Start(key string,clients  map[*websocket.Conn]bool)     //启动
	Newconnect(ws *websocket.Conn) //新连接时处理
}

//频道信息
type Channel struct{
	Key string
	Clients map[*websocket.Conn]bool
	Hander channelbHander
}

func (this *Channel)NewClient(ws *websocket.Conn){
	fmt.Println("NEW CLIENT")
	this.Clients[ws] =true
	this.Hander.Newconnect(ws)
}
func (this *Channel)RemoveClient(ws *websocket.Conn){
	delete(this.Clients,ws)
}

func init(){

	Channelist["btc_usd_1min"]= &Channel{Key:"btc_usd_1min",Clients:make(map[*websocket.Conn]bool),Hander:new(KlineListen)}
	go Channelist["btc_usd_1min"].Hander.Start("btc_usd_1min",Channelist["btc_usd_1min"].Clients)

}