# Simple File Server by Golang
Because there is no password required when uploading or downloading, so only for temporary use.

We already generate the runnable binary for linux-amd64 by `GOOS=linux GOARCH=amd64 go build -o vfile_amd64`.

## Server

### Install & Run
```
go install github.com/shevacn/vfile@latest

vfile -p 8080
```

## Client

### Upload
```
curl -F "uploads=@my-file.txt" http://your-server-ip:8080/upload
```
or by ui

http://your-server-ip:8080/uploadUi

### Download
```
Open http://your-server-ip:8080/ls/tmp/
```
