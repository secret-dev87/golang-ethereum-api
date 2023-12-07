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

// TypesOrder is an auto generated low-level Go binding around an user-defined struct.
type TypesOrder struct {
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
}

// EscrowMetaData contains all meta data concerning the Escrow contract.
var EscrowMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ochestrator\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"ApproveRejectedOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fulfiledTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"orderStatus\",\"type\":\"uint8\"}],\"name\":\"ClosedOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"orderStatus\",\"type\":\"uint8\"}],\"name\":\"OpenOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"OrderFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"RejectedOrder\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_orderIndex\",\"type\":\"uint256\"}],\"name\":\"closeOpenOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_orderIndex\",\"type\":\"uint256\"}],\"name\":\"consentOrderRejected\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"}],\"name\":\"getOpenOrdersOf\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"}],\"name\":\"getOrderById\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"orderStatus\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fulfiledTime\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.Order\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_orderType\",\"type\":\"uint8\"}],\"name\":\"newOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"}],\"name\":\"numberOfOpenOrders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ochestrator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"openOrders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"orders\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"orderStatus\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fulfiledTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"}],\"name\":\"rejectOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usdcContractAddress\",\"type\":\"address\"}],\"name\":\"setUsdcTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdcToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawFeesEarned\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801562000010575f80fd5b5060405162001fcf38038062001fcf8339818101604052810190620000369190620002e4565b8060025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505f6040518061014001604052805f73ffffffffffffffffffffffffffffffffffffffff1681526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f81526020015f81526020015f81526020015f60ff168152602001600160ff1681526020014281526020014281525090505f81908060018154018082558091505060019003905f5260205f2090600902015f909190919091505f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151816002015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550606082015181600301556080820151816004015560a0820151816005015560c0820151816006015f6101000a81548160ff021916908360ff16021790555060e08201518160060160016101000a81548160ff021916908360ff160217905550610100820151816007015561012082015181600801555050505062000314565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f620002ae8262000283565b9050919050565b620002c081620002a2565b8114620002cb575f80fd5b50565b5f81519050620002de81620002b5565b92915050565b5f60208284031215620002fc57620002fb6200027f565b5b5f6200030b84828501620002ce565b91505092915050565b611cad80620003225f395ff3fe608060405234801561000f575f80fd5b50600436106100cd575f3560e01c80636dee20321161008a578063b00f215c11610064578063b00f215c1461022c578063b1fab21c14610248578063ef3bbbff14610278578063f3fd622f14610294576100cd565b80636dee2032146101a55780639c811d43146101d5578063a85c38ef146101f3576100cd565b806311a21b2c146100d157806311eac855146100db5780631e7cea80146100f9578063240f41601461012957806332c011171461014557806342b50a7a14610175575b5f80fd5b6100d96102b0565b005b6100e361049c565b6040516100f0919061130b565b60405180910390f35b610113600480360381019061010e9190611385565b6104c1565b60405161012091906113d2565b60405180910390f35b610143600480360381019061013e9190611385565b6104ec565b005b61015f600480360381019061015a91906113eb565b6105d6565b60405161016c91906113d2565b60405180910390f35b61018f600480360381019061018a9190611416565b61061f565b60405161019c9190611544565b60405180910390f35b6101bf60048036038101906101ba91906113eb565b6107c2565b6040516101cc9190611606565b60405180910390f35b6101dd610855565b6040516101ea919061130b565b60405180910390f35b61020d60048036038101906102089190611416565b61087a565b6040516102239a99989796959493929190611635565b60405180910390f35b61024660048036038101906102419190611385565b61094e565b005b610262600480360381019061025d91906116f9565b610be1565b60405161026f91906113d2565b60405180910390f35b610292600480360381019061028d9190611385565b610f9c565b005b6102ae60048036038101906102a991906113eb565b6110fb565b005b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461033f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610336906117f0565b60405180910390fd5b5f60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b815260040161039a919061130b565b602060405180830381865afa1580156103b5573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103d99190611822565b905060035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b815260040161045892919061184d565b6020604051808303815f875af1158015610474573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061049891906118a9565b5050565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6001602052815f5260405f2081815481106104da575f80fd5b905f5260205f20015f91509150505481565b8173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461055a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105519061191e565b60405180910390fd5b5f80828154811061056e5761056d61193c565b5b905f5260205f209060090201905060048160060160016101000a81548160ff021916908360ff1602179055507fbd551ca6c38e2ad16afa16c81fcb4b8e891315559bfb80f551ab6b66cb583dbd826040516105c991906113d2565b60405180910390a1505050565b5f60015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20805490509050919050565b61062761123b565b5f828154811061063a5761063961193c565b5b905f5260205f209060090201604051806101400160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600382015481526020016004820154815260200160058201548152602001600682015f9054906101000a900460ff1660ff1660ff1681526020016006820160019054906101000a900460ff1660ff1660ff168152602001600782015481526020016008820154815250509050919050565b606060015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2080548060200260200160405190810160405280929190818152602001828054801561084957602002820191905f5260205f20905b815481526020019060010190808311610835575b50505050509050919050565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f8181548110610888575f80fd5b905f5260205f2090600902015f91509050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806002015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806003015490806004015490806005015490806006015f9054906101000a900460ff16908060060160019054906101000a900460ff1690806007015490806008015490508a565b8173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146109bc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109b39061191e565b60405180910390fd5b5f60015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208281548110610a0c57610a0b61193c565b5b905f5260205f200154905060015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208281548110610a6657610a6561193c565b5b905f5260205f20015f90555f808281548110610a8557610a8461193c565b5b905f5260205f209060090201905042816008018190555060018160060160016101000a81548160ff021916908360ff160217905550806002015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16816001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16825f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f7296e5ad5c0ce57fe4cbceb3362886c7162b073357dd6b39025453458f6f848285856003015486600401548760050154886006015f9054906101000a900460ff1689600801548a60060160019054906101000a900460ff16604051610bd39796959493929190611969565b60405180910390a450505050565b5f8673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610c50576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c4790611a20565b60405180910390fd5b5f8511610c92576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c8990611a88565b60405180910390fd5b5f6040518061014001604052808a73ffffffffffffffffffffffffffffffffffffffff1681526020018973ffffffffffffffffffffffffffffffffffffffff1681526020018873ffffffffffffffffffffffffffffffffffffffff1681526020018781526020018681526020018581526020018460ff1681526020015f60ff1681526020014281526020015f81525090505f81908060018154018082558091505060019003905f5260205f2090600902015f909190919091505f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151816002015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550606082015181600301556080820151816004015560a0820151816005015560c0820151816006015f6101000a81548160ff021916908360ff16021790555060e08201518160060160016101000a81548160ff021916908360ff1602179055506101008201518160070155610120820151816008015550505f60015f80549050610ea49190611ad3565b905060015f8b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081908060018154018082558091505060019003905f5260205f20015f90919091909150558773ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff168b73ffffffffffffffffffffffffffffffffffffffff167e23c8dda5a5d69bbdf1f810b48e0b0fff6827cca652b18c81b25d8759bf67f2848b8b8b8b5f604051610f8496959493929190611b48565b60405180910390a48092505050979650505050505050565b8173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461100a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161100190611bf1565b60405180910390fd5b5f60015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20828154811061105a5761105961193c565b5b905f5260205f200154905060015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2082815481106110b4576110b361193c565b5b905f5260205f20015f90557f6526538f9bd98d105bdfc50096f40172dea60988b982f7776fa03d816600eac8816040516110ee91906113d2565b60405180910390a1505050565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461118a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611181906117f0565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036111f8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111ef90611c59565b60405180910390fd5b8060035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6040518061014001604052805f73ffffffffffffffffffffffffffffffffffffffff1681526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f81526020015f81526020015f81526020015f60ff1681526020015f60ff1681526020015f81526020015f81525090565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6112f5826112cc565b9050919050565b611305816112eb565b82525050565b5f60208201905061131e5f8301846112fc565b92915050565b5f80fd5b611331816112eb565b811461133b575f80fd5b50565b5f8135905061134c81611328565b92915050565b5f819050919050565b61136481611352565b811461136e575f80fd5b50565b5f8135905061137f8161135b565b92915050565b5f806040838503121561139b5761139a611324565b5b5f6113a88582860161133e565b92505060206113b985828601611371565b9150509250929050565b6113cc81611352565b82525050565b5f6020820190506113e55f8301846113c3565b92915050565b5f60208284031215611400576113ff611324565b5b5f61140d8482850161133e565b91505092915050565b5f6020828403121561142b5761142a611324565b5b5f61143884828501611371565b91505092915050565b61144a816112eb565b82525050565b61145981611352565b82525050565b5f60ff82169050919050565b6114748161145f565b82525050565b61014082015f82015161148f5f850182611441565b5060208201516114a26020850182611441565b5060408201516114b56040850182611441565b5060608201516114c86060850182611450565b5060808201516114db6080850182611450565b5060a08201516114ee60a0850182611450565b5060c082015161150160c085018261146b565b5060e082015161151460e085018261146b565b50610100820151611529610100850182611450565b5061012082015161153e610120850182611450565b50505050565b5f610140820190506115585f83018461147a565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f6115928383611450565b60208301905092915050565b5f602082019050919050565b5f6115b48261155e565b6115be8185611568565b93506115c983611578565b805f5b838110156115f95781516115e08882611587565b97506115eb8361159e565b9250506001810190506115cc565b5085935050505092915050565b5f6020820190508181035f83015261161e81846115aa565b905092915050565b61162f8161145f565b82525050565b5f610140820190506116495f83018d6112fc565b611656602083018c6112fc565b611663604083018b6112fc565b611670606083018a6113c3565b61167d60808301896113c3565b61168a60a08301886113c3565b61169760c0830187611626565b6116a460e0830186611626565b6116b26101008301856113c3565b6116c06101208301846113c3565b9b9a5050505050505050505050565b6116d88161145f565b81146116e2575f80fd5b50565b5f813590506116f3816116cf565b92915050565b5f805f805f805f60e0888a03121561171457611713611324565b5b5f6117218a828b0161133e565b97505060206117328a828b0161133e565b96505060406117438a828b0161133e565b95505060606117548a828b01611371565b94505060806117658a828b01611371565b93505060a06117768a828b01611371565b92505060c06117878a828b016116e5565b91505092959891949750929550565b5f82825260208201905092915050565b7f463a206f6e6c79206f63686573747261746f72000000000000000000000000005f82015250565b5f6117da601383611796565b91506117e5826117a6565b602082019050919050565b5f6020820190508181035f830152611807816117ce565b9050919050565b5f8151905061181c8161135b565b92915050565b5f6020828403121561183757611836611324565b5b5f6118448482850161180e565b91505092915050565b5f6040820190506118605f8301856112fc565b61186d60208301846113c3565b9392505050565b5f8115159050919050565b61188881611874565b8114611892575f80fd5b50565b5f815190506118a38161187f565b92915050565b5f602082840312156118be576118bd611324565b5b5f6118cb84828501611895565b91505092915050565b7f433a206f6e6c792073656c6c65720000000000000000000000000000000000005f82015250565b5f611908600e83611796565b9150611913826118d4565b602082019050919050565b5f6020820190508181035f830152611935816118fc565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f60e08201905061197c5f83018a6113c3565b61198960208301896113c3565b61199660408301886113c3565b6119a360608301876113c3565b6119b06080830186611626565b6119bd60a08301856113c3565b6119ca60c0830184611626565b98975050505050505050565b7f433a20637573746f6d6572206f6e6c79000000000000000000000000000000005f82015250565b5f611a0a601083611796565b9150611a15826119d6565b602082019050919050565b5f6020820190508181035f830152611a37816119fe565b9050919050565b7f433a20696e76616c6964206f72646572000000000000000000000000000000005f82015250565b5f611a72601083611796565b9150611a7d82611a3e565b602082019050919050565b5f6020820190508181035f830152611a9f81611a66565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611add82611352565b9150611ae883611352565b9250828203905081811115611b0057611aff611aa6565b5b92915050565b5f819050919050565b5f819050919050565b5f611b32611b2d611b2884611b06565b611b0f565b61145f565b9050919050565b611b4281611b18565b82525050565b5f60c082019050611b5b5f8301896113c3565b611b6860208301886113c3565b611b7560408301876113c3565b611b8260608301866113c3565b611b8f6080830185611626565b611b9c60a0830184611b39565b979650505050505050565b7f433a206f6e6c79206275796572000000000000000000000000000000000000005f82015250565b5f611bdb600d83611796565b9150611be682611ba7565b602082019050919050565b5f6020820190508181035f830152611c0881611bcf565b9050919050565b7f463a20696e76616c6964206164647265737300000000000000000000000000005f82015250565b5f611c43601283611796565b9150611c4e82611c0f565b602082019050919050565b5f6020820190508181035f830152611c7081611c37565b905091905056fea26469706673582212202eb2ed58458269e7514129ea2c14eb6c165c7a15856d3576a8ae1c928e419b9164736f6c63430008170033",
}

// EscrowABI is the input ABI used to generate the binding from.
// Deprecated: Use EscrowMetaData.ABI instead.
var EscrowABI = EscrowMetaData.ABI

// EscrowBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EscrowMetaData.Bin instead.
var EscrowBin = EscrowMetaData.Bin

// DeployEscrow deploys a new Ethereum contract, binding an instance of Escrow to it.
func DeployEscrow(auth *bind.TransactOpts, backend bind.ContractBackend, _ochestrator common.Address) (common.Address, *types.Transaction, *Escrow, error) {
	parsed, err := EscrowMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EscrowBin), backend, _ochestrator)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Escrow{EscrowCaller: EscrowCaller{contract: contract}, EscrowTransactor: EscrowTransactor{contract: contract}, EscrowFilterer: EscrowFilterer{contract: contract}}, nil
}

// Escrow is an auto generated Go binding around an Ethereum contract.
type Escrow struct {
	EscrowCaller     // Read-only binding to the contract
	EscrowTransactor // Write-only binding to the contract
	EscrowFilterer   // Log filterer for contract events
}

// EscrowCaller is an auto generated read-only Go binding around an Ethereum contract.
type EscrowCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EscrowTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EscrowTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EscrowFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EscrowFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EscrowSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EscrowSession struct {
	Contract     *Escrow           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EscrowCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EscrowCallerSession struct {
	Contract *EscrowCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EscrowTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EscrowTransactorSession struct {
	Contract     *EscrowTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EscrowRaw is an auto generated low-level Go binding around an Ethereum contract.
type EscrowRaw struct {
	Contract *Escrow // Generic contract binding to access the raw methods on
}

// EscrowCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EscrowCallerRaw struct {
	Contract *EscrowCaller // Generic read-only contract binding to access the raw methods on
}

// EscrowTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EscrowTransactorRaw struct {
	Contract *EscrowTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEscrow creates a new instance of Escrow, bound to a specific deployed contract.
func NewEscrow(address common.Address, backend bind.ContractBackend) (*Escrow, error) {
	contract, err := bindEscrow(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Escrow{EscrowCaller: EscrowCaller{contract: contract}, EscrowTransactor: EscrowTransactor{contract: contract}, EscrowFilterer: EscrowFilterer{contract: contract}}, nil
}

// NewEscrowCaller creates a new read-only instance of Escrow, bound to a specific deployed contract.
func NewEscrowCaller(address common.Address, caller bind.ContractCaller) (*EscrowCaller, error) {
	contract, err := bindEscrow(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EscrowCaller{contract: contract}, nil
}

// NewEscrowTransactor creates a new write-only instance of Escrow, bound to a specific deployed contract.
func NewEscrowTransactor(address common.Address, transactor bind.ContractTransactor) (*EscrowTransactor, error) {
	contract, err := bindEscrow(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EscrowTransactor{contract: contract}, nil
}

// NewEscrowFilterer creates a new log filterer instance of Escrow, bound to a specific deployed contract.
func NewEscrowFilterer(address common.Address, filterer bind.ContractFilterer) (*EscrowFilterer, error) {
	contract, err := bindEscrow(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EscrowFilterer{contract: contract}, nil
}

// bindEscrow binds a generic wrapper to an already deployed contract.
func bindEscrow(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EscrowMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Escrow *EscrowRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Escrow.Contract.EscrowCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Escrow *EscrowRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Escrow.Contract.EscrowTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Escrow *EscrowRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Escrow.Contract.EscrowTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Escrow *EscrowCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Escrow.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Escrow *EscrowTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Escrow.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Escrow *EscrowTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Escrow.Contract.contract.Transact(opts, method, params...)
}

// GetOpenOrdersOf is a free data retrieval call binding the contract method 0x6dee2032.
//
// Solidity: function getOpenOrdersOf(address _seller) view returns(uint256[])
func (_Escrow *EscrowCaller) GetOpenOrdersOf(opts *bind.CallOpts, _seller common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "getOpenOrdersOf", _seller)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetOpenOrdersOf is a free data retrieval call binding the contract method 0x6dee2032.
//
// Solidity: function getOpenOrdersOf(address _seller) view returns(uint256[])
func (_Escrow *EscrowSession) GetOpenOrdersOf(_seller common.Address) ([]*big.Int, error) {
	return _Escrow.Contract.GetOpenOrdersOf(&_Escrow.CallOpts, _seller)
}

// GetOpenOrdersOf is a free data retrieval call binding the contract method 0x6dee2032.
//
// Solidity: function getOpenOrdersOf(address _seller) view returns(uint256[])
func (_Escrow *EscrowCallerSession) GetOpenOrdersOf(_seller common.Address) ([]*big.Int, error) {
	return _Escrow.Contract.GetOpenOrdersOf(&_Escrow.CallOpts, _seller)
}

// GetOrderById is a free data retrieval call binding the contract method 0x42b50a7a.
//
// Solidity: function getOrderById(uint256 _orderId) view returns((address,address,address,uint256,uint256,uint256,uint8,uint8,uint256,uint256))
func (_Escrow *EscrowCaller) GetOrderById(opts *bind.CallOpts, _orderId *big.Int) (TypesOrder, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "getOrderById", _orderId)

	if err != nil {
		return *new(TypesOrder), err
	}

	out0 := *abi.ConvertType(out[0], new(TypesOrder)).(*TypesOrder)

	return out0, err

}

// GetOrderById is a free data retrieval call binding the contract method 0x42b50a7a.
//
// Solidity: function getOrderById(uint256 _orderId) view returns((address,address,address,uint256,uint256,uint256,uint8,uint8,uint256,uint256))
func (_Escrow *EscrowSession) GetOrderById(_orderId *big.Int) (TypesOrder, error) {
	return _Escrow.Contract.GetOrderById(&_Escrow.CallOpts, _orderId)
}

// GetOrderById is a free data retrieval call binding the contract method 0x42b50a7a.
//
// Solidity: function getOrderById(uint256 _orderId) view returns((address,address,address,uint256,uint256,uint256,uint8,uint8,uint256,uint256))
func (_Escrow *EscrowCallerSession) GetOrderById(_orderId *big.Int) (TypesOrder, error) {
	return _Escrow.Contract.GetOrderById(&_Escrow.CallOpts, _orderId)
}

// NumberOfOpenOrders is a free data retrieval call binding the contract method 0x32c01117.
//
// Solidity: function numberOfOpenOrders(address _seller) view returns(uint256)
func (_Escrow *EscrowCaller) NumberOfOpenOrders(opts *bind.CallOpts, _seller common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "numberOfOpenOrders", _seller)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumberOfOpenOrders is a free data retrieval call binding the contract method 0x32c01117.
//
// Solidity: function numberOfOpenOrders(address _seller) view returns(uint256)
func (_Escrow *EscrowSession) NumberOfOpenOrders(_seller common.Address) (*big.Int, error) {
	return _Escrow.Contract.NumberOfOpenOrders(&_Escrow.CallOpts, _seller)
}

// NumberOfOpenOrders is a free data retrieval call binding the contract method 0x32c01117.
//
// Solidity: function numberOfOpenOrders(address _seller) view returns(uint256)
func (_Escrow *EscrowCallerSession) NumberOfOpenOrders(_seller common.Address) (*big.Int, error) {
	return _Escrow.Contract.NumberOfOpenOrders(&_Escrow.CallOpts, _seller)
}

// Ochestrator is a free data retrieval call binding the contract method 0x9c811d43.
//
// Solidity: function ochestrator() view returns(address)
func (_Escrow *EscrowCaller) Ochestrator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "ochestrator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ochestrator is a free data retrieval call binding the contract method 0x9c811d43.
//
// Solidity: function ochestrator() view returns(address)
func (_Escrow *EscrowSession) Ochestrator() (common.Address, error) {
	return _Escrow.Contract.Ochestrator(&_Escrow.CallOpts)
}

// Ochestrator is a free data retrieval call binding the contract method 0x9c811d43.
//
// Solidity: function ochestrator() view returns(address)
func (_Escrow *EscrowCallerSession) Ochestrator() (common.Address, error) {
	return _Escrow.Contract.Ochestrator(&_Escrow.CallOpts)
}

// OpenOrders is a free data retrieval call binding the contract method 0x1e7cea80.
//
// Solidity: function openOrders(address , uint256 ) view returns(uint256)
func (_Escrow *EscrowCaller) OpenOrders(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "openOrders", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OpenOrders is a free data retrieval call binding the contract method 0x1e7cea80.
//
// Solidity: function openOrders(address , uint256 ) view returns(uint256)
func (_Escrow *EscrowSession) OpenOrders(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Escrow.Contract.OpenOrders(&_Escrow.CallOpts, arg0, arg1)
}

// OpenOrders is a free data retrieval call binding the contract method 0x1e7cea80.
//
// Solidity: function openOrders(address , uint256 ) view returns(uint256)
func (_Escrow *EscrowCallerSession) OpenOrders(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Escrow.Contract.OpenOrders(&_Escrow.CallOpts, arg0, arg1)
}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(address seller, address buyer, address receiver, uint256 amount, uint256 rate, uint256 fee, uint8 orderType, uint8 orderStatus, uint256 startTime, uint256 fulfiledTime)
func (_Escrow *EscrowCaller) Orders(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	err := _Escrow.contract.Call(opts, &out, "orders", arg0)

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
func (_Escrow *EscrowSession) Orders(arg0 *big.Int) (struct {
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
	return _Escrow.Contract.Orders(&_Escrow.CallOpts, arg0)
}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(address seller, address buyer, address receiver, uint256 amount, uint256 rate, uint256 fee, uint8 orderType, uint8 orderStatus, uint256 startTime, uint256 fulfiledTime)
func (_Escrow *EscrowCallerSession) Orders(arg0 *big.Int) (struct {
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
	return _Escrow.Contract.Orders(&_Escrow.CallOpts, arg0)
}

// UsdcToken is a free data retrieval call binding the contract method 0x11eac855.
//
// Solidity: function usdcToken() view returns(address)
func (_Escrow *EscrowCaller) UsdcToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "usdcToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UsdcToken is a free data retrieval call binding the contract method 0x11eac855.
//
// Solidity: function usdcToken() view returns(address)
func (_Escrow *EscrowSession) UsdcToken() (common.Address, error) {
	return _Escrow.Contract.UsdcToken(&_Escrow.CallOpts)
}

// UsdcToken is a free data retrieval call binding the contract method 0x11eac855.
//
// Solidity: function usdcToken() view returns(address)
func (_Escrow *EscrowCallerSession) UsdcToken() (common.Address, error) {
	return _Escrow.Contract.UsdcToken(&_Escrow.CallOpts)
}

// CloseOpenOrder is a paid mutator transaction binding the contract method 0xb00f215c.
//
// Solidity: function closeOpenOrder(address _seller, uint256 _orderIndex) returns()
func (_Escrow *EscrowTransactor) CloseOpenOrder(opts *bind.TransactOpts, _seller common.Address, _orderIndex *big.Int) (*types.Transaction, error) {
	return _Escrow.contract.Transact(opts, "closeOpenOrder", _seller, _orderIndex)
}

// CloseOpenOrder is a paid mutator transaction binding the contract method 0xb00f215c.
//
// Solidity: function closeOpenOrder(address _seller, uint256 _orderIndex) returns()
func (_Escrow *EscrowSession) CloseOpenOrder(_seller common.Address, _orderIndex *big.Int) (*types.Transaction, error) {
	return _Escrow.Contract.CloseOpenOrder(&_Escrow.TransactOpts, _seller, _orderIndex)
}

// CloseOpenOrder is a paid mutator transaction binding the contract method 0xb00f215c.
//
// Solidity: function closeOpenOrder(address _seller, uint256 _orderIndex) returns()
func (_Escrow *EscrowTransactorSession) CloseOpenOrder(_seller common.Address, _orderIndex *big.Int) (*types.Transaction, error) {
	return _Escrow.Contract.CloseOpenOrder(&_Escrow.TransactOpts, _seller, _orderIndex)
}

// ConsentOrderRejected is a paid mutator transaction binding the contract method 0xef3bbbff.
//
// Solidity: function consentOrderRejected(address _buyer, uint256 _orderIndex) returns()
func (_Escrow *EscrowTransactor) ConsentOrderRejected(opts *bind.TransactOpts, _buyer common.Address, _orderIndex *big.Int) (*types.Transaction, error) {
	return _Escrow.contract.Transact(opts, "consentOrderRejected", _buyer, _orderIndex)
}

// ConsentOrderRejected is a paid mutator transaction binding the contract method 0xef3bbbff.
//
// Solidity: function consentOrderRejected(address _buyer, uint256 _orderIndex) returns()
func (_Escrow *EscrowSession) ConsentOrderRejected(_buyer common.Address, _orderIndex *big.Int) (*types.Transaction, error) {
	return _Escrow.Contract.ConsentOrderRejected(&_Escrow.TransactOpts, _buyer, _orderIndex)
}

// ConsentOrderRejected is a paid mutator transaction binding the contract method 0xef3bbbff.
//
// Solidity: function consentOrderRejected(address _buyer, uint256 _orderIndex) returns()
func (_Escrow *EscrowTransactorSession) ConsentOrderRejected(_buyer common.Address, _orderIndex *big.Int) (*types.Transaction, error) {
	return _Escrow.Contract.ConsentOrderRejected(&_Escrow.TransactOpts, _buyer, _orderIndex)
}

// NewOrder is a paid mutator transaction binding the contract method 0xb1fab21c.
//
// Solidity: function newOrder(address _seller, address _buyer, address _receiver, uint256 _amount, uint256 _rate, uint256 _fee, uint8 _orderType) returns(uint256)
func (_Escrow *EscrowTransactor) NewOrder(opts *bind.TransactOpts, _seller common.Address, _buyer common.Address, _receiver common.Address, _amount *big.Int, _rate *big.Int, _fee *big.Int, _orderType uint8) (*types.Transaction, error) {
	return _Escrow.contract.Transact(opts, "newOrder", _seller, _buyer, _receiver, _amount, _rate, _fee, _orderType)
}

// NewOrder is a paid mutator transaction binding the contract method 0xb1fab21c.
//
// Solidity: function newOrder(address _seller, address _buyer, address _receiver, uint256 _amount, uint256 _rate, uint256 _fee, uint8 _orderType) returns(uint256)
func (_Escrow *EscrowSession) NewOrder(_seller common.Address, _buyer common.Address, _receiver common.Address, _amount *big.Int, _rate *big.Int, _fee *big.Int, _orderType uint8) (*types.Transaction, error) {
	return _Escrow.Contract.NewOrder(&_Escrow.TransactOpts, _seller, _buyer, _receiver, _amount, _rate, _fee, _orderType)
}

// NewOrder is a paid mutator transaction binding the contract method 0xb1fab21c.
//
// Solidity: function newOrder(address _seller, address _buyer, address _receiver, uint256 _amount, uint256 _rate, uint256 _fee, uint8 _orderType) returns(uint256)
func (_Escrow *EscrowTransactorSession) NewOrder(_seller common.Address, _buyer common.Address, _receiver common.Address, _amount *big.Int, _rate *big.Int, _fee *big.Int, _orderType uint8) (*types.Transaction, error) {
	return _Escrow.Contract.NewOrder(&_Escrow.TransactOpts, _seller, _buyer, _receiver, _amount, _rate, _fee, _orderType)
}

// RejectOrder is a paid mutator transaction binding the contract method 0x240f4160.
//
// Solidity: function rejectOrder(address _seller, uint256 _orderId) returns()
func (_Escrow *EscrowTransactor) RejectOrder(opts *bind.TransactOpts, _seller common.Address, _orderId *big.Int) (*types.Transaction, error) {
	return _Escrow.contract.Transact(opts, "rejectOrder", _seller, _orderId)
}

// RejectOrder is a paid mutator transaction binding the contract method 0x240f4160.
//
// Solidity: function rejectOrder(address _seller, uint256 _orderId) returns()
func (_Escrow *EscrowSession) RejectOrder(_seller common.Address, _orderId *big.Int) (*types.Transaction, error) {
	return _Escrow.Contract.RejectOrder(&_Escrow.TransactOpts, _seller, _orderId)
}

// RejectOrder is a paid mutator transaction binding the contract method 0x240f4160.
//
// Solidity: function rejectOrder(address _seller, uint256 _orderId) returns()
func (_Escrow *EscrowTransactorSession) RejectOrder(_seller common.Address, _orderId *big.Int) (*types.Transaction, error) {
	return _Escrow.Contract.RejectOrder(&_Escrow.TransactOpts, _seller, _orderId)
}

// SetUsdcTokenAddress is a paid mutator transaction binding the contract method 0xf3fd622f.
//
// Solidity: function setUsdcTokenAddress(address usdcContractAddress) returns()
func (_Escrow *EscrowTransactor) SetUsdcTokenAddress(opts *bind.TransactOpts, usdcContractAddress common.Address) (*types.Transaction, error) {
	return _Escrow.contract.Transact(opts, "setUsdcTokenAddress", usdcContractAddress)
}

// SetUsdcTokenAddress is a paid mutator transaction binding the contract method 0xf3fd622f.
//
// Solidity: function setUsdcTokenAddress(address usdcContractAddress) returns()
func (_Escrow *EscrowSession) SetUsdcTokenAddress(usdcContractAddress common.Address) (*types.Transaction, error) {
	return _Escrow.Contract.SetUsdcTokenAddress(&_Escrow.TransactOpts, usdcContractAddress)
}

// SetUsdcTokenAddress is a paid mutator transaction binding the contract method 0xf3fd622f.
//
// Solidity: function setUsdcTokenAddress(address usdcContractAddress) returns()
func (_Escrow *EscrowTransactorSession) SetUsdcTokenAddress(usdcContractAddress common.Address) (*types.Transaction, error) {
	return _Escrow.Contract.SetUsdcTokenAddress(&_Escrow.TransactOpts, usdcContractAddress)
}

// WithdrawFeesEarned is a paid mutator transaction binding the contract method 0x11a21b2c.
//
// Solidity: function withdrawFeesEarned() returns()
func (_Escrow *EscrowTransactor) WithdrawFeesEarned(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Escrow.contract.Transact(opts, "withdrawFeesEarned")
}

// WithdrawFeesEarned is a paid mutator transaction binding the contract method 0x11a21b2c.
//
// Solidity: function withdrawFeesEarned() returns()
func (_Escrow *EscrowSession) WithdrawFeesEarned() (*types.Transaction, error) {
	return _Escrow.Contract.WithdrawFeesEarned(&_Escrow.TransactOpts)
}

// WithdrawFeesEarned is a paid mutator transaction binding the contract method 0x11a21b2c.
//
// Solidity: function withdrawFeesEarned() returns()
func (_Escrow *EscrowTransactorSession) WithdrawFeesEarned() (*types.Transaction, error) {
	return _Escrow.Contract.WithdrawFeesEarned(&_Escrow.TransactOpts)
}

// EscrowApproveRejectedOrderIterator is returned from FilterApproveRejectedOrder and is used to iterate over the raw logs and unpacked data for ApproveRejectedOrder events raised by the Escrow contract.
type EscrowApproveRejectedOrderIterator struct {
	Event *EscrowApproveRejectedOrder // Event containing the contract specifics and raw log

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
func (it *EscrowApproveRejectedOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EscrowApproveRejectedOrder)
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
		it.Event = new(EscrowApproveRejectedOrder)
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
func (it *EscrowApproveRejectedOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EscrowApproveRejectedOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EscrowApproveRejectedOrder represents a ApproveRejectedOrder event raised by the Escrow contract.
type EscrowApproveRejectedOrder struct {
	OrderId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproveRejectedOrder is a free log retrieval operation binding the contract event 0x6526538f9bd98d105bdfc50096f40172dea60988b982f7776fa03d816600eac8.
//
// Solidity: event ApproveRejectedOrder(uint256 orderId)
func (_Escrow *EscrowFilterer) FilterApproveRejectedOrder(opts *bind.FilterOpts) (*EscrowApproveRejectedOrderIterator, error) {

	logs, sub, err := _Escrow.contract.FilterLogs(opts, "ApproveRejectedOrder")
	if err != nil {
		return nil, err
	}
	return &EscrowApproveRejectedOrderIterator{contract: _Escrow.contract, event: "ApproveRejectedOrder", logs: logs, sub: sub}, nil
}

// WatchApproveRejectedOrder is a free log subscription operation binding the contract event 0x6526538f9bd98d105bdfc50096f40172dea60988b982f7776fa03d816600eac8.
//
// Solidity: event ApproveRejectedOrder(uint256 orderId)
func (_Escrow *EscrowFilterer) WatchApproveRejectedOrder(opts *bind.WatchOpts, sink chan<- *EscrowApproveRejectedOrder) (event.Subscription, error) {

	logs, sub, err := _Escrow.contract.WatchLogs(opts, "ApproveRejectedOrder")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EscrowApproveRejectedOrder)
				if err := _Escrow.contract.UnpackLog(event, "ApproveRejectedOrder", log); err != nil {
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
func (_Escrow *EscrowFilterer) ParseApproveRejectedOrder(log types.Log) (*EscrowApproveRejectedOrder, error) {
	event := new(EscrowApproveRejectedOrder)
	if err := _Escrow.contract.UnpackLog(event, "ApproveRejectedOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EscrowClosedOrderIterator is returned from FilterClosedOrder and is used to iterate over the raw logs and unpacked data for ClosedOrder events raised by the Escrow contract.
type EscrowClosedOrderIterator struct {
	Event *EscrowClosedOrder // Event containing the contract specifics and raw log

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
func (it *EscrowClosedOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EscrowClosedOrder)
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
		it.Event = new(EscrowClosedOrder)
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
func (it *EscrowClosedOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EscrowClosedOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EscrowClosedOrder represents a ClosedOrder event raised by the Escrow contract.
type EscrowClosedOrder struct {
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
func (_Escrow *EscrowFilterer) FilterClosedOrder(opts *bind.FilterOpts, seller []common.Address, buyer []common.Address, receiver []common.Address) (*EscrowClosedOrderIterator, error) {

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

	logs, sub, err := _Escrow.contract.FilterLogs(opts, "ClosedOrder", sellerRule, buyerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &EscrowClosedOrderIterator{contract: _Escrow.contract, event: "ClosedOrder", logs: logs, sub: sub}, nil
}

// WatchClosedOrder is a free log subscription operation binding the contract event 0x7296e5ad5c0ce57fe4cbceb3362886c7162b073357dd6b39025453458f6f8482.
//
// Solidity: event ClosedOrder(uint256 orderId, address indexed seller, address indexed buyer, address indexed receiver, uint256 amount, uint256 rate, uint256 fee, uint8 orderType, uint256 fulfiledTime, uint8 orderStatus)
func (_Escrow *EscrowFilterer) WatchClosedOrder(opts *bind.WatchOpts, sink chan<- *EscrowClosedOrder, seller []common.Address, buyer []common.Address, receiver []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Escrow.contract.WatchLogs(opts, "ClosedOrder", sellerRule, buyerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EscrowClosedOrder)
				if err := _Escrow.contract.UnpackLog(event, "ClosedOrder", log); err != nil {
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
func (_Escrow *EscrowFilterer) ParseClosedOrder(log types.Log) (*EscrowClosedOrder, error) {
	event := new(EscrowClosedOrder)
	if err := _Escrow.contract.UnpackLog(event, "ClosedOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EscrowOpenOrderIterator is returned from FilterOpenOrder and is used to iterate over the raw logs and unpacked data for OpenOrder events raised by the Escrow contract.
type EscrowOpenOrderIterator struct {
	Event *EscrowOpenOrder // Event containing the contract specifics and raw log

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
func (it *EscrowOpenOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EscrowOpenOrder)
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
		it.Event = new(EscrowOpenOrder)
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
func (it *EscrowOpenOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EscrowOpenOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EscrowOpenOrder represents a OpenOrder event raised by the Escrow contract.
type EscrowOpenOrder struct {
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
func (_Escrow *EscrowFilterer) FilterOpenOrder(opts *bind.FilterOpts, seller []common.Address, buyer []common.Address, receiver []common.Address) (*EscrowOpenOrderIterator, error) {

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

	logs, sub, err := _Escrow.contract.FilterLogs(opts, "OpenOrder", sellerRule, buyerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &EscrowOpenOrderIterator{contract: _Escrow.contract, event: "OpenOrder", logs: logs, sub: sub}, nil
}

// WatchOpenOrder is a free log subscription operation binding the contract event 0x0023c8dda5a5d69bbdf1f810b48e0b0fff6827cca652b18c81b25d8759bf67f2.
//
// Solidity: event OpenOrder(uint256 orderId, address indexed seller, address indexed buyer, address indexed receiver, uint256 amount, uint256 rate, uint256 fee, uint8 orderType, uint8 orderStatus)
func (_Escrow *EscrowFilterer) WatchOpenOrder(opts *bind.WatchOpts, sink chan<- *EscrowOpenOrder, seller []common.Address, buyer []common.Address, receiver []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Escrow.contract.WatchLogs(opts, "OpenOrder", sellerRule, buyerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EscrowOpenOrder)
				if err := _Escrow.contract.UnpackLog(event, "OpenOrder", log); err != nil {
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
func (_Escrow *EscrowFilterer) ParseOpenOrder(log types.Log) (*EscrowOpenOrder, error) {
	event := new(EscrowOpenOrder)
	if err := _Escrow.contract.UnpackLog(event, "OpenOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EscrowOrderFulfilledIterator is returned from FilterOrderFulfilled and is used to iterate over the raw logs and unpacked data for OrderFulfilled events raised by the Escrow contract.
type EscrowOrderFulfilledIterator struct {
	Event *EscrowOrderFulfilled // Event containing the contract specifics and raw log

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
func (it *EscrowOrderFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EscrowOrderFulfilled)
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
		it.Event = new(EscrowOrderFulfilled)
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
func (it *EscrowOrderFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EscrowOrderFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EscrowOrderFulfilled represents a OrderFulfilled event raised by the Escrow contract.
type EscrowOrderFulfilled struct {
	OrderId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOrderFulfilled is a free log retrieval operation binding the contract event 0xd276fbc5e545783825cac5416bad154c47e39daebf44589255a65b4f56cf3e0e.
//
// Solidity: event OrderFulfilled(uint256 orderId)
func (_Escrow *EscrowFilterer) FilterOrderFulfilled(opts *bind.FilterOpts) (*EscrowOrderFulfilledIterator, error) {

	logs, sub, err := _Escrow.contract.FilterLogs(opts, "OrderFulfilled")
	if err != nil {
		return nil, err
	}
	return &EscrowOrderFulfilledIterator{contract: _Escrow.contract, event: "OrderFulfilled", logs: logs, sub: sub}, nil
}

// WatchOrderFulfilled is a free log subscription operation binding the contract event 0xd276fbc5e545783825cac5416bad154c47e39daebf44589255a65b4f56cf3e0e.
//
// Solidity: event OrderFulfilled(uint256 orderId)
func (_Escrow *EscrowFilterer) WatchOrderFulfilled(opts *bind.WatchOpts, sink chan<- *EscrowOrderFulfilled) (event.Subscription, error) {

	logs, sub, err := _Escrow.contract.WatchLogs(opts, "OrderFulfilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EscrowOrderFulfilled)
				if err := _Escrow.contract.UnpackLog(event, "OrderFulfilled", log); err != nil {
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
func (_Escrow *EscrowFilterer) ParseOrderFulfilled(log types.Log) (*EscrowOrderFulfilled, error) {
	event := new(EscrowOrderFulfilled)
	if err := _Escrow.contract.UnpackLog(event, "OrderFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EscrowRejectedOrderIterator is returned from FilterRejectedOrder and is used to iterate over the raw logs and unpacked data for RejectedOrder events raised by the Escrow contract.
type EscrowRejectedOrderIterator struct {
	Event *EscrowRejectedOrder // Event containing the contract specifics and raw log

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
func (it *EscrowRejectedOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EscrowRejectedOrder)
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
		it.Event = new(EscrowRejectedOrder)
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
func (it *EscrowRejectedOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EscrowRejectedOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EscrowRejectedOrder represents a RejectedOrder event raised by the Escrow contract.
type EscrowRejectedOrder struct {
	OrderId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRejectedOrder is a free log retrieval operation binding the contract event 0xbd551ca6c38e2ad16afa16c81fcb4b8e891315559bfb80f551ab6b66cb583dbd.
//
// Solidity: event RejectedOrder(uint256 orderId)
func (_Escrow *EscrowFilterer) FilterRejectedOrder(opts *bind.FilterOpts) (*EscrowRejectedOrderIterator, error) {

	logs, sub, err := _Escrow.contract.FilterLogs(opts, "RejectedOrder")
	if err != nil {
		return nil, err
	}
	return &EscrowRejectedOrderIterator{contract: _Escrow.contract, event: "RejectedOrder", logs: logs, sub: sub}, nil
}

// WatchRejectedOrder is a free log subscription operation binding the contract event 0xbd551ca6c38e2ad16afa16c81fcb4b8e891315559bfb80f551ab6b66cb583dbd.
//
// Solidity: event RejectedOrder(uint256 orderId)
func (_Escrow *EscrowFilterer) WatchRejectedOrder(opts *bind.WatchOpts, sink chan<- *EscrowRejectedOrder) (event.Subscription, error) {

	logs, sub, err := _Escrow.contract.WatchLogs(opts, "RejectedOrder")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EscrowRejectedOrder)
				if err := _Escrow.contract.UnpackLog(event, "RejectedOrder", log); err != nil {
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
func (_Escrow *EscrowFilterer) ParseRejectedOrder(log types.Log) (*EscrowRejectedOrder, error) {
	event := new(EscrowRejectedOrder)
	if err := _Escrow.contract.UnpackLog(event, "RejectedOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
