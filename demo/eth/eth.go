package eth

import (
	"context"
	"math/big"
	"storage/config"
	"storage/log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GenEthclient() (cli *ethclient.Client, err error) {
	return dial(config.NodeUrl)
}

func GenWsEthclient() (cli *ethclient.Client, err error) {
	return dial(config.NodeWsUrl)
}

func dial(uri string) (cli *ethclient.Client, err error) {
	cli, err = ethclient.Dial(config.NodeWsUrl)
	if err != nil {
		log.Logger.Error("Failed to dial connects a client:", err)
		return
	}
	return
}

func GenAuth(cli *ethclient.Client, privateKey string) (auth *bind.TransactOpts, err error) {
	// 查询链ID
	chainID, err := cli.ChainID(context.Background())
	if err != nil {
		log.Logger.Error(err)
		return
	}

	_privateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Logger.Error(err)
		return
	}

	fromAddress := crypto.PubkeyToAddress(_privateKey.PublicKey)
	nonce, err := cli.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Logger.Error(err)
		return
	}
	// gasPrice, err := cli.SuggestGasPrice(context.Background())
	// if err != nil {
	// 	log.Logger.Error(err)
	// 	return
	// }
	auth, err = bind.NewKeyedTransactorWithChainID(_privateKey, chainID)
	auth.Nonce = big.NewInt(int64(nonce))
	// auth.Value = big.NewInt(0)     // in wei
	// auth.GasLimit = uint64(300000) // in units
	// auth.GasPrice = gasPrice
	return
}
