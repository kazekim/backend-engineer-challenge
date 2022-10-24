package besqlx

import (
	"database/sql"
	"github.com/rs/xid"
	"time"
)

type IDModel struct {
	Id        string       `db:"id"`
	CreatedAt sql.NullTime `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

func (u *IDModel) BeforeCreate() (err error) {

	if u.Id == "" {
		id := xid.New().String()
		u.Id = id
	}

	t := time.Now()
	if !u.CreatedAt.Valid {
		u.CreatedAt = sql.NullTime{
			Time:  t,
			Valid: true,
		}
	}

	if !u.UpdatedAt.Valid {
		u.UpdatedAt = sql.NullTime{
			Time:  t,
			Valid: true,
		}
	}

	return
}
