package entity

import "database/sql"

type Profile struct {
	Id   int           `json:"id"`
	Name string        `json:"name"`
	Age  sql.NullInt64 `json:"age"`
}
