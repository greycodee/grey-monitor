package main

import (
	"fmt"
	"grey-monitor/api"
	"grey-monitor/ws"
	"html/template"
	"net/http"
	"strings"
)

type handRequest struct {

}

func main() {
	serverStart(":8988")
}

func serverStart(addr string)  {
	// 接口路由处理
	http.Handle("/",distribute())

	// 静态文件路由处理
	//指定相对路径./static 为文件服务路径
	staticHandle := http.FileServer(http.Dir("./views"))
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
	api:=httpApi[upath]
	if api!=nil {
		stringSlice := strings.Split(upath,"/")
		if stringSlice!=nil && strings.Compare(stringSlice[1],"ws")==0{
			api(w,r,stringSlice[2])
		}else {

			api(w,r,"")
		}

	}else {
		return
	}
	fmt.Println(upath)
}

func index(w http.ResponseWriter, r *http.Request,e string)  {
	t,_:=template.ParseFiles("views/index.html")
	t.Execute(w,"")
}


var httpApi = map[string]func(w http.ResponseWriter, r *http.Request,e string){
	"/"				:		index,
	"/ws/mem"		:		ws.Client,
	"/ws/cpu"		:		ws.Client,
	"/ws/ps"		:		ws.Client,
	"/api/disk"		:		api.DiskInfo,
	"/api/diskPart"	:		api.DiskPart,
	"/api/diskPath" :		api.DiskInfoOfPath,
}