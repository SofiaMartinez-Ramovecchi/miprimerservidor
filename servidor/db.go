//base de datos de los usuarios

package main

import "net"

type userConnection struct {
	connection net.Conn
	id         int
}

var arrayUser []userConnection
