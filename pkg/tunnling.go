package pkg
// So it basiclly. The client connect to the server. Let's say my service litent on port 8080 and forward traffic on my localhost port 7070 to where very that target server is so any request on port 7070 will be forwarded to the target server. On the target server am going to have mu client so my client will revice that message and forward it to the local machine on the compromise system the compromise system will response to my client so i take that response and forward it to my server am i correct


import (
	"io"
	"log"
	"log/slog"
	"net"
)

type Tunnling struct {
	dest string
	src string
};

func forward(src net.Conn ,tun *Tunnling) {

    // Connect to the destination (localhost:4444) just for testing
	// The destnation here is the internal services for example SQL
	// Connecting on port 3306
    dest, err := net.Dial("tcp", tun.dest)
    if err != nil {
        log.Printf("Failed to connect to destination: %v", err)
        src.Close()
        return
    }
    defer dest.Close()

    go func() { io.Copy(dest, src) }() // Client -> Destination
    io.Copy(src, dest)                 // Destination -> Client
}

func InitTunnling() *Tunnling {
	return &Tunnling{};
}



func (tun *Tunnling) SetDest(dest string) *Tunnling {
	tun.dest = dest;
	return tun;
}


func (tun *Tunnling) SetSrc(src string ) *Tunnling {
	tun.src = src;
	return tun;
}

func (tun *Tunnling) RunTun() {

    listener, err := net.Listen("tcp",tun.src)
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
