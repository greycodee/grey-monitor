package ws

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     checkOrigin,
}

func checkOrigin(r *http.Request) bool {
	return true
}

func Client(w http.ResponseWriter, r *http.Request, e string){
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	event := wsEvent[e]
	event(c,err)
}

func sleep()  {
	time.Sleep(1*1e9)
}

var wsEvent = map[string] func(c *websocket.Conn,err error){
	"cpu"		:		CpuInfo,
	"mem"		:		MemInfo,
	"memPercent":		MemPercent,
	"cpuPercentSingle": 	CpuPercentSingle,
	"cpuPercentAll"	  : 	CpuPercentAll,
}
