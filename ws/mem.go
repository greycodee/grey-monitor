package ws

import (
	"github.com/gorilla/websocket"
	"github.com/shirou/gopsutil/v3/mem"
	"log"
)


func MemInfo(c *websocket.Conn,err error)  {

	for {
		sleep()
		v, _ := mem.VirtualMemory()
		err = c.WriteMessage(1, []byte(v.String()))
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}


