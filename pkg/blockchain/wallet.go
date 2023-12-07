package blockchain

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/secret-dev87/golang-ethereum-api/pkg/ethereum"
)

var WalletInstance *Wallet

var (
	WalletAddress string
)

type Order struct {
	Seller   string `json:"seller"`
	Receiver string `json:"receiver"`
	Amount   int64  `json:"amount"`
	Rate     int64  `json:"rate"`
	Fee      int64  `json:"fee"`
}

type Wallet struct {
	Address  common.Address
	Instance *ethereum.CustodianWalletLogic
}

func (clientCon ClientConnection) newWallet(address string) *ethereum.CustodianWalletLogic {
	WalletInstance = new(Wallet)

	contractAddress := common.HexToAddress(address)

	WalletInstance.Address = contractAddress

	instance, err := ethereum.NewCustodianWalletLogic(contractAddress, clientCon.Client)

	if err != nil {
		log.Fatalln("Cannot get wallet address at ", address, " due to: ", err)
	}

	return instance
}

func (clientCon ClientConnection) OpenBuyOrder(order Order) (*types.Transaction, error) {
	trx, err := getWalletLogic().NewBuyOrder(
		clientCon.trxOpts,
		common.HexToAddress(order.Seller),
		common.HexToAddress(order.Receiver),
		big.NewInt(order.Amount),
		big.NewInt(order.Rate),
		big.NewInt(order.Fee),
	)

	if err != nil {
		log.Printf("Cannot open order : %v", err)

		return new(types.Transaction), err
	}

	clientCon.postTransact()

	return trx, nil
}

func (clentCon ClientConnection) GetUSDCBalance() (*big.Int, error) {
	balance, err := getWalletLogic().GetTotalBalance(clentCon.callOpts)

	if err != nil {
		log.Printf("Cannot open order: %v", err)

		return new(big.Int), err
	}

	return balance, nil
}

func getWalletLogic() *ethereum.CustodianWalletLogic {
	return CurrentConnection.newWallet(WalletAddress)
}
