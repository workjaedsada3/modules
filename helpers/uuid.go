package helpers

import (
	mssql "github.com/microsoft/go-mssqldb"
)

type UUID struct {
	mssql.UniqueIdentifier
}

func (u UUID) MarshalText() (text []byte, err error) {
	text = []byte(u.String())
	return
}
