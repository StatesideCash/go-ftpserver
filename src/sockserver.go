package main

import (
	"log"
	"net"
)

func FTPServ(iface, port string) {
	/// Start listener
	serv, err := net.Listen("tcp", iface+":"+port)
	if err != nil {
		log.Fatal(err)
	}

	for {
		/// Accept server connections
		conn, err := serv.Accept()
		if err != nil {
			log.Fatal(err)
		}

		/// Start thread to handle the connection
		go HandleCommandChannel(conn)
	}
}

func HandleCommandChannel(conn net.Conn) {

}
