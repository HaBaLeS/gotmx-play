package task

import "strconv"

func (d *TaskDomain) ListTasks() []TaskEntity {
	retVal := make([]TaskEntity, 0)
	err := d.db.DB.Select(&retVal, "select * from task order by order_num")
	if err != nil {
		panic(err)
	}
	return retVal
}

func (d *TaskDomain) UpdateOrder(newOrder []string) {
	for k, v := range newOrder {
		d.db.DB.MustExec("update task set order_num=$1 where id=$2", k, v)
	}
}

func (d *TaskDomain) AddTask(data AddTask) {
	d.db.DB.MustExec("insert into task (title,description, order_num) values ( $1,$2, 0);", data.Title, data.Description)
}

func (d *TaskDomain) DeleteTask(id string) {
	if intId, err := strconv.Atoi(id); err != nil {
		panic(err)
	} else {
		d.db.DB.MustExec("delete from task where id = $1", intId)
	}
}
