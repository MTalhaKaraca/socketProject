package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
)

func SendName(conn net.Conn, username string) {
	nameLen := uint32(len(username))
	nameHeader := make([]byte, 4+nameLen)
	binary.LittleEndian.PutUint32(nameHeader[0:], nameLen)
	copy(nameHeader[4:], username)
	_, err := conn.Write(nameHeader)
	if err != nil {
		log.Println("cant send the message:", err)
		conn.Close()
	}
}

func ReadName(conn net.Conn) (uint32, string) {
	//gelen usernamennin ilk 4 bytesini okur.
	header := make([]byte, 4)
	_, err := conn.Read(header[:])
	if err != nil {
		log.Println("cant read:", err)
		conn.Close()
	}
	//4 byteyi uint32 olan nameLen atar.
	nameLen := binary.LittleEndian.Uint32(header[0:])
	//nameLen uzunluğuna göre bir []byte make eder.
	userName := make([]byte, nameLen)
	_, err = conn.Read(userName[:])
	if err != nil {
		fmt.Println("cant read conn:", err)
	}
	return nameLen, string(userName)

}

func GetUserName() string {
	var username string
	fmt.Printf("enter your name: ")
	_, err := fmt.Scanf("%v", &username)
	if err != nil {
		log.Println("cant get username:", err)
	}
	return username
}
