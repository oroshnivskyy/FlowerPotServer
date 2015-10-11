package http_helpers

import (
	"net/http"
	"strings"
)

func GetClientIP(r *http.Request) string {
	clientIP := strings.TrimSpace(r.Header.Get("X-Real-IP"))
	if len(clientIP) > 0 {
		return clientIP
	}
	clientIP = r.Header.Get("X-Forwarded-For")
	clientIP = strings.TrimSpace(strings.Split(clientIP, ",")[0])
	if len(clientIP) > 0 {
		return clientIP
	}
	return strings.TrimSpace(r.RemoteAddr)
}
