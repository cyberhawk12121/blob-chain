package node

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"strings"
)

type Transport interface {
	StartListening() error
	SendMessage(msg Message, target string)
}

type udpTransport struct {
	localAddr string
	service   Service
}

func NewUDPTransport(localAddr string, service Service) Transport {
	return &udpTransport{
		localAddr: localAddr,
		service:   service,
	}
}

func (t *udpTransport) StartListening() error {
	udpAddr, err := net.ResolveUDPAddr("udp", t.localAddr)
	if err != nil {
		return fmt.Errorf("failed to resolve UDP addr: %w", err)
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		return fmt.Errorf("failed to listen on UDP: %w", err)
	}

	log.Printf("Listening on %s", t.localAddr)

	go func() {
		defer conn.Close()
		buf := make([]byte, 2048)
		for {
			rlen, remoteAddr, err := conn.ReadFromUDP(buf)
			if err != nil {
				log.Printf("Error reading UDP: %v", err)
				continue
			}

			msgStr := strings.TrimSpace(string(bytes.TrimSpace(buf[:rlen])))
			t.service.HandleIncomingMessages(Message{Content: msgStr}, remoteAddr.String())
		}
	}()

	return nil
}

func (t *udpTransport) SendMessage(msg Message, target string) {
	udpAddr, err := net.ResolveUDPAddr("udp", target)
	if err != nil {
		log.Printf("Error resolving target %s: %v", target, err)
		return
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		log.Printf("Error connecting to %s: %v", target, err)
		return
	}
	defer conn.Close()

	_, err = conn.Write([]byte(msg.Content))
	if err != nil {
		log.Printf("Error sending message to %s: %v", target, err)
	}
}
