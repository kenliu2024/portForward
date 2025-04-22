package main

import (
    "io"
    "log"
    "net"
)

func PortForward() {
    // 本地监听端口
    const listenPort = ":8001"
    // 目标地址（例如：127.0.0.1:8000）
    const targetAddr = "yz.testapp.shop:8000"

    // 监听本地端口
    listener, err := net.Listen("tcp", listenPort)
    if err != nil {
        log.Fatalf("无法监听端口 %s: %v", listenPort, err)
    }
    defer listener.Close()

    log.Printf("正在监听 %s，转发到 %s", listenPort, targetAddr)

    for {
        // 接受客户端连接
        clientConn, err := listener.Accept()
        if err != nil {
            log.Printf("接受连接失败: %v", err)
            continue
        }

        // 连接到目标地址
        targetConn, err := net.Dial("tcp", targetAddr)
        if err != nil {
            log.Printf("连接目标地址 %s 失败: %v", targetAddr, err)
            clientConn.Close()
            continue
        }

        // 启动 goroutine 转发数据
        go forward(clientConn, targetConn)
    }
}

// 转发数据
func forward(src, dst net.Conn) {
    defer src.Close()
    defer dst.Close()

    // 从 src 读取数据并写入 dst
    go func() {
        if _, err := io.Copy(dst, src); err != nil {
            log.Printf("从 src 到 dst 复制数据出错: %v", err)
        }
    }()

    // 从 dst 读取数据并写入 src
    if _, err := io.Copy(src, dst); err != nil {
        log.Printf("从 dst 到 src 复制数据出错: %v", err)
    }
}

func main() {
    PortForward()
}