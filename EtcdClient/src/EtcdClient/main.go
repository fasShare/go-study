package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/client"
	"log"
	"time"
)

func main() {
	cfg := client.Config{
		Endpoints:               []string{"http://127.0.0.1:2379"},
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := client.New(cfg)
	if err != nil {
		log.Fatalf("New CLient err:%v\n", err)
	}
	keys_api := client.NewKeysAPI(c)
	set_res, err := keys_api.Set(context.Background(), "/my_api_client", "Hello api client", nil)
	if err != nil {
		log.Println("Set my_api_client err:%v\n", err)
	} else {
		log.Printf("Set my_api_client succeed and the metadata:%q\n", set_res)
	}
	get_res, err := keys_api.Get(context.Background(), "/", nil)
	if err != nil {
		log.Fatalf("Get my_api_client er:%v\n", err)
	} else {
		for _, node := range get_res.Node.Nodes {
			fmt.Println(node.Key)
		}
	}

	watcher := keys_api.Watcher("my_api_client", nil)

	if watcher_res, err := watcher.Next(context.Background()); err != nil {
		log.Fatalf("watcher Next return err:%v\n", err)
	} else {
		log.Println(watcher_res)
		get_res, err := keys_api.Get(context.Background(), "my_api_client", nil)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(get_res)
	}
}
