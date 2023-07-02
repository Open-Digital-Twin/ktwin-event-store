package db

import (
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
)

type DBConnection interface {
	GetManyWithParameters(queryParameters QueryParameters, returnObject interface{}) error
	GetOneWithParameters(queryParameters QueryParameters, returnObject interface{}) error
	InsertQueryDB(table string, columns []string, insertInterface interface{}) error
	DeleteTableByColumn(table string, whereConditions []qb.Cmp, deleteInterface interface{}) error
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
func (db *dbConnection) GetManyWithParameters(queryParameters QueryParameters, returnObject interface{}) error {
	session, err := gocqlx.WrapSession(db.dbCluster.CreateSession())

	if err != nil {
		return err
	}

	query := qb.Select(queryParameters.GetTable())
	statement, names := query.ToCql()

	q := session.Query(statement, names)

	if queryParameters.GetParametersValues() != nil {
		q = q.BindMap(queryParameters.GetParametersValues())
	}

	err = q.Select(returnObject)

	if err != nil {
		return err
	}

	return err
}

func (db *dbConnection) GetOneWithParameters(queryParameters QueryParameters, returnObject interface{}) error {
	session, err := gocqlx.WrapSession(db.dbCluster.CreateSession())

	if err != nil {
		return err
	}

	query := qb.Select(queryParameters.GetTable())
	statement, names := query.ToCql()

	q := session.Query(statement, names)

	if queryParameters.GetParametersValues() != nil {
		q = q.BindMap(queryParameters.GetParametersValues())
	}

	err = q.Get(returnObject)

	if err != nil {
		return err
	}

	return err
}

func (db *dbConnection) InsertQueryDB(table string, columns []string, insertInterface interface{}) error {
	session, err := gocqlx.WrapSession(db.dbCluster.CreateSession())

	if err != nil {
		return err
	}

	insertQuery := qb.Insert(table).Columns(columns...).Query(session)
	insertQuery.BindStruct(insertInterface)

	if err := insertQuery.ExecRelease(); err != nil {
		return err
	}

	return err
}

func (db *dbConnection) DeleteTableByColumn(table string, whereConditions []qb.Cmp, deleteInterface interface{}) error {
	session, err := gocqlx.WrapSession(db.dbCluster.CreateSession())

	if err != nil {
		return err
	}

	deleteQuery := qb.Delete(table).Where(whereConditions...).Query(session)
	deleteQuery.BindStruct(deleteInterface)

	if err := deleteQuery.ExecRelease(); err != nil {
		return err
	}

	return err
}
