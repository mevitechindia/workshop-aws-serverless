package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func getIPAddress(r *http.Request) string {
	// Check the X-Forwarded-For header for proxies
	xff := r.Header.Get("X-Forwarded-For")
	if xff != "" {
		ips := strings.Split(xff, ",")
		return strings.TrimSpace(ips[0])
	}

	// Check the X-Real-IP header
	xri := r.Header.Get("X-Real-IP")
	if xri != "" {
		return xri
	}

	// Default to using the remote address
	ip := r.RemoteAddr
	if colon := strings.LastIndex(ip, ":"); colon != -1 {
		ip = ip[:colon]
	}
	return ip
}

func ipHandler(w http.ResponseWriter, r *http.Request) {
	ip := getIPAddress(r)
	fmt.Fprintf(w, "Your IP address is: %s\n", ip)
}

func main() {
	http.HandleFunc("/ipaddress", ipHandler)
	port := ":8080"
	log.Printf("Server is running on port %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
