package srv

import (
	"github.com/habales/gotmx-play/data/apptpl"
	"github.com/habales/gotmx-play/data/static"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type App struct {
	devMode bool
	tplRoot *template.Template
}

func (a *App) RunApp() {
	r := gin.Default()
	a.tplRoot = template.New("root")
	if _, e := os.Open("data"); e == nil {
		a.devMode = true
		a.tplRoot, e = a.tplRoot.ParseGlob("data/apptpl/**")
		if e != nil {
			panic(e)
		}
		r.StaticFS("/static", http.Dir("data/static"))
	} else {
		a.tplRoot, e = a.tplRoot.ParseFS(apptpl.Embedded, "*")
		if e != nil {
			panic(e)
		}
		r.StaticFS("/static", http.FS(static.Embedded))
	}

	r.GET("/", func(c *gin.Context) {
		tpl := a.GetTemplate("main.gohtmx")
		tpl.Execute(c.Writer, nil)
	})

	r.POST("/clicked", func(c *gin.Context) {
		tpl := a.GetTemplate("click_result.gohtmx")
		tpl.Execute(c.Writer, nil)
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func (a *App) GetTemplate(name string) *template.Template {
	if a.devMode {
		var err error
		a.tplRoot = template.New("root") //FIXME this will not include any custom functions, also might be slow when there are manny templates
		a.tplRoot, err = a.tplRoot.ParseGlob("data/apptpl/**")
		if err != nil {
			log.Printf("could not parse a template %v", err)
		}
	}
	return a.tplRoot.Lookup(name)
}
