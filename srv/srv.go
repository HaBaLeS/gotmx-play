package srv

import (
	"github.com/habales/gotmx-play/data/apptpl"
	"github.com/habales/gotmx-play/data/static"
	"html/template"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func RunApp() {
	r := gin.Default()
	tplRoot := template.New("root")
	if _, e := os.Open("data"); e == nil {
		tplRoot, e = tplRoot.ParseGlob("data/apptpl/**")
		if e != nil {
			panic(e)
		}
		r.StaticFS("/static", http.Dir("data/static"))
	} else {
		tplRoot, e = tplRoot.ParseFS(apptpl.Embedded, "*")
		if e != nil {
			panic(e)
		}
		r.StaticFS("/static", http.FS(static.Embedded))
	}

	r.GET("/", func(c *gin.Context) {
		tpl := tplRoot.Lookup("main.gohtmx")
		tpl.Execute(c.Writer, nil)
	})

	r.POST("/clicked", func(c *gin.Context) {
		tpl := tplRoot.Lookup("click_result.gohtmx")
		tpl.Execute(c.Writer, nil)
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
