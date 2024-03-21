package static

import (
	"github.com/gin-gonic/gin"
	"github.com/habales/gotmx-play/data/static"
	"github.com/habales/gotmx-play/srv/templ"
	"net/http"
)

func InitStatic(tpls *templ.Tpls, router gin.IRouter) {

	if tpls.DevMode {
		router.StaticFS("/static", http.Dir("data/static"))
	} else {
		router.StaticFS("/static", http.FS(static.Embedded))
	}
}
