go env -w GOARCH=amd64
go env -w GOOS=windows
go build -ldflags '-s -w'