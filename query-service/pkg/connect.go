package pkg

import (
	"log"
	"time"

	"github.com/couchbase/gocb/v2"
)

func ConnectCB() *gocb.Scope {
	cluster, err := gocb.Connect("couchbase://couchbase", gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: "Administrator",
			Password: "123456",
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	bucket := cluster.Bucket("my-db")
	err = bucket.WaitUntilReady(5*time.Second, nil)
	if err != nil {
		log.Fatal(err)
	}

	return bucket.Scope("test")

}
