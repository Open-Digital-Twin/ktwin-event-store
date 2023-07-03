package db

import (
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scylladb/gocqlx/v2/table"
)

type QueryParameters interface {
	GetTable() *table.Table
	GetParametersValues() qb.M
}

type qParameters struct {
	table            *table.Table
	parametersValues qb.M
}

func NewQueryParameters(
	table *table.Table,
	parametersValues qb.M,
) QueryParameters {
	return &qParameters{
		table:            table,
		parametersValues: parametersValues,
	}
}

func (q *qParameters) GetTable() *table.Table {
	return q.table
}

func (q *qParameters) GetParametersValues() qb.M {
	return q.parametersValues
}
