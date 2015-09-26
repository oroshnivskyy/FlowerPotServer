package routing
import(
	"github.com/go-martini/martini"
	"http_handlers/humidity"
	"http_handlers/index"
)

func Configure(m *martini.ClassicMartini){
	m.Group("/humidity", func(r martini.Router){
		r.Post("", humidity.Post)
		r.Get("", humidity.Get)
	})
	m.Group("/", func(r martini.Router){
		r.Get("", index.Get)
	})
}