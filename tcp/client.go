package main
 
import (
    "fmt"
    "net"
    "time"
)
 
func main() {
    conn, err := net.Dial("tcp", "127.0.0.1:8000") // 프로토콜, IP 주소, 포트 번호를 설정하여 서버에 연결
    if err != nil {
        fmt.Println("Failed to Dial : ", err)
    }
    defer conn.Close() // TCP 연결을 닫음
 
    go func(c net.Conn) {
        send := "Hello"
		
        for {
            _, err = c.Write([]byte(send)) // server로 데이터를 보냄
            if err != nil {
                fmt.Println("Failed to write data : ", err)
                break;
            }
 
            time.Sleep(1 * time.Second)
        }
    }(conn)
 
    go func(c net.Conn) {
        recv := make([]byte, 4096)
 
        for {
            n, err := c.Read(recv) // server에게 받은 데이터를 읽음
            if err != nil {
                fmt.Println("Failed to read data : ", err)
                break
            }
 
            fmt.Println(string(recv[:n]))
        }
    }(conn)
 
    fmt.Scanln()
}