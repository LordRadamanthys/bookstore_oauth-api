package cassandra

import (
	"github.com/gocql/gocql"
)

var (
	session *gocql.Session
)

func init() {
	// connect to cassandra cluster
	// 	Run commands in cqlsh like this:

	// docker exec -it cassandra cqlsh
	// Stop the service like this:

	// docker stop cassandra
	// Start it again like this:

	// docker start cassandra
	cluster := gocql.NewCluster("localhost")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum
	var err error
	if session, err = cluster.CreateSession(); err != nil {
		panic(err)
	}
}

func GetSession() *gocql.Session {
	return session
}
