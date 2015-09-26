package humidity
import (
	"net/http"
	"github.com/op/go-logging"
)

func Get(log *logging.Logger, w http.ResponseWriter, r *http.Request) (int, string) {
	r.ParseForm()
	log.Info("[%s] %s %s %s \"%s\" %d", getClientIP(r), r.Method, r.URL, r.Proto, r.Header.Get("user-agent"), r.ContentLength)
	return 200 ,"Hello world"
}