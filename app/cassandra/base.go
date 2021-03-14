package cassandra

import (
	"log"
	"os"

	"chat/app"
	"github.com/gocql/gocql"
)

var Session *gocql.Session

func init() {
	app.LoadEnv()
	cluster := gocql.NewCluster(os.Getenv("DB_HOST"))
	cluster.Keyspace = os.Getenv("KEYSPACE")
	var err error
	Session, err = cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
}
