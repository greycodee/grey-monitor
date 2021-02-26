package api

import (
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"net/http"
)

func CpuInfo(w http.ResponseWriter, r *http.Request,e string)  {
	cpuInfo, _ := cpu.Info()
	j,_:=json.Marshal(cpuInfo)
	_, _ = fmt.Fprint(w, string(j))
}
