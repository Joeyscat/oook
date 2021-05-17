#! /bin/bash
go env -w GOARCH=amd64
go env -w GOOS=linux
go build -ldflags '-s -w'