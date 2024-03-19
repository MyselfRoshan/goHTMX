package models

import (
	"database/sql"
	"time"
)

type Todo struct {
	Id          string
	Title       string
	CreatedOn   time.Time
	CompletedOn sql.NullTime
}
