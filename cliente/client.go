package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	serverCnn, err := net.Dial("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	//escuchando lo que manda el servidor

	go readMessages(serverCnn)
	// envia mensajes
	writeMessage(serverCnn)
}

func writeMessage(conn net.Conn) {
	var scanner = bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		//una vez cada vez que se preisona enter
		var texto = scanner.Text()

		_, err := conn.Write([]byte(texto))
		if err != nil {
			panic(err)
		}
	}
}

func readMessages(cnn net.Conn) {

	for {
		//lee todos los mensajes
		var message = make([]byte, 2048)

		_, err := cnn.Read(message)
		if err != nil {
			panic(err)
		}
		//print
		fmt.Println("El mensaje del servidor es: ", string(message))
	}

}
