package main

import (
	"example/server/pkg/server"
	"flag"
	"strconv"
)

func main() {
	var listAddr = flag.String("listenAddr", "0.0.0.0", "specify address for the server to listen")
	var listPort = flag.Int("listenPort", 8070, "specify port for the server to listen")
	flag.Parse()

	//checkListenadrress(*listAddr)

	server.StartServer(*listAddr + ":" + strconv.Itoa(*listPort))

}
