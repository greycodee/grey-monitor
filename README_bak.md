# grey-monitor
基于go编写的系统监控可视化工具

## 使用的技术
- 前端：html+css+javascript+Echarts
- 后端：go
- go第三方库：gopsutil+websocket+go-bindata+go-bindata-assetfs

## go-bindata使用方法
1. 下载包(在项目外执行命令)
```shell script
    go get github.com/jteeuwen/go-bindata/...
    go get github.com/elazarl/go-bindata-assetfs/...
```
2. 设置环境变量
```shell script
    // 根据本地的目录来设置
    export PATH=$PATH:$HOME/go/bin/  
```
3. 打包静态文件
```shell script
    go-bindata views/...
```
此时根目录就会出现`bindata.go`文件
> debug模式(-debug)：可以不用频繁构建静态文件，使代码直接读取源静态文件，部署的时候在重新按上面步骤构建一下静态文件
```shell script
    go-bindata -debug views/...
```