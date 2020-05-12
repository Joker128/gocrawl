package main

import (
	"flag"
	"fmt"
	"github.com/olivere/elastic"
	"gocrawl/crawl_zhenai3/distributed/config"
	"gocrawl/crawl_zhenai3/distributed/persist"
	"gocrawl/crawl_zhenai3/distributed/rpcsupport"
	"log"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port ... ")
		return
	}

	//如果发生错误，Fatal()会强制退出。。
	log.Fatal(serveRpc(fmt.Sprintf(":%d", *port), config.ElasticIndex))

}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))

	if err != nil {
		return err
	}
	return rpcsupport.ServeRpc(host, &persist.ItemSaveService{
		Client: client,
		Index:  index,
	})
}
