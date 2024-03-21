package srv

import (
	"github.com/gin-gonic/gin"
	"github.com/habales/gotmx-play/srv/database"
	"github.com/habales/gotmx-play/srv/domain/static"
	"github.com/habales/gotmx-play/srv/domain/task"
	"github.com/habales/gotmx-play/srv/templ"
)

type App struct {
	Router *gin.Engine
	td     *task.TaskDomain
}

func (a *App) RunApp() {
	a.Router = gin.Default()

	db := database.NewDao()
	tpls := templ.NewTpls()

	//Init Static Ressource Serving
	static.InitStatic(tpls, a.Router)

	//Initi Domain handling Tasks
	a.td = task.NewTaskDomain(db, a.Router, tpls)

	//Entrypoint
	a.Router.GET("/", func(c *gin.Context) {
		tpl := tpls.GetTemplate("main.gohtmx")
		tpl.Execute(c.Writer, nil)
	})

	/*a.Router.POST("/clicked", func(c *gin.Context) {
		tpl := a.GetTemplate("click_result.gohtmx")
		tpl.Execute(c.Writer, nil)
	})*/

	a.Router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
