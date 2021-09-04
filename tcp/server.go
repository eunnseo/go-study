package main

import (
    "fmt"
    "io"
    "net"
)
 
func handler(conn net.Conn) {
	fmt.Println("=== Run handler")
    recv := make([]byte, 4096)
 
    for {
        n, err := conn.Read(recv) // client에게 받은 데이터를 읽음
        if err != nil {
            if err == io.EOF {
                fmt.Println("connection is closed from client : ", conn.RemoteAddr().String())
            }
            fmt.Println("Failed to receive data : ", err)
            break
        }
 
        if n > 0 {
            fmt.Println(string(recv[:n]))
            conn.Write(recv[:n]) // 받은 데이터를 client에게 다시 보냄
        }
    }
}
 
func main() {
    l, err := net.Listen("tcp", ":8000") // 프로토콜, IP 주소, 포트 번호를 설정하여 네트워크 연결을 대기
    if err != nil {
        fmt.Println("Failed to Listen : ", err)
    }
    defer l.Close() // TCP 연결 대기를 닫음
 
    for {
        conn, err := l.Accept() // 클라이언트가 연결되면 커넥션을 리턴
        if err != nil {
            fmt.Println("Failed to Accept : ", err)
            continue
        }
 
        go handler(conn)
    }
}