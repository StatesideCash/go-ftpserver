package main

import (
	"log"
)

func main() {
	/// Set log output info
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	FTPServ("127.0.0.1", "3124")
}
