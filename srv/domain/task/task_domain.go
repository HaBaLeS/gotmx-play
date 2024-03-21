package task

import (
	"github.com/gin-gonic/gin"
	"github.com/habales/gotmx-play/srv/database"
	"github.com/habales/gotmx-play/srv/templ"
)

type TaskDomain struct {
	db     *database.Db
	router gin.IRouter
	tpls   *templ.Tpls
}

func (d TaskDomain) init() {

	task := d.router.Group("/task")

	task.GET("/tasklist", func(c *gin.Context) {
		tpl := d.tpls.GetTemplate("tasklist.gohtmx")
		tpl.Execute(c.Writer, d.ListTasks())
	})
	task.POST("/tasklist", func(c *gin.Context) {
		//update order
		tpl := d.tpls.GetTemplate("tasklist.gohtmx")
		tpl.Execute(c.Writer, d.ListTasks())
	})

}

func NewTaskDomain(db *database.Db, router gin.IRouter, tpls *templ.Tpls) *TaskDomain {

	domain := &TaskDomain{
		db:     db,
		router: router,
		tpls:   tpls,
	}

	domain.init()
	return domain
}
