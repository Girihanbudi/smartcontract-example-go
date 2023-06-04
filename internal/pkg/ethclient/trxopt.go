package ethclient

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func (c *Client) GenerateKey(privateKeyAddress string) (addr common.Address, privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey, err error) {
	privateKey, err = crypto.HexToECDSA(privateKeyAddress)
	if err != nil {
		err = fmt.Errorf("failed to parse address as private key: %s", err.Error())
		return
	}

	cryptoPublicKey := privateKey.Public()
	publicKey, ok := cryptoPublicKey.(*ecdsa.PublicKey)
	if !ok {
		err = fmt.Errorf("invalid provided key")
		return
	}

	addr = crypto.PubkeyToAddress(*publicKey)
	return
}

func (c *Client) GetNonce(ctx context.Context, address common.Address) (nonce uint64, err error) {
	//fetch the last use nonce of account
	nonce, err = c.Client.PendingNonceAt(ctx, address)
	if err != nil {
		err = fmt.Errorf("failed to get pending nonce, %s", err.Error())
		return
	}

	return
}

func (c *Client) GetNonceFromPrivateKey(ctx context.Context, privateKeyAddress string) (nonce uint64, err error) {
	addr, _, _, err := c.GenerateKey(privateKeyAddress)
	if err != nil {
		return
	}

	//fetch the last use nonce of account
	nonce, err = c.Client.PendingNonceAt(ctx, addr)
	if err != nil {
		err = fmt.Errorf("failed to get pending nonce, %s", err.Error())
		return
	}

	return
}

func (c *Client) GetTrxOpts(ctx context.Context, privateKeyAddress string) (trxOpts *bind.TransactOpts, err error) {
	addr, privateKey, _, err := c.GenerateKey(privateKeyAddress)
	if err != nil {
		return
	}

	nonce, err := c.GetNonce(ctx, addr)
	if err != nil {
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

	trxOpts.Nonce = big.NewInt(int64(nonce))

	return trxOpts, nil
}

func (c *Client) SetTrxOptsVal(trxOpts *bind.TransactOpts, fund int64, gasLimit uint64, gasPrice int64) {
	if trxOpts != nil {
		trxOpts.Value = big.NewInt(c.DefaultFund) // in wei
		trxOpts.GasLimit = c.DefaultGasLimit      // in units
		trxOpts.GasPrice = big.NewInt(c.DefaultGasPrice)
	}
}

func (c *Client) SetWithDefaultTrxOptsVal(trxOpts *bind.TransactOpts) {
	c.SetTrxOptsVal(trxOpts, c.Options.DefaultFund, c.Options.DefaultGasLimit, c.Options.DefaultGasPrice)
}
