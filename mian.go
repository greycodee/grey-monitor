package main

import (
	"fmt"
	"grey-monitor/api"
	"grey-monitor/ws"
	"html/template"
	"net/http"
	"os"
	"strings"
)

type handRequest struct {

}

func main() {
	//获取命令行参数
	port:=os.Args[1]
	fmt.Printf(port)
	if port==""{
		// 默认端口
		port=":8989"
	}else {
		port=":"+port
	}
	serverStart(port)
}

func serverStart(addr string)  {
	// 接口路由处理
	http.Handle("/",distribute())

	staticHandle := http.FileServer(assetFS())
	//将/js/路径下的请求匹配到 ./views/js/下
	http.Handle("/js/", staticHandle)

	// 开启http服务
	_=http.ListenAndServe(addr,nil)
}

func distribute()  http.Handler{
	return &handRequest{}
}

func (s *handRequest) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	upath := r.URL.Path
	hApi:=httpApi[upath]
	if hApi!=nil {
		stringSlice := strings.Split(upath,"/")
		if stringSlice!=nil && strings.Compare(stringSlice[1],"ws")==0{
			hApi(w,r,stringSlice[2])
		}else {

			hApi(w,r,"")
		}

	}else {
		return
	}
	fmt.Println(upath)
}

func index(w http.ResponseWriter, r *http.Request,e string)  {
	indexPage, _ :=Asset("views/index.html")
	t,_:=template.New("index").Parse(string(indexPage))
	_ = t.Execute(w, "")
}


var httpApi = map[string]func(w http.ResponseWriter, r *http.Request,e string){
	"/"				:		index,
	"/ws/mem"		:		ws.Client,
	"/ws/memPercent"		:		ws.Client,
	"/ws/cpuPercentSingle"	:		ws.Client,
	"/ws/cpuPercentAll"		:		ws.Client,
	"/api/cpuInfo"	:		api.CpuInfo,
	"/api/disk"		:		api.DiskInfo,
	"/api/diskPart"	:		api.DiskPart,
	"/api/diskPath" :		api.DiskInfoOfPath,

}