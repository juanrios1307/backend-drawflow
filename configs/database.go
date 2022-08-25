package configs

import (
	"log"

	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
)


func NewClient() *dgo.Dgraph {
	// Dial a gRPC connection. The address to dial to can be configured when
	// setting up the dgraph cluster.
	d, err := dgo.DialSlashEndpoint("blue-surf-591387.grpc.us-east-1.aws.cloud.dgraph.io:443","YzQ5ZDI2ZjUyNTMzNzIyYTJkYTVkNzNlNTNiNGE4M2M=")
	if err != nil {
		log.Fatal(err)
	}

	return dgo.NewDgraphClient(
		api.NewDgraphClient(d),
	)
}