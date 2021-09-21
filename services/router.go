package services

import (
	"github.com/deso-protocol/rosetta-deso/deso"
	"net/http"

	"github.com/coinbase/rosetta-sdk-go/asserter"
	"github.com/coinbase/rosetta-sdk-go/server"
)

func NewBlockchainRouter(
	config *deso.Config,
	node *deso.Node,
	asserter *asserter.Asserter,
) http.Handler {
	networkAPIService := NewNetworkAPIService(config, node)
	networkAPIController := server.NewNetworkAPIController(networkAPIService, asserter)

	blockAPIService := NewBlockAPIService(config, node)
	blockAPIController := server.NewBlockAPIController(blockAPIService, asserter)

	accountAPIService := NewAccountAPIService(config, node)
	accountAPIController := server.NewAccountAPIController(accountAPIService, asserter)

	constructionAPIService := NewConstructionAPIService(config, node)
	constructionAPIController := server.NewConstructionAPIController(constructionAPIService, asserter)

	mempoolAPIService := NewMempoolAPIService(config, node)
	mempoolAPIController := server.NewMempoolAPIController(mempoolAPIService, asserter)

	return server.NewRouter(
		networkAPIController,
		blockAPIController,
		accountAPIController,
		constructionAPIController,
		mempoolAPIController,
	)
}
