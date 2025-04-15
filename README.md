# portForward
端口转发工具，实现仅允许本地访问的端口开放在0.0.0.0

用法：./portForward <想要开放在0.0.0.0的端口> <127.0.0.1:本地开放的端口>

如下，将仅允许本地访问的8080端口开放到0.0.0.0:8081

![image](https://github.com/user-attachments/assets/f4c2474a-81d8-47a9-a515-173c7af8e919)



```
编译指令：
$env:GOOS="linux"; $env:GOARCH="amd64"; go build -o portforward-linux-amd64 portforward.go
$env:GOOS="darwin"; $env:GOARCH="amd64"; go build -o portforward-darwin-amd64 portforward.go
$env:GOOS="windows"; $env:GOARCH="amd64"; go build -o portforward-windows-amd64.exe portforward.go
$env:GOOS="linux"; $env:GOARCH="arm"; go build -o portforward-linux-arm portforward.go
$env:GOOS="linux"; $env:GOARCH="386"; go build -o portforward-linux-386 portforward.go
$env:GOOS="darwin"; $env:GOARCH="arm64"; go build -o portforward-darwin-arm64 portforward.go
```

