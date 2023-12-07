package blockchain

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/secret-dev87/golang-ethereum-api/pkg/ethereum"
)

var FactoryInstance *Factory

const (
	FactoryAddress = "0xc2Ce51479F3545d28c184400895A129d8c33c9cD"
)

type Factory struct {
	Address  common.Address
	Instance *ethereum.Factory
}

func (clientCon ClientConnection) newFactory(address string) *ethereum.Factory {
	FactoryInstance = new(Factory)

	contractAddress := common.HexToAddress(address)

	FactoryInstance.Address = contractAddress

	instance, err := ethereum.NewFactory(contractAddress, clientCon.Client)
	if err != nil {
		log.Fatalln("Cannot get Factory contract at address ", address, " due to: ", err)
	}

	return instance
}

// GetLogicAddress 0xB69da002264F6985F65379d1867360aC85A55507
func (clientCon ClientConnection) GetLogicAddress() common.Address {

	address, err := getFactory().CustodianWalletLogic(clientCon.callOpts)

	if err != nil {
		log.Fatalln("Cannot make call to Factory at ", FactoryInstance.Address, "due to: ", err)
	}

	return address
}

// GetEscrowAddress 0xB6E0D71771A3Adfa75760D49f0c1581e65071F1F
func (clientCon ClientConnection) GetEscrowAddress() common.Address {

	address, err := getFactory().EscrowContractAddress(clientCon.callOpts)

	if err != nil {
		log.Fatalln("Cannot make call to Factory at ", FactoryInstance.Address, "due to: ", err)
	}

	clientCon.postTransact()

	return address
}

func (clientCon ClientConnection) NewWallet(uuid string) (*types.Transaction, error) {

	trx, err := getFactory().NewCustodian(clientCon.trxOpts, uuid)

	if err != nil {
		log.Printf("Cannot create new wallet %v due to error: %v", uuid, err)

		return new(types.Transaction), err
	}

	clientCon.postTransact()

	return trx, nil
}

func (clientCon ClientConnection) GetAccountByUUID(uuid string) (*common.Address, error) {

	address, err := getFactory().Accounts(clientCon.callOpts, uuid)

	if err != nil {
		log.Printf("Cannot get account: %v due to error %v", uuid, err)

		return new(common.Address), err
	}

	return &address, nil
}

func getFactory() *ethereum.Factory {
	return CurrentConnection.newFactory(FactoryAddress)
}
