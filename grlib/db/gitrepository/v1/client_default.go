package grgitrepositorydb

import (
	"github.com/kazekim/backend-engineer-challenge/grlib/besqlx"
)

type defaultClient struct {
	db besqlx.Client
}

//NewClient new git repository db client
func NewClient(db besqlx.Client) Client {
	return &defaultClient{
		db: db,
	}
}
