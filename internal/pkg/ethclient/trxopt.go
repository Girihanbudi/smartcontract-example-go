package ethclient

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
)

func (c *Client) GetTrxOpts(ctx context.Context, address string) (trxOpts *bind.TransactOpts, nonce uint64, err error) {
	privateKey, err := crypto.HexToECDSA(address)
	if err != nil {
		err = fmt.Errorf("failed to parse address as private key: %s", err.Error())
		return
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		err = fmt.Errorf("invalid provided key")
		return
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	//fetch the last use nonce of account
	nonce, err = c.Client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		err = fmt.Errorf("failed to get pending nonce, %s", err.Error())
		return
	}

	chainID, err := c.Client.ChainID(ctx)
	if err != nil {
		err = fmt.Errorf("failed to get chain id, %s", err.Error())
		return
	}

	trxOpts, err = bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		err = fmt.Errorf("failed to get authorization data, %s", err.Error())
		return
	}

	return trxOpts, nonce, nil
}

func (c *Client) SetTrxOptsVal(trxOpts *bind.TransactOpts, nonce uint64, fund int64, gasLimit uint64, gasPrice int64) *bind.TransactOpts {
	if trxOpts != nil {
		trxOpts.Nonce = big.NewInt(int64(nonce))
		trxOpts.Value = big.NewInt(c.DefaultFund) // in wei
		trxOpts.GasLimit = c.DefaultGasLimit      // in units
		trxOpts.GasPrice = big.NewInt(c.DefaultGasPrice)
	}

	return trxOpts
}

func (c *Client) SetWithDefaultTrxOptsVal(trxOpts *bind.TransactOpts, nonce uint64) *bind.TransactOpts {
	return c.SetTrxOptsVal(trxOpts, nonce, c.Options.DefaultFund, c.Options.DefaultGasLimit, c.Options.DefaultGasPrice)
}
