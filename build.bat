@echo off

go-winres make --arch amd64,386,arm64

set GOARCH=amd64
go build -trimpath -o KakaoGuard.exe -ldflags "-H windowsgui" .\cmd\main.go

set GOARCH=386
go build -trimpath -o KakaoGuard_i386.exe -ldflags "-H windowsgui" .\cmd\main.go

set GOARCH=arm64
go build -trimpath -o KakaoGuard_arm64.exe -ldflags "-H windowsgui" .\cmd\main.go
