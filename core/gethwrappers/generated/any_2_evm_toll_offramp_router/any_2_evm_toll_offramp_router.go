// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package any_2_evm_toll_offramp_router

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated"
)

var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

type CCIPAny2EVMMessageFromSender struct {
	SourceChainId *big.Int
	Sender        []byte
	Receiver      common.Address
	Data          []byte
	DestTokens    []common.Address
	DestPools     []common.Address
	Amounts       []*big.Int
	GasLimit      *big.Int
}

var Any2EVMTollOffRampRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractBaseOffRampInterface[]\",\"name\":\"offRamps\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"contractBaseOffRampInterface\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"AlreadyConfigured\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"MustCallFromOffRamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoOffRampsConfigured\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractBaseOffRampInterface\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"OffRampNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractBaseOffRampInterface\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"OffRampAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractBaseOffRampInterface\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"OffRampRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractBaseOffRampInterface\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"addOffRamp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOffRamps\",\"outputs\":[{\"internalType\":\"contractBaseOffRampInterface[]\",\"name\":\"offRamps\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractBaseOffRampInterface\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"isOffRamp\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractBaseOffRampInterface\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"removeOffRamp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"destTokens\",\"type\":\"address[]\"},{\"internalType\":\"contractPoolInterface[]\",\"name\":\"destPools\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMMessageFromSender\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"routeMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620016cf380380620016cf8339810160408190526200003491620002f9565b8033806000816200008c5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000bf57620000bf816200019a565b50508151620000d79150600390602084019062000245565b5060005b815181101562000191576040518060400160405280826001600160601b0316815260200160011515815250600260008484815181106200011f576200011f620003cb565b6020908102919091018101516001600160a01b031682528181019290925260400160002082518154939092015115156c01000000000000000000000000026001600160681b03199093166001600160601b03909216919091179190911790556200018981620003e1565b9050620000db565b50505062000409565b336001600160a01b03821603620001f45760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000083565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b8280548282559060005260206000209081019282156200029d579160200282015b828111156200029d57825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019062000266565b50620002ab929150620002af565b5090565b5b80821115620002ab5760008155600101620002b0565b634e487b7160e01b600052604160045260246000fd5b80516001600160a01b0381168114620002f457600080fd5b919050565b600060208083850312156200030d57600080fd5b82516001600160401b03808211156200032557600080fd5b818501915085601f8301126200033a57600080fd5b8151818111156200034f576200034f620002c6565b8060051b604051601f19603f83011681018181108582111715620003775762000377620002c6565b6040529182528482019250838101850191888311156200039657600080fd5b938501935b82851015620003bf57620003af85620002dc565b845293850193928501926200039b565b98975050505050505050565b634e487b7160e01b600052603260045260246000fd5b6000600182016200040257634e487b7160e01b600052601160045260246000fd5b5060010190565b6112b680620004196000396000f3fe608060405234801561001057600080fd5b50600436106100a35760003560e01c80638da5cb5b11610076578063a40e69c71161005b578063a40e69c7146101ab578063adb9f71b146101c0578063f2fde38b146101d357600080fd5b80638da5cb5b14610170578063991f65431461019857600080fd5b8063181f5a77146100a85780631d7a74a0146100fa5780635c9fa11e1461015357806379ba509714610166575b600080fd5b6100e46040518060400160405280601e81526020017f416e793245564d546f6c6c4f666652616d70526f7574657220312e302e30000081525081565b6040516100f19190610cbc565b60405180910390f35b610143610108366004610cf8565b73ffffffffffffffffffffffffffffffffffffffff166000908152600260205260409020546c01000000000000000000000000900460ff1690565b60405190151581526020016100f1565b610143610161366004610d15565b6101e6565b61016e610305565b005b60005460405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100f1565b61016e6101a6366004610cf8565b610402565b6101b3610785565b6040516100f19190610d51565b61016e6101ce366004610cf8565b6107f4565b61016e6101e1366004610cf8565b610a04565b336000818152600260205260408120549091906c01000000000000000000000000900460ff16610249576040517fa2c8bfb60000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b600061025c6102578561100c565b610a18565b9050600063a0c6df1560e01b82604051602401610279919061114a565b60408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009094169390931790925291506102fc9060e0870135906102f49060608901908901610cf8565b600084610aab565b95945050505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610386576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610240565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b61040a610af7565b6003546000819003610448576040517f22babb3200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82166000908152600260209081526040918290208251808401909352546bffffffffffffffffffffffff811683526c01000000000000000000000000900460ff1615159082018190526104f4576040517f8c97f12200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84166004820152602401610240565b6000600361050360018561120d565b815481106105135761051361124b565b60009182526020909120015482516003805473ffffffffffffffffffffffffffffffffffffffff9093169350916bffffffffffffffffffffffff90911690811061055f5761055f61124b565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16600361058e60018661120d565b8154811061059e5761059e61124b565b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600383600001516bffffffffffffffffffffffff168154811061060c5761060c61124b565b600091825260208083209190910180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff94851617905584519284168252600290526040902080547fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166bffffffffffffffffffffffff90921691909117905560038054806106b3576106b361127a565b6000828152602080822083017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590920190925573ffffffffffffffffffffffffffffffffffffffff86168083526002909152604080832080547fffffffffffffffffffffffffffffffffffffff000000000000000000000000001690555190917fcf91daec21e3510e2f2aea4b09d08c235d5c6844980be709f282ef591dbf420c91a250505050565b606060038054806020026020016040519081016040528092919081815260200182805480156107ea57602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff1681526001909101906020018083116107bf575b5050505050905090565b6107fc610af7565b73ffffffffffffffffffffffffffffffffffffffff8116610849576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff81166000908152600260209081526040918290208251808401909352546bffffffffffffffffffffffff811683526c01000000000000000000000000900460ff16158015918301919091526108f7576040517f3a4406b500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610240565b60016020808301828152600380546bffffffffffffffffffffffff908116865273ffffffffffffffffffffffffffffffffffffffff871660008181526002909552604080862088518154965115156c01000000000000000000000000027fffffffffffffffffffffffffffffffffffffff0000000000000000000000000090971694169390931794909417909155815494850182559083527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b90930180547fffffffffffffffffffffffff00000000000000000000000000000000000000001684179055517f78f53b26906785548b265fa08f4197f9f3fff73fe0d504d30400aacb527f4ce09190a25050565b610a0c610af7565b610a1581610b7a565b50565b610a4a6040518060a0016040528060008152602001606081526020016060815260200160608152602001606081525090565b6040518060a00160405280836000015181526020018360200151604051602001610a749190610cbc565b604051602081830303815290604052815260200183606001518152602001836080015181526020018360c001518152509050919050565b60005a611388811015610abd57600080fd5b611388810390508560408204820311610ad557600080fd5b50833b610ae157600080fd5b60008083516020850186888af195945050505050565b60005473ffffffffffffffffffffffffffffffffffffffff163314610b78576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610240565b565b3373ffffffffffffffffffffffffffffffffffffffff821603610bf9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610240565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000815180845260005b81811015610c9557602081850181015186830182015201610c79565b81811115610ca7576000602083870101525b50601f01601f19169290920160200192915050565b602081526000610ccf6020830184610c6f565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff81168114610a1557600080fd5b600060208284031215610d0a57600080fd5b8135610ccf81610cd6565b600060208284031215610d2757600080fd5b813567ffffffffffffffff811115610d3e57600080fd5b82016101008185031215610ccf57600080fd5b6020808252825182820181905260009190848201906040850190845b81811015610d9f57835173ffffffffffffffffffffffffffffffffffffffff1683529284019291840191600101610d6d565b50909695505050505050565b8035610db681610cd6565b919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051610100810167ffffffffffffffff81118282101715610e0e57610e0e610dbb565b60405290565b604051601f8201601f1916810167ffffffffffffffff81118282101715610e3d57610e3d610dbb565b604052919050565b600082601f830112610e5657600080fd5b813567ffffffffffffffff811115610e7057610e70610dbb565b610e836020601f19601f84011601610e14565b818152846020838601011115610e9857600080fd5b816020850160208301376000918101602001919091529392505050565b600067ffffffffffffffff821115610ecf57610ecf610dbb565b5060051b60200190565b600082601f830112610eea57600080fd5b81356020610eff610efa83610eb5565b610e14565b82815260059290921b84018101918181019086841115610f1e57600080fd5b8286015b84811015610f42578035610f3581610cd6565b8352918301918301610f22565b509695505050505050565b600082601f830112610f5e57600080fd5b81356020610f6e610efa83610eb5565b82815260059290921b84018101918181019086841115610f8d57600080fd5b8286015b84811015610f42578035610fa481610cd6565b8352918301918301610f91565b600082601f830112610fc257600080fd5b81356020610fd2610efa83610eb5565b82815260059290921b84018101918181019086841115610ff157600080fd5b8286015b84811015610f425780358352918301918301610ff5565b6000610100823603121561101f57600080fd5b611027610dea565b82358152602083013567ffffffffffffffff8082111561104657600080fd5b61105236838701610e45565b602084015261106360408601610dab565b6040840152606085013591508082111561107c57600080fd5b61108836838701610e45565b606084015260808501359150808211156110a157600080fd5b6110ad36838701610ed9565b608084015260a08501359150808211156110c657600080fd5b6110d236838701610f4d565b60a084015260c08501359150808211156110eb57600080fd5b506110f836828601610fb1565b60c08301525060e092830135928101929092525090565b600081518084526020808501945080840160005b8381101561113f57815187529582019590820190600101611123565b509495945050505050565b600060208083528351818401528084015160a0604085015261116f60c0850182610c6f565b90506040850151601f198086840301606087015261118d8383610c6f565b6060880151878203830160808901528051808352908601945060009350908501905b808410156111e557845173ffffffffffffffffffffffffffffffffffffffff1682529385019360019390930192908501906111af565b5060808801519450818782030160a0880152611201818661110f565b98975050505050505050565b600082821015611246577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea164736f6c634300080f000a",
}

var Any2EVMTollOffRampRouterABI = Any2EVMTollOffRampRouterMetaData.ABI

var Any2EVMTollOffRampRouterBin = Any2EVMTollOffRampRouterMetaData.Bin

func DeployAny2EVMTollOffRampRouter(auth *bind.TransactOpts, backend bind.ContractBackend, offRamps []common.Address) (common.Address, *types.Transaction, *Any2EVMTollOffRampRouter, error) {
	parsed, err := Any2EVMTollOffRampRouterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Any2EVMTollOffRampRouterBin), backend, offRamps)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Any2EVMTollOffRampRouter{Any2EVMTollOffRampRouterCaller: Any2EVMTollOffRampRouterCaller{contract: contract}, Any2EVMTollOffRampRouterTransactor: Any2EVMTollOffRampRouterTransactor{contract: contract}, Any2EVMTollOffRampRouterFilterer: Any2EVMTollOffRampRouterFilterer{contract: contract}}, nil
}

type Any2EVMTollOffRampRouter struct {
	address common.Address
	abi     abi.ABI
	Any2EVMTollOffRampRouterCaller
	Any2EVMTollOffRampRouterTransactor
	Any2EVMTollOffRampRouterFilterer
}

type Any2EVMTollOffRampRouterCaller struct {
	contract *bind.BoundContract
}

type Any2EVMTollOffRampRouterTransactor struct {
	contract *bind.BoundContract
}

type Any2EVMTollOffRampRouterFilterer struct {
	contract *bind.BoundContract
}

type Any2EVMTollOffRampRouterSession struct {
	Contract     *Any2EVMTollOffRampRouter
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type Any2EVMTollOffRampRouterCallerSession struct {
	Contract *Any2EVMTollOffRampRouterCaller
	CallOpts bind.CallOpts
}

type Any2EVMTollOffRampRouterTransactorSession struct {
	Contract     *Any2EVMTollOffRampRouterTransactor
	TransactOpts bind.TransactOpts
}

type Any2EVMTollOffRampRouterRaw struct {
	Contract *Any2EVMTollOffRampRouter
}

type Any2EVMTollOffRampRouterCallerRaw struct {
	Contract *Any2EVMTollOffRampRouterCaller
}

type Any2EVMTollOffRampRouterTransactorRaw struct {
	Contract *Any2EVMTollOffRampRouterTransactor
}

func NewAny2EVMTollOffRampRouter(address common.Address, backend bind.ContractBackend) (*Any2EVMTollOffRampRouter, error) {
	abi, err := abi.JSON(strings.NewReader(Any2EVMTollOffRampRouterABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindAny2EVMTollOffRampRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampRouter{address: address, abi: abi, Any2EVMTollOffRampRouterCaller: Any2EVMTollOffRampRouterCaller{contract: contract}, Any2EVMTollOffRampRouterTransactor: Any2EVMTollOffRampRouterTransactor{contract: contract}, Any2EVMTollOffRampRouterFilterer: Any2EVMTollOffRampRouterFilterer{contract: contract}}, nil
}

func NewAny2EVMTollOffRampRouterCaller(address common.Address, caller bind.ContractCaller) (*Any2EVMTollOffRampRouterCaller, error) {
	contract, err := bindAny2EVMTollOffRampRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampRouterCaller{contract: contract}, nil
}

func NewAny2EVMTollOffRampRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*Any2EVMTollOffRampRouterTransactor, error) {
	contract, err := bindAny2EVMTollOffRampRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampRouterTransactor{contract: contract}, nil
}

func NewAny2EVMTollOffRampRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*Any2EVMTollOffRampRouterFilterer, error) {
	contract, err := bindAny2EVMTollOffRampRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampRouterFilterer{contract: contract}, nil
}

func bindAny2EVMTollOffRampRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Any2EVMTollOffRampRouterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Any2EVMTollOffRampRouter.Contract.Any2EVMTollOffRampRouterCaller.contract.Call(opts, result, method, params...)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Any2EVMTollOffRampRouter.Contract.Any2EVMTollOffRampRouterTransactor.contract.Transfer(opts)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Any2EVMTollOffRampRouter.Contract.Any2EVMTollOffRampRouterTransactor.contract.Transact(opts, method, params...)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Any2EVMTollOffRampRouter.Contract.contract.Call(opts, result, method, params...)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Any2EVMTollOffRampRouter.Contract.contract.Transfer(opts)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Any2EVMTollOffRampRouter.Contract.contract.Transact(opts, method, params...)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterCaller) GetOffRamps(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Any2EVMTollOffRampRouter.contract.Call(opts, &out, "getOffRamps")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterSession) GetOffRamps() ([]common.Address, error) {
	return _Any2EVMTollOffRampRouter.Contract.GetOffRamps(&_Any2EVMTollOffRampRouter.CallOpts)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterCallerSession) GetOffRamps() ([]common.Address, error) {
	return _Any2EVMTollOffRampRouter.Contract.GetOffRamps(&_Any2EVMTollOffRampRouter.CallOpts)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterCaller) IsOffRamp(opts *bind.CallOpts, offRamp common.Address) (bool, error) {
	var out []interface{}
	err := _Any2EVMTollOffRampRouter.contract.Call(opts, &out, "isOffRamp", offRamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterSession) IsOffRamp(offRamp common.Address) (bool, error) {
	return _Any2EVMTollOffRampRouter.Contract.IsOffRamp(&_Any2EVMTollOffRampRouter.CallOpts, offRamp)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterCallerSession) IsOffRamp(offRamp common.Address) (bool, error) {
	return _Any2EVMTollOffRampRouter.Contract.IsOffRamp(&_Any2EVMTollOffRampRouter.CallOpts, offRamp)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Any2EVMTollOffRampRouter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterSession) Owner() (common.Address, error) {
	return _Any2EVMTollOffRampRouter.Contract.Owner(&_Any2EVMTollOffRampRouter.CallOpts)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterCallerSession) Owner() (common.Address, error) {
	return _Any2EVMTollOffRampRouter.Contract.Owner(&_Any2EVMTollOffRampRouter.CallOpts)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Any2EVMTollOffRampRouter.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterSession) TypeAndVersion() (string, error) {
	return _Any2EVMTollOffRampRouter.Contract.TypeAndVersion(&_Any2EVMTollOffRampRouter.CallOpts)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterCallerSession) TypeAndVersion() (string, error) {
	return _Any2EVMTollOffRampRouter.Contract.TypeAndVersion(&_Any2EVMTollOffRampRouter.CallOpts)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Any2EVMTollOffRampRouter.contract.Transact(opts, "acceptOwnership")
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterSession) AcceptOwnership() (*types.Transaction, error) {
	return _Any2EVMTollOffRampRouter.Contract.AcceptOwnership(&_Any2EVMTollOffRampRouter.TransactOpts)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Any2EVMTollOffRampRouter.Contract.AcceptOwnership(&_Any2EVMTollOffRampRouter.TransactOpts)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterTransactor) AddOffRamp(opts *bind.TransactOpts, offRamp common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRampRouter.contract.Transact(opts, "addOffRamp", offRamp)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterSession) AddOffRamp(offRamp common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRampRouter.Contract.AddOffRamp(&_Any2EVMTollOffRampRouter.TransactOpts, offRamp)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterTransactorSession) AddOffRamp(offRamp common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRampRouter.Contract.AddOffRamp(&_Any2EVMTollOffRampRouter.TransactOpts, offRamp)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterTransactor) RemoveOffRamp(opts *bind.TransactOpts, offRamp common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRampRouter.contract.Transact(opts, "removeOffRamp", offRamp)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterSession) RemoveOffRamp(offRamp common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRampRouter.Contract.RemoveOffRamp(&_Any2EVMTollOffRampRouter.TransactOpts, offRamp)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterTransactorSession) RemoveOffRamp(offRamp common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRampRouter.Contract.RemoveOffRamp(&_Any2EVMTollOffRampRouter.TransactOpts, offRamp)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterTransactor) RouteMessage(opts *bind.TransactOpts, message CCIPAny2EVMMessageFromSender) (*types.Transaction, error) {
	return _Any2EVMTollOffRampRouter.contract.Transact(opts, "routeMessage", message)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterSession) RouteMessage(message CCIPAny2EVMMessageFromSender) (*types.Transaction, error) {
	return _Any2EVMTollOffRampRouter.Contract.RouteMessage(&_Any2EVMTollOffRampRouter.TransactOpts, message)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterTransactorSession) RouteMessage(message CCIPAny2EVMMessageFromSender) (*types.Transaction, error) {
	return _Any2EVMTollOffRampRouter.Contract.RouteMessage(&_Any2EVMTollOffRampRouter.TransactOpts, message)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRampRouter.contract.Transact(opts, "transferOwnership", to)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRampRouter.Contract.TransferOwnership(&_Any2EVMTollOffRampRouter.TransactOpts, to)
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _Any2EVMTollOffRampRouter.Contract.TransferOwnership(&_Any2EVMTollOffRampRouter.TransactOpts, to)
}

type Any2EVMTollOffRampRouterOffRampAddedIterator struct {
	Event *Any2EVMTollOffRampRouterOffRampAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *Any2EVMTollOffRampRouterOffRampAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Any2EVMTollOffRampRouterOffRampAdded)
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

	select {
	case log := <-it.logs:
		it.Event = new(Any2EVMTollOffRampRouterOffRampAdded)
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

func (it *Any2EVMTollOffRampRouterOffRampAddedIterator) Error() error {
	return it.fail
}

func (it *Any2EVMTollOffRampRouterOffRampAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type Any2EVMTollOffRampRouterOffRampAdded struct {
	OffRamp common.Address
	Raw     types.Log
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterFilterer) FilterOffRampAdded(opts *bind.FilterOpts, offRamp []common.Address) (*Any2EVMTollOffRampRouterOffRampAddedIterator, error) {

	var offRampRule []interface{}
	for _, offRampItem := range offRamp {
		offRampRule = append(offRampRule, offRampItem)
	}

	logs, sub, err := _Any2EVMTollOffRampRouter.contract.FilterLogs(opts, "OffRampAdded", offRampRule)
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampRouterOffRampAddedIterator{contract: _Any2EVMTollOffRampRouter.contract, event: "OffRampAdded", logs: logs, sub: sub}, nil
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterFilterer) WatchOffRampAdded(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampRouterOffRampAdded, offRamp []common.Address) (event.Subscription, error) {

	var offRampRule []interface{}
	for _, offRampItem := range offRamp {
		offRampRule = append(offRampRule, offRampItem)
	}

	logs, sub, err := _Any2EVMTollOffRampRouter.contract.WatchLogs(opts, "OffRampAdded", offRampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(Any2EVMTollOffRampRouterOffRampAdded)
				if err := _Any2EVMTollOffRampRouter.contract.UnpackLog(event, "OffRampAdded", log); err != nil {
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

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterFilterer) ParseOffRampAdded(log types.Log) (*Any2EVMTollOffRampRouterOffRampAdded, error) {
	event := new(Any2EVMTollOffRampRouterOffRampAdded)
	if err := _Any2EVMTollOffRampRouter.contract.UnpackLog(event, "OffRampAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type Any2EVMTollOffRampRouterOffRampRemovedIterator struct {
	Event *Any2EVMTollOffRampRouterOffRampRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *Any2EVMTollOffRampRouterOffRampRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Any2EVMTollOffRampRouterOffRampRemoved)
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

	select {
	case log := <-it.logs:
		it.Event = new(Any2EVMTollOffRampRouterOffRampRemoved)
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

func (it *Any2EVMTollOffRampRouterOffRampRemovedIterator) Error() error {
	return it.fail
}

func (it *Any2EVMTollOffRampRouterOffRampRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type Any2EVMTollOffRampRouterOffRampRemoved struct {
	OffRamp common.Address
	Raw     types.Log
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterFilterer) FilterOffRampRemoved(opts *bind.FilterOpts, offRamp []common.Address) (*Any2EVMTollOffRampRouterOffRampRemovedIterator, error) {

	var offRampRule []interface{}
	for _, offRampItem := range offRamp {
		offRampRule = append(offRampRule, offRampItem)
	}

	logs, sub, err := _Any2EVMTollOffRampRouter.contract.FilterLogs(opts, "OffRampRemoved", offRampRule)
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampRouterOffRampRemovedIterator{contract: _Any2EVMTollOffRampRouter.contract, event: "OffRampRemoved", logs: logs, sub: sub}, nil
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterFilterer) WatchOffRampRemoved(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampRouterOffRampRemoved, offRamp []common.Address) (event.Subscription, error) {

	var offRampRule []interface{}
	for _, offRampItem := range offRamp {
		offRampRule = append(offRampRule, offRampItem)
	}

	logs, sub, err := _Any2EVMTollOffRampRouter.contract.WatchLogs(opts, "OffRampRemoved", offRampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(Any2EVMTollOffRampRouterOffRampRemoved)
				if err := _Any2EVMTollOffRampRouter.contract.UnpackLog(event, "OffRampRemoved", log); err != nil {
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

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterFilterer) ParseOffRampRemoved(log types.Log) (*Any2EVMTollOffRampRouterOffRampRemoved, error) {
	event := new(Any2EVMTollOffRampRouterOffRampRemoved)
	if err := _Any2EVMTollOffRampRouter.contract.UnpackLog(event, "OffRampRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type Any2EVMTollOffRampRouterOwnershipTransferRequestedIterator struct {
	Event *Any2EVMTollOffRampRouterOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *Any2EVMTollOffRampRouterOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Any2EVMTollOffRampRouterOwnershipTransferRequested)
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

	select {
	case log := <-it.logs:
		it.Event = new(Any2EVMTollOffRampRouterOwnershipTransferRequested)
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

func (it *Any2EVMTollOffRampRouterOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *Any2EVMTollOffRampRouterOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type Any2EVMTollOffRampRouterOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Any2EVMTollOffRampRouterOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Any2EVMTollOffRampRouter.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampRouterOwnershipTransferRequestedIterator{contract: _Any2EVMTollOffRampRouter.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampRouterOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Any2EVMTollOffRampRouter.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(Any2EVMTollOffRampRouterOwnershipTransferRequested)
				if err := _Any2EVMTollOffRampRouter.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterFilterer) ParseOwnershipTransferRequested(log types.Log) (*Any2EVMTollOffRampRouterOwnershipTransferRequested, error) {
	event := new(Any2EVMTollOffRampRouterOwnershipTransferRequested)
	if err := _Any2EVMTollOffRampRouter.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type Any2EVMTollOffRampRouterOwnershipTransferredIterator struct {
	Event *Any2EVMTollOffRampRouterOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *Any2EVMTollOffRampRouterOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Any2EVMTollOffRampRouterOwnershipTransferred)
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

	select {
	case log := <-it.logs:
		it.Event = new(Any2EVMTollOffRampRouterOwnershipTransferred)
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

func (it *Any2EVMTollOffRampRouterOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *Any2EVMTollOffRampRouterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type Any2EVMTollOffRampRouterOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Any2EVMTollOffRampRouterOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Any2EVMTollOffRampRouter.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Any2EVMTollOffRampRouterOwnershipTransferredIterator{contract: _Any2EVMTollOffRampRouter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampRouterOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Any2EVMTollOffRampRouter.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(Any2EVMTollOffRampRouterOwnershipTransferred)
				if err := _Any2EVMTollOffRampRouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouterFilterer) ParseOwnershipTransferred(log types.Log) (*Any2EVMTollOffRampRouterOwnershipTransferred, error) {
	event := new(Any2EVMTollOffRampRouterOwnershipTransferred)
	if err := _Any2EVMTollOffRampRouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouter) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _Any2EVMTollOffRampRouter.abi.Events["OffRampAdded"].ID:
		return _Any2EVMTollOffRampRouter.ParseOffRampAdded(log)
	case _Any2EVMTollOffRampRouter.abi.Events["OffRampRemoved"].ID:
		return _Any2EVMTollOffRampRouter.ParseOffRampRemoved(log)
	case _Any2EVMTollOffRampRouter.abi.Events["OwnershipTransferRequested"].ID:
		return _Any2EVMTollOffRampRouter.ParseOwnershipTransferRequested(log)
	case _Any2EVMTollOffRampRouter.abi.Events["OwnershipTransferred"].ID:
		return _Any2EVMTollOffRampRouter.ParseOwnershipTransferred(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (Any2EVMTollOffRampRouterOffRampAdded) Topic() common.Hash {
	return common.HexToHash("0x78f53b26906785548b265fa08f4197f9f3fff73fe0d504d30400aacb527f4ce0")
}

func (Any2EVMTollOffRampRouterOffRampRemoved) Topic() common.Hash {
	return common.HexToHash("0xcf91daec21e3510e2f2aea4b09d08c235d5c6844980be709f282ef591dbf420c")
}

func (Any2EVMTollOffRampRouterOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (Any2EVMTollOffRampRouterOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_Any2EVMTollOffRampRouter *Any2EVMTollOffRampRouter) Address() common.Address {
	return _Any2EVMTollOffRampRouter.address
}

type Any2EVMTollOffRampRouterInterface interface {
	GetOffRamps(opts *bind.CallOpts) ([]common.Address, error)

	IsOffRamp(opts *bind.CallOpts, offRamp common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddOffRamp(opts *bind.TransactOpts, offRamp common.Address) (*types.Transaction, error)

	RemoveOffRamp(opts *bind.TransactOpts, offRamp common.Address) (*types.Transaction, error)

	RouteMessage(opts *bind.TransactOpts, message CCIPAny2EVMMessageFromSender) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterOffRampAdded(opts *bind.FilterOpts, offRamp []common.Address) (*Any2EVMTollOffRampRouterOffRampAddedIterator, error)

	WatchOffRampAdded(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampRouterOffRampAdded, offRamp []common.Address) (event.Subscription, error)

	ParseOffRampAdded(log types.Log) (*Any2EVMTollOffRampRouterOffRampAdded, error)

	FilterOffRampRemoved(opts *bind.FilterOpts, offRamp []common.Address) (*Any2EVMTollOffRampRouterOffRampRemovedIterator, error)

	WatchOffRampRemoved(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampRouterOffRampRemoved, offRamp []common.Address) (event.Subscription, error)

	ParseOffRampRemoved(log types.Log) (*Any2EVMTollOffRampRouterOffRampRemoved, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Any2EVMTollOffRampRouterOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampRouterOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*Any2EVMTollOffRampRouterOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Any2EVMTollOffRampRouterOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Any2EVMTollOffRampRouterOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*Any2EVMTollOffRampRouterOwnershipTransferred, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
