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

type AddTask struct {
	Title       string `form:"title"`
	Description string `form:"description"`
}

func (d *TaskDomain) init() {

	task := d.router.Group("/task")

	task.GET("/tasklist", func(c *gin.Context) {
		tpl := d.tpls.GetTemplate("tasklist.gohtmx")
		tpl.Execute(c.Writer, d.ListTasks())
	})
	task.POST("/tasklist", func(c *gin.Context) {
		//update order
		tpl := d.tpls.GetTemplate("tasklist.gohtmx")
		if newOrder, ok := c.GetPostFormArray("item"); ok {
			d.UpdateOrder(newOrder)
		}
		tpl.Execute(c.Writer, d.ListTasks())
	})

	task.GET("/create", func(c *gin.Context) {
		d.tpls.GetTemplate("taskedit.gohtmx").Execute(c.Writer, nil)
	})

	task.PUT("/create", func(c *gin.Context) {
		var data AddTask
		if e := c.ShouldBind(&data); e != nil {
			panic(e)
		}
		d.AddTask(data)
		c.Header("HX-Trigger", "newTask")
		c.Writer.WriteString("<button hx-get=\"/task/create\" class=\"button is-primary\" hx-swap=\"outerHTML\">Add Task</button>")
	})

	task.DELETE("/:id", func(c *gin.Context) {
		idToDelete := c.Param("id")
		d.DeleteTask(idToDelete)
		c.Header("HX-Trigger", "newTask") //fixme add a task delete or list update tirgger
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
