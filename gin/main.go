package main

import (
	"net"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/client-ip", getClientIP)

	router.Run()
}

func getClientIP(c *gin.Context) {
	ip := getPublicIP(c.Request)
	c.String(http.StatusOK, ip)
}

func getPublicIP(req *http.Request) string {
	xff := req.Header.Get("X-Forwarded-For")
	if xff != "" {
		return strings.TrimSpace(strings.Split(xff, ",")[0])
	}

	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		return ""
	}

	return ip
}
