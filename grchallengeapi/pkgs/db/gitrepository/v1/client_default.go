package grcgitrepositorydb

import (
	"github.com/kazekim/backend-engineer-challenge/grlib/besqlx"
)

type defaultClient struct {
	db besqlx.Client
}

func NewClient(db besqlx.Client) Client {
	return &defaultClient{
		db: db,
	}
}
