package db

import (
	"consumer-service/event"
	"log"
	"time"

	"github.com/couchbase/gocb/v2"
)

var UserAccount event.UserAccountEvent
var UserAccountList []event.UserAccountEvent

func ConnectCB() *gocb.Collection {
	time.Sleep(time.Second * 10)

	cluster, err := gocb.Connect("couchbase://couchbase", gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: "Administrator",
			Password: "123456",
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	err = cluster.Buckets().CreateBucket(gocb.CreateBucketSettings{
		BucketSettings: gocb.BucketSettings{
			Name:       "my-db",
			RAMQuotaMB: 101,
			BucketType: gocb.EphemeralBucketType,
		}}, nil)

	bucket := cluster.Bucket("my-db")

	// err = bucket.WaitUntilReady(5*time.Second, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	time.Sleep(time.Second * 3)
	_, err = cluster.Query("CREATE SCOPE `my-db`.test IF NOT EXISTS", nil)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Second * 3)
	_, err = cluster.Query("CREATE COLLECTION `my-db`.`test`.useraccounts IF NOT EXISTS", nil)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Second * 3)
	query := "CREATE PRIMARY INDEX ON `my-db`.`test`.`useraccounts`;"
	_, err = bucket.Scope("test").Query(query, nil)
	if err != nil {
		panic(err.Error())
	}

	return bucket.Scope("test").Collection("useraccounts")

}
