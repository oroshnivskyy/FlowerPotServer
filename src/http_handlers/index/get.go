package index
import (
	"net/http"
	"github.com/martini-contrib/render"
)

func Get(w http.ResponseWriter, r *http.Request, render render.Render)  {
	render.HTML(200, "index", "index")
}