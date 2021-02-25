package api

import (
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/v3/disk"
	"io/ioutil"
	"net/http"
)

type PathParam struct {
	Path string `json:"path"`
}

func DiskInfo(w http.ResponseWriter, r *http.Request,e string)  {
	var result []*disk.UsageStat
	allDiskInfo,_:=disk.Partitions(true)
	var m string
	for value := range allDiskInfo{
		m=allDiskInfo[value].Mountpoint
		u,_:=disk.Usage(m)
		result=append(result, u)
	}
	j,_:=json.Marshal(result)
	_, _ = fmt.Fprint(w, string(j))
}

func DiskInfoOfPath(w http.ResponseWriter, r *http.Request,e string)  {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
		return
	}
	var p PathParam
	err =json.Unmarshal(body,&p)
	if err != nil {
		fmt.Printf("Unmarshal body err, %v\n", err)
		return
	}

	u,_:=disk.Usage(p.Path)
	j,_:=json.Marshal(u)
	_, _ = fmt.Fprint(w, string(j))
}