package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	//inicia server
	server, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)

	}

	//clientes
	var i int
	for {

		fmt.Println("Esperando un cliente")

		//se conecta un cliente
		client, err := server.Accept()
		if err != nil {
			panic(err)
		}
		//manager connection
		go managerconnection(client, i)
		i++
	}
}

func managerconnection(cliente net.Conn, i int) {
	//creando una variable de struct para enviar al array arrayUser
	c := userConnection{
		connection: cliente,
		id:         i,
	}
	//agregando variale struct al array arrayUser
	arrayUser = append(arrayUser, c)

	for {
		var buff = make([]byte, 2048)
		//leyendo el mensaje
		_, err := cliente.Read(buff)
		if err != nil {
			panic(err)
		}

		//declarando nombre archivo pasandolo a string
		nom := string(buff)
		//cortando el mensaje a partir de .fin
		var nombre []string = strings.Split(nom, ".fin")
		//abriendo archivo enviado
		file, errOpen := os.Open(nombre[0]) // For read access.
		if errOpen != nil {
			fmt.Println("ERROR AL ABRIR EL ARCHIVO")
			log.Fatal(errOpen)

		}

		//leyendo archivo
		data, errRead := os.ReadFile(nombre[0])
		if errRead != nil {
			fmt.Println("ERROR AL LEER EL ARCHIVO")
			log.Fatal(errRead)

		}

		//cerrando archivo
		file.Close()

		//editando nombre para cambiar la direccion del archivo
		nombreArchivoCrear := strings.Split(nombre[0], "./")

		//crear un archivo en DIALTCP/archivos con el mismo nombre del archivo enviado
		emptyFile, errCreate := os.Create("../archivos/" + nombreArchivoCrear[1])
		if errCreate != nil {
			fmt.Println("ERROR AL crear el archivo")
			log.Fatal(errCreate)

		} else {
			fmt.Println("se creo el archivo exitosamente")
		}

		//escribiendo el archivo creado
		errWrite := os.WriteFile("../archivos/"+nombreArchivoCrear[1], data, 0666)
		if errWrite != nil {
			fmt.Println("ERROR AL ESCRIBIR EL ARCHIVO")
			log.Fatal(errWrite)

		} else {
			fmt.Println("Se escribio el archivo exitosamente")
		}

		//cerrando copia de archivo
		emptyFile.Close()
		//se lo manda a todos los usuarios
		writeMessageUser(data)

	}

}

func writeMessageUser(message []byte) {
	for _, c := range arrayUser {
		//recorro el array de struct arrayUser
		//si el id coincide con el numero enviado, envio el mensaje a esa coneccion
		if c.id == 1 {
			c.connection.Write(message)
		}

	}
}
