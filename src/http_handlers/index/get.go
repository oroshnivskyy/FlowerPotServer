package index
import (
	"github.com/martini-contrib/render"
)

func Get(r render.Render)  {
	r.HTML(200, "index", "rosh")
}