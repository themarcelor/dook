package main

import (
	"context"
	"fmt"
	"log"

	"github.com/FerretDB/FerretDB/ferretdb"
)

func main() {
	f, err := ferretdb.New(&ferretdb.Config{
		Listener: ferretdb.ListenerConfig{
			TCP: "127.0.0.1:17027",
		},
		Handler:       "pg",
		PostgreSQLURL: "postgres://127.0.0.1:5432/ferretdb",
	})
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	done := make(chan struct{})

	go func() {
		log.Print(f.Run(ctx))
		close(done)
	}()

	uri := f.MongoDBURI()
	fmt.Println(uri)

	// Use MongoDB URI as usual. For example:
	//
	// import "go.mongodb.org/mongo-driver/mongo"
	//
	// [...]
	//
	// mongo.Connect(ctx, options.Client().ApplyURI(uri))

	cancel()
	<-done

}

