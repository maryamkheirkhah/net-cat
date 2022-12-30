package server

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
)

var (
	Guest        = map[string]net.Conn{}
	NumbTerminal []int
	MaxGuest     = int(10)
	messages     string
	mutex        sync.Mutex
	file         *os.File
	C            = Colors{
		Reset:      "\033[0m",
		DarkGrey:   "\033[1;30m",
		Red:        "\033[31m",
		LightRed:   "\033[1;31m",
		Green:      "\033[32m",
		LightGreen: "\033[1;32m",
		Orange:     "\033[0;33m",
		Yellow:     "\033[1;33m",
		Blue:       "\033[0;34m",
		LightBlue:  "\033[1;34m",
		Purple:     "\033[0;35m",
		Black:      "\033[0;30m",
		Cyan:       "\033[36m",
		Magenta:    "\033[35m",
	}
)

/*
saveData takes an input string and writes it to the chat-history text file created on
server startup.
*/
func saveData(s string) {
	_, err2 := file.Write([]byte(s))
	if err2 != nil {
		fmt.Println(err2)
		log.Fatal(err2)
	}
}

/*
send takes the user's name and message as input strings, along with a save boolean,
and adds the message string to all other users' message
*/
func send(activeName, msg, msgHistory string, saveMsg bool) {
	for guestName := range Guest {
		if guestName != activeName {
			Guest[guestName].Write([]byte("\n" + msg))
		}
		// new line
		sb := fmt.Sprintf("[%s][%s]:", GetTime(), guestName)
		Guest[guestName].Write([]byte(sb))
	}
	if saveMsg {
		messages += msgHistory
	}
}

/*
sendChat takes a net.Conn object, a name string and a color string as inputs and sets
up a bufio scanner object for the net.Conn object (which implements the io.Reader interface)
and a mutex method allowing only one chatRoom go-routine (individual user terminals)
to access the chat message string and chat history / logs at any one time.
*/
func sendChat(conn net.Conn, name string, color string) string {
	scanner := bufio.NewScanner(conn)
	var msg string
	var msgHistory string
	Name := name

	// Infinite loop for listening, until terminated
	for scanner.Scan() {
		text := strings.Trim(scanner.Text(), "\n\t\r")

		if len(text) != 0 {
			if text == "/cn" {
				Name = changeMyName(conn, Name)
			} else {
				colorname := color + Name + C.Reset
				mutex.Lock()
				// Create message string for saving to chat history / logs
				data := fmt.Sprintf("[%s][%s]:%s\n", GetTime(), Name, text)
				saveData(data)
				// Create color-formatted message string for publishing on chat
				s := fmt.Sprintf("\033[%dA", 1)
				msg = fmt.Sprintf(s+"[%s][%s]:", GetTime(), colorname) + color + text + C.Reset + "\n"
				msgHistory = fmt.Sprintf("[%s][%s]:", GetTime(), colorname) + color + text + C.Reset + "\n"
				send(Name, msg, msgHistory, true)
				mutex.Unlock()
			}
		}
	}
	return Name
}

/*
chatRoom runs as a go-routine within the StartServer function for each connection made to
the netcat TCP server. It takes the net.Conn identifier, a logo in the form of a byte slice,
and the index of the chat / terminal connected to the
*/
func chatRoom(conn net.Conn, logo []byte, userNbr int) {
	NumbTerminal = append(NumbTerminal, userNbr)
	conn.Write(logo)
	conn.Write([]byte("\n"))
	name := addName(conn) // Lets guest add guest name
	fmt.Printf("New user ( %v ) has joined the chat\n", name)
	fmt.Println("Number of users in chat Room : ", len(NumbTerminal))
	color := AssignColor(FindIndex(userNbr))

	// sendChat sets up listening on connection, with mutexes etc.
	name = sendChat(conn, name, color)

	// Only upon sendChat terminating, does the logout sequence progress
	send(name, fmt.Sprintf("%s has left our chat ...\n", name), "void", false)
	fmt.Printf("User ( %v ) has left the chat\n", name)
	RemoveElement(userNbr)
	delete(Guest, name) // Delete guest name who has leave from chat
	conn.Close()
}

/*
addName is a function that takes a net.Conn object as input and prompts for a new
username to be added to the global Guest struct within the related chat / TCP connection.
This new username is also returned.
*/
func addName(conn net.Conn) string {
	conn.Write([]byte("[ENTER YOUR NAME]:"))
	scanner := bufio.NewScanner(conn)
	var name string
	for scanner.Scan() {
		name = strings.TrimSpace(scanner.Text())
		if len(name) == 0 {
			conn.Write([]byte("[PLEASE, ENTER YOUR NAME]:"))
		} else if _, exists := Guest[name]; exists {
			conn.Write([]byte("[PLEASE, ENTER A DIFFERENT NAME]:"))
		} else {
			Guest[name] = conn
			break
		}
	}
	conn.Write([]byte(messages))
	JoinMsg := fmt.Sprintf("%s has joined our chat...\n", name)
	send(name, JoinMsg, "void", false)
	return name
}

/*
changeMyName is a function that takes a net.Conn object and existing username string
as inputs and prompts for a replacement for the existing one in the global Guest struct
(the new username is returned as a string).
*/
func changeMyName(conn net.Conn, name string) string {
	conn.Write([]byte("[ENTER YOUR NAME]:"))
	scanner := bufio.NewScanner(conn)
	var newName string
	for scanner.Scan() {
		newName = strings.TrimSpace(scanner.Text())
		if len(newName) == 0 {
			conn.Write([]byte("[PLEASE, ENTER YOUR NAME]:"))
		} else if _, ok := Guest[newName]; ok {
			conn.Write([]byte("[PLEASE, ENTER A DIFFERENT NAME]:"))
		} else {
			delete(Guest, name)
			Guest[newName] = conn
			break
		}
	}
	changeNameMsg := fmt.Sprintf("%s has changed name to %s.\n", name, newName)
	send(name, changeNameMsg, changeNameMsg, true)
	return newName
}

/*
CheckPort takes an input string and checks whether it corresponds to a
valid positive integer as well as lying within the accpetable port range.
*/
func CheckPort(port string) (string, error) {
	// Check for invalid characters (non-integers)
	for _, r := range port {
		if r < '0' || r > '9' {
			return "", errors.New(C.Yellow + "port number " + C.Red + port +
				C.Yellow + " not a valid integer" + C.Reset)
		}
	}
	// Check that port exists within valid range
	portNbr := PortAtoi(port)
	if portNbr < 1024 && portNbr > 65535 {
		return "", errors.New(C.Yellow + "port number " + C.Red + port + C.Yellow +
			" not within acceptable range 1024 - 65535" + C.Reset)
	}
	return port, nil
}

/*
StartServer is a function which takes a Host and Port input string, as well as a logo
in the form of a slice of bytes. It establishes a TCP chat connection
*/
func StartServer(port string, logo []byte) {
	// Announce server listening on given port on local network
	address := ":" + port
	listening, errListen := net.Listen("tcp", address)
	if errListen != nil {
		log.Fatal(errListen)
	}

	// Create log file for recording of client activities / chat history
	var errFileCreate error
	file, errFileCreate = os.Create("Data.txt")
	if errFileCreate != nil {
		log.Fatal(errFileCreate)
	}
	defer file.Close()

	// Prepare to accept connection requests
	userNbr := 0
	for {
		conn, err := listening.Accept()
		if err != nil {
			log.Printf("error: %s", err)
		} else {
			if len(Guest) >= MaxGuest {
				conn.Write([]byte("Sorry, all 10 chat slots are in use, please come back later."))
				conn.Close()
			} else {
				// Run each connection as a seperate go routine
				userNbr++
				go chatRoom(conn, logo, userNbr)
			}
		}
	}
}
