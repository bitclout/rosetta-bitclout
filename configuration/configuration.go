package configuration

import (
	"errors"
	"strings"

	"github.com/bitclout/core/lib"
	"github.com/bitclout/rosetta-bitclout/bitclout"
	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/spf13/viper"
)

type Mode string
type Network string

const (
	Online  Mode = "ONLINE"
	Offline Mode = "OFFLINE"

	Mainnet Network = "MAINNET"
	Testnet Network = "TESTNET"
)

type Configuration struct {
	Mode                   Mode
	Network                *types.NetworkIdentifier
	Params                 *lib.BitCloutParams
	Currency               *types.Currency
	GenesisBlockIdentifier *types.BlockIdentifier
	Port                   int
	NodePort               int
	DataDirectory          string
	MinerPublicKeys        []string
}

func LoadConfiguration() (*Configuration, error) {
	result := Configuration{}

	switch result.Mode = Mode(strings.ToUpper(viper.GetString("mode"))); result.Mode {
	case Online, Offline:
	default:
		return nil, errors.New("unknown mode")
	}

	result.Currency = &bitclout.Currency

	switch network := Network(strings.ToUpper(viper.GetString("network"))); network {
	case Mainnet:
		result.Params = &lib.BitCloutMainnetParams
	case Testnet:
		result.Params = &lib.BitCloutTestnetParams
		result.Currency.Symbol = "t" + result.Currency.Symbol
	default:
		return nil, errors.New("unknown network")
	}

	result.Network = &types.NetworkIdentifier{
		Blockchain: "Bitclout",
		Network:    result.Params.NetworkType.String(),
	}

	result.GenesisBlockIdentifier = &types.BlockIdentifier{
		Hash: result.Params.GenesisBlockHashHex,
	}

	result.DataDirectory = viper.GetString("data-directory")
	result.Port = viper.GetInt("port")
	result.NodePort = viper.GetInt("node-port")
	result.MinerPublicKeys = viper.GetStringSlice("miner-public-keys")

	return &result, nil
}
