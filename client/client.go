package main

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	service := "127.0.0.1:9000"
	
	remoteAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, remoteAddr)
	checkError(err)
	defer conn.Close()

	data := make([]byte, 4)
	binary.BigEndian.PutUint32(data, 10)
	data = append(data, "hello\x00"...)
	fmt.Println("length of input data - ", len(data))

	_, err = conn.Write(data)
	checkError(err)

	result, err := ioutil.ReadAll(conn)
	checkError(err)

	fmt.Println(string(result))
}

func checkError(err error) {
	if err != nil {
		fmt.Printf("Fatal error : %s", err.Error())
		os.Exit(1)
	}
}