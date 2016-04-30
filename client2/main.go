//Jacob Daniel
//Client 2

/*
This is one of the clients for my chat server
*/

package main
 
import (
        "fmt"
        "log"
        "net"
		"bufio"
		"os"
)
 
func main() {
        conn, err := net.Dial("tcp", "localhost:8080")
        if err != nil {
                log.Fatalln(err)
        }

		go listenConnection(conn, err)
		writeMessage(conn, err)

}

//Searches for a connection and receives a message if one is available
func listenConnection(conn net.Conn, err error){
	for {
                buffer := make([]byte, 1400)
                dataSize, err := conn.Read(buffer)
                if err != nil {
                        fmt.Println("CONNECTION CLOSED")
                        return
                }
 
                data := buffer[:dataSize]
                fmt.Print(string(data))
        }
}

//Writes a message to be sent to the server
func writeMessage(conn net.Conn, err error){
	for {
 		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = "Dude 2: " + text

        _, err = conn.Write([]byte(text))
        if err != nil {
                log.Fatalln(err)
        }
	}
}