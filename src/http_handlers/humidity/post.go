package humidity
import (
	"net/http"
	"github.com/op/go-logging"
	"github.com/dancannon/gorethink"
	"models"
	"http_helpers"

)

func Post(session *gorethink.Session,log *logging.Logger, w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Info("[%s] %s %s %s %d %#v", r.Method, r.URL, r.Proto, r.Header.Get("user-agent"), r.ContentLength, r.Form)
	humidity, err :=models.NewRecord(r.Form["data"], r.Form.Get("name"), http_helpers.GetClientIP(r))
	if err !=nil{
		log.Error("Error writing to database: %s", err)
		return
	}
	_, err = gorethink.Table("humidity").Insert(humidity).RunWrite(session)
	if err !=nil{
		log.Error("Error writing to database: %s", err)
		return
	}
}

