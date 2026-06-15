package main

import (
	"e-wallet-transaction-service/cmd"
	"e-wallet-transaction-service/helpers"
)

func main() {

	// Setup logger
	helpers.SetupLogger()
	// Setup config
	helpers.SetupConfig()
	// load database
	helpers.SetupMySql()

	// start http server
	cmd.ServerHttp()

	// start grpc server
	//cmd.ServerGRPC()

}
