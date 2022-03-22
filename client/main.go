package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	time2 "time"
)

const (
	msgTypeJSON = 1
	msgTypeText = 2
	msgTypeXML  = 3
)

var count = 0

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	//send own name and receive the opposite username.
	username := GetUserName()
	SendName(conn, username)
	_, name := ReadName(conn)

	fmt.Println("------------MESSAGİNG APPLİCATİON------------")
	fmt.Println("The person you messaging with:", name)

	go func() {
		//read message from server.
		for {
			msgType, msgLen, data := ReadMessage(conn)
			printMessage(msgType, msgLen, data, name)

		}
	}()

	//send message to user.
	var message string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message = scanner.Text()
		CreateMessage(conn, 2, message)
	}

}

func printMessage(msgType uint32, msgLen uint32, message string, name string) {
	time := time2.Now().Format("15:04")
	fmt.Printf("#%d [%s]%s>> %s   \n  Type:%d Len:%d\n\n", count, time, name, message, msgType, msgLen)
	count++
}
