package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

const (
	join = iota
	username
	message
	messages
	file
	endSession
)

type server struct {
	members  map[string]client
	Messages []Message
}

type client struct {
	Username string
	Port     int
}

type request struct {
	ClientId string
	Action   int
	Username string
	Message  string
	File     File
}

type response struct {
	Message  string
	Messages []Message
	File     File
}

type Message struct {
	From client
	Text string
	Date time.Time
}

type File struct {
	Bytes    []byte
	Length   int
	Filename string
}

func newServer() *server {
	return &server{
		members: make(map[string]client),
	}
}

func (s *server) decode(conn net.Conn) request {
	var req request
	err := gob.NewDecoder(conn).Decode(&req)
	if err != nil {
		log.Printf("Unable to decode request: %s", err.Error())
	}
	return req
}

func (s *server) encode(conn net.Conn, data interface{}) {
	err := gob.NewEncoder(conn).Encode(data)
	if err != nil {
		log.Printf("Unable to encode request: %s", err.Error())
	}
}

func (s *server) run() {
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("Unable to start server: %s", err.Error())
		return
	}
	defer listener.Close()
	log.Println("Server started on port 5000")

	for {
		c, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %s", err.Error())
			continue
		}
		go s.clientHandler(c)
	}
}

func (s *server) clientHandler(conn net.Conn) {
	defer conn.Close()
	req := s.decode(conn)

	if req.Action == join {
		s.newClient(conn, req.ClientId)
	} else if req.Action == username {
		s.setupUsername(conn, req.ClientId, req.Username)
	} else if req.Action == message {
		s.msg(req.Username, req.Message, conn)
	} else if req.Action == messages {
		s.sendAllMessagesToClient(req.Username, conn)
	} else if req.Action == file {
		s.file(req.Username, req.File, conn)
	} else if req.Action == endSession {
		s.endClient(req.Username)
	}
}

func (s *server) newClient(conn net.Conn, randid string) {
	c := client{
		Username: randid,
		Port:     s.getNewPort(),
	}
	s.join(c, conn)
}

func (s *server) join(c client, conn net.Conn) {
	log.Println("New client has joined")
	s.members[c.Username] = c
	s.encode(conn, response{strconv.Itoa(c.Port), []Message{}, File{}})
}

func (s *server) setupUsername(conn net.Conn, clientId string, newUsername string) {
	c := s.getClientByUsername(clientId)
	if c.Username == "" {
		log.Fatalln("Couldn't find client with id", clientId)
		return
	}
	s.sendUsernameValidationToClient(c, newUsername, conn)
}

func (s *server) sendUsernameValidationToClient(c client, username string, conn net.Conn) {
	usernames := s.getUsernames()
	if existsInSlice(username, usernames) {
		log.Printf("Username %s taken", username)
		s.encode(conn, response{"0", []Message{}, File{}})
	} else {
		s.encode(conn, response{"1", []Message{}, File{}})
		prevUsername := c.Username
		c.Username = username
		s.members[username] = c
		delete(s.members, prevUsername)
		log.Printf("Username %s available", username)
		s.logCurrentUsers()
	}
}

func (s *server) msg(from string, msg string, conn net.Conn) {
	log.Println(from, ":", msg)
	c := s.getClientByUsername(from)
	s.saveMessage(Message{From: c, Text: msg, Date: time.Now()})
	s.broadcastMessage(from, msg, conn)
}

func (s *server) file(from string, file File, conn net.Conn) {
	log.Println(from, "sent", file.Filename)
	c := s.getClientByUsername(from)
	s.saveMessage(Message{From: c, Text: file.Filename, Date: time.Now()})
	s.broadcastFile(from, file, conn)
}

func (s *server) broadcastMessage(from string, msg string, conn net.Conn) {
	clients := s.getClients()

	for _, c := range clients {
		if from != c.Username {
			cc, err := net.Dial("tcp", ":"+strconv.Itoa(c.Port))
			if err != nil {
				log.Fatalln("Something went wrong broadcasting message", err.Error())
				return
			}
			s.encode(cc, response{from + ": " + msg, []Message{}, File{}})
			cc.Close()
		}
	}
}

func (s *server) broadcastFile(from string, file File, conn net.Conn) {
	clients := s.getClients()

	for _, c := range clients {
		if from != c.Username {
			cc, err := net.Dial("tcp", ":"+strconv.Itoa(c.Port))
			if err != nil {
				log.Fatalln("Something went wrong broadcasting message", err.Error())
				return
			}
			s.encode(cc, response{from + " envio " + file.Filename, []Message{}, file})
			cc.Close()
		}
	}
}

func (s *server) backupMessages() {
	delimiter := "|"
	messages := []string{}
	for _, m := range s.Messages {
		messages = append(messages, m.Date.Format("06-Jan-02")+delimiter+m.From.Username+delimiter+m.Text)
	}
	saveToFile(messages, "messages.txt")
	log.Println("Messages saved to messages.txt")
}

func (s *server) endClient(username string) {
	log.Println("User", username, "has disconnected from the server")
	delete(s.members, username)
	s.logCurrentUsers()
}

func (s *server) logCurrentUsers() {
	log.Printf("Current users in chat %s", s.getUsernames())
}

func (s *server) saveMessage(msg Message) {
	s.Messages = append(s.Messages, msg)
}

func (s *server) sendAllMessagesToClient(username string, conn net.Conn) {
	s.encode(conn, response{"", s.getAllMessagesForClient(username), File{}})
}

func (s *server) getAllMessagesForClient(username string) []Message {
	messages := []Message{}
	for _, m := range s.Messages {
		messages = append(messages, m)
	}
	return messages
}

func (s *server) quit(c *client) {

}

func (s *server) getClients() []client {
	var clients []client
	for _, c := range s.members {
		clients = append(clients, c)
	}
	return clients
}

func (s *server) getUsernames() []string {
	var usernames []string
	for _, c := range s.members {
		usernames = append(usernames, c.Username)
	}
	return usernames
}

func (s *server) getClientByUsername(u string) client {
	c := client{}
	for username, cli := range s.members {
		if username == u {
			c = cli
		}
	}
	return c
}

func (s *server) getNewPort() int {
	port := 5001
	for _, c := range s.members {
		port = c.Port
	}
	return port + 1
}

func (s *server) displayAllMessages() {
	fmt.Printf("\n\n\n\n\n\n\n\n")
	fmt.Printf("------------------------\n")
	fmt.Println("-> Todos los mensajes")
	fmt.Print("------------------------")
	for _, m := range s.Messages {
		printMessage(m)
	}
	fmt.Println("------------------------")
	fmt.Printf("\n\n\n\n\n\n")
}

func printMessage(m Message) {
	fmt.Printf("\n\nDe ")
	fmt.Println(m.From.Username)
	fmt.Println(m.Text)
	fmt.Printf("el ")
	fmt.Print(m.Date.Format("06-Jan-02"))
	fmt.Printf("\n\n")
}

func main() {
	s := newServer()
	go s.run()

	killServer := false
	for !killServer {
		displayMenu()
		option := getIntFromUser()
		if option == 1 {
			s.displayAllMessages()
		} else if option == 2 {
			s.backupMessages()
		} else if option == 3 {
			killServer = true
		}

	}
}

func displayMenu() {
	fmt.Println("1. Mostrar mensajes/archivos")
	fmt.Println("2. Hacer backup de mensajes")
	fmt.Println("3. Terminar servidor")
	fmt.Println("Opcion: ")
}

func getStringFromUser() string {
	var line string

	fmt.Scan(&line)

	return line
}

func getIntFromUser() int64 {
	var op int64
	fmt.Scan(&op)

	return op
}

func existsInSlice(s string, slice []string) bool {
	for _, a := range slice {
		if a == s {
			return true
		}
	}
	return false
}

func saveToFile(strings []string, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	for _, s := range strings {
		file.WriteString(s + "\n")
	}
}
