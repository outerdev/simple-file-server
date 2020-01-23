package main

import (
	. "fmt"

	"flag"
	"log"
	"net"
	"net/http"
	"os"
)

const (
	defaultPort = "8123"
	defaultDir  = "."
)

func GetLocalIP() string {

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Error:: " + err.Error() + "\n")
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}

	return ""
}

func main() {
	port := flag.String("p", defaultPort, "static file server port")
	dir := flag.String("d", defaultDir, "the static files directory")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*dir)))

	localIP := GetLocalIP()
	if localIP == "" {
		Println("Error: Could not get local IP!")
		os.Exit(1)
	}

	Printf("Serving directory %s\n    ==> localhost:%s\n    ==> %s:%s\n", *dir, *port, localIP, *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
