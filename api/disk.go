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

func DiskPart(w http.ResponseWriter, r *http.Request,e string)  {
	allDiskInfo,_:=disk.Partitions(true)
	j,_:=json.Marshal(allDiskInfo)
	_, _ = fmt.Fprint(w, string(j))
}

func DiskInfoOfPath(w http.ResponseWriter, r *http.Request,e string)  {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
		return
	}
	var p PathParam
	fmt.Printf("请求参数：%s",body)
	err =json.Unmarshal(body,&p)
	if err != nil {
		fmt.Printf("Unmarshal body err, %v\n", err)
		return
	}

	u,_:=disk.Usage(p.Path)
	j,_:=json.Marshal(u)
	header:=w.Header()
	header.Set("Content-Type","application/json")
	_, _ = w.Write(j)
}