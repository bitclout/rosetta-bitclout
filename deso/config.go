package deso

import (
	"errors"
	"strings"

	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/deso-protocol/core/lib"
	"github.com/spf13/viper"
)

type Config struct {
	Mode                   Mode
	Network                *types.NetworkIdentifier
	Params                 *lib.DeSoParams
	Currency               *types.Currency
	GenesisBlockIdentifier *types.BlockIdentifier
	Port                   int
	NodePort               int
	DataDirectory          string
	MinerPublicKeys        []string
	TXIndex                bool
	Regtest                bool
}

func LoadConfig() (*Config, error) {
	result := Config{}

	switch result.Mode = Mode(strings.ToUpper(viper.GetString("mode"))); result.Mode {
	case Online, Offline:
	default:
		return nil, errors.New("unknown mode")
	}

	result.Currency = &Currency

	switch network := Network(strings.ToUpper(viper.GetString("network"))); network {
	case Mainnet:
		result.Params = &lib.DeSoMainnetParams
	case Testnet:
		result.Params = &lib.DeSoTestnetParams
		result.Currency.Symbol = "t" + result.Currency.Symbol
	default:
		return nil, errors.New("unknown network")
	}

	result.Network = &types.NetworkIdentifier{
		Blockchain: "DeSo",
		Network:    result.Params.NetworkType.String(),
	}

	result.GenesisBlockIdentifier = &types.BlockIdentifier{
		Hash: result.Params.GenesisBlockHashHex,
	}

	result.DataDirectory = viper.GetString("data-directory")
	result.Port = viper.GetInt("port")
	result.NodePort = viper.GetInt("node-port")
	result.MinerPublicKeys = viper.GetStringSlice("miner-public-keys")
	result.TXIndex = viper.GetBool("txindex")
	result.Regtest = viper.GetBool("regtest")

	return &result, nil
}
