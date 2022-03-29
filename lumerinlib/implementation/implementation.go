// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package implementation

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
)

// ImplementationMetaData contains all meta data concerning the Implementation contract.
var ImplementationMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newCipherText\",\"type\":\"string\"}],\"name\":\"cipherTextUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"contractClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"}],\"name\":\"contractPurchased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"purchaseInfoUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"buyer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractState\",\"outputs\":[{\"internalType\":\"enumImplementation.ContractState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"encryptedPoolData\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"escrow_purchaser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"escrow_seller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPublicVariables\",\"outputs\":[{\"internalType\":\"enumImplementation.ContractState\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_speed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_lmn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_cloneFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"length\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"limit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receivedTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"seller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"closeOutType\",\"type\":\"uint256\"}],\"name\":\"setContractCloseOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_encryptedPoolData\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"}],\"name\":\"setPurchaseContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_newEncryptedPoolData\",\"type\":\"string\"}],\"name\":\"setUpdateMiningInformation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_speed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_closeoutType\",\"type\":\"uint256\"}],\"name\":\"setUpdatePurchaseInformation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"speed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startingBlockTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50600180819055506126a7806100276000396000f3fe608060405234801561001057600080fd5b50600436106101215760003560e01c80638b7e4b13116100ad578063c5095d6811610071578063c5095d68146102c4578063ca3225fa146102e2578063ce0c722a14610308578063ddcb1bf214610326578063f1a6e3011461034257610121565b80638b7e4b13146102305780638e2e6d5d1461024e578063a035b1fe1461026a578063a4d66daf14610288578063c20906ac146102a657610121565b80631f7b6d32116100f45780631f7b6d321461019e5780633a5d4213146101bc5780637150d8ae146101d8578063719e6b5b146101f657806385209ee01461021257610121565b806308551a5314610126578063089aa8a2146101445780630a61e2d91461016257806316713b3714610180575b600080fd5b61012e61035e565b60405161013b9190611ba6565b60405180910390f35b61014c610384565b6040516101599190611ba6565b60405180910390f35b61016a6103aa565b6040516101779190611e5b565b60405180910390f35b6101886103b0565b6040516101959190611e5b565b60405180910390f35b6101a66103b6565b6040516101b39190611e5b565b60405180910390f35b6101d660048036038101906101d19190611848565b6103bc565b005b6101e06105b0565b6040516101ed9190611ba6565b60405180910390f35b610210600480360381019061020b9190611749565b6105d6565b005b61021a61072c565b6040516102279190611bea565b60405180910390f35b61023861073f565b6040516102459190611c99565b60405180910390f35b610268600480360381019061026391906117ee565b6107cd565b005b610272610ad5565b60405161027f9190611e5b565b60405180910390f35b610290610adb565b60405161029d9190611e5b565b60405180910390f35b6102ae610ae1565b6040516102bb9190611e5b565b60405180910390f35b6102cc610ae7565b6040516102d99190611e5b565b60405180910390f35b6102ea610aed565b6040516102ff99989796959493929190611c05565b60405180910390f35b610310610c0b565b60405161031d9190611ba6565b60405180910390f35b610340600480360381019061033b9190611792565b610c31565b005b61035c600480360381019061035791906118fe565b610e59565b005b600d60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60045481565b60055481565b600a5481565b600060019054906101000a900460ff16806103e2575060008054906101000a900460ff16155b610421576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161041890611dbb565b60405180910390fd5b60008060019054906101000a900460ff161590508015610471576001600060016101000a81548160ff02191690831515021790555060016000806101000a81548160ff0219169083151502179055505b88600781905550876008819055508660098190555085600a8190555084600d60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600e60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600f60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600660146101000a81548160ff0219169083600181111561057657610575612117565b5b02179055506105848461105c565b80156105a55760008060016101000a81548160ff0219169083151502179055505b505050505050505050565b600c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610666576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161065d90611cbb565b60405180910390fd5b60018081111561067957610678612117565b5b600660149054906101000a900460ff16600181111561069b5761069a612117565b5b146106db576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106d290611dfb565b60405180910390fd5b80601090805190602001906106f19291906115b5565b507f2301ef7d9f42b857543faf9e285b5807e028d4ae99810ea7fe0aadda3a717e9d816040516107219190611c99565b60405180910390a150565b600660149054906101000a900460ff1681565b6010805461074c90612056565b80601f016020809104026020016040519081016040528092919081815260200182805461077890612056565b80156107c55780601f1061079a576101008083540402835291602001916107c5565b820191906000526020600020905b8154815290600101906020018083116107a857829003601f168201915b505050505081565b600081141561091b57600c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16148061087f5750600f60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b6108be576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108b590611d3b565b60405180910390fd5b60006108c86110a0565b90506108e1816007546108db9190611f73565b826110e1565b6108e96112e7565b7ff5e1a452bb76d7335225182a97ad694be2c7b4b5d75dcffb67ddf15db95f484460405160405180910390a150610ad2565b60018114156109d657600d60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146109b4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109ab90611e1b565b60405180910390fd5b6109d16109bf6110a0565b6007546109cc9190611f73565b61139d565b610ad1565b60028114806109e55750600381145b15610a8c57600a54600b54426109fb9190611f73565b1015610a3c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a3390611cdb565b60405180910390fd5b6003811415610a5357610a5260075460006110e1565b5b610a5b6112e7565b7ff5e1a452bb76d7335225182a97ad694be2c7b4b5d75dcffb67ddf15db95f484460405160405180910390a1610ad0565b60048110610acf576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ac690611d9b565b60405180910390fd5b5b5b5b50565b60075481565b60085481565b60095481565b600b5481565b6000806000806000806000806060600660149054906101000a900460ff16600754600854600954600a54600b54600c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600d60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166010808054610b6e90612056565b80601f0160208091040260200160405190810160405280929190818152602001828054610b9a90612056565b8015610be75780601f10610bbc57610100808354040283529160200191610be7565b820191906000526020600020905b815481529060010190602001808311610bca57829003601f168201915b50505050509050985098509850985098509850985098509850909192939495969798565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006001811115610c4557610c44612117565b5b600660149054906101000a900460ff166001811115610c6757610c66612117565b5b14610ca7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c9e90611d5b565b60405180910390fd5b600e60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610d37576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d2e90611d7b565b60405180910390fd5b8160109080519060200190610d4d9291906115b5565b5080600c60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555042600b819055506001600660146101000a81548160ff02191690836001811115610dbc57610dbb612117565b5b0217905550610e12600d60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600754611527565b3373ffffffffffffffffffffffffffffffffffffffff167f0c00d1d6cea0bd55f7d3b6e92ef60237b117b050185fc2816c708fd45f45e5bb60405160405180910390a25050565b6000600b5442610e699190611f73565b9050600d60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610efb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ef290611ddb565b60405180910390fd5b600180811115610f0e57610f0d612117565b5b600660149054906101000a900460ff166001811115610f3057610f2f612117565b5b14610f70576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f6790611cfb565b60405180910390fd5b6002821480610f7f5750600382145b610fbe576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fb590611d1b565b60405180910390fd5b600a54811015611003576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ffa90611cdb565b60405180910390fd5b61100c826107cd565b85600781905550846008819055508360098190555082600a819055507f03e052767f275c0c51cc93a76255d42498341feb7a5beef7cc11fd57c5b6681860405160405180910390a1505050505050565b80600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600080600b54426110b19190611f73565b9050600a5481600a546110c49190611f73565b6007546110d19190611f19565b6110db9190611ee8565b91505090565b60026001541415611127576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161111e90611e3b565b60405180910390fd5b6002600181905550600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846040518363ffffffff1660e01b81526004016111ae929190611bc1565b602060405180830381600087803b1580156111c857600080fd5b505af11580156111dc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611200919061171c565b50600081146112dc57600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b8152600401611288929190611bc1565b602060405180830381600087803b1580156112a257600080fd5b505af11580156112b6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112da919061171c565b505b600180819055505050565b600d60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600c60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550604051806020016040528060008152506010908051906020019061136f9291906115b5565b506000600660146101000a81548160ff0219169083600181111561139657611395612117565b5b0217905550565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1683600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b815260040161145a9190611ba6565b60206040518083038186803b15801561147257600080fd5b505afa158015611486573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114aa919061181b565b6114b49190611f73565b6040518363ffffffff1660e01b81526004016114d1929190611bc1565b602060405180830381600087803b1580156114eb57600080fd5b505af11580156114ff573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611523919061171c565b5050565b82600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600481905550505050565b8280546115c190612056565b90600052602060002090601f0160209004810192826115e3576000855561162a565b82601f106115fc57805160ff191683800117855561162a565b8280016001018555821561162a579182015b8281111561162957825182559160200191906001019061160e565b5b509050611637919061163b565b5090565b5b8082111561165457600081600090555060010161163c565b5090565b600061166b61166684611e9b565b611e76565b905082815260208101848484011115611687576116866121a9565b5b611692848285612014565b509392505050565b6000813590506116a98161262c565b92915050565b6000815190506116be81612643565b92915050565b600082601f8301126116d9576116d86121a4565b5b81356116e9848260208601611658565b91505092915050565b6000813590506117018161265a565b92915050565b6000815190506117168161265a565b92915050565b600060208284031215611732576117316121b3565b5b6000611740848285016116af565b91505092915050565b60006020828403121561175f5761175e6121b3565b5b600082013567ffffffffffffffff81111561177d5761177c6121ae565b5b611789848285016116c4565b91505092915050565b600080604083850312156117a9576117a86121b3565b5b600083013567ffffffffffffffff8111156117c7576117c66121ae565b5b6117d3858286016116c4565b92505060206117e48582860161169a565b9150509250929050565b600060208284031215611804576118036121b3565b5b6000611812848285016116f2565b91505092915050565b600060208284031215611831576118306121b3565b5b600061183f84828501611707565b91505092915050565b600080600080600080600080610100898b031215611869576118686121b3565b5b60006118778b828c016116f2565b98505060206118888b828c016116f2565b97505060406118998b828c016116f2565b96505060606118aa8b828c016116f2565b95505060806118bb8b828c0161169a565b94505060a06118cc8b828c0161169a565b93505060c06118dd8b828c0161169a565b92505060e06118ee8b828c0161169a565b9150509295985092959890939650565b600080600080600060a0868803121561191a576119196121b3565b5b6000611928888289016116f2565b9550506020611939888289016116f2565b945050604061194a888289016116f2565b935050606061195b888289016116f2565b925050608061196c888289016116f2565b9150509295509295909350565b61198281611fa7565b82525050565b61199181612002565b82525050565b60006119a282611ecc565b6119ac8185611ed7565b93506119bc818560208601612023565b6119c5816121b8565b840191505092915050565b60006119dd604383611ed7565b91506119e8826121c9565b606082019050919050565b6000611a00602a83611ed7565b9150611a0b8261223e565b604082019050919050565b6000611a23602b83611ed7565b9150611a2e8261228d565b604082019050919050565b6000611a46602883611ed7565b9150611a51826122dc565b604082019050919050565b6000611a69603b83611ed7565b9150611a748261232b565b604082019050919050565b6000611a8c602583611ed7565b9150611a978261237a565b604082019050919050565b6000611aaf603a83611ed7565b9150611aba826123c9565b604082019050919050565b6000611ad2602983611ed7565b9150611add82612418565b604082019050919050565b6000611af5602e83611ed7565b9150611b0082612467565b604082019050919050565b6000611b18604383611ed7565b9150611b23826124b6565b606082019050919050565b6000611b3b602883611ed7565b9150611b468261252b565b604082019050919050565b6000611b5e604183611ed7565b9150611b698261257a565b606082019050919050565b6000611b81601f83611ed7565b9150611b8c826125ef565b602082019050919050565b611ba081611ff8565b82525050565b6000602082019050611bbb6000830184611979565b92915050565b6000604082019050611bd66000830185611979565b611be36020830184611b97565b9392505050565b6000602082019050611bff6000830184611988565b92915050565b600061012082019050611c1b600083018c611988565b611c28602083018b611b97565b611c35604083018a611b97565b611c426060830189611b97565b611c4f6080830188611b97565b611c5c60a0830187611b97565b611c6960c0830186611979565b611c7660e0830185611979565b818103610100830152611c898184611997565b90509a9950505050505050505050565b60006020820190508181036000830152611cb38184611997565b905092915050565b60006020820190508181036000830152611cd4816119d0565b9050919050565b60006020820190508181036000830152611cf4816119f3565b9050919050565b60006020820190508181036000830152611d1481611a16565b9050919050565b60006020820190508181036000830152611d3481611a39565b9050919050565b60006020820190508181036000830152611d5481611a5c565b9050919050565b60006020820190508181036000830152611d7481611a7f565b9050919050565b60006020820190508181036000830152611d9481611aa2565b9050919050565b60006020820190508181036000830152611db481611ac5565b9050919050565b60006020820190508181036000830152611dd481611ae8565b9050919050565b60006020820190508181036000830152611df481611b0b565b9050919050565b60006020820190508181036000830152611e1481611b2e565b9050919050565b60006020820190508181036000830152611e3481611b51565b9050919050565b60006020820190508181036000830152611e5481611b74565b9050919050565b6000602082019050611e706000830184611b97565b92915050565b6000611e80611e91565b9050611e8c8282612088565b919050565b6000604051905090565b600067ffffffffffffffff821115611eb657611eb5612175565b5b611ebf826121b8565b9050602081019050919050565b600081519050919050565b600082825260208201905092915050565b6000611ef382611ff8565b9150611efe83611ff8565b925082611f0e57611f0d6120e8565b5b828204905092915050565b6000611f2482611ff8565b9150611f2f83611ff8565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611f6857611f676120b9565b5b828202905092915050565b6000611f7e82611ff8565b9150611f8983611ff8565b925082821015611f9c57611f9b6120b9565b5b828203905092915050565b6000611fb282611fd8565b9050919050565b60008115159050919050565b6000819050611fd382612618565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600061200d82611fc5565b9050919050565b82818337600083830152505050565b60005b83811015612041578082015181840152602081019050612026565b83811115612050576000848401525b50505050565b6000600282049050600182168061206e57607f821691505b6020821081141561208257612081612146565b5b50919050565b612091826121b8565b810181811067ffffffffffffffff821117156120b0576120af612175565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f74686973206163636f756e74206973206e6f7420617574686f72697a6564207460008201527f6f2075706461746520746865206369706865727465787420696e666f726d617460208201527f696f6e0000000000000000000000000000000000000000000000000000000000604082015250565b7f74686520636f6e7472616374206861732079657420746f20626520636172726960008201527f656420746f207465726d00000000000000000000000000000000000000000000602082015250565b7f74686973206973206163636f756e74206973206e6f7420696e2074686520727560008201527f6e6e696e67207374617465000000000000000000000000000000000000000000602082015250565b7f796f752063616e206f6e6c792075736520636c6f73656f7574206f7074696f6e60008201527f732032206f722033000000000000000000000000000000000000000000000000602082015250565b7f74686973206163636f756e74206973206e6f7420617574686f72697a6564207460008201527f6f207472696767657220616e206561726c7920636c6f73656f75740000000000602082015250565b7f636f6e7472616374206973206e6f7420696e20616e20617661696c61626c652060008201527f7374617465000000000000000000000000000000000000000000000000000000602082015250565b7f746869732061646472657373206973206e6f7420617070726f76656420746f2060008201527f63616c6c207468652070757263686173652066756e6374696f6e000000000000602082015250565b7f796f75206d757374206d616b6520612073656c656374696f6e2062657477656560008201527f6e203020616e6420330000000000000000000000000000000000000000000000602082015250565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b7f74686973206973206163636f756e74206973206e6f7420617574686f72697a6560008201527f6420746f207570646174652074686520636f6e747261637420706172616d657460208201527f6572730000000000000000000000000000000000000000000000000000000000604082015250565b7f74686520636f6e7472616374206973206e6f7420696e207468652072756e6e6960008201527f6e67207374617465000000000000000000000000000000000000000000000000602082015250565b7f74686973206163636f756e74206973206e6f7420617574686f72697a6564207460008201527f6f20747269676765722061206d69642d636f6e747261637420636c6f73656f7560208201527f7400000000000000000000000000000000000000000000000000000000000000604082015250565b7f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00600082015250565b6002811061262957612628612117565b5b50565b61263581611fa7565b811461264057600080fd5b50565b61264c81611fb9565b811461265757600080fd5b50565b61266381611ff8565b811461266e57600080fd5b5056fea26469706673582212207bce95be2d6ab00e3674a5952406bf0d8068d97ae278ee826079136d1d3d603464736f6c63430008070033",
}

// ImplementationABI is the input ABI used to generate the binding from.
// Deprecated: Use ImplementationMetaData.ABI instead.
var ImplementationABI = ImplementationMetaData.ABI

// ImplementationBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ImplementationMetaData.Bin instead.
var ImplementationBin = ImplementationMetaData.Bin

// DeployImplementation deploys a new Ethereum contract, binding an instance of Implementation to it.
func DeployImplementation(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Implementation, error) {
	parsed, err := ImplementationMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ImplementationBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Implementation{ImplementationCaller: ImplementationCaller{contract: contract}, ImplementationTransactor: ImplementationTransactor{contract: contract}, ImplementationFilterer: ImplementationFilterer{contract: contract}}, nil
}

// Implementation is an auto generated Go binding around an Ethereum contract.
type Implementation struct {
	ImplementationCaller     // Read-only binding to the contract
	ImplementationTransactor // Write-only binding to the contract
	ImplementationFilterer   // Log filterer for contract events
}

// ImplementationCaller is an auto generated read-only Go binding around an Ethereum contract.
type ImplementationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ImplementationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ImplementationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ImplementationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ImplementationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ImplementationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ImplementationSession struct {
	Contract     *Implementation   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ImplementationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ImplementationCallerSession struct {
	Contract *ImplementationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ImplementationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ImplementationTransactorSession struct {
	Contract     *ImplementationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ImplementationRaw is an auto generated low-level Go binding around an Ethereum contract.
type ImplementationRaw struct {
	Contract *Implementation // Generic contract binding to access the raw methods on
}

// ImplementationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ImplementationCallerRaw struct {
	Contract *ImplementationCaller // Generic read-only contract binding to access the raw methods on
}

// ImplementationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ImplementationTransactorRaw struct {
	Contract *ImplementationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewImplementation creates a new instance of Implementation, bound to a specific deployed contract.
func NewImplementation(address common.Address, backend bind.ContractBackend) (*Implementation, error) {
	contract, err := bindImplementation(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Implementation{ImplementationCaller: ImplementationCaller{contract: contract}, ImplementationTransactor: ImplementationTransactor{contract: contract}, ImplementationFilterer: ImplementationFilterer{contract: contract}}, nil
}

// NewImplementationCaller creates a new read-only instance of Implementation, bound to a specific deployed contract.
func NewImplementationCaller(address common.Address, caller bind.ContractCaller) (*ImplementationCaller, error) {
	contract, err := bindImplementation(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ImplementationCaller{contract: contract}, nil
}

// NewImplementationTransactor creates a new write-only instance of Implementation, bound to a specific deployed contract.
func NewImplementationTransactor(address common.Address, transactor bind.ContractTransactor) (*ImplementationTransactor, error) {
	contract, err := bindImplementation(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ImplementationTransactor{contract: contract}, nil
}

// NewImplementationFilterer creates a new log filterer instance of Implementation, bound to a specific deployed contract.
func NewImplementationFilterer(address common.Address, filterer bind.ContractFilterer) (*ImplementationFilterer, error) {
	contract, err := bindImplementation(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ImplementationFilterer{contract: contract}, nil
}

// bindImplementation binds a generic wrapper to an already deployed contract.
func bindImplementation(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ImplementationABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Implementation *ImplementationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Implementation.Contract.ImplementationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Implementation *ImplementationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Implementation.Contract.ImplementationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Implementation *ImplementationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Implementation.Contract.ImplementationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Implementation *ImplementationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Implementation.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Implementation *ImplementationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Implementation.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Implementation *ImplementationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Implementation.Contract.contract.Transact(opts, method, params...)
}

// Buyer is a free data retrieval call binding the contract method 0x7150d8ae.
//
// Solidity: function buyer() view returns(address)
func (_Implementation *ImplementationCaller) Buyer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "buyer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Buyer is a free data retrieval call binding the contract method 0x7150d8ae.
//
// Solidity: function buyer() view returns(address)
func (_Implementation *ImplementationSession) Buyer() (common.Address, error) {
	return _Implementation.Contract.Buyer(&_Implementation.CallOpts)
}

// Buyer is a free data retrieval call binding the contract method 0x7150d8ae.
//
// Solidity: function buyer() view returns(address)
func (_Implementation *ImplementationCallerSession) Buyer() (common.Address, error) {
	return _Implementation.Contract.Buyer(&_Implementation.CallOpts)
}

// ContractState is a free data retrieval call binding the contract method 0x85209ee0.
//
// Solidity: function contractState() view returns(uint8)
func (_Implementation *ImplementationCaller) ContractState(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "contractState")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ContractState is a free data retrieval call binding the contract method 0x85209ee0.
//
// Solidity: function contractState() view returns(uint8)
func (_Implementation *ImplementationSession) ContractState() (uint8, error) {
	return _Implementation.Contract.ContractState(&_Implementation.CallOpts)
}

// ContractState is a free data retrieval call binding the contract method 0x85209ee0.
//
// Solidity: function contractState() view returns(uint8)
func (_Implementation *ImplementationCallerSession) ContractState() (uint8, error) {
	return _Implementation.Contract.ContractState(&_Implementation.CallOpts)
}

// ContractTotal is a free data retrieval call binding the contract method 0x0a61e2d9.
//
// Solidity: function contractTotal() view returns(uint256)
func (_Implementation *ImplementationCaller) ContractTotal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "contractTotal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ContractTotal is a free data retrieval call binding the contract method 0x0a61e2d9.
//
// Solidity: function contractTotal() view returns(uint256)
func (_Implementation *ImplementationSession) ContractTotal() (*big.Int, error) {
	return _Implementation.Contract.ContractTotal(&_Implementation.CallOpts)
}

// ContractTotal is a free data retrieval call binding the contract method 0x0a61e2d9.
//
// Solidity: function contractTotal() view returns(uint256)
func (_Implementation *ImplementationCallerSession) ContractTotal() (*big.Int, error) {
	return _Implementation.Contract.ContractTotal(&_Implementation.CallOpts)
}

// EncryptedPoolData is a free data retrieval call binding the contract method 0x8b7e4b13.
//
// Solidity: function encryptedPoolData() view returns(string)
func (_Implementation *ImplementationCaller) EncryptedPoolData(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "encryptedPoolData")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// EncryptedPoolData is a free data retrieval call binding the contract method 0x8b7e4b13.
//
// Solidity: function encryptedPoolData() view returns(string)
func (_Implementation *ImplementationSession) EncryptedPoolData() (string, error) {
	return _Implementation.Contract.EncryptedPoolData(&_Implementation.CallOpts)
}

// EncryptedPoolData is a free data retrieval call binding the contract method 0x8b7e4b13.
//
// Solidity: function encryptedPoolData() view returns(string)
func (_Implementation *ImplementationCallerSession) EncryptedPoolData() (string, error) {
	return _Implementation.Contract.EncryptedPoolData(&_Implementation.CallOpts)
}

// EscrowPurchaser is a free data retrieval call binding the contract method 0x089aa8a2.
//
// Solidity: function escrow_purchaser() view returns(address)
func (_Implementation *ImplementationCaller) EscrowPurchaser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "escrow_purchaser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EscrowPurchaser is a free data retrieval call binding the contract method 0x089aa8a2.
//
// Solidity: function escrow_purchaser() view returns(address)
func (_Implementation *ImplementationSession) EscrowPurchaser() (common.Address, error) {
	return _Implementation.Contract.EscrowPurchaser(&_Implementation.CallOpts)
}

// EscrowPurchaser is a free data retrieval call binding the contract method 0x089aa8a2.
//
// Solidity: function escrow_purchaser() view returns(address)
func (_Implementation *ImplementationCallerSession) EscrowPurchaser() (common.Address, error) {
	return _Implementation.Contract.EscrowPurchaser(&_Implementation.CallOpts)
}

// EscrowSeller is a free data retrieval call binding the contract method 0xce0c722a.
//
// Solidity: function escrow_seller() view returns(address)
func (_Implementation *ImplementationCaller) EscrowSeller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "escrow_seller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EscrowSeller is a free data retrieval call binding the contract method 0xce0c722a.
//
// Solidity: function escrow_seller() view returns(address)
func (_Implementation *ImplementationSession) EscrowSeller() (common.Address, error) {
	return _Implementation.Contract.EscrowSeller(&_Implementation.CallOpts)
}

// EscrowSeller is a free data retrieval call binding the contract method 0xce0c722a.
//
// Solidity: function escrow_seller() view returns(address)
func (_Implementation *ImplementationCallerSession) EscrowSeller() (common.Address, error) {
	return _Implementation.Contract.EscrowSeller(&_Implementation.CallOpts)
}

// GetPublicVariables is a free data retrieval call binding the contract method 0xca3225fa.
//
// Solidity: function getPublicVariables() view returns(uint8, uint256, uint256, uint256, uint256, uint256, address, address, string)
func (_Implementation *ImplementationCaller) GetPublicVariables(opts *bind.CallOpts) (uint8, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, common.Address, common.Address, string, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "getPublicVariables")

	if err != nil {
		return *new(uint8), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(common.Address), *new(common.Address), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	out6 := *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	out7 := *abi.ConvertType(out[7], new(common.Address)).(*common.Address)
	out8 := *abi.ConvertType(out[8], new(string)).(*string)

	return out0, out1, out2, out3, out4, out5, out6, out7, out8, err

}

// GetPublicVariables is a free data retrieval call binding the contract method 0xca3225fa.
//
// Solidity: function getPublicVariables() view returns(uint8, uint256, uint256, uint256, uint256, uint256, address, address, string)
func (_Implementation *ImplementationSession) GetPublicVariables() (uint8, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, common.Address, common.Address, string, error) {
	return _Implementation.Contract.GetPublicVariables(&_Implementation.CallOpts)
}

// GetPublicVariables is a free data retrieval call binding the contract method 0xca3225fa.
//
// Solidity: function getPublicVariables() view returns(uint8, uint256, uint256, uint256, uint256, uint256, address, address, string)
func (_Implementation *ImplementationCallerSession) GetPublicVariables() (uint8, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, common.Address, common.Address, string, error) {
	return _Implementation.Contract.GetPublicVariables(&_Implementation.CallOpts)
}

// Length is a free data retrieval call binding the contract method 0x1f7b6d32.
//
// Solidity: function length() view returns(uint256)
func (_Implementation *ImplementationCaller) Length(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "length")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Length is a free data retrieval call binding the contract method 0x1f7b6d32.
//
// Solidity: function length() view returns(uint256)
func (_Implementation *ImplementationSession) Length() (*big.Int, error) {
	return _Implementation.Contract.Length(&_Implementation.CallOpts)
}

// Length is a free data retrieval call binding the contract method 0x1f7b6d32.
//
// Solidity: function length() view returns(uint256)
func (_Implementation *ImplementationCallerSession) Length() (*big.Int, error) {
	return _Implementation.Contract.Length(&_Implementation.CallOpts)
}

// Limit is a free data retrieval call binding the contract method 0xa4d66daf.
//
// Solidity: function limit() view returns(uint256)
func (_Implementation *ImplementationCaller) Limit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "limit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Limit is a free data retrieval call binding the contract method 0xa4d66daf.
//
// Solidity: function limit() view returns(uint256)
func (_Implementation *ImplementationSession) Limit() (*big.Int, error) {
	return _Implementation.Contract.Limit(&_Implementation.CallOpts)
}

// Limit is a free data retrieval call binding the contract method 0xa4d66daf.
//
// Solidity: function limit() view returns(uint256)
func (_Implementation *ImplementationCallerSession) Limit() (*big.Int, error) {
	return _Implementation.Contract.Limit(&_Implementation.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_Implementation *ImplementationCaller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_Implementation *ImplementationSession) Price() (*big.Int, error) {
	return _Implementation.Contract.Price(&_Implementation.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_Implementation *ImplementationCallerSession) Price() (*big.Int, error) {
	return _Implementation.Contract.Price(&_Implementation.CallOpts)
}

// ReceivedTotal is a free data retrieval call binding the contract method 0x16713b37.
//
// Solidity: function receivedTotal() view returns(uint256)
func (_Implementation *ImplementationCaller) ReceivedTotal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "receivedTotal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReceivedTotal is a free data retrieval call binding the contract method 0x16713b37.
//
// Solidity: function receivedTotal() view returns(uint256)
func (_Implementation *ImplementationSession) ReceivedTotal() (*big.Int, error) {
	return _Implementation.Contract.ReceivedTotal(&_Implementation.CallOpts)
}

// ReceivedTotal is a free data retrieval call binding the contract method 0x16713b37.
//
// Solidity: function receivedTotal() view returns(uint256)
func (_Implementation *ImplementationCallerSession) ReceivedTotal() (*big.Int, error) {
	return _Implementation.Contract.ReceivedTotal(&_Implementation.CallOpts)
}

// Seller is a free data retrieval call binding the contract method 0x08551a53.
//
// Solidity: function seller() view returns(address)
func (_Implementation *ImplementationCaller) Seller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "seller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Seller is a free data retrieval call binding the contract method 0x08551a53.
//
// Solidity: function seller() view returns(address)
func (_Implementation *ImplementationSession) Seller() (common.Address, error) {
	return _Implementation.Contract.Seller(&_Implementation.CallOpts)
}

// Seller is a free data retrieval call binding the contract method 0x08551a53.
//
// Solidity: function seller() view returns(address)
func (_Implementation *ImplementationCallerSession) Seller() (common.Address, error) {
	return _Implementation.Contract.Seller(&_Implementation.CallOpts)
}

// Speed is a free data retrieval call binding the contract method 0xc20906ac.
//
// Solidity: function speed() view returns(uint256)
func (_Implementation *ImplementationCaller) Speed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "speed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Speed is a free data retrieval call binding the contract method 0xc20906ac.
//
// Solidity: function speed() view returns(uint256)
func (_Implementation *ImplementationSession) Speed() (*big.Int, error) {
	return _Implementation.Contract.Speed(&_Implementation.CallOpts)
}

// Speed is a free data retrieval call binding the contract method 0xc20906ac.
//
// Solidity: function speed() view returns(uint256)
func (_Implementation *ImplementationCallerSession) Speed() (*big.Int, error) {
	return _Implementation.Contract.Speed(&_Implementation.CallOpts)
}

// StartingBlockTimestamp is a free data retrieval call binding the contract method 0xc5095d68.
//
// Solidity: function startingBlockTimestamp() view returns(uint256)
func (_Implementation *ImplementationCaller) StartingBlockTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "startingBlockTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartingBlockTimestamp is a free data retrieval call binding the contract method 0xc5095d68.
//
// Solidity: function startingBlockTimestamp() view returns(uint256)
func (_Implementation *ImplementationSession) StartingBlockTimestamp() (*big.Int, error) {
	return _Implementation.Contract.StartingBlockTimestamp(&_Implementation.CallOpts)
}

// StartingBlockTimestamp is a free data retrieval call binding the contract method 0xc5095d68.
//
// Solidity: function startingBlockTimestamp() view returns(uint256)
func (_Implementation *ImplementationCallerSession) StartingBlockTimestamp() (*big.Int, error) {
	return _Implementation.Contract.StartingBlockTimestamp(&_Implementation.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x3a5d4213.
//
// Solidity: function initialize(uint256 _price, uint256 _limit, uint256 _speed, uint256 _length, address _seller, address _lmn, address _cloneFactory, address _validator) returns()
func (_Implementation *ImplementationTransactor) Initialize(opts *bind.TransactOpts, _price *big.Int, _limit *big.Int, _speed *big.Int, _length *big.Int, _seller common.Address, _lmn common.Address, _cloneFactory common.Address, _validator common.Address) (*types.Transaction, error) {
	return _Implementation.contract.Transact(opts, "initialize", _price, _limit, _speed, _length, _seller, _lmn, _cloneFactory, _validator)
}

// Initialize is a paid mutator transaction binding the contract method 0x3a5d4213.
//
// Solidity: function initialize(uint256 _price, uint256 _limit, uint256 _speed, uint256 _length, address _seller, address _lmn, address _cloneFactory, address _validator) returns()
func (_Implementation *ImplementationSession) Initialize(_price *big.Int, _limit *big.Int, _speed *big.Int, _length *big.Int, _seller common.Address, _lmn common.Address, _cloneFactory common.Address, _validator common.Address) (*types.Transaction, error) {
	return _Implementation.Contract.Initialize(&_Implementation.TransactOpts, _price, _limit, _speed, _length, _seller, _lmn, _cloneFactory, _validator)
}

// Initialize is a paid mutator transaction binding the contract method 0x3a5d4213.
//
// Solidity: function initialize(uint256 _price, uint256 _limit, uint256 _speed, uint256 _length, address _seller, address _lmn, address _cloneFactory, address _validator) returns()
func (_Implementation *ImplementationTransactorSession) Initialize(_price *big.Int, _limit *big.Int, _speed *big.Int, _length *big.Int, _seller common.Address, _lmn common.Address, _cloneFactory common.Address, _validator common.Address) (*types.Transaction, error) {
	return _Implementation.Contract.Initialize(&_Implementation.TransactOpts, _price, _limit, _speed, _length, _seller, _lmn, _cloneFactory, _validator)
}

// SetContractCloseOut is a paid mutator transaction binding the contract method 0x8e2e6d5d.
//
// Solidity: function setContractCloseOut(uint256 closeOutType) returns()
func (_Implementation *ImplementationTransactor) SetContractCloseOut(opts *bind.TransactOpts, closeOutType *big.Int) (*types.Transaction, error) {
	return _Implementation.contract.Transact(opts, "setContractCloseOut", closeOutType)
}

// SetContractCloseOut is a paid mutator transaction binding the contract method 0x8e2e6d5d.
//
// Solidity: function setContractCloseOut(uint256 closeOutType) returns()
func (_Implementation *ImplementationSession) SetContractCloseOut(closeOutType *big.Int) (*types.Transaction, error) {
	return _Implementation.Contract.SetContractCloseOut(&_Implementation.TransactOpts, closeOutType)
}

// SetContractCloseOut is a paid mutator transaction binding the contract method 0x8e2e6d5d.
//
// Solidity: function setContractCloseOut(uint256 closeOutType) returns()
func (_Implementation *ImplementationTransactorSession) SetContractCloseOut(closeOutType *big.Int) (*types.Transaction, error) {
	return _Implementation.Contract.SetContractCloseOut(&_Implementation.TransactOpts, closeOutType)
}

// SetPurchaseContract is a paid mutator transaction binding the contract method 0xddcb1bf2.
//
// Solidity: function setPurchaseContract(string _encryptedPoolData, address _buyer) returns()
func (_Implementation *ImplementationTransactor) SetPurchaseContract(opts *bind.TransactOpts, _encryptedPoolData string, _buyer common.Address) (*types.Transaction, error) {
	return _Implementation.contract.Transact(opts, "setPurchaseContract", _encryptedPoolData, _buyer)
}

// SetPurchaseContract is a paid mutator transaction binding the contract method 0xddcb1bf2.
//
// Solidity: function setPurchaseContract(string _encryptedPoolData, address _buyer) returns()
func (_Implementation *ImplementationSession) SetPurchaseContract(_encryptedPoolData string, _buyer common.Address) (*types.Transaction, error) {
	return _Implementation.Contract.SetPurchaseContract(&_Implementation.TransactOpts, _encryptedPoolData, _buyer)
}

// SetPurchaseContract is a paid mutator transaction binding the contract method 0xddcb1bf2.
//
// Solidity: function setPurchaseContract(string _encryptedPoolData, address _buyer) returns()
func (_Implementation *ImplementationTransactorSession) SetPurchaseContract(_encryptedPoolData string, _buyer common.Address) (*types.Transaction, error) {
	return _Implementation.Contract.SetPurchaseContract(&_Implementation.TransactOpts, _encryptedPoolData, _buyer)
}

// SetUpdateMiningInformation is a paid mutator transaction binding the contract method 0x719e6b5b.
//
// Solidity: function setUpdateMiningInformation(string _newEncryptedPoolData) returns()
func (_Implementation *ImplementationTransactor) SetUpdateMiningInformation(opts *bind.TransactOpts, _newEncryptedPoolData string) (*types.Transaction, error) {
	return _Implementation.contract.Transact(opts, "setUpdateMiningInformation", _newEncryptedPoolData)
}

// SetUpdateMiningInformation is a paid mutator transaction binding the contract method 0x719e6b5b.
//
// Solidity: function setUpdateMiningInformation(string _newEncryptedPoolData) returns()
func (_Implementation *ImplementationSession) SetUpdateMiningInformation(_newEncryptedPoolData string) (*types.Transaction, error) {
	return _Implementation.Contract.SetUpdateMiningInformation(&_Implementation.TransactOpts, _newEncryptedPoolData)
}

// SetUpdateMiningInformation is a paid mutator transaction binding the contract method 0x719e6b5b.
//
// Solidity: function setUpdateMiningInformation(string _newEncryptedPoolData) returns()
func (_Implementation *ImplementationTransactorSession) SetUpdateMiningInformation(_newEncryptedPoolData string) (*types.Transaction, error) {
	return _Implementation.Contract.SetUpdateMiningInformation(&_Implementation.TransactOpts, _newEncryptedPoolData)
}

// SetUpdatePurchaseInformation is a paid mutator transaction binding the contract method 0xf1a6e301.
//
// Solidity: function setUpdatePurchaseInformation(uint256 _price, uint256 _limit, uint256 _speed, uint256 _length, uint256 _closeoutType) returns()
func (_Implementation *ImplementationTransactor) SetUpdatePurchaseInformation(opts *bind.TransactOpts, _price *big.Int, _limit *big.Int, _speed *big.Int, _length *big.Int, _closeoutType *big.Int) (*types.Transaction, error) {
	return _Implementation.contract.Transact(opts, "setUpdatePurchaseInformation", _price, _limit, _speed, _length, _closeoutType)
}

// SetUpdatePurchaseInformation is a paid mutator transaction binding the contract method 0xf1a6e301.
//
// Solidity: function setUpdatePurchaseInformation(uint256 _price, uint256 _limit, uint256 _speed, uint256 _length, uint256 _closeoutType) returns()
func (_Implementation *ImplementationSession) SetUpdatePurchaseInformation(_price *big.Int, _limit *big.Int, _speed *big.Int, _length *big.Int, _closeoutType *big.Int) (*types.Transaction, error) {
	return _Implementation.Contract.SetUpdatePurchaseInformation(&_Implementation.TransactOpts, _price, _limit, _speed, _length, _closeoutType)
}

// SetUpdatePurchaseInformation is a paid mutator transaction binding the contract method 0xf1a6e301.
//
// Solidity: function setUpdatePurchaseInformation(uint256 _price, uint256 _limit, uint256 _speed, uint256 _length, uint256 _closeoutType) returns()
func (_Implementation *ImplementationTransactorSession) SetUpdatePurchaseInformation(_price *big.Int, _limit *big.Int, _speed *big.Int, _length *big.Int, _closeoutType *big.Int) (*types.Transaction, error) {
	return _Implementation.Contract.SetUpdatePurchaseInformation(&_Implementation.TransactOpts, _price, _limit, _speed, _length, _closeoutType)
}

// ImplementationCipherTextUpdatedIterator is returned from FilterCipherTextUpdated and is used to iterate over the raw logs and unpacked data for CipherTextUpdated events raised by the Implementation contract.
type ImplementationCipherTextUpdatedIterator struct {
	Event *ImplementationCipherTextUpdated // Event containing the contract specifics and raw log

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
func (it *ImplementationCipherTextUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ImplementationCipherTextUpdated)
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
		it.Event = new(ImplementationCipherTextUpdated)
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
func (it *ImplementationCipherTextUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ImplementationCipherTextUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ImplementationCipherTextUpdated represents a CipherTextUpdated event raised by the Implementation contract.
type ImplementationCipherTextUpdated struct {
	NewCipherText string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCipherTextUpdated is a free log retrieval operation binding the contract event 0x2301ef7d9f42b857543faf9e285b5807e028d4ae99810ea7fe0aadda3a717e9d.
//
// Solidity: event cipherTextUpdated(string newCipherText)
func (_Implementation *ImplementationFilterer) FilterCipherTextUpdated(opts *bind.FilterOpts) (*ImplementationCipherTextUpdatedIterator, error) {

	logs, sub, err := _Implementation.contract.FilterLogs(opts, "cipherTextUpdated")
	if err != nil {
		return nil, err
	}
	return &ImplementationCipherTextUpdatedIterator{contract: _Implementation.contract, event: "cipherTextUpdated", logs: logs, sub: sub}, nil
}

// WatchCipherTextUpdated is a free log subscription operation binding the contract event 0x2301ef7d9f42b857543faf9e285b5807e028d4ae99810ea7fe0aadda3a717e9d.
//
// Solidity: event cipherTextUpdated(string newCipherText)
func (_Implementation *ImplementationFilterer) WatchCipherTextUpdated(opts *bind.WatchOpts, sink chan<- *ImplementationCipherTextUpdated) (event.Subscription, error) {

	logs, sub, err := _Implementation.contract.WatchLogs(opts, "cipherTextUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ImplementationCipherTextUpdated)
				if err := _Implementation.contract.UnpackLog(event, "cipherTextUpdated", log); err != nil {
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

// ParseCipherTextUpdated is a log parse operation binding the contract event 0x2301ef7d9f42b857543faf9e285b5807e028d4ae99810ea7fe0aadda3a717e9d.
//
// Solidity: event cipherTextUpdated(string newCipherText)
func (_Implementation *ImplementationFilterer) ParseCipherTextUpdated(log types.Log) (*ImplementationCipherTextUpdated, error) {
	event := new(ImplementationCipherTextUpdated)
	if err := _Implementation.contract.UnpackLog(event, "cipherTextUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ImplementationContractClosedIterator is returned from FilterContractClosed and is used to iterate over the raw logs and unpacked data for ContractClosed events raised by the Implementation contract.
type ImplementationContractClosedIterator struct {
	Event *ImplementationContractClosed // Event containing the contract specifics and raw log

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
func (it *ImplementationContractClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ImplementationContractClosed)
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
		it.Event = new(ImplementationContractClosed)
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
func (it *ImplementationContractClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ImplementationContractClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ImplementationContractClosed represents a ContractClosed event raised by the Implementation contract.
type ImplementationContractClosed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterContractClosed is a free log retrieval operation binding the contract event 0xf5e1a452bb76d7335225182a97ad694be2c7b4b5d75dcffb67ddf15db95f4844.
//
// Solidity: event contractClosed()
func (_Implementation *ImplementationFilterer) FilterContractClosed(opts *bind.FilterOpts) (*ImplementationContractClosedIterator, error) {

	logs, sub, err := _Implementation.contract.FilterLogs(opts, "contractClosed")
	if err != nil {
		return nil, err
	}
	return &ImplementationContractClosedIterator{contract: _Implementation.contract, event: "contractClosed", logs: logs, sub: sub}, nil
}

// WatchContractClosed is a free log subscription operation binding the contract event 0xf5e1a452bb76d7335225182a97ad694be2c7b4b5d75dcffb67ddf15db95f4844.
//
// Solidity: event contractClosed()
func (_Implementation *ImplementationFilterer) WatchContractClosed(opts *bind.WatchOpts, sink chan<- *ImplementationContractClosed) (event.Subscription, error) {

	logs, sub, err := _Implementation.contract.WatchLogs(opts, "contractClosed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ImplementationContractClosed)
				if err := _Implementation.contract.UnpackLog(event, "contractClosed", log); err != nil {
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

// ParseContractClosed is a log parse operation binding the contract event 0xf5e1a452bb76d7335225182a97ad694be2c7b4b5d75dcffb67ddf15db95f4844.
//
// Solidity: event contractClosed()
func (_Implementation *ImplementationFilterer) ParseContractClosed(log types.Log) (*ImplementationContractClosed, error) {
	event := new(ImplementationContractClosed)
	if err := _Implementation.contract.UnpackLog(event, "contractClosed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ImplementationContractPurchasedIterator is returned from FilterContractPurchased and is used to iterate over the raw logs and unpacked data for ContractPurchased events raised by the Implementation contract.
type ImplementationContractPurchasedIterator struct {
	Event *ImplementationContractPurchased // Event containing the contract specifics and raw log

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
func (it *ImplementationContractPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ImplementationContractPurchased)
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
		it.Event = new(ImplementationContractPurchased)
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
func (it *ImplementationContractPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ImplementationContractPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ImplementationContractPurchased represents a ContractPurchased event raised by the Implementation contract.
type ImplementationContractPurchased struct {
	Buyer common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterContractPurchased is a free log retrieval operation binding the contract event 0x0c00d1d6cea0bd55f7d3b6e92ef60237b117b050185fc2816c708fd45f45e5bb.
//
// Solidity: event contractPurchased(address indexed _buyer)
func (_Implementation *ImplementationFilterer) FilterContractPurchased(opts *bind.FilterOpts, _buyer []common.Address) (*ImplementationContractPurchasedIterator, error) {

	var _buyerRule []interface{}
	for _, _buyerItem := range _buyer {
		_buyerRule = append(_buyerRule, _buyerItem)
	}

	logs, sub, err := _Implementation.contract.FilterLogs(opts, "contractPurchased", _buyerRule)
	if err != nil {
		return nil, err
	}
	return &ImplementationContractPurchasedIterator{contract: _Implementation.contract, event: "contractPurchased", logs: logs, sub: sub}, nil
}

// WatchContractPurchased is a free log subscription operation binding the contract event 0x0c00d1d6cea0bd55f7d3b6e92ef60237b117b050185fc2816c708fd45f45e5bb.
//
// Solidity: event contractPurchased(address indexed _buyer)
func (_Implementation *ImplementationFilterer) WatchContractPurchased(opts *bind.WatchOpts, sink chan<- *ImplementationContractPurchased, _buyer []common.Address) (event.Subscription, error) {

	var _buyerRule []interface{}
	for _, _buyerItem := range _buyer {
		_buyerRule = append(_buyerRule, _buyerItem)
	}

	logs, sub, err := _Implementation.contract.WatchLogs(opts, "contractPurchased", _buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ImplementationContractPurchased)
				if err := _Implementation.contract.UnpackLog(event, "contractPurchased", log); err != nil {
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

// ParseContractPurchased is a log parse operation binding the contract event 0x0c00d1d6cea0bd55f7d3b6e92ef60237b117b050185fc2816c708fd45f45e5bb.
//
// Solidity: event contractPurchased(address indexed _buyer)
func (_Implementation *ImplementationFilterer) ParseContractPurchased(log types.Log) (*ImplementationContractPurchased, error) {
	event := new(ImplementationContractPurchased)
	if err := _Implementation.contract.UnpackLog(event, "contractPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ImplementationPurchaseInfoUpdatedIterator is returned from FilterPurchaseInfoUpdated and is used to iterate over the raw logs and unpacked data for PurchaseInfoUpdated events raised by the Implementation contract.
type ImplementationPurchaseInfoUpdatedIterator struct {
	Event *ImplementationPurchaseInfoUpdated // Event containing the contract specifics and raw log

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
func (it *ImplementationPurchaseInfoUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ImplementationPurchaseInfoUpdated)
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
		it.Event = new(ImplementationPurchaseInfoUpdated)
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
func (it *ImplementationPurchaseInfoUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ImplementationPurchaseInfoUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ImplementationPurchaseInfoUpdated represents a PurchaseInfoUpdated event raised by the Implementation contract.
type ImplementationPurchaseInfoUpdated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPurchaseInfoUpdated is a free log retrieval operation binding the contract event 0x03e052767f275c0c51cc93a76255d42498341feb7a5beef7cc11fd57c5b66818.
//
// Solidity: event purchaseInfoUpdated()
func (_Implementation *ImplementationFilterer) FilterPurchaseInfoUpdated(opts *bind.FilterOpts) (*ImplementationPurchaseInfoUpdatedIterator, error) {

	logs, sub, err := _Implementation.contract.FilterLogs(opts, "purchaseInfoUpdated")
	if err != nil {
		return nil, err
	}
	return &ImplementationPurchaseInfoUpdatedIterator{contract: _Implementation.contract, event: "purchaseInfoUpdated", logs: logs, sub: sub}, nil
}

// WatchPurchaseInfoUpdated is a free log subscription operation binding the contract event 0x03e052767f275c0c51cc93a76255d42498341feb7a5beef7cc11fd57c5b66818.
//
// Solidity: event purchaseInfoUpdated()
func (_Implementation *ImplementationFilterer) WatchPurchaseInfoUpdated(opts *bind.WatchOpts, sink chan<- *ImplementationPurchaseInfoUpdated) (event.Subscription, error) {

	logs, sub, err := _Implementation.contract.WatchLogs(opts, "purchaseInfoUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ImplementationPurchaseInfoUpdated)
				if err := _Implementation.contract.UnpackLog(event, "purchaseInfoUpdated", log); err != nil {
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

// ParsePurchaseInfoUpdated is a log parse operation binding the contract event 0x03e052767f275c0c51cc93a76255d42498341feb7a5beef7cc11fd57c5b66818.
//
// Solidity: event purchaseInfoUpdated()
func (_Implementation *ImplementationFilterer) ParsePurchaseInfoUpdated(log types.Log) (*ImplementationPurchaseInfoUpdated, error) {
	event := new(ImplementationPurchaseInfoUpdated)
	if err := _Implementation.contract.UnpackLog(event, "purchaseInfoUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
