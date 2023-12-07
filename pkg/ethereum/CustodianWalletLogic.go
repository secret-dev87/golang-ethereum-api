// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethereum

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// CustodianWalletLogicMetaData contains all meta data concerning the CustodianWalletLogic contract.
var CustodianWalletLogicMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"ApproveRejectedOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fulfiledTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"orderStatus\",\"type\":\"uint8\"}],\"name\":\"ClosedOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"orderStatus\",\"type\":\"uint8\"}],\"name\":\"OpenOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"OrderFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"RejectedOrder\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_openOrderIndex\",\"type\":\"uint256\"}],\"name\":\"approveOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"availBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_openOrderIndex\",\"type\":\"uint256\"}],\"name\":\"consentOrderRejected\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOpenOrders\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"newBuyOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"newSellOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"openOrders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"orders\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"orderStatus\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fulfiledTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"}],\"name\":\"rejectOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506120728061001d5f395ff3fe608060405234801561000f575f80fd5b50600436106100a7575f3560e01c806391c024151161006f57806391c0241514610175578063a85c38ef14610191578063bd4b8a04146101ca578063c45a0155146101e6578063dbe5bab514610204578063fdf61e6814610222576100a7565b806312b58349146100ab5780631971e03b146100c95780631e7cea80146100f957806324f63894146101295780636282682a14610145575b5f80fd5b6100b3610240565b6040516100c09190611274565b60405180910390f35b6100e360048036038101906100de9190611322565b61024e565b6040516100f09190611274565b60405180910390f35b610113600480360381019061010e9190611399565b6104ff565b6040516101209190611274565b60405180910390f35b610143600480360381019061013e91906113d7565b61052a565b005b61015f600480360381019061015a9190611322565b610691565b60405161016c9190611274565b60405180910390f35b61018f600480360381019061018a91906113d7565b61072a565b005b6101ab60048036038101906101a691906113d7565b6108bd565b6040516101c19a9998979695949392919061142c565b60405180910390f35b6101e460048036038101906101df91906113d7565b610991565b005b6101ee610bb4565b6040516101fb91906114c6565b60405180910390f35b61020c610bd9565b6040516102199190611596565b60405180910390f35b61022a610be8565b6040516102379190611274565b60405180910390f35b5f610249610bf6565b905090565b5f80610258610ce7565b73ffffffffffffffffffffffffffffffffffffffff166311eac8556040518163ffffffff1660e01b8152600401602060405180830381865afa1580156102a0573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102c491906115ca565b90505f73ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff1603610334576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161032b9061164f565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036103a2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610399906116b7565b60405180910390fd5b82856103ae9190611702565b8173ffffffffffffffffffffffffffffffffffffffff166370a08231896040518263ffffffff1660e01b81526004016103e791906114c6565b602060405180830381865afa158015610402573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104269190611749565b1015610467576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161045e906117be565b60405180910390fd5b61046f610ce7565b73ffffffffffffffffffffffffffffffffffffffff1663b1fab21c8830898989895f6040518863ffffffff1660e01b81526004016104b3979695949392919061181e565b6020604051808303815f875af11580156104cf573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104f39190611749565b91505095945050505050565b6001602052815f5260405f208181548110610518575f80fd5b905f5260205f20015f91509150505481565b5f610533610ce7565b73ffffffffffffffffffffffffffffffffffffffff166342b50a7a836040518263ffffffff1660e01b815260040161056b9190611274565b61014060405180830381865afa158015610587573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105ab9190611a34565b90503073ffffffffffffffffffffffffffffffffffffffff16815f015173ffffffffffffffffffffffffffffffffffffffff161461061e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161061590611aaa565b60405180910390fd5b610626610ce7565b73ffffffffffffffffffffffffffffffffffffffff1663240f416030846040518363ffffffff1660e01b8152600401610660929190611ac8565b5f604051808303815f87803b158015610677575f80fd5b505af1158015610689573d5f803e3d5ffd5b505050505050565b5f61069a610ce7565b73ffffffffffffffffffffffffffffffffffffffff1663b1fab21c30888888888860016040518863ffffffff1660e01b81526004016106df9796959493929190611b28565b6020604051808303815f875af11580156106fb573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061071f9190611749565b905095945050505050565b5f610733610d7b565b90505f81838151811061074957610748611b95565b5b602002602001015190505f61075c610ce7565b73ffffffffffffffffffffffffffffffffffffffff166342b50a7a836040518263ffffffff1660e01b81526004016107949190611274565b61014060405180830381865afa1580156107b0573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107d49190611a34565b90503073ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff1614610848576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161083f90611c0c565b60405180910390fd5b610850610ce7565b73ffffffffffffffffffffffffffffffffffffffff1663ef3bbbff30866040518363ffffffff1660e01b815260040161088a929190611ac8565b5f604051808303815f87803b1580156108a1575f80fd5b505af11580156108b3573d5f803e3d5ffd5b5050505050505050565b5f81815481106108cb575f80fd5b905f5260205f2090600902015f91509050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806002015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806003015490806004015490806005015490806006015f9054906101000a900460ff16908060060160019054906101000a900460ff1690806007015490806008015490508a565b5f61099a610d7b565b90505f8151116109df576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109d690611c74565b60405180910390fd5b5f8183815181106109f3576109f2611b95565b5b602002602001015190505f610a06610ce7565b73ffffffffffffffffffffffffffffffffffffffff166342b50a7a836040518263ffffffff1660e01b8152600401610a3e9190611274565b61014060405180830381865afa158015610a5a573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a7e9190611a34565b90503073ffffffffffffffffffffffffffffffffffffffff16815f015173ffffffffffffffffffffffffffffffffffffffff1614610af1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ae890611cdc565b60405180910390fd5b610af9610ce7565b73ffffffffffffffffffffffffffffffffffffffff1663b00f215c30866040518363ffffffff1660e01b8152600401610b33929190611ac8565b5f604051808303815f87803b158015610b4a575f80fd5b505af1158015610b5c573d5f803e3d5ffd5b50505050610b77816040015182606001518360a00151610e04565b7fd276fbc5e545783825cac5416bad154c47e39daebf44589255a65b4f56cf3e0e82604051610ba69190611274565b60405180910390a150505050565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6060610be3610d7b565b905090565b5f610bf1611156565b905090565b5f610bff610ce7565b73ffffffffffffffffffffffffffffffffffffffff166311eac8556040518163ffffffff1660e01b8152600401602060405180830381865afa158015610c47573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c6b91906115ca565b73ffffffffffffffffffffffffffffffffffffffff166370a08231336040518263ffffffff1660e01b8152600401610ca391906114c6565b602060405180830381865afa158015610cbe573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ce29190611749565b905090565b5f60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d226ae086040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d52573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d7691906115ca565b905090565b6060610d85610ce7565b73ffffffffffffffffffffffffffffffffffffffff16636dee2032306040518263ffffffff1660e01b8152600401610dbd91906114c6565b5f60405180830381865afa158015610dd7573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f82011682018060405250810190610dff9190611dc2565b905090565b3073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610e72576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e6990611e53565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610ee0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ed790611ebb565b60405180910390fd5b5f8211610f22576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f1990611f23565b60405180910390fd5b81610f2b611156565b1015610f6c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f6390611f8b565b60405180910390fd5b610f74610ce7565b73ffffffffffffffffffffffffffffffffffffffff166311eac8556040518163ffffffff1660e01b8152600401602060405180830381865afa158015610fbc573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610fe091906115ca565b73ffffffffffffffffffffffffffffffffffffffff1663a9059cbb84846040518363ffffffff1660e01b815260040161101a929190611ac8565b6020604051808303815f875af1158015611036573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061105a9190611fde565b50611063610ce7565b73ffffffffffffffffffffffffffffffffffffffff166311eac8556040518163ffffffff1660e01b8152600401602060405180830381865afa1580156110ab573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110cf91906115ca565b73ffffffffffffffffffffffffffffffffffffffff1663a9059cbb6110f2610ce7565b836040518363ffffffff1660e01b8152600401611110929190611ac8565b6020604051808303815f875af115801561112c573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111509190611fde565b50505050565b5f80611160610d7b565b90505f61116b610bf6565b90505f5b6001835161117d9190612009565b811015611253575f61118d610ce7565b73ffffffffffffffffffffffffffffffffffffffff166342b50a7a8584815181106111bb576111ba611b95565b5b60200260200101516040518263ffffffff1660e01b81526004016111df9190611274565b61014060405180830381865afa1580156111fb573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061121f9190611a34565b90508060600151836112319190612009565b92508060a00151836112439190612009565b925050808060010191505061116f565b50809250505090565b5f819050919050565b61126e8161125c565b82525050565b5f6020820190506112875f830184611265565b92915050565b5f604051905090565b5f80fd5b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6112c78261129e565b9050919050565b6112d7816112bd565b81146112e1575f80fd5b50565b5f813590506112f2816112ce565b92915050565b6113018161125c565b811461130b575f80fd5b50565b5f8135905061131c816112f8565b92915050565b5f805f805f60a0868803121561133b5761133a611296565b5b5f611348888289016112e4565b9550506020611359888289016112e4565b945050604061136a8882890161130e565b935050606061137b8882890161130e565b925050608061138c8882890161130e565b9150509295509295909350565b5f80604083850312156113af576113ae611296565b5b5f6113bc858286016112e4565b92505060206113cd8582860161130e565b9150509250929050565b5f602082840312156113ec576113eb611296565b5b5f6113f98482850161130e565b91505092915050565b61140b816112bd565b82525050565b5f60ff82169050919050565b61142681611411565b82525050565b5f610140820190506114405f83018d611402565b61144d602083018c611402565b61145a604083018b611402565b611467606083018a611265565b6114746080830189611265565b61148160a0830188611265565b61148e60c083018761141d565b61149b60e083018661141d565b6114a9610100830185611265565b6114b7610120830184611265565b9b9a5050505050505050505050565b5f6020820190506114d95f830184611402565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b6115118161125c565b82525050565b5f6115228383611508565b60208301905092915050565b5f602082019050919050565b5f611544826114df565b61154e81856114e9565b9350611559836114f9565b805f5b838110156115895781516115708882611517565b975061157b8361152e565b92505060018101905061155c565b5085935050505092915050565b5f6020820190508181035f8301526115ae818461153a565b905092915050565b5f815190506115c4816112ce565b92915050565b5f602082840312156115df576115de611296565b5b5f6115ec848285016115b6565b91505092915050565b5f82825260208201905092915050565b7f43574c3a2073656c6c6572206e6f7420736574000000000000000000000000005f82015250565b5f6116396013836115f5565b915061164482611605565b602082019050919050565b5f6020820190508181035f8301526116668161162d565b9050919050565b7f43574c3a207573646320746f6b656e206e6f74207365740000000000000000005f82015250565b5f6116a16017836115f5565b91506116ac8261166d565b602082019050919050565b5f6020820190508181035f8301526116ce81611695565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61170c8261125c565b91506117178361125c565b925082820190508082111561172f5761172e6116d5565b5b92915050565b5f81519050611743816112f8565b92915050565b5f6020828403121561175e5761175d611296565b5b5f61176b84828501611735565b91505092915050565b7f433a206e6f7420656e6f756768205553440000000000000000000000000000005f82015250565b5f6117a86011836115f5565b91506117b382611774565b602082019050919050565b5f6020820190508181035f8301526117d58161179c565b9050919050565b5f819050919050565b5f819050919050565b5f6118086118036117fe846117dc565b6117e5565b611411565b9050919050565b611818816117ee565b82525050565b5f60e0820190506118315f83018a611402565b61183e6020830189611402565b61184b6040830188611402565b6118586060830187611265565b6118656080830186611265565b61187260a0830185611265565b61187f60c083018461180f565b98975050505050505050565b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6118d58261188f565b810181811067ffffffffffffffff821117156118f4576118f361189f565b5b80604052505050565b5f61190661128d565b905061191282826118cc565b919050565b61192081611411565b811461192a575f80fd5b50565b5f8151905061193b81611917565b92915050565b5f61014082840312156119575761195661188b565b5b6119626101406118fd565b90505f611971848285016115b6565b5f830152506020611984848285016115b6565b6020830152506040611998848285016115b6565b60408301525060606119ac84828501611735565b60608301525060806119c084828501611735565b60808301525060a06119d484828501611735565b60a08301525060c06119e88482850161192d565b60c08301525060e06119fc8482850161192d565b60e083015250610100611a1184828501611735565b61010083015250610120611a2784828501611735565b6101208301525092915050565b5f6101408284031215611a4a57611a49611296565b5b5f611a5784828501611941565b91505092915050565b7f43574c3a206f6e6c792073656c6c6572000000000000000000000000000000005f82015250565b5f611a946010836115f5565b9150611a9f82611a60565b602082019050919050565b5f6020820190508181035f830152611ac181611a88565b9050919050565b5f604082019050611adb5f830185611402565b611ae86020830184611265565b9392505050565b5f819050919050565b5f611b12611b0d611b0884611aef565b6117e5565b611411565b9050919050565b611b2281611af8565b82525050565b5f60e082019050611b3b5f83018a611402565b611b486020830189611402565b611b556040830188611402565b611b626060830187611265565b611b6f6080830186611265565b611b7c60a0830185611265565b611b8960c0830184611b19565b98975050505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f43574c3a206f6e6c7920627579657200000000000000000000000000000000005f82015250565b5f611bf6600f836115f5565b9150611c0182611bc2565b602082019050919050565b5f6020820190508181035f830152611c2381611bea565b9050919050565b7f43574c3a206e6f206f70656e206f7264657273000000000000000000000000005f82015250565b5f611c5e6013836115f5565b9150611c6982611c2a565b602082019050919050565b5f6020820190508181035f830152611c8b81611c52565b9050919050565b7f43574c3a20696e76616c6964206f7264657200000000000000000000000000005f82015250565b5f611cc66012836115f5565b9150611cd182611c92565b602082019050919050565b5f6020820190508181035f830152611cf381611cba565b9050919050565b5f80fd5b5f67ffffffffffffffff821115611d1857611d1761189f565b5b602082029050602081019050919050565b5f80fd5b5f611d3f611d3a84611cfe565b6118fd565b90508083825260208201905060208402830185811115611d6257611d61611d29565b5b835b81811015611d8b5780611d778882611735565b845260208401935050602081019050611d64565b5050509392505050565b5f82601f830112611da957611da8611cfa565b5b8151611db9848260208601611d2d565b91505092915050565b5f60208284031215611dd757611dd6611296565b5b5f82015167ffffffffffffffff811115611df457611df361129a565b5b611e0084828501611d95565b91505092915050565b7f43574c3a2073656c6620666f7262696464656e000000000000000000000000005f82015250565b5f611e3d6013836115f5565b9150611e4882611e09565b602082019050919050565b5f6020820190508181035f830152611e6a81611e31565b9050919050565b7f43574c3a20696e76616c696420746f20616464726573730000000000000000005f82015250565b5f611ea56017836115f5565b9150611eb082611e71565b602082019050919050565b5f6020820190508181035f830152611ed281611e99565b9050919050565b7f43574c3a20616d6f756e742063616e6e6f7420657175616c20300000000000005f82015250565b5f611f0d601a836115f5565b9150611f1882611ed9565b602082019050919050565b5f6020820190508181035f830152611f3a81611f01565b9050919050565b7f43574c3a20696e73756666696369656e742066756e64730000000000000000005f82015250565b5f611f756017836115f5565b9150611f8082611f41565b602082019050919050565b5f6020820190508181035f830152611fa281611f69565b9050919050565b5f8115159050919050565b611fbd81611fa9565b8114611fc7575f80fd5b50565b5f81519050611fd881611fb4565b92915050565b5f60208284031215611ff357611ff2611296565b5b5f61200084828501611fca565b91505092915050565b5f6120138261125c565b915061201e8361125c565b9250828203905081811115612036576120356116d5565b5b9291505056fea264697066735822122071b7e2d2ebf38593d9cebc74a8b22c8e62ff94cf0fa3f05640a313d7a688fa3b64736f6c63430008170033",
}

// CustodianWalletLogicABI is the input ABI used to generate the binding from.
// Deprecated: Use CustodianWalletLogicMetaData.ABI instead.
var CustodianWalletLogicABI = CustodianWalletLogicMetaData.ABI

// CustodianWalletLogicBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CustodianWalletLogicMetaData.Bin instead.
var CustodianWalletLogicBin = CustodianWalletLogicMetaData.Bin

// DeployCustodianWalletLogic deploys a new Ethereum contract, binding an instance of CustodianWalletLogic to it.
func DeployCustodianWalletLogic(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CustodianWalletLogic, error) {
	parsed, err := CustodianWalletLogicMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CustodianWalletLogicBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CustodianWalletLogic{CustodianWalletLogicCaller: CustodianWalletLogicCaller{contract: contract}, CustodianWalletLogicTransactor: CustodianWalletLogicTransactor{contract: contract}, CustodianWalletLogicFilterer: CustodianWalletLogicFilterer{contract: contract}}, nil
}

// CustodianWalletLogic is an auto generated Go binding around an Ethereum contract.
type CustodianWalletLogic struct {
	CustodianWalletLogicCaller     // Read-only binding to the contract
	CustodianWalletLogicTransactor // Write-only binding to the contract
	CustodianWalletLogicFilterer   // Log filterer for contract events
}

// CustodianWalletLogicCaller is an auto generated read-only Go binding around an Ethereum contract.
type CustodianWalletLogicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CustodianWalletLogicTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CustodianWalletLogicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CustodianWalletLogicFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CustodianWalletLogicFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CustodianWalletLogicSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CustodianWalletLogicSession struct {
	Contract     *CustodianWalletLogic // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CustodianWalletLogicCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CustodianWalletLogicCallerSession struct {
	Contract *CustodianWalletLogicCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// CustodianWalletLogicTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CustodianWalletLogicTransactorSession struct {
	Contract     *CustodianWalletLogicTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// CustodianWalletLogicRaw is an auto generated low-level Go binding around an Ethereum contract.
type CustodianWalletLogicRaw struct {
	Contract *CustodianWalletLogic // Generic contract binding to access the raw methods on
}

// CustodianWalletLogicCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CustodianWalletLogicCallerRaw struct {
	Contract *CustodianWalletLogicCaller // Generic read-only contract binding to access the raw methods on
}

// CustodianWalletLogicTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CustodianWalletLogicTransactorRaw struct {
	Contract *CustodianWalletLogicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCustodianWalletLogic creates a new instance of CustodianWalletLogic, bound to a specific deployed contract.
func NewCustodianWalletLogic(address common.Address, backend bind.ContractBackend) (*CustodianWalletLogic, error) {
	contract, err := bindCustodianWalletLogic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CustodianWalletLogic{CustodianWalletLogicCaller: CustodianWalletLogicCaller{contract: contract}, CustodianWalletLogicTransactor: CustodianWalletLogicTransactor{contract: contract}, CustodianWalletLogicFilterer: CustodianWalletLogicFilterer{contract: contract}}, nil
}

// NewCustodianWalletLogicCaller creates a new read-only instance of CustodianWalletLogic, bound to a specific deployed contract.
func NewCustodianWalletLogicCaller(address common.Address, caller bind.ContractCaller) (*CustodianWalletLogicCaller, error) {
	contract, err := bindCustodianWalletLogic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CustodianWalletLogicCaller{contract: contract}, nil
}

// NewCustodianWalletLogicTransactor creates a new write-only instance of CustodianWalletLogic, bound to a specific deployed contract.
func NewCustodianWalletLogicTransactor(address common.Address, transactor bind.ContractTransactor) (*CustodianWalletLogicTransactor, error) {
	contract, err := bindCustodianWalletLogic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CustodianWalletLogicTransactor{contract: contract}, nil
}

// NewCustodianWalletLogicFilterer creates a new log filterer instance of CustodianWalletLogic, bound to a specific deployed contract.
func NewCustodianWalletLogicFilterer(address common.Address, filterer bind.ContractFilterer) (*CustodianWalletLogicFilterer, error) {
	contract, err := bindCustodianWalletLogic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CustodianWalletLogicFilterer{contract: contract}, nil
}

// bindCustodianWalletLogic binds a generic wrapper to an already deployed contract.
func bindCustodianWalletLogic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CustodianWalletLogicMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CustodianWalletLogic *CustodianWalletLogicRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CustodianWalletLogic.Contract.CustodianWalletLogicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CustodianWalletLogic *CustodianWalletLogicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CustodianWalletLogic.Contract.CustodianWalletLogicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CustodianWalletLogic *CustodianWalletLogicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CustodianWalletLogic.Contract.CustodianWalletLogicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CustodianWalletLogic *CustodianWalletLogicCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CustodianWalletLogic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CustodianWalletLogic *CustodianWalletLogicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CustodianWalletLogic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CustodianWalletLogic *CustodianWalletLogicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CustodianWalletLogic.Contract.contract.Transact(opts, method, params...)
}

// AvailBalance is a free data retrieval call binding the contract method 0xfdf61e68.
//
// Solidity: function availBalance() view returns(uint256)
func (_CustodianWalletLogic *CustodianWalletLogicCaller) AvailBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CustodianWalletLogic.contract.Call(opts, &out, "availBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AvailBalance is a free data retrieval call binding the contract method 0xfdf61e68.
//
// Solidity: function availBalance() view returns(uint256)
func (_CustodianWalletLogic *CustodianWalletLogicSession) AvailBalance() (*big.Int, error) {
	return _CustodianWalletLogic.Contract.AvailBalance(&_CustodianWalletLogic.CallOpts)
}

// AvailBalance is a free data retrieval call binding the contract method 0xfdf61e68.
//
// Solidity: function availBalance() view returns(uint256)
func (_CustodianWalletLogic *CustodianWalletLogicCallerSession) AvailBalance() (*big.Int, error) {
	return _CustodianWalletLogic.Contract.AvailBalance(&_CustodianWalletLogic.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_CustodianWalletLogic *CustodianWalletLogicCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CustodianWalletLogic.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_CustodianWalletLogic *CustodianWalletLogicSession) Factory() (common.Address, error) {
	return _CustodianWalletLogic.Contract.Factory(&_CustodianWalletLogic.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_CustodianWalletLogic *CustodianWalletLogicCallerSession) Factory() (common.Address, error) {
	return _CustodianWalletLogic.Contract.Factory(&_CustodianWalletLogic.CallOpts)
}

// GetOpenOrders is a free data retrieval call binding the contract method 0xdbe5bab5.
//
// Solidity: function getOpenOrders() view returns(uint256[])
func (_CustodianWalletLogic *CustodianWalletLogicCaller) GetOpenOrders(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _CustodianWalletLogic.contract.Call(opts, &out, "getOpenOrders")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetOpenOrders is a free data retrieval call binding the contract method 0xdbe5bab5.
//
// Solidity: function getOpenOrders() view returns(uint256[])
func (_CustodianWalletLogic *CustodianWalletLogicSession) GetOpenOrders() ([]*big.Int, error) {
	return _CustodianWalletLogic.Contract.GetOpenOrders(&_CustodianWalletLogic.CallOpts)
}

// GetOpenOrders is a free data retrieval call binding the contract method 0xdbe5bab5.
//
// Solidity: function getOpenOrders() view returns(uint256[])
func (_CustodianWalletLogic *CustodianWalletLogicCallerSession) GetOpenOrders() ([]*big.Int, error) {
	return _CustodianWalletLogic.Contract.GetOpenOrders(&_CustodianWalletLogic.CallOpts)
}

// GetTotalBalance is a free data retrieval call binding the contract method 0x12b58349.
//
// Solidity: function getTotalBalance() view returns(uint256)
func (_CustodianWalletLogic *CustodianWalletLogicCaller) GetTotalBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CustodianWalletLogic.contract.Call(opts, &out, "getTotalBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalBalance is a free data retrieval call binding the contract method 0x12b58349.
//
// Solidity: function getTotalBalance() view returns(uint256)
func (_CustodianWalletLogic *CustodianWalletLogicSession) GetTotalBalance() (*big.Int, error) {
	return _CustodianWalletLogic.Contract.GetTotalBalance(&_CustodianWalletLogic.CallOpts)
}

// GetTotalBalance is a free data retrieval call binding the contract method 0x12b58349.
//
// Solidity: function getTotalBalance() view returns(uint256)
func (_CustodianWalletLogic *CustodianWalletLogicCallerSession) GetTotalBalance() (*big.Int, error) {
	return _CustodianWalletLogic.Contract.GetTotalBalance(&_CustodianWalletLogic.CallOpts)
}

// OpenOrders is a free data retrieval call binding the contract method 0x1e7cea80.
//
// Solidity: function openOrders(address , uint256 ) view returns(uint256)
func (_CustodianWalletLogic *CustodianWalletLogicCaller) OpenOrders(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CustodianWalletLogic.contract.Call(opts, &out, "openOrders", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OpenOrders is a free data retrieval call binding the contract method 0x1e7cea80.
//
// Solidity: function openOrders(address , uint256 ) view returns(uint256)
func (_CustodianWalletLogic *CustodianWalletLogicSession) OpenOrders(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _CustodianWalletLogic.Contract.OpenOrders(&_CustodianWalletLogic.CallOpts, arg0, arg1)
}

// OpenOrders is a free data retrieval call binding the contract method 0x1e7cea80.
//
// Solidity: function openOrders(address , uint256 ) view returns(uint256)
func (_CustodianWalletLogic *CustodianWalletLogicCallerSession) OpenOrders(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _CustodianWalletLogic.Contract.OpenOrders(&_CustodianWalletLogic.CallOpts, arg0, arg1)
}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(address seller, address buyer, address receiver, uint256 amount, uint256 rate, uint256 fee, uint8 orderType, uint8 orderStatus, uint256 startTime, uint256 fulfiledTime)
func (_CustodianWalletLogic *CustodianWalletLogicCaller) Orders(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Seller       common.Address
	Buyer        common.Address
	Receiver     common.Address
	Amount       *big.Int
	Rate         *big.Int
	Fee          *big.Int
	OrderType    uint8
	OrderStatus  uint8
	StartTime    *big.Int
	FulfiledTime *big.Int
}, error) {
	var out []interface{}
	err := _CustodianWalletLogic.contract.Call(opts, &out, "orders", arg0)

	outstruct := new(struct {
		Seller       common.Address
		Buyer        common.Address
		Receiver     common.Address
		Amount       *big.Int
		Rate         *big.Int
		Fee          *big.Int
		OrderType    uint8
		OrderStatus  uint8
		StartTime    *big.Int
		FulfiledTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Seller = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Buyer = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Receiver = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Rate = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Fee = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.OrderType = *abi.ConvertType(out[6], new(uint8)).(*uint8)
	outstruct.OrderStatus = *abi.ConvertType(out[7], new(uint8)).(*uint8)
	outstruct.StartTime = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.FulfiledTime = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(address seller, address buyer, address receiver, uint256 amount, uint256 rate, uint256 fee, uint8 orderType, uint8 orderStatus, uint256 startTime, uint256 fulfiledTime)
func (_CustodianWalletLogic *CustodianWalletLogicSession) Orders(arg0 *big.Int) (struct {
	Seller       common.Address
	Buyer        common.Address
	Receiver     common.Address
	Amount       *big.Int
	Rate         *big.Int
	Fee          *big.Int
	OrderType    uint8
	OrderStatus  uint8
	StartTime    *big.Int
	FulfiledTime *big.Int
}, error) {
	return _CustodianWalletLogic.Contract.Orders(&_CustodianWalletLogic.CallOpts, arg0)
}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(address seller, address buyer, address receiver, uint256 amount, uint256 rate, uint256 fee, uint8 orderType, uint8 orderStatus, uint256 startTime, uint256 fulfiledTime)
func (_CustodianWalletLogic *CustodianWalletLogicCallerSession) Orders(arg0 *big.Int) (struct {
	Seller       common.Address
	Buyer        common.Address
	Receiver     common.Address
	Amount       *big.Int
	Rate         *big.Int
	Fee          *big.Int
	OrderType    uint8
	OrderStatus  uint8
	StartTime    *big.Int
	FulfiledTime *big.Int
}, error) {
	return _CustodianWalletLogic.Contract.Orders(&_CustodianWalletLogic.CallOpts, arg0)
}

// ApproveOrder is a paid mutator transaction binding the contract method 0xbd4b8a04.
//
// Solidity: function approveOrder(uint256 _openOrderIndex) returns()
func (_CustodianWalletLogic *CustodianWalletLogicTransactor) ApproveOrder(opts *bind.TransactOpts, _openOrderIndex *big.Int) (*types.Transaction, error) {
	return _CustodianWalletLogic.contract.Transact(opts, "approveOrder", _openOrderIndex)
}

// ApproveOrder is a paid mutator transaction binding the contract method 0xbd4b8a04.
//
// Solidity: function approveOrder(uint256 _openOrderIndex) returns()
func (_CustodianWalletLogic *CustodianWalletLogicSession) ApproveOrder(_openOrderIndex *big.Int) (*types.Transaction, error) {
	return _CustodianWalletLogic.Contract.ApproveOrder(&_CustodianWalletLogic.TransactOpts, _openOrderIndex)
}

// ApproveOrder is a paid mutator transaction binding the contract method 0xbd4b8a04.
//
// Solidity: function approveOrder(uint256 _openOrderIndex) returns()
func (_CustodianWalletLogic *CustodianWalletLogicTransactorSession) ApproveOrder(_openOrderIndex *big.Int) (*types.Transaction, error) {
	return _CustodianWalletLogic.Contract.ApproveOrder(&_CustodianWalletLogic.TransactOpts, _openOrderIndex)
}

// ConsentOrderRejected is a paid mutator transaction binding the contract method 0x91c02415.
//
// Solidity: function consentOrderRejected(uint256 _openOrderIndex) returns()
func (_CustodianWalletLogic *CustodianWalletLogicTransactor) ConsentOrderRejected(opts *bind.TransactOpts, _openOrderIndex *big.Int) (*types.Transaction, error) {
	return _CustodianWalletLogic.contract.Transact(opts, "consentOrderRejected", _openOrderIndex)
}

// ConsentOrderRejected is a paid mutator transaction binding the contract method 0x91c02415.
//
// Solidity: function consentOrderRejected(uint256 _openOrderIndex) returns()
func (_CustodianWalletLogic *CustodianWalletLogicSession) ConsentOrderRejected(_openOrderIndex *big.Int) (*types.Transaction, error) {
	return _CustodianWalletLogic.Contract.ConsentOrderRejected(&_CustodianWalletLogic.TransactOpts, _openOrderIndex)
}

// ConsentOrderRejected is a paid mutator transaction binding the contract method 0x91c02415.
//
// Solidity: function consentOrderRejected(uint256 _openOrderIndex) returns()
func (_CustodianWalletLogic *CustodianWalletLogicTransactorSession) ConsentOrderRejected(_openOrderIndex *big.Int) (*types.Transaction, error) {
	return _CustodianWalletLogic.Contract.ConsentOrderRejected(&_CustodianWalletLogic.TransactOpts, _openOrderIndex)
}

// NewBuyOrder is a paid mutator transaction binding the contract method 0x1971e03b.
//
// Solidity: function newBuyOrder(address _seller, address _receiver, uint256 _amount, uint256 _rate, uint256 _fee) returns(uint256)
func (_CustodianWalletLogic *CustodianWalletLogicTransactor) NewBuyOrder(opts *bind.TransactOpts, _seller common.Address, _receiver common.Address, _amount *big.Int, _rate *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _CustodianWalletLogic.contract.Transact(opts, "newBuyOrder", _seller, _receiver, _amount, _rate, _fee)
}

// NewBuyOrder is a paid mutator transaction binding the contract method 0x1971e03b.
//
// Solidity: function newBuyOrder(address _seller, address _receiver, uint256 _amount, uint256 _rate, uint256 _fee) returns(uint256)
func (_CustodianWalletLogic *CustodianWalletLogicSession) NewBuyOrder(_seller common.Address, _receiver common.Address, _amount *big.Int, _rate *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _CustodianWalletLogic.Contract.NewBuyOrder(&_CustodianWalletLogic.TransactOpts, _seller, _receiver, _amount, _rate, _fee)
}

// NewBuyOrder is a paid mutator transaction binding the contract method 0x1971e03b.
//
// Solidity: function newBuyOrder(address _seller, address _receiver, uint256 _amount, uint256 _rate, uint256 _fee) returns(uint256)
func (_CustodianWalletLogic *CustodianWalletLogicTransactorSession) NewBuyOrder(_seller common.Address, _receiver common.Address, _amount *big.Int, _rate *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _CustodianWalletLogic.Contract.NewBuyOrder(&_CustodianWalletLogic.TransactOpts, _seller, _receiver, _amount, _rate, _fee)
}

// NewSellOrder is a paid mutator transaction binding the contract method 0x6282682a.
//
// Solidity: function newSellOrder(address _buyer, address _receiver, uint256 _amount, uint256 _rate, uint256 _fee) returns(uint256)
func (_CustodianWalletLogic *CustodianWalletLogicTransactor) NewSellOrder(opts *bind.TransactOpts, _buyer common.Address, _receiver common.Address, _amount *big.Int, _rate *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _CustodianWalletLogic.contract.Transact(opts, "newSellOrder", _buyer, _receiver, _amount, _rate, _fee)
}

// NewSellOrder is a paid mutator transaction binding the contract method 0x6282682a.
//
// Solidity: function newSellOrder(address _buyer, address _receiver, uint256 _amount, uint256 _rate, uint256 _fee) returns(uint256)
func (_CustodianWalletLogic *CustodianWalletLogicSession) NewSellOrder(_buyer common.Address, _receiver common.Address, _amount *big.Int, _rate *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _CustodianWalletLogic.Contract.NewSellOrder(&_CustodianWalletLogic.TransactOpts, _buyer, _receiver, _amount, _rate, _fee)
}

// NewSellOrder is a paid mutator transaction binding the contract method 0x6282682a.
//
// Solidity: function newSellOrder(address _buyer, address _receiver, uint256 _amount, uint256 _rate, uint256 _fee) returns(uint256)
func (_CustodianWalletLogic *CustodianWalletLogicTransactorSession) NewSellOrder(_buyer common.Address, _receiver common.Address, _amount *big.Int, _rate *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _CustodianWalletLogic.Contract.NewSellOrder(&_CustodianWalletLogic.TransactOpts, _buyer, _receiver, _amount, _rate, _fee)
}

// RejectOrder is a paid mutator transaction binding the contract method 0x24f63894.
//
// Solidity: function rejectOrder(uint256 _orderId) returns()
func (_CustodianWalletLogic *CustodianWalletLogicTransactor) RejectOrder(opts *bind.TransactOpts, _orderId *big.Int) (*types.Transaction, error) {
	return _CustodianWalletLogic.contract.Transact(opts, "rejectOrder", _orderId)
}

// RejectOrder is a paid mutator transaction binding the contract method 0x24f63894.
//
// Solidity: function rejectOrder(uint256 _orderId) returns()
func (_CustodianWalletLogic *CustodianWalletLogicSession) RejectOrder(_orderId *big.Int) (*types.Transaction, error) {
	return _CustodianWalletLogic.Contract.RejectOrder(&_CustodianWalletLogic.TransactOpts, _orderId)
}

// RejectOrder is a paid mutator transaction binding the contract method 0x24f63894.
//
// Solidity: function rejectOrder(uint256 _orderId) returns()
func (_CustodianWalletLogic *CustodianWalletLogicTransactorSession) RejectOrder(_orderId *big.Int) (*types.Transaction, error) {
	return _CustodianWalletLogic.Contract.RejectOrder(&_CustodianWalletLogic.TransactOpts, _orderId)
}

// CustodianWalletLogicApproveRejectedOrderIterator is returned from FilterApproveRejectedOrder and is used to iterate over the raw logs and unpacked data for ApproveRejectedOrder events raised by the CustodianWalletLogic contract.
type CustodianWalletLogicApproveRejectedOrderIterator struct {
	Event *CustodianWalletLogicApproveRejectedOrder // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CustodianWalletLogicApproveRejectedOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustodianWalletLogicApproveRejectedOrder)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CustodianWalletLogicApproveRejectedOrder)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CustodianWalletLogicApproveRejectedOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CustodianWalletLogicApproveRejectedOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CustodianWalletLogicApproveRejectedOrder represents a ApproveRejectedOrder event raised by the CustodianWalletLogic contract.
type CustodianWalletLogicApproveRejectedOrder struct {
	OrderId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproveRejectedOrder is a free log retrieval operation binding the contract event 0x6526538f9bd98d105bdfc50096f40172dea60988b982f7776fa03d816600eac8.
//
// Solidity: event ApproveRejectedOrder(uint256 orderId)
func (_CustodianWalletLogic *CustodianWalletLogicFilterer) FilterApproveRejectedOrder(opts *bind.FilterOpts) (*CustodianWalletLogicApproveRejectedOrderIterator, error) {

	logs, sub, err := _CustodianWalletLogic.contract.FilterLogs(opts, "ApproveRejectedOrder")
	if err != nil {
		return nil, err
	}
	return &CustodianWalletLogicApproveRejectedOrderIterator{contract: _CustodianWalletLogic.contract, event: "ApproveRejectedOrder", logs: logs, sub: sub}, nil
}

// WatchApproveRejectedOrder is a free log subscription operation binding the contract event 0x6526538f9bd98d105bdfc50096f40172dea60988b982f7776fa03d816600eac8.
//
// Solidity: event ApproveRejectedOrder(uint256 orderId)
func (_CustodianWalletLogic *CustodianWalletLogicFilterer) WatchApproveRejectedOrder(opts *bind.WatchOpts, sink chan<- *CustodianWalletLogicApproveRejectedOrder) (event.Subscription, error) {

	logs, sub, err := _CustodianWalletLogic.contract.WatchLogs(opts, "ApproveRejectedOrder")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CustodianWalletLogicApproveRejectedOrder)
				if err := _CustodianWalletLogic.contract.UnpackLog(event, "ApproveRejectedOrder", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproveRejectedOrder is a log parse operation binding the contract event 0x6526538f9bd98d105bdfc50096f40172dea60988b982f7776fa03d816600eac8.
//
// Solidity: event ApproveRejectedOrder(uint256 orderId)
func (_CustodianWalletLogic *CustodianWalletLogicFilterer) ParseApproveRejectedOrder(log types.Log) (*CustodianWalletLogicApproveRejectedOrder, error) {
	event := new(CustodianWalletLogicApproveRejectedOrder)
	if err := _CustodianWalletLogic.contract.UnpackLog(event, "ApproveRejectedOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CustodianWalletLogicClosedOrderIterator is returned from FilterClosedOrder and is used to iterate over the raw logs and unpacked data for ClosedOrder events raised by the CustodianWalletLogic contract.
type CustodianWalletLogicClosedOrderIterator struct {
	Event *CustodianWalletLogicClosedOrder // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CustodianWalletLogicClosedOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustodianWalletLogicClosedOrder)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CustodianWalletLogicClosedOrder)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CustodianWalletLogicClosedOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CustodianWalletLogicClosedOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CustodianWalletLogicClosedOrder represents a ClosedOrder event raised by the CustodianWalletLogic contract.
type CustodianWalletLogicClosedOrder struct {
	OrderId      *big.Int
	Seller       common.Address
	Buyer        common.Address
	Receiver     common.Address
	Amount       *big.Int
	Rate         *big.Int
	Fee          *big.Int
	OrderType    uint8
	FulfiledTime *big.Int
	OrderStatus  uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterClosedOrder is a free log retrieval operation binding the contract event 0x7296e5ad5c0ce57fe4cbceb3362886c7162b073357dd6b39025453458f6f8482.
//
// Solidity: event ClosedOrder(uint256 orderId, address indexed seller, address indexed buyer, address indexed receiver, uint256 amount, uint256 rate, uint256 fee, uint8 orderType, uint256 fulfiledTime, uint8 orderStatus)
func (_CustodianWalletLogic *CustodianWalletLogicFilterer) FilterClosedOrder(opts *bind.FilterOpts, seller []common.Address, buyer []common.Address, receiver []common.Address) (*CustodianWalletLogicClosedOrderIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _CustodianWalletLogic.contract.FilterLogs(opts, "ClosedOrder", sellerRule, buyerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &CustodianWalletLogicClosedOrderIterator{contract: _CustodianWalletLogic.contract, event: "ClosedOrder", logs: logs, sub: sub}, nil
}

// WatchClosedOrder is a free log subscription operation binding the contract event 0x7296e5ad5c0ce57fe4cbceb3362886c7162b073357dd6b39025453458f6f8482.
//
// Solidity: event ClosedOrder(uint256 orderId, address indexed seller, address indexed buyer, address indexed receiver, uint256 amount, uint256 rate, uint256 fee, uint8 orderType, uint256 fulfiledTime, uint8 orderStatus)
func (_CustodianWalletLogic *CustodianWalletLogicFilterer) WatchClosedOrder(opts *bind.WatchOpts, sink chan<- *CustodianWalletLogicClosedOrder, seller []common.Address, buyer []common.Address, receiver []common.Address) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _CustodianWalletLogic.contract.WatchLogs(opts, "ClosedOrder", sellerRule, buyerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CustodianWalletLogicClosedOrder)
				if err := _CustodianWalletLogic.contract.UnpackLog(event, "ClosedOrder", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClosedOrder is a log parse operation binding the contract event 0x7296e5ad5c0ce57fe4cbceb3362886c7162b073357dd6b39025453458f6f8482.
//
// Solidity: event ClosedOrder(uint256 orderId, address indexed seller, address indexed buyer, address indexed receiver, uint256 amount, uint256 rate, uint256 fee, uint8 orderType, uint256 fulfiledTime, uint8 orderStatus)
func (_CustodianWalletLogic *CustodianWalletLogicFilterer) ParseClosedOrder(log types.Log) (*CustodianWalletLogicClosedOrder, error) {
	event := new(CustodianWalletLogicClosedOrder)
	if err := _CustodianWalletLogic.contract.UnpackLog(event, "ClosedOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CustodianWalletLogicOpenOrderIterator is returned from FilterOpenOrder and is used to iterate over the raw logs and unpacked data for OpenOrder events raised by the CustodianWalletLogic contract.
type CustodianWalletLogicOpenOrderIterator struct {
	Event *CustodianWalletLogicOpenOrder // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CustodianWalletLogicOpenOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustodianWalletLogicOpenOrder)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CustodianWalletLogicOpenOrder)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CustodianWalletLogicOpenOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CustodianWalletLogicOpenOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CustodianWalletLogicOpenOrder represents a OpenOrder event raised by the CustodianWalletLogic contract.
type CustodianWalletLogicOpenOrder struct {
	OrderId     *big.Int
	Seller      common.Address
	Buyer       common.Address
	Receiver    common.Address
	Amount      *big.Int
	Rate        *big.Int
	Fee         *big.Int
	OrderType   uint8
	OrderStatus uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOpenOrder is a free log retrieval operation binding the contract event 0x0023c8dda5a5d69bbdf1f810b48e0b0fff6827cca652b18c81b25d8759bf67f2.
//
// Solidity: event OpenOrder(uint256 orderId, address indexed seller, address indexed buyer, address indexed receiver, uint256 amount, uint256 rate, uint256 fee, uint8 orderType, uint8 orderStatus)
func (_CustodianWalletLogic *CustodianWalletLogicFilterer) FilterOpenOrder(opts *bind.FilterOpts, seller []common.Address, buyer []common.Address, receiver []common.Address) (*CustodianWalletLogicOpenOrderIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _CustodianWalletLogic.contract.FilterLogs(opts, "OpenOrder", sellerRule, buyerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &CustodianWalletLogicOpenOrderIterator{contract: _CustodianWalletLogic.contract, event: "OpenOrder", logs: logs, sub: sub}, nil
}

// WatchOpenOrder is a free log subscription operation binding the contract event 0x0023c8dda5a5d69bbdf1f810b48e0b0fff6827cca652b18c81b25d8759bf67f2.
//
// Solidity: event OpenOrder(uint256 orderId, address indexed seller, address indexed buyer, address indexed receiver, uint256 amount, uint256 rate, uint256 fee, uint8 orderType, uint8 orderStatus)
func (_CustodianWalletLogic *CustodianWalletLogicFilterer) WatchOpenOrder(opts *bind.WatchOpts, sink chan<- *CustodianWalletLogicOpenOrder, seller []common.Address, buyer []common.Address, receiver []common.Address) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _CustodianWalletLogic.contract.WatchLogs(opts, "OpenOrder", sellerRule, buyerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CustodianWalletLogicOpenOrder)
				if err := _CustodianWalletLogic.contract.UnpackLog(event, "OpenOrder", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOpenOrder is a log parse operation binding the contract event 0x0023c8dda5a5d69bbdf1f810b48e0b0fff6827cca652b18c81b25d8759bf67f2.
//
// Solidity: event OpenOrder(uint256 orderId, address indexed seller, address indexed buyer, address indexed receiver, uint256 amount, uint256 rate, uint256 fee, uint8 orderType, uint8 orderStatus)
func (_CustodianWalletLogic *CustodianWalletLogicFilterer) ParseOpenOrder(log types.Log) (*CustodianWalletLogicOpenOrder, error) {
	event := new(CustodianWalletLogicOpenOrder)
	if err := _CustodianWalletLogic.contract.UnpackLog(event, "OpenOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CustodianWalletLogicOrderFulfilledIterator is returned from FilterOrderFulfilled and is used to iterate over the raw logs and unpacked data for OrderFulfilled events raised by the CustodianWalletLogic contract.
type CustodianWalletLogicOrderFulfilledIterator struct {
	Event *CustodianWalletLogicOrderFulfilled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CustodianWalletLogicOrderFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustodianWalletLogicOrderFulfilled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CustodianWalletLogicOrderFulfilled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CustodianWalletLogicOrderFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CustodianWalletLogicOrderFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CustodianWalletLogicOrderFulfilled represents a OrderFulfilled event raised by the CustodianWalletLogic contract.
type CustodianWalletLogicOrderFulfilled struct {
	OrderId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOrderFulfilled is a free log retrieval operation binding the contract event 0xd276fbc5e545783825cac5416bad154c47e39daebf44589255a65b4f56cf3e0e.
//
// Solidity: event OrderFulfilled(uint256 orderId)
func (_CustodianWalletLogic *CustodianWalletLogicFilterer) FilterOrderFulfilled(opts *bind.FilterOpts) (*CustodianWalletLogicOrderFulfilledIterator, error) {

	logs, sub, err := _CustodianWalletLogic.contract.FilterLogs(opts, "OrderFulfilled")
	if err != nil {
		return nil, err
	}
	return &CustodianWalletLogicOrderFulfilledIterator{contract: _CustodianWalletLogic.contract, event: "OrderFulfilled", logs: logs, sub: sub}, nil
}

// WatchOrderFulfilled is a free log subscription operation binding the contract event 0xd276fbc5e545783825cac5416bad154c47e39daebf44589255a65b4f56cf3e0e.
//
// Solidity: event OrderFulfilled(uint256 orderId)
func (_CustodianWalletLogic *CustodianWalletLogicFilterer) WatchOrderFulfilled(opts *bind.WatchOpts, sink chan<- *CustodianWalletLogicOrderFulfilled) (event.Subscription, error) {

	logs, sub, err := _CustodianWalletLogic.contract.WatchLogs(opts, "OrderFulfilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CustodianWalletLogicOrderFulfilled)
				if err := _CustodianWalletLogic.contract.UnpackLog(event, "OrderFulfilled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOrderFulfilled is a log parse operation binding the contract event 0xd276fbc5e545783825cac5416bad154c47e39daebf44589255a65b4f56cf3e0e.
//
// Solidity: event OrderFulfilled(uint256 orderId)
func (_CustodianWalletLogic *CustodianWalletLogicFilterer) ParseOrderFulfilled(log types.Log) (*CustodianWalletLogicOrderFulfilled, error) {
	event := new(CustodianWalletLogicOrderFulfilled)
	if err := _CustodianWalletLogic.contract.UnpackLog(event, "OrderFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CustodianWalletLogicRejectedOrderIterator is returned from FilterRejectedOrder and is used to iterate over the raw logs and unpacked data for RejectedOrder events raised by the CustodianWalletLogic contract.
type CustodianWalletLogicRejectedOrderIterator struct {
	Event *CustodianWalletLogicRejectedOrder // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CustodianWalletLogicRejectedOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustodianWalletLogicRejectedOrder)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CustodianWalletLogicRejectedOrder)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CustodianWalletLogicRejectedOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CustodianWalletLogicRejectedOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CustodianWalletLogicRejectedOrder represents a RejectedOrder event raised by the CustodianWalletLogic contract.
type CustodianWalletLogicRejectedOrder struct {
	OrderId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRejectedOrder is a free log retrieval operation binding the contract event 0xbd551ca6c38e2ad16afa16c81fcb4b8e891315559bfb80f551ab6b66cb583dbd.
//
// Solidity: event RejectedOrder(uint256 orderId)
func (_CustodianWalletLogic *CustodianWalletLogicFilterer) FilterRejectedOrder(opts *bind.FilterOpts) (*CustodianWalletLogicRejectedOrderIterator, error) {

	logs, sub, err := _CustodianWalletLogic.contract.FilterLogs(opts, "RejectedOrder")
	if err != nil {
		return nil, err
	}
	return &CustodianWalletLogicRejectedOrderIterator{contract: _CustodianWalletLogic.contract, event: "RejectedOrder", logs: logs, sub: sub}, nil
}

// WatchRejectedOrder is a free log subscription operation binding the contract event 0xbd551ca6c38e2ad16afa16c81fcb4b8e891315559bfb80f551ab6b66cb583dbd.
//
// Solidity: event RejectedOrder(uint256 orderId)
func (_CustodianWalletLogic *CustodianWalletLogicFilterer) WatchRejectedOrder(opts *bind.WatchOpts, sink chan<- *CustodianWalletLogicRejectedOrder) (event.Subscription, error) {

	logs, sub, err := _CustodianWalletLogic.contract.WatchLogs(opts, "RejectedOrder")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CustodianWalletLogicRejectedOrder)
				if err := _CustodianWalletLogic.contract.UnpackLog(event, "RejectedOrder", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRejectedOrder is a log parse operation binding the contract event 0xbd551ca6c38e2ad16afa16c81fcb4b8e891315559bfb80f551ab6b66cb583dbd.
//
// Solidity: event RejectedOrder(uint256 orderId)
func (_CustodianWalletLogic *CustodianWalletLogicFilterer) ParseRejectedOrder(log types.Log) (*CustodianWalletLogicRejectedOrder, error) {
	event := new(CustodianWalletLogicRejectedOrder)
	if err := _CustodianWalletLogic.contract.UnpackLog(event, "RejectedOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
