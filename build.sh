#!/bin/bash

go mod download
go-bindata views/...
# 在linux环境下编译Mac Windows Linux的包
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o grey-monitor_mac_x64 .
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o grey-monitor_win_x64.exe .
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o grey-monitor_linux_x64 .
