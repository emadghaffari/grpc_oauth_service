package cassandra

import (
	"github.com/gocql/gocql"
)

var (
	sesstion *gocql.Session
)

func init()  {
	cluster := gocql.NewCluster("cassandra")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum
	var err error 
	if sesstion, err = cluster.CreateSession(); err != nil {
		panic(err)
	}
}

// GetSesstion func
func GetSesstion() *gocql.Session {
	return sesstion
}