
set GOOS=windows
set GOARCH=amd64
go build -o ./build/autoConnectVPN.exe -ldflags="-w -s" ./main.go
