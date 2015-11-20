package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

// FTPServ is the main handler for the FTP Server. It binds to the specified
// port and performs the server loop
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

// HandleCommandChannel processes the user's input to the server through the
// FTP command channel, where they can change the state of thier session. This
// function also processes the user's input through the correct handler.
func HandleCommandChannel(conn net.Conn) {
	var err error                   // Error checks
	state := new(ConnState)         // The state for the user's connection
	reader := bufio.NewReader(conn) // Reader to simplify reading input

	// Set the CWD for the session
	state.Directory, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	for {
		// Reads until a newline. I use \n for the delimiter instead of \r\n in case
		// some program does not follow the carridge return spec.
		command, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		action := strings.Split(string(command), " ")[0]
		action = strings.ToUpper(action)
		action = strings.TrimSpace(action)

		// The arguments provided to the command. It is up to the handler functions
		// to make sure they are proverly formatted and such.
		var args string
		if len(strings.TrimSpace(command)) <= len(strings.TrimSpace(action)) {
			args = ""
		} else {
			args = strings.TrimSpace(string(command[len(action)+1 : len(command)-1]))
		}

		// Selects the right handler to use
		switch action {
		case "USER":
			err = HandleUser(args, state)
		case "PASS":
			err = HandlePassword(args, state)
		case "CWD":
			err = HandleCWD(args, state)
		}

		//TODO Remove debug statement
		fmt.Println(action + ": " + args)
		fmt.Printf("%#v\n", state)

		if _, ok := err.(*os.PathError); ok {
			// TODO Handle invalid directory error
			log.Fatal(err)
		} else if err != nil {
			log.Fatal(err)
		}
	}
}
