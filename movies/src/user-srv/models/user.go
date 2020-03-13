package models

import "database/sql"

type User struct {
	Id   int64  `json:"user_id" db:"id"`
	Name string `json:"user_name" db:"name"`
	Password string `json:"password" db:"password"`
	CreateTime string `json:"create_at" db:"create_time"`
	Email    sql.NullString `json:"email" db:"email"`
	Phone    sql.NullString `json:"phone" db:"phone"`
	Address sql.NullString `json:"address" db:"address"`
	UpdateTime string `json:"updatetime" db:"update_time"`
}


