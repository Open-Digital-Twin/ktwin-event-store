package db

import (
	"fmt"
	"time"

	"github.com/Open-Digital-Twin/ktwin-event-store/internal/app/config"

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
	CloseSession()
}

func NewDBConnection() DBConnection {
	host := config.GetConfig("DB_HOST")

	dbCluster := gocql.NewCluster(host)
	dbCluster.Keyspace = config.GetConfig("DB_KEYSPACE")
	timeoutString := config.GetConfig("TIMEOUT")
	timeDuration, err := time.ParseDuration(timeoutString + "ms")

	if err != nil {
		fmt.Println("Error while parsing timeout: ", timeoutString)
	}

	if err == nil && timeDuration != 0 {
		fmt.Printf("Setting timeout value: %d\n", timeDuration)
		dbCluster.Timeout = timeDuration
	}

	return &dbConnection{
		dbCluster: dbCluster,
	}
}

type dbConnection struct {
	dbCluster *gocql.ClusterConfig
	session   *gocqlx.Session
}

func (db *dbConnection) loadSession() (*gocqlx.Session, error) {
	if db.session == nil {
		newSession, err := gocqlx.WrapSession(db.dbCluster.CreateSession())
		if err != nil {
			return nil, err
		}
		db.session = &newSession
	}
	return db.session, nil
}

// No support for pagination
func (db *dbConnection) GetManyWithoutParameters(table *table.Table, conditions qb.M, returnObject interface{}) error {
	session, err := db.loadSession()

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
	session, err := db.loadSession()

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
	session, err := db.loadSession()

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
	session, err := db.loadSession()

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
	session, err := db.loadSession()

	if err != nil {
		return err
	}

	q := session.Query(table.Delete()).BindMap(conditions)

	if err := q.ExecRelease(); err != nil {
		return err
	}

	return err
}

func (db *dbConnection) CloseSession() {
	defer db.session.Close()
}
