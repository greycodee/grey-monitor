package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/shirou/gopsutil/v3/mem"
	"log"
	"strconv"
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

func MemPercent(c *websocket.Conn,err error)  {

	for {
		sleep()
		v, _ := mem.VirtualMemory()
		memPercent:=float64(v.Used)/float64(v.Total)*100
		m, _ := strconv.ParseFloat(fmt.Sprintf("%.0f", memPercent), 64)
		err=c.WriteJSON(m)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}


