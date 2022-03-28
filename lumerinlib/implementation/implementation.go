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
	Bin: "0x608060405234801561001057600080fd5b50612619806100206000396000f3fe608060405234801561001057600080fd5b50600436106101215760003560e01c80638b7e4b13116100ad578063c5095d6811610071578063c5095d68146102c4578063ca3225fa146102e2578063ce0c722a14610308578063ddcb1bf214610326578063f1a6e3011461034257610121565b80638b7e4b13146102305780638e2e6d5d1461024e578063a035b1fe1461026a578063a4d66daf14610288578063c20906ac146102a657610121565b80631f7b6d32116100f45780631f7b6d321461019e5780633a5d4213146101bc5780637150d8ae146101d8578063719e6b5b146101f657806385209ee01461021257610121565b806308551a5314610126578063089aa8a2146101445780630a61e2d91461016257806316713b3714610180575b600080fd5b61012e61035e565b60405161013b9190611b61565b60405180910390f35b61014c610384565b6040516101599190611b61565b60405180910390f35b61016a6103aa565b6040516101779190611df6565b60405180910390f35b6101886103b0565b6040516101959190611df6565b60405180910390f35b6101a66103b6565b6040516101b39190611df6565b60405180910390f35b6101d660048036038101906101d19190611826565b6103bc565b005b6101e06105b0565b6040516101ed9190611b61565b60405180910390f35b610210600480360381019061020b9190611727565b6105d6565b005b61021a61072c565b6040516102279190611ba5565b60405180910390f35b61023861073f565b6040516102459190611c54565b60405180910390f35b610268600480360381019061026391906117cc565b6107cd565b005b610272610b47565b60405161027f9190611df6565b60405180910390f35b610290610b4d565b60405161029d9190611df6565b60405180910390f35b6102ae610b53565b6040516102bb9190611df6565b60405180910390f35b6102cc610b59565b6040516102d99190611df6565b60405180910390f35b6102ea610b5f565b6040516102ff99989796959493929190611bc0565b60405180910390f35b610310610c7d565b60405161031d9190611b61565b60405180910390f35b610340600480360381019061033b9190611770565b610ca3565b005b61035c600480360381019061035791906118dc565b610ecb565b005b600c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60035481565b60045481565b60095481565b600060019054906101000a900460ff16806103e2575060008054906101000a900460ff16155b610421576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161041890611d76565b60405180910390fd5b60008060019054906101000a900460ff161590508015610471576001600060016101000a81548160ff02191690831515021790555060016000806101000a81548160ff0219169083151502179055505b8860068190555087600781905550866008819055508560098190555084600c60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600d60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600e60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600560146101000a81548160ff02191690836001811115610576576105756120b2565b5b021790555061058484611076565b80156105a55760008060016101000a81548160ff0219169083151502179055505b505050505050505050565b600b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610666576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161065d90611c76565b60405180910390fd5b600180811115610679576106786120b2565b5b600560149054906101000a900460ff16600181111561069b5761069a6120b2565b5b146106db576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106d290611db6565b60405180910390fd5b80600f90805190602001906106f1929190611593565b507f2301ef7d9f42b857543faf9e285b5807e028d4ae99810ea7fe0aadda3a717e9d816040516107219190611c54565b60405180910390a150565b600560149054906101000a900460ff1681565b600f805461074c90611ff1565b80601f016020809104026020016040519081016040528092919081815260200182805461077890611ff1565b80156107c55780601f1061079a576101008083540402835291602001916107c5565b820191906000526020600020905b8154815290600101906020018083116107a857829003601f168201915b505050505081565b600081141561095457600b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16148061087f5750600e60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b6108be576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108b590611cf6565b60405180910390fd5b6000600a54426108ce9190611f0e565b90506000600954826009546108e39190611f0e565b6006546108f09190611eb4565b6108fa9190611e83565b905060008160065461090c9190611f0e565b9050610918818361111d565b6109206112c5565b7ff5e1a452bb76d7335225182a97ad694be2c7b4b5d75dcffb67ddf15db95f484460405160405180910390a1505050610b44565b6001811415610a3b576000600a544261096d9190611f0e565b90506000600954826009546109829190611f0e565b60065461098f9190611eb4565b6109999190611e83565b9050600c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a2b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a2290611dd6565b60405180910390fd5b610a348161137b565b5050610b43565b6002811480610a4a5750600381145b15610afe576000600a5442610a5f9190611f0e565b9050600954811015610aa6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a9d90611c96565b60405180910390fd5b6003821415610ac4576000806006549050610ac1818361111d565b50505b610acc6112c5565b7ff5e1a452bb76d7335225182a97ad694be2c7b4b5d75dcffb67ddf15db95f484460405160405180910390a150610b42565b60048110610b41576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b3890611d56565b60405180910390fd5b5b5b5b50565b60065481565b60075481565b60085481565b600a5481565b6000806000806000806000806060600560149054906101000a900460ff16600654600754600854600954600a54600b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600f808054610be090611ff1565b80601f0160208091040260200160405190810160405280929190818152602001828054610c0c90611ff1565b8015610c595780601f10610c2e57610100808354040283529160200191610c59565b820191906000526020600020905b815481529060010190602001808311610c3c57829003601f168201915b50505050509050985098509850985098509850985098509850909192939495969798565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006001811115610cb757610cb66120b2565b5b600560149054906101000a900460ff166001811115610cd957610cd86120b2565b5b14610d19576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d1090611d16565b60405180910390fd5b600d60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610da9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610da090611d36565b60405180910390fd5b81600f9080519060200190610dbf929190611593565b5080600b60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555042600a819055506001600560146101000a81548160ff02191690836001811115610e2e57610e2d6120b2565b5b0217905550610e84600c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600654611505565b3373ffffffffffffffffffffffffffffffffffffffff167f0c00d1d6cea0bd55f7d3b6e92ef60237b117b050185fc2816c708fd45f45e5bb60405160405180910390a25050565b600c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610f5b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f5290611d96565b60405180910390fd5b600180811115610f6e57610f6d6120b2565b5b600560149054906101000a900460ff166001811115610f9057610f8f6120b2565b5b14610fd0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fc790611cb6565b60405180910390fd5b6002811480610fdf5750600381145b61101e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161101590611cd6565b60405180910390fd5b84600681905550836007819055508260088190555081600981905550611043816107cd565b7f03e052767f275c0c51cc93a76255d42498341feb7a5beef7cc11fd57c5b6681860405160405180910390a15050505050565b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846040518363ffffffff1660e01b815260040161119c929190611b7c565b602060405180830381600087803b1580156111b657600080fd5b505af11580156111ca573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111ee91906116fa565b50600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b815260040161126e929190611b7c565b602060405180830381600087803b15801561128857600080fd5b505af115801561129c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112c091906116fa565b505050565b600c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600b60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060405180602001604052806000815250600f908051906020019061134d929190611593565b506000600560146101000a81548160ff02191690836001811115611374576113736120b2565b5b0217905550565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1683600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016114389190611b61565b60206040518083038186803b15801561145057600080fd5b505afa158015611464573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061148891906117f9565b6114929190611f0e565b6040518363ffffffff1660e01b81526004016114af929190611b7c565b602060405180830381600087803b1580156114c957600080fd5b505af11580156114dd573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061150191906116fa565b5050565b82600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600060026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600381905550505050565b82805461159f90611ff1565b90600052602060002090601f0160209004810192826115c15760008555611608565b82601f106115da57805160ff1916838001178555611608565b82800160010185558215611608579182015b828111156116075782518255916020019190600101906115ec565b5b5090506116159190611619565b5090565b5b8082111561163257600081600090555060010161161a565b5090565b600061164961164484611e36565b611e11565b90508281526020810184848401111561166557611664612144565b5b611670848285611faf565b509392505050565b6000813590506116878161259e565b92915050565b60008151905061169c816125b5565b92915050565b600082601f8301126116b7576116b661213f565b5b81356116c7848260208601611636565b91505092915050565b6000813590506116df816125cc565b92915050565b6000815190506116f4816125cc565b92915050565b6000602082840312156117105761170f61214e565b5b600061171e8482850161168d565b91505092915050565b60006020828403121561173d5761173c61214e565b5b600082013567ffffffffffffffff81111561175b5761175a612149565b5b611767848285016116a2565b91505092915050565b600080604083850312156117875761178661214e565b5b600083013567ffffffffffffffff8111156117a5576117a4612149565b5b6117b1858286016116a2565b92505060206117c285828601611678565b9150509250929050565b6000602082840312156117e2576117e161214e565b5b60006117f0848285016116d0565b91505092915050565b60006020828403121561180f5761180e61214e565b5b600061181d848285016116e5565b91505092915050565b600080600080600080600080610100898b0312156118475761184661214e565b5b60006118558b828c016116d0565b98505060206118668b828c016116d0565b97505060406118778b828c016116d0565b96505060606118888b828c016116d0565b95505060806118998b828c01611678565b94505060a06118aa8b828c01611678565b93505060c06118bb8b828c01611678565b92505060e06118cc8b828c01611678565b9150509295985092959890939650565b600080600080600060a086880312156118f8576118f761214e565b5b6000611906888289016116d0565b9550506020611917888289016116d0565b9450506040611928888289016116d0565b9350506060611939888289016116d0565b925050608061194a888289016116d0565b9150509295509295909350565b61196081611f42565b82525050565b61196f81611f9d565b82525050565b600061198082611e67565b61198a8185611e72565b935061199a818560208601611fbe565b6119a381612153565b840191505092915050565b60006119bb604383611e72565b91506119c682612164565b606082019050919050565b60006119de602a83611e72565b91506119e9826121d9565b604082019050919050565b6000611a01602b83611e72565b9150611a0c82612228565b604082019050919050565b6000611a24602883611e72565b9150611a2f82612277565b604082019050919050565b6000611a47603b83611e72565b9150611a52826122c6565b604082019050919050565b6000611a6a602583611e72565b9150611a7582612315565b604082019050919050565b6000611a8d603a83611e72565b9150611a9882612364565b604082019050919050565b6000611ab0602983611e72565b9150611abb826123b3565b604082019050919050565b6000611ad3602e83611e72565b9150611ade82612402565b604082019050919050565b6000611af6604383611e72565b9150611b0182612451565b606082019050919050565b6000611b19602883611e72565b9150611b24826124c6565b604082019050919050565b6000611b3c604183611e72565b9150611b4782612515565b606082019050919050565b611b5b81611f93565b82525050565b6000602082019050611b766000830184611957565b92915050565b6000604082019050611b916000830185611957565b611b9e6020830184611b52565b9392505050565b6000602082019050611bba6000830184611966565b92915050565b600061012082019050611bd6600083018c611966565b611be3602083018b611b52565b611bf0604083018a611b52565b611bfd6060830189611b52565b611c0a6080830188611b52565b611c1760a0830187611b52565b611c2460c0830186611957565b611c3160e0830185611957565b818103610100830152611c448184611975565b90509a9950505050505050505050565b60006020820190508181036000830152611c6e8184611975565b905092915050565b60006020820190508181036000830152611c8f816119ae565b9050919050565b60006020820190508181036000830152611caf816119d1565b9050919050565b60006020820190508181036000830152611ccf816119f4565b9050919050565b60006020820190508181036000830152611cef81611a17565b9050919050565b60006020820190508181036000830152611d0f81611a3a565b9050919050565b60006020820190508181036000830152611d2f81611a5d565b9050919050565b60006020820190508181036000830152611d4f81611a80565b9050919050565b60006020820190508181036000830152611d6f81611aa3565b9050919050565b60006020820190508181036000830152611d8f81611ac6565b9050919050565b60006020820190508181036000830152611daf81611ae9565b9050919050565b60006020820190508181036000830152611dcf81611b0c565b9050919050565b60006020820190508181036000830152611def81611b2f565b9050919050565b6000602082019050611e0b6000830184611b52565b92915050565b6000611e1b611e2c565b9050611e278282612023565b919050565b6000604051905090565b600067ffffffffffffffff821115611e5157611e50612110565b5b611e5a82612153565b9050602081019050919050565b600081519050919050565b600082825260208201905092915050565b6000611e8e82611f93565b9150611e9983611f93565b925082611ea957611ea8612083565b5b828204905092915050565b6000611ebf82611f93565b9150611eca83611f93565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611f0357611f02612054565b5b828202905092915050565b6000611f1982611f93565b9150611f2483611f93565b925082821015611f3757611f36612054565b5b828203905092915050565b6000611f4d82611f73565b9050919050565b60008115159050919050565b6000819050611f6e8261258a565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000611fa882611f60565b9050919050565b82818337600083830152505050565b60005b83811015611fdc578082015181840152602081019050611fc1565b83811115611feb576000848401525b50505050565b6000600282049050600182168061200957607f821691505b6020821081141561201d5761201c6120e1565b5b50919050565b61202c82612153565b810181811067ffffffffffffffff8211171561204b5761204a612110565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f74686973206163636f756e74206973206e6f7420617574686f72697a6564207460008201527f6f2075706461746520746865206369706865727465787420696e666f726d617460208201527f696f6e0000000000000000000000000000000000000000000000000000000000604082015250565b7f74686520636f6e7472616374206861732079657420746f20626520636172726960008201527f656420746f207465726d00000000000000000000000000000000000000000000602082015250565b7f74686973206973206163636f756e74206973206e6f7420696e2074686520727560008201527f6e6e696e67207374617465000000000000000000000000000000000000000000602082015250565b7f796f752063616e206f6e6c792075736520636c6f73656f7574206f7074696f6e60008201527f732032206f722033000000000000000000000000000000000000000000000000602082015250565b7f74686973206163636f756e74206973206e6f7420617574686f72697a6564207460008201527f6f207472696767657220616e206561726c7920636c6f73656f75740000000000602082015250565b7f636f6e7472616374206973206e6f7420696e20616e20617661696c61626c652060008201527f7374617465000000000000000000000000000000000000000000000000000000602082015250565b7f746869732061646472657373206973206e6f7420617070726f76656420746f2060008201527f63616c6c207468652070757263686173652066756e6374696f6e000000000000602082015250565b7f796f75206d757374206d616b6520612073656c656374696f6e2062657477656560008201527f6e203020616e6420330000000000000000000000000000000000000000000000602082015250565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b7f74686973206973206163636f756e74206973206e6f7420617574686f72697a6560008201527f6420746f207570646174652074686520636f6e747261637420706172616d657460208201527f6572730000000000000000000000000000000000000000000000000000000000604082015250565b7f74686520636f6e7472616374206973206e6f7420696e207468652072756e6e6960008201527f6e67207374617465000000000000000000000000000000000000000000000000602082015250565b7f74686973206163636f756e74206973206e6f7420617574686f72697a6564207460008201527f6f20747269676765722061206d69642d636f6e747261637420636c6f73656f7560208201527f7400000000000000000000000000000000000000000000000000000000000000604082015250565b6002811061259b5761259a6120b2565b5b50565b6125a781611f42565b81146125b257600080fd5b50565b6125be81611f54565b81146125c957600080fd5b50565b6125d581611f93565b81146125e057600080fd5b5056fea26469706673582212200649dffda7f488c7e68a704123d19e2d86269852d188929c80057fea8e736a1264736f6c63430008070033",
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
