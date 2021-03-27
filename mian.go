package main

import (
	"embed"
	"flag"
	"fmt"
	"github.com/greycodee/grey-monitor/api"
	"github.com/greycodee/grey-monitor/ws"
	"io/fs"
	"net/http"
	"strings"
)

type handRequest struct {

}
var port = flag.String("p", "8989", "http端口")
//go:embed views
var htmlFile embed.FS

func main() {
	//获取命令行参数
	flag.Parse()
	p:=":"+*port
	fmt.Printf(p)
	serverStart(p)
}

func serverStart(addr string)  {


	fsys, _ := fs.Sub(htmlFile, "views")
	staticHandle := http.FileServer(http.FS(fsys))
	//将/js/路径下的请求匹配到 ./views/js/下
	// 接口路由处理
	http.Handle("/",staticHandle)
	http.Handle("/js/", staticHandle)
	http.Handle("/archive/", staticHandle)
	http.Handle("/service/",distribute())

	fmt.Println("http服务器端口："+addr)
	// 开启http服务
	e:=http.ListenAndServe(addr,nil)
	if e!=nil {
		fmt.Println(e)
	}
}


func distribute()  http.Handler{
	return &handRequest{}
}

/*
	http路由
*/
func (s *handRequest) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	upath := r.URL.Path
	hApi:=httpApi[upath]
	if hApi!=nil {
		stringSlice := strings.Split(upath,"/")
		if stringSlice!=nil && strings.Compare(stringSlice[2],"ws")==0{
			hApi(w,r,stringSlice[3])
		}else if stringSlice!=nil && strings.Compare(stringSlice[2],"api")==0 {
			hApi(w,r,"")
		}

	}else {
		return
	}
	fmt.Println(upath)
}



var httpApi = map[string]func(w http.ResponseWriter, r *http.Request,e string){
	"/service/ws/mem"				:		ws.Client,
	"/service/ws/memPercent"		:		ws.Client,
	"/service/ws/cpuPercentSingle"	:		ws.Client,
	"/service/ws/cpuPercentAll"		:		ws.Client,
	"/service/api/cpuInfo"			:		api.CpuInfo,
	"/service/api/disk"				:		api.DiskInfo,
	"/service/api/diskPart"			:		api.DiskPart,
	"/service/api/diskPath" 		:		api.DiskInfoOfPath,

}