package ws

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/shirou/gopsutil/v3/cpu"
	"log"
)

func CpuInfo(c *websocket.Conn,err error)  {
	for {
		sleep()
		cpuInfo, _ := cpu.Info()
		j,_:=json.Marshal(cpuInfo)
		err = c.WriteMessage(1, j)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func CpuPercentSingle(c *websocket.Conn,err error)  {
	for{
		sleep()
		info,_:=cpu.Percent(0,false)
		j,_:=json.Marshal(info)
		_ = c.WriteMessage(1, j)
	}
}

func CpuPercentAll(c *websocket.Conn,err error)  {
	for{
		sleep()
		info,_:=cpu.Percent(0,true)
		j,_:=json.Marshal(info)
		_ = c.WriteMessage(1, j)
	}
}