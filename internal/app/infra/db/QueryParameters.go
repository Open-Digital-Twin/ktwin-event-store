package db

import "github.com/scylladb/gocqlx/v2/qb"

type QueryParameters interface {
	GetTable() string
	GetWhereConditions() []qb.Cmp
	GetParametersValues() map[string]interface{}
}

type qParameters struct {
	table            string
	whereConditions  []qb.Cmp
	parametersValues map[string]interface{}
}

func NewQueryParameters(
	table string,
	whereConditions []qb.Cmp,
	parametersValues map[string]interface{},
) QueryParameters {
	return &qParameters{
		table:           table,
		whereConditions: whereConditions,
	}
}

func (q *qParameters) GetTable() string {
	return q.table
}

func (q *qParameters) GetWhereConditions() []qb.Cmp {
	return q.whereConditions
}

func (q *qParameters) GetParametersValues() map[string]interface{} {
	return q.parametersValues
}
