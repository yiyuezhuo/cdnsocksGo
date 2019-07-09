package main

import (
	"fmt"
	"log"
	"net"
	"net/url"
	"sync"

	"github.com/gorilla/websocket"
)

func main() {
	config := LoadConfig("config-client.json")
	config.Print()

	remote_ip_port := config.RemoteIp + ":" + config.RemotePort

	var wg sync.WaitGroup
	for _, route := range config.Routes {
		ip_port := route.LocalIp + ":" + route.LocalPort
		wg.Add(1)
		go func(ip_port, remote_ip_port, url string) {
			defer wg.Done()
			listener(ip_port, remote_ip_port, url)
		}(ip_port, remote_ip_port, route.Url)
	}

	fmt.Println("Goroutine start finished, wait to exit")
	wg.Wait()
	fmt.Println("Normal Exit")
}

func listener(ip_port, remote_ip_port, remoteUrl string) error {
	l, err := net.Listen(ip_port, remoteUrl)
	if err != nil {
		log.Println("start listen fail", ip_port, remoteUrl, err)
		return err
	}
	defer l.Close()

	log.Println("Listen to:", ip_port, " link to:", remoteUrl)

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("Error: Accept fail", ip_port, remoteUrl, err)
			continue
		}

		go handleRequest(conn, remote_ip_port, remoteUrl)
	}

}

func handleRequest(localConn net.Conn, remote_ip_port, remoteUrl string) {
	defer localConn.Close()

	if remoteUrl[0] != '/' {
		remoteUrl = "/" + remoteUrl
	}
	wsUrl := url.URL{Scheme: "ws", Host: remote_ip_port, Path: remoteUrl}

	fmt.Println("Connect to", wsUrl)

	remoteConn, _, err := websocket.DefaultDialer.Dial(wsUrl.String(), nil)
	if err != nil {
		log.Println("Error Websocket dial fail", err)
		return
	}
	defer remoteConn.Close()

	go websocket_to_socket(localConn, remoteConn)
	socket_to_websocket(localConn, remoteConn)

	log.Println("Disconnect:", remote_ip_port, remoteUrl)
}

func websocket_to_socket(localConn net.Conn, remoteConn *websocket.Conn) {
	for {
		_, message, err := remoteConn.ReadMessage()
		if err != nil {
			log.Println("Error: remoteConn.ReadMessage fail", err)
			return
		}
		fmt.Println("remote to local", len(message), message[:10])

		localConn.Write(message)
		// Write return (int, err), Should I use those information to do something?
	}
}

func socket_to_websocket(localConn net.Conn, remoteConn *websocket.Conn) {
	buf := make([]byte, 2048)

	for {
		readLen, err := localConn.Read(buf)
		if err != nil {
			log.Println("Error Read from local fail", err)
		}
		fmt.Println("local to remote", readLen, buf[:10])

		remoteConn.WriteMessage(websocket.BinaryMessage, buf[:readLen])
		// WriteMessage return error, Should I use the information?
	}
}
