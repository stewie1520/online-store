// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package domain

import (
	"database/sql"
)

type Category struct {
	ID        int64
	Name      string
	Slug      string
	Author    sql.NullInt32
	Isactive  sql.NullBool
	Isdeleted sql.NullBool
}
