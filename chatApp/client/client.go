package main
import (
	"net"
	"fmt"
	"bufio"
	"sync"
	"os"
)
func outgoing(sock net.Conn){
	for {
		fmt.Println("message : ")
		message,_:=bufio.NewReader(os.Stdin).ReadString('\n')
		fmt.Fprintf(sock, message+"\n")
}
}
func ingoing(sock net.Conn){
	for {
		message,_:=bufio.NewReader(sock).ReadString('\n')
	    fmt.Println(message)
	}
}
func main(){
	sock,_:=net.Dial("tcp","127.0.0.1:8080")
	fmt.Println("connection Accepted")
	// writer:=bufio.NewWriter(sock)
	// writer.WriteString("message from\n")
	// writer.Flush()
	// fmt.Fprintf(sock,"message from\n")
	var wg sync.WaitGroup
	wg.Add(1)
	go outgoing(sock)
	go ingoing(sock)
	wg.Wait()
}