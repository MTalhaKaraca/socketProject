package main

import (
	"bufio"
	"fmt"
	"log"
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

	listener, err := net.Listen("tcp", ":8080")
	fmt.Println("Connection is ready!")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		defer conn.Close()
		if err != nil {
			log.Println(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	//send name to opposite user.
	username := GetUserName()
	SendName(conn, username)

	//receive the name of opposite user.
	_, userBeingTexted := ReadName(conn)
	fmt.Println("------------MESSAGİNG APPLİCATİON------------")
	fmt.Println("The person you messaging with:", userBeingTexted)

	go func() {
		//send message to user.
		var message string
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			message = scanner.Text()
			CreateMessage(conn, 2, message)
		}
	}()

	//read message from server.
	for {
		msgType, msgLen, data := ReadMessage(conn)
		printMessage(msgType, msgLen, data, userBeingTexted)
	}
}

func printMessage(msgType uint32, msgLen uint32, message string, name string) {
	time := time2.Now().Format("15:04")
	fmt.Printf("#%d [%s]%s>> %s   \n  Type:%d Len:%d\n\n", count, time, name, message, msgType, msgLen)
	count++
}
