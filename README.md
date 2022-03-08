# Simple File Server by Golang
Because there is no password required when uploading or downloading, so only for temporary use.

## Server

### Install & Run
```
go install github.com/shevacn/vfile@latest

vfile
```

## Client

### Upload
```
curl -F "uploads=@my-file.txt" http://your-server-ip:8080/upload
```

### Download
```
Open http://your-server-ip:8080/ls/tmp/
```
