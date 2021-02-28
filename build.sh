#!/bin/bash

go mod download
go-bindata-assetfs views/...
go build .