package main

import (
	"fmt"
	"log"
	"netcat/server"
	"os"
)

/*
INSTRUCTIONS:
This project consists on recreating the NetCat command-line utility in a Server-Client Architecture
that can run in a server mode on a specified port listening for incoming connections,
and it can be used in client mode, trying to connect to a specified port and
transmitting information to the server. The Netcat command-line utility reads and writes data across
network connections using TCP or UDP. It can be used for anything involving TCP, UDP, or
UNIX-domain sockets, it is able to open TCP connections, send UDP packages,
listen on arbitrary TCP and UDP ports and many more.

The project must work in a similar way that the original NetCat works, in other words,
it must create a group chat. The project must have the following features :

 01. TCP connection between server and multiple clients (relation of 1 to many).
 02. A name requirement to the client.
 03. Control connections quantity.
 04. Clients must be able to send messages to the chat.
 05. Do not broadcast EMPTY messages from a client.
 06. Messages sent, must be identified by the time that was sent
    and the user name of who sent the message,
    example : [2020-01-20 15:48:41][client.name]:[client.message]
 07. If a Client joins the chat, all the previous messages sent to the chat
    must be uploaded to the new Client.
 08. If a Client connects to the server, the rest of the Clients
    must be informed by the server that the Client joined the group.
 09. If a Client exits the chat, the rest of the Clients
    must be informed by the server that the Client left.
 10. All Clients must receive the messages sent by other Clients.
 11. If a Client leaves the chat, the rest of the Clients must
    not disconnect.
 12. If there is no port specified, then set as default the port 8989.
    Otherwise, program must respond with usage message: [USAGE]: ./TCPChat $port
*/
func main() {
	var port string
	input := os.Args[1:]

	// No additional NetCat flags have been implemented
	// e.g. -l, -L, -u, -p, -e, -n, -z, -w, -v, -vv
	switch len(input) {
	// If no input argument is given, assign default port number
	case 0:
		port = "8989"
	// If port number is given, check validity
	case 1:
		var err error
		port, err = server.CheckPort(input[0])
		if err != nil {
			log.Fatal(err)
		}
	// If more than one argument is entered, return usage instructions
	default:
		log.Fatal(server.C.Yellow +
			"[USAGE]: ./TCPChat $port\n\nAlternatively, enter " +
			server.C.Red + "<go run .>" + server.C.Yellow + " or " +
			server.C.Red + "<go run . <your port>" + server.C.Yellow +
			" in the terminal to start the server" + server.C.Reset + "\n")
	}

	// Check for presence of linux logo text file
	logo, err := os.ReadFile("logo.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(server.C.Green+"Listening on port :%s\n"+server.C.Reset, port)
	server.StartServer(port, logo)
}
