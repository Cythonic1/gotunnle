package pkg

// So it basically. The client connects to the server. Let's say my service listens on port 8080 and forwards traffic on my localhost port 7070 to where very that target server is so any request on port 7070 will be forwarded to the target server. On the target server am going to have my client so my client will receive that message and forward it to the local machine on the compromised system the compromised system will respond to my client so i take that response and forward it to my server am i correct

import (
	"io"
	"log"
	"log/slog"
	"net"
)

type Tunnling struct {
	// Localport which port gonna be listen for connections from clients
	localPort string
	bindPort  string
}

//TODO : Implement client action.

// Still need some modification in here.
// i need to make it listen on the bindPort or get traffic from there.
// and then forward it to the client
func forward(client net.Conn, tun *Tunnling) {
	// Connect to the bindPortination (localhost:4444) just for testing
	// The bindPortination here is the internal services for example SQL
	// Connecting on port 3306
	bindPort, err := net.Dial("tcp", tun.bindPort)
	if err != nil {
		log.Printf("Failed to connect to bindPortination: %v", err)
		client.Close()
		return
	}
	defer bindPort.Close()

	go func() { io.Copy(bindPort, client) }() // Client -> Destination
	io.Copy(client, bindPort)                 // Destination -> Client
}

func InitTunnling() *Tunnling {
	return &Tunnling{}
}

func (tun *Tunnling) SetBindPort(bindPort string) *Tunnling {
	tun.bindPort = bindPort
	return tun
}

func (tun *Tunnling) SetLocalPort(localPort string) *Tunnling {
	tun.localPort = localPort
	return tun
}

func (tun *Tunnling) RunTun() {

	listener, err := net.Listen("tcp", tun.localPort)
	if err != nil {
		slog.Error("Failed to start server", "error", err)
		return
	}
	defer listener.Close()

	slog.Info("Server started", "addr", listener.Addr().String())

	for {
		client, err := listener.Accept()
		if err != nil {
			slog.Error("Failed to accept client", "error", err)
			continue
		}
		slog.Info("New client connected", "addr", client.RemoteAddr())

		// Forward traffic to localhost:4444
		go forward(client, tun)
	}
}
