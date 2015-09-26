package humidity
import (
	"net/http"
	"github.com/op/go-logging"
	"strings"
)

func Post(log *logging.Logger, w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Info("[%s] %s %s %s \"%s\" %d %#v", getClientIP(r), r.Method, r.URL, r.Proto, r.Header.Get("user-agent"), r.ContentLength, r.Form)
}

func getClientIP(r *http.Request) string {
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