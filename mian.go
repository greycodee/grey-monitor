package main

import (
	"flag"
	"fmt"
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"grey-monitor/api"
	"grey-monitor/ws"
	"html/template"
	"net/http"
	"os"
	"strings"
)

type handRequest struct {

}
var port = flag.String("p", "8989", "http端口")
func main() {
	//获取命令行参数
	flag.Parse()
	p:=":"+*port
	fmt.Printf(p)
	serverStart(p)
}

func serverStart(addr string)  {
	// 接口路由处理
	http.Handle("/",distribute())

	staticHandle := http.FileServer(assetFS())
	//将/js/路径下的请求匹配到 ./views/js/下
	http.Handle("/js/", staticHandle)
	http.Handle("/archive/", staticHandle)

	fmt.Println("http服务器端口："+addr)
	// 开启http服务
	e:=http.ListenAndServe(addr,nil)
	if e!=nil {
		fmt.Println(e)
	}
}

/*
	go-bindata-assetfs静态文件路由
*/
func assetFS() *assetfs.AssetFS {
	assetInfo := func(path string) (os.FileInfo, error) {
		return os.Stat(path)
	}
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: assetInfo, Prefix: k}
	}
	panic("unreachable")
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
/*
	主页
*/
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