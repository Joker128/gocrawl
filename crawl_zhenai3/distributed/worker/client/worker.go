package client

import (
	"gocrawl/crawl_zhenai3/distributed/config"
	"gocrawl/crawl_zhenai3/distributed/worker"
	"gocrawl/crawl_zhenai3/engine"
	"net/rpc"
)

func CreateProcessor(clientChan chan *rpc.Client) (engine.Processor) {
	//client, err := rpcsupport.NewClient(fmt.Sprintf(":%d", config.WorkerPort0))
	//if err != nil {
	//	return nil, err
	//}

	return func(req engine.Request) (engine.ParseResult, error) {
		sReq := worker.SerializeRequest(req)
		var sResult worker.ParseResult

		client := <-clientChan

		err := client.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(sResult), nil
	}
}
