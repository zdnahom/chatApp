package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
)

var socks = []net.Conn{}

func outgoing(sock net.Conn) {
	for {
		fmt.Println("message: ")
		message, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		fmt.Fprintf(sock, message+"\n")
	}
}
func ingoing(sock net.Conn, index int) {
	for {
		server_msg, _ := bufio.NewReader(sock).ReadString('\n')

		// for i := 0; i < len(socks); i++ {
		// 	fmt.Fprintf(socks[i],"Client " + strconv.Itoa(index) + " : " + server_msg + "\n")
		// }
		for _, conn := range socks {
			if conn == sock {
				continue
			}
			fmt.Fprintf(sock, "Client "+strconv.Itoa(index)+" : "+server_msg+"\n")
		}
	}
}
func main() {
	l, _ := net.Listen("tcp", "127.0.0.1:8080")
	fmt.Println("server started at 127.0.0.1:8080.....")
	for {
		sock, _ := l.Accept()
		fmt.Println("User joined")
		socks = append(socks, sock)
		go ingoing(sock, len(socks))
	}

}

// message,_:=bufio.NewReader(sock).ReadString('\n')
// fmt.Println("message is ",message)
// sock.Close()
