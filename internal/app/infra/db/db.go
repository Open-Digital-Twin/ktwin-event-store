package db

import "github.com/gocql/gocql"

func NewDatabaseConnection() *gocql.ClusterConfig {
	cluster := gocql.NewCluster("localhost")
	cluster.Keyspace = "ktwin"
	return cluster
}
