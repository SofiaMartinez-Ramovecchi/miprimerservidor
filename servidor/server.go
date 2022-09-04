package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	//inicia server
	server, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)

	}

	//clientes

	for {
		fmt.Println("Esperando un cliente")

		//se conecta un cliente
		client, err := server.Accept()
		if err != nil {
			panic(err)
		}
		//manager connection
		go managerconnection(client)
	}
}

func managerconnection(cliente net.Conn) {

	userConnection = append(userConnection, cliente)

	for {
		var buff = make([]byte, 2048)
		//leyendo el mensaje
		_, err := cliente.Read(buff)
		if err != nil {
			panic(err)
		}

		//abriendo archivo enviado
		file, err := os.Open(string(buff)) // For read access.
		if err != nil {
			log.Fatal(err)
		}

		//leyendo archivo
		data, err := os.ReadFile(string(buff))
		if err != nil {
			log.Fatal(err)
		}
		os.Stdout.Write(data)
		//cerrando archivo

		file.Close()

		//crear un archivo en DIALTCP/archivos con el mismo nombre del archivo enviado
		emptyFile, err := os.Create(string(buff))
		if err != nil {
			log.Fatal(err)
		}
		log.Println(emptyFile)

		//escribiendo el archivo creado
		errwrite := os.WriteFile(string(buff), data, 0666)
		if errwrite != nil {
			log.Fatal(errwrite)
		}

		//cerrando copia de archivo
		emptyFile.Close()
		//se lo manda a todos los usuarios
		writeMessageAllUsers(buff)

	}

}

func writeMessageAllUsers(message []byte) {
	for _, c := range userConnection {
		c.Write(message)
	}
}
