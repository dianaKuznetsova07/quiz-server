package db

import (
	"github.com/Masterminds/squirrel"
	"github.com/elgris/stom"
)

const SelectTag = "db"
const InsertTag = "insert"

type Entity interface {
	GetTableName() string
}

func InsertEntities(qb squirrel.InsertBuilder, st stom.ToMapper, cols []string, row Entity) squirrel.InsertBuilder {
	m, err := st.ToMap(row)
	if err != nil {
		panic(err)
	}

	values := make([]interface{}, len(cols))
	for i, c := range cols {
		values[i] = m[c]
	}

	return qb.Values(values...)
}
