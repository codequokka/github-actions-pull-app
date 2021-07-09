package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

const APP_VER = "1.0.0"

// GetIpAddr gets one of the IP address except 127.0.0.1.
func GetIpAddr() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}

	var IpAddr string
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				IpAddr = ipnet.IP.String()
			}
		}
	}

	return IpAddr
}

// HelloServer responds to requests with app ver, ip address and the given URL path.
func HelloServer(w http.ResponseWriter, r *http.Request) {
	var IpAddr string = GetIpAddr()
	fmt.Fprintf(w, "App ver: %s, Ip: %s, Path: %s", APP_VER, IpAddr, r.URL.Path)
	log.Printf("App ver: %s, Ip: %s, Path: %s", APP_VER, IpAddr, r.URL.Path)
}

func main() {
	var addr string = ":8180"
	handler := http.HandlerFunc(HelloServer)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatalf("Could not listen on port %s %v", addr, err)
	}
}
