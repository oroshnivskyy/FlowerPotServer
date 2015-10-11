package humidity
import (
	"net/http"
	"github.com/op/go-logging"
	"http_helpers"
)

func Get(log *logging.Logger, w http.ResponseWriter, r *http.Request) (int, string) {
	r.ParseForm()
	log.Info("[%s] %s %s %s \"%s\" %d", http_helpers.GetClientIP(r), r.Method, r.URL, r.Proto, r.Header.Get("user-agent"), r.ContentLength)
	return 200 ,"Hello world"
}