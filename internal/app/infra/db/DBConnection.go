package db

import (
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scylladb/gocqlx/v2/table"
)

type DBConnection interface {
	GetManyWithoutParameters(table *table.Table, conditions qb.M, returnObject interface{}) error
	GetManyWithParameters(table *table.Table, conditions qb.M, returnObject interface{}) error
	GetOneWithParameters(table *table.Table, conditions qb.M, returnObject interface{}) error
	InsertQueryDB(table *table.Table, insertInterface interface{}) error
	DeleteQuery(table *table.Table, conditions qb.M) error
}

func NewDBConnection() DBConnection {
	host := "localhost"

	dbCluster := gocql.NewCluster(host)
	dbCluster.Keyspace = "ktwin"

	return &dbConnection{
		dbCluster: dbCluster,
	}
}

type dbConnection struct {
	dbCluster *gocql.ClusterConfig
}

// No support for pagination
func (db *dbConnection) GetManyWithoutParameters(table *table.Table, conditions qb.M, returnObject interface{}) error {
	session, err := gocqlx.WrapSession(db.dbCluster.CreateSession())

	if err != nil {
		return err
	}

	q := session.Query(table.SelectAll()).BindMap(conditions)

	if err := q.SelectRelease(returnObject); err != nil {
		return err
	}

	return err
}

func (db *dbConnection) GetManyWithParameters(table *table.Table, conditions qb.M, returnObject interface{}) error {
	session, err := gocqlx.WrapSession(db.dbCluster.CreateSession())

	if err != nil {
		return err
	}

	q := session.Query(table.Select()).BindMap(conditions)

	if err := q.SelectRelease(returnObject); err != nil {
		return err
	}

	return err
}

func (db *dbConnection) GetOneWithParameters(table *table.Table, conditions qb.M, returnObject interface{}) error {
	session, err := gocqlx.WrapSession(db.dbCluster.CreateSession())

	if err != nil {
		return err
	}

	q := session.Query(table.Get()).BindMap(conditions)

	if err := q.GetRelease(returnObject); err != nil {
		return err
	}

	return err
}

func (db *dbConnection) InsertQueryDB(table *table.Table, insertInterface interface{}) error {
	session, err := gocqlx.WrapSession(db.dbCluster.CreateSession())

	if err != nil {
		return err
	}

	q := session.Query(table.Insert()).BindStruct(insertInterface)

	if err := q.ExecRelease(); err != nil {
		return err
	}

	return err
}

func (db *dbConnection) DeleteQuery(table *table.Table, conditions qb.M) error {
	session, err := gocqlx.WrapSession(db.dbCluster.CreateSession())

	if err != nil {
		return err
	}

	q := session.Query(table.Delete()).BindMap(conditions)

	if err := q.ExecRelease(); err != nil {
		return err
	}

	return err
}
