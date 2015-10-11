package routing
import(
	"github.com/go-martini/martini"
	"http_handlers/humidity"
	"http_handlers/current_humidity"
	"http_handlers/index"
	"net/http"
)

func Configure(m *martini.ClassicMartini){
	m.Group("/humidity", func(r martini.Router){
		r.Post("", humidity.Post)
		r.Get("", humidity.Get)
	})
	m.Get("/current/humidity", current_humidity.Get)

	m.Group("/", func(r martini.Router){
		r.Get("", index.Get)
	})

	static := martini.Static("templates", martini.StaticOptions{Fallback: "/index.tmpl", Exclude: "/api/v"})
	m.NotFound(static, http.NotFound)
	m.Use(martini.Static("static"))
}