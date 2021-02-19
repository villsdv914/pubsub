package main

import (
	"pubsub/exchange"
	"pubsub/handeler"
	"pubsub/publog"
	"pubsub/sqlutils"
)

func main() {
	publog.Logrs.Info("Starting.....")
	sqlutils.SqliteMigrate()
	exchange.Send()
	data, err := exchange.Receive()
	if err != nil{
		panic(err)
	}
	handeler.HandleData(data)
}
