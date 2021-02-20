package main

import (
	"pubsub/exchange"

	"pubsub/handler"
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
	handler.HandleData(data)
}
