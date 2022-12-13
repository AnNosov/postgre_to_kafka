package entity

import (
	"gopkg.in/guregu/null.v3"
)

type Profile struct {
	Id   int      `json:"id"`
	Name string   `json:"name"`
	Age  null.Int `json:"age"`
}
