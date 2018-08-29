package main

import (
	"ocrawler/distributed_service/rpcsupport"
	"ocrawler/distributed_service/persist"
	"github.com/olivere/elastic"
	"flag"
	"fmt"
)

var port = flag.Int("port", 0, "the port for me to listen on")


func main() {
	flag.Parse()
	err := serveRpc(fmt.Sprintf(":%d", *port), "dating_profile")
	if err != nil {
		panic(err)
	}
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}

	return rpcsupport.ServeRpc(host,
		&persist.ItemSaverService{
			Client: client,
			Index: index,
	})
}
