//Jacob Daniel
//Chat Server

/*
This is the server for my chat client.
It accepts messages from 2 clients and send them to the appropriate client.
*/
package main

import (
	"bufio"
	"net"
	"fmt"
)

var conn, conn2 net.Conn

func main() {
	ln, _ := net.Listen("tcp", ":3000")
	conn, _ = ln.Accept()
	ln2, _ := net.Listen("tcp", ":8080")
	conn2, _ = ln2.Accept()
	for {
		go connHandler()
		connHandler2()
	}
}

//This reads client 1's messages and sends them to the other client
func connHandler(){
	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println("Message Received: \n" + string(message))
	conn2.Write([]byte(message))
}

//This reads client 2's messages and sends them to the other client
func connHandler2(){
	message, _ := bufio.NewReader(conn2).ReadString('\n')
	fmt.Println("Message Received: \n" + string(message))
	
	conn.Write([]byte(message))
}