package main

import (
	"encoding/binary"
	"log"
	"net"
	"os"
)

func CreateMessage(conn net.Conn, msgType uint32, data string) {
	msgHeader := make([]byte, 8)
	msgLen := uint32(len(data[:]))

	//ilk 8bytesine msg type ve dataLen koyar.
	binary.LittleEndian.PutUint32(msgHeader[0:], msgType)
	binary.LittleEndian.PutUint32(msgHeader[4:], msgLen)

	//gönderilecek olan mesajın kendisi.  msgHeader + data -> message
	var message []byte
	message = append(message, msgHeader...)
	message = append(message, data...)

	//data gönderilir.
	_, err := conn.Write(message)
	if err != nil {
		log.Println("message cant send:", err)
		conn.Close()
		os.Exit(1)
	}
}

func ReadMessage(conn net.Conn) (uint32, uint32, string) {
	//Gelen mesajın ilk 8 bytesini okur.
	header := make([]byte, 8)
	_, err := conn.Read(header[:])
	if err != nil {
		log.Println("cant read the message: ", err)
		conn.Close()
		os.Exit(1)
	}

	//ilk 8 bytesine karşılık gelen verileri atar.
	msgType := binary.LittleEndian.Uint32(header[0:])
	msgLen := binary.LittleEndian.Uint32(header[4:])

	//msgLen e göre data make eder.
	data := make([]byte, msgLen)

	//connectiondan mesajın geri kalanını okur. dataya atar.
	_, err = conn.Read(data[:])
	if err != nil {
		log.Println("cant read the message: ", err)
		conn.Close()
		os.Exit(1)
	}
	//fmt.Printf("Type:%d  Len:%d   msg: %s",msgType,msgLen,string(data))
	return msgType, msgLen, string(data[:])
}
