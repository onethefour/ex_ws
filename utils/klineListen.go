package utils

import (
	"strings"
	"runtime"
	"fmt"
	"github.com/gorilla/websocket"
)

type KlineListen struct{
	Key string   //订阅的频道
	Data []string //已发送数据
	Clients map[*websocket.Conn]bool
}
func (this *KlineListen)Start(key string,clients map[*websocket.Conn]bool){
	this.Key = key
	this.Clients = clients
	this.init()
	//redis轮询

	for {
		Len :=Redis.LLen(key).Val()
		ret,_:=Redis.LIndex(key,Len-1).Result()

		if ret != this.Data[len(this.Data)-1] && ret != ""{
			fmt.Println("发现新数据",ret)
			this.broadcast() //向订阅客户端发送消息
		}
		runtime.Gosched()
	}
}

//初始化
func (this *KlineListen)init(){
	var key = this.Key
	Len :=Redis.LLen(key).Val()
	ret,_:=Redis.LRange(key,0,Len-1).Result()
	//fmt.Println(ret)
	this.Data = ret
}
//广播订阅消息
func (this *KlineListen)broadcast(){
	Len :=Redis.LLen(this.Key).Val()
	var offset int64
	for i:=Len-2;i>=0;i--{
		ret,_:=Redis.LIndex(this.Key,i).Result()
		if ret == this.Data[len(this.Data)-1]{
			offset = i
			break
		}
	}
	ret,_:=Redis.LRange(this.Key,offset+1,Len-1).Result()
	this.sendAll(ret)
	this.init()
}

func (this *KlineListen)sendAll(msg []string){
	data := make(map[string]interface{})

	data["channel"]=this.Key
	data["error"]="0"
	data["msg"]=""
	data["data"]=msg
	str :=`{"error":"`+data["error"].(string)+`","msg":"`+data["msg"].(string)+`","channel":"`+data["channel"].(string)+`","data":[`+strings.Join(msg,",")+`]}`
	text:= []byte(str)
	for client,_:=range(this.Clients){
		if _,ok:=Clients[client];!ok{ //连接已关闭要删除链接
			continue
		}
		client.WriteMessage(websocket.TextMessage,text)
	}
}

func (this *KlineListen)Newconnect(ws *websocket.Conn){
	data := make(map[string]interface{})

	data["channel"]=this.Key
	data["error"]="0"
	data["msg"]=""
	data["data"]=this.Data

	str :=`{"error":"`+data["error"].(string)+`","msg":"`+data["msg"].(string)+`","channel":"`+data["channel"].(string)+`","data":[`+strings.Join(this.Data,",")+`]}`
	//ws.WriteJSON(str)
	ws.WriteMessage(websocket.TextMessage,[]byte(str))
}