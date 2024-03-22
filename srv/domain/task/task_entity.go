package task

import "database/sql"

type TaskEntity struct {
	Id       int            `db:"id"`
	Title    string         `db:"title"`
	Desc     sql.NullString `db:"description"`
	OrderNum int            `db:"order_num"`
}
