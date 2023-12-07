package blockchain

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/secret-dev87/golang-ethereum-api/pkg/ethereum"
)

var EscrowInstance *Escrow

const (
	EscrowAddress = "0xB6E0D71771A3Adfa75760D49f0c1581e65071F1F"
)

type Escrow struct {
	Address  common.Address
	Instance *ethereum.Escrow
}

func (clientCon ClientConnection) newEscrow(address string) *ethereum.Escrow {
	EscrowInstance = new(Escrow)

	contractAddress := common.HexToAddress(address)

	EscrowInstance.Address = contractAddress

	instance, err := ethereum.NewEscrow(contractAddress, clientCon.Client)

	if err != nil {
		log.Fatalln("Cannot get Factory contract at address ", address, " due to: ", err)
	}

	return instance
}

func (clientCon ClientConnection) SetUSDCAddress(address string) error {
	_, err := getEscrow().SetUsdcTokenAddress(clientCon.trxOpts, common.HexToAddress(address))

	if err != nil {
		log.Printf("Cannot set new USDC token on Escrow due to: %v", err)

		return err
	}

	return nil
}

func (clientCon ClientConnection) GetUSDCAddress() (*common.Address, error) {
	address, err := getEscrow().UsdcToken(clientCon.callOpts)

	if err != nil {
		log.Printf("Cannot get USDC token on Escrow due to: %v", err)

		return new(common.Address), err
	}

	return &address, nil
}

func getEscrow() *ethereum.Escrow {
	return CurrentConnection.newEscrow(EscrowAddress)
}
