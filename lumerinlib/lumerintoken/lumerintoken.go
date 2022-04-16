// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lumerintoken

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

// LumerintokenMetaData contains all meta data concerning the Lumerintoken contract.
var LumerintokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040518060400160405280600781526020017f4c756d6572696e000000000000000000000000000000000000000000000000008152506040518060400160405280600381526020017f4c4d520000000000000000000000000000000000000000000000000000000000815250816003908051906020019062000096929190620003e5565b508060049080519060200190620000af929190620003e5565b5050506000600560006101000a81548160ff021916908315150217905550620000ed620000e16200010d60201b60201c565b6200011560201b60201c565b620001073367016345785d8a0000620001db60201b60201c565b620006b3565b600033905090565b6000600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600560016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156200024e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620002459062000516565b60405180910390fd5b62000262600083836200035460201b60201c565b806002600082825462000276919062000566565b92505081905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254620002cd919062000566565b925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405162000334919062000538565b60405180910390a36200035060008383620003c460201b60201c565b5050565b62000364620003c960201b60201c565b15620003a7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200039e90620004f4565b60405180910390fd5b620003bf838383620003e060201b62000bc91760201c565b505050565b505050565b6000600560009054906101000a900460ff16905090565b505050565b828054620003f390620005cd565b90600052602060002090601f01602090048101928262000417576000855562000463565b82601f106200043257805160ff191683800117855562000463565b8280016001018555821562000463579182015b828111156200046257825182559160200191906001019062000445565b5b50905062000472919062000476565b5090565b5b808211156200049157600081600090555060010162000477565b5090565b6000620004a460108362000555565b9150620004b18262000661565b602082019050919050565b6000620004cb601f8362000555565b9150620004d8826200068a565b602082019050919050565b620004ee81620005c3565b82525050565b600060208201905081810360008301526200050f8162000495565b9050919050565b600060208201905081810360008301526200053181620004bc565b9050919050565b60006020820190506200054f6000830184620004e3565b92915050565b600082825260208201905092915050565b60006200057382620005c3565b91506200058083620005c3565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115620005b857620005b762000603565b5b828201905092915050565b6000819050919050565b60006002820490506001821680620005e657607f821691505b60208210811415620005fd57620005fc62000632565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f5061757361626c653a2070617573656400000000000000000000000000000000600082015250565b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b61207d80620006c36000396000f3fe608060405234801561001057600080fd5b50600436106101215760003560e01c806370a08231116100ad57806395d89b411161007157806395d89b41146102d2578063a457c2d7146102f0578063a9059cbb14610320578063dd62ed3e14610350578063f2fde38b1461038057610121565b806370a0823114610254578063715018a61461028457806379cc67901461028e5780638456cb59146102aa5780638da5cb5b146102b457610121565b8063313ce567116100f4578063313ce567146101c257806339509351146101e05780633f4ba83a1461021057806342966c681461021a5780635c975abb1461023657610121565b806306fdde0314610126578063095ea7b31461014457806318160ddd1461017457806323b872dd14610192575b600080fd5b61012e61039c565b60405161013b919061184d565b60405180910390f35b61015e6004803603810190610159919061154b565b61042e565b60405161016b9190611832565b60405180910390f35b61017c61044c565b6040516101899190611a2f565b60405180910390f35b6101ac60048036038101906101a791906114f8565b610456565b6040516101b99190611832565b60405180910390f35b6101ca61054e565b6040516101d79190611a4a565b60405180910390f35b6101fa60048036038101906101f5919061154b565b610557565b6040516102079190611832565b60405180910390f35b610218610603565b005b610234600480360381019061022f919061158b565b610689565b005b61023e61069d565b60405161024b9190611832565b60405180910390f35b61026e6004803603810190610269919061148b565b6106b4565b60405161027b9190611a2f565b60405180910390f35b61028c6106fc565b005b6102a860048036038101906102a3919061154b565b610784565b005b6102b26107ff565b005b6102bc610885565b6040516102c99190611817565b60405180910390f35b6102da6108af565b6040516102e7919061184d565b60405180910390f35b61030a6004803603810190610305919061154b565b610941565b6040516103179190611832565b60405180910390f35b61033a6004803603810190610335919061154b565b610a2c565b6040516103479190611832565b60405180910390f35b61036a600480360381019061036591906114b8565b610a4a565b6040516103779190611a2f565b60405180910390f35b61039a6004803603810190610395919061148b565b610ad1565b005b6060600380546103ab90611b93565b80601f01602080910402602001604051908101604052809291908181526020018280546103d790611b93565b80156104245780601f106103f957610100808354040283529160200191610424565b820191906000526020600020905b81548152906001019060200180831161040757829003601f168201915b5050505050905090565b600061044261043b610bce565b8484610bd6565b6001905092915050565b6000600254905090565b6000610463848484610da1565b6000600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006104ae610bce565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490508281101561052e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105259061194f565b60405180910390fd5b6105428561053a610bce565b858403610bd6565b60019150509392505050565b60006012905090565b60006105f9610564610bce565b848460016000610572610bce565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546105f49190611a81565b610bd6565b6001905092915050565b61060b610bce565b73ffffffffffffffffffffffffffffffffffffffff16610629610885565b73ffffffffffffffffffffffffffffffffffffffff161461067f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106769061196f565b60405180910390fd5b610687611022565b565b61069a610694610bce565b826110c4565b50565b6000600560009054906101000a900460ff16905090565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b610704610bce565b73ffffffffffffffffffffffffffffffffffffffff16610722610885565b73ffffffffffffffffffffffffffffffffffffffff1614610778576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161076f9061196f565b60405180910390fd5b610782600061129b565b565b600061079783610792610bce565b610a4a565b9050818110156107dc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107d39061198f565b60405180910390fd5b6107f0836107e8610bce565b848403610bd6565b6107fa83836110c4565b505050565b610807610bce565b73ffffffffffffffffffffffffffffffffffffffff16610825610885565b73ffffffffffffffffffffffffffffffffffffffff161461087b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108729061196f565b60405180910390fd5b610883611361565b565b6000600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6060600480546108be90611b93565b80601f01602080910402602001604051908101604052809291908181526020018280546108ea90611b93565b80156109375780601f1061090c57610100808354040283529160200191610937565b820191906000526020600020905b81548152906001019060200180831161091a57829003601f168201915b5050505050905090565b60008060016000610950610bce565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905082811015610a0d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a0490611a0f565b60405180910390fd5b610a21610a18610bce565b85858403610bd6565b600191505092915050565b6000610a40610a39610bce565b8484610da1565b6001905092915050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b610ad9610bce565b73ffffffffffffffffffffffffffffffffffffffff16610af7610885565b73ffffffffffffffffffffffffffffffffffffffff1614610b4d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b449061196f565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610bbd576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bb4906118cf565b60405180910390fd5b610bc68161129b565b50565b505050565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610c46576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c3d906119ef565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610cb6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cad906118ef565b60405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92583604051610d949190611a2f565b60405180910390a3505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610e11576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e08906119cf565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610e81576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e789061186f565b60405180910390fd5b610e8c838383611404565b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015610f12576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f099061190f565b60405180910390fd5b8181036000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610fa59190611a81565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516110099190611a2f565b60405180910390a361101c84848461145c565b50505050565b61102a61069d565b611069576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110609061188f565b60405180910390fd5b6000600560006101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6110ad610bce565b6040516110ba9190611817565b60405180910390a1565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611134576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161112b906119af565b60405180910390fd5b61114082600083611404565b60008060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050818110156111c6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111bd906118af565b60405180910390fd5b8181036000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550816002600082825461121d9190611ad7565b92505081905550600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516112829190611a2f565b60405180910390a36112968360008461145c565b505050565b6000600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600560016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b61136961069d565b156113a9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113a09061192f565b60405180910390fd5b6001600560006101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586113ed610bce565b6040516113fa9190611817565b60405180910390a1565b61140c61069d565b1561144c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114439061192f565b60405180910390fd5b611457838383610bc9565b505050565b505050565b60008135905061147081612019565b92915050565b60008135905061148581612030565b92915050565b6000602082840312156114a1576114a0611c23565b5b60006114af84828501611461565b91505092915050565b600080604083850312156114cf576114ce611c23565b5b60006114dd85828601611461565b92505060206114ee85828601611461565b9150509250929050565b60008060006060848603121561151157611510611c23565b5b600061151f86828701611461565b935050602061153086828701611461565b925050604061154186828701611476565b9150509250925092565b6000806040838503121561156257611561611c23565b5b600061157085828601611461565b925050602061158185828601611476565b9150509250929050565b6000602082840312156115a1576115a0611c23565b5b60006115af84828501611476565b91505092915050565b6115c181611b0b565b82525050565b6115d081611b1d565b82525050565b60006115e182611a65565b6115eb8185611a70565b93506115fb818560208601611b60565b61160481611c28565b840191505092915050565b600061161c602383611a70565b915061162782611c39565b604082019050919050565b600061163f601483611a70565b915061164a82611c88565b602082019050919050565b6000611662602283611a70565b915061166d82611cb1565b604082019050919050565b6000611685602683611a70565b915061169082611d00565b604082019050919050565b60006116a8602283611a70565b91506116b382611d4f565b604082019050919050565b60006116cb602683611a70565b91506116d682611d9e565b604082019050919050565b60006116ee601083611a70565b91506116f982611ded565b602082019050919050565b6000611711602883611a70565b915061171c82611e16565b604082019050919050565b6000611734602083611a70565b915061173f82611e65565b602082019050919050565b6000611757602483611a70565b915061176282611e8e565b604082019050919050565b600061177a602183611a70565b915061178582611edd565b604082019050919050565b600061179d602583611a70565b91506117a882611f2c565b604082019050919050565b60006117c0602483611a70565b91506117cb82611f7b565b604082019050919050565b60006117e3602583611a70565b91506117ee82611fca565b604082019050919050565b61180281611b49565b82525050565b61181181611b53565b82525050565b600060208201905061182c60008301846115b8565b92915050565b600060208201905061184760008301846115c7565b92915050565b6000602082019050818103600083015261186781846115d6565b905092915050565b600060208201905081810360008301526118888161160f565b9050919050565b600060208201905081810360008301526118a881611632565b9050919050565b600060208201905081810360008301526118c881611655565b9050919050565b600060208201905081810360008301526118e881611678565b9050919050565b600060208201905081810360008301526119088161169b565b9050919050565b60006020820190508181036000830152611928816116be565b9050919050565b60006020820190508181036000830152611948816116e1565b9050919050565b6000602082019050818103600083015261196881611704565b9050919050565b6000602082019050818103600083015261198881611727565b9050919050565b600060208201905081810360008301526119a88161174a565b9050919050565b600060208201905081810360008301526119c88161176d565b9050919050565b600060208201905081810360008301526119e881611790565b9050919050565b60006020820190508181036000830152611a08816117b3565b9050919050565b60006020820190508181036000830152611a28816117d6565b9050919050565b6000602082019050611a4460008301846117f9565b92915050565b6000602082019050611a5f6000830184611808565b92915050565b600081519050919050565b600082825260208201905092915050565b6000611a8c82611b49565b9150611a9783611b49565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611acc57611acb611bc5565b5b828201905092915050565b6000611ae282611b49565b9150611aed83611b49565b925082821015611b0057611aff611bc5565b5b828203905092915050565b6000611b1682611b29565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b60005b83811015611b7e578082015181840152602081019050611b63565b83811115611b8d576000848401525b50505050565b60006002820490506001821680611bab57607f821691505b60208210811415611bbf57611bbe611bf4565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600080fd5b6000601f19601f8301169050919050565b7f45524332303a207472616e7366657220746f20746865207a65726f206164647260008201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b7f5061757361626c653a206e6f7420706175736564000000000000000000000000600082015250565b7f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60008201527f6365000000000000000000000000000000000000000000000000000000000000602082015250565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a20617070726f766520746f20746865207a65726f20616464726560008201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206260008201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b7f5061757361626c653a2070617573656400000000000000000000000000000000600082015250565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206160008201527f6c6c6f77616e6365000000000000000000000000000000000000000000000000602082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b7f45524332303a206275726e20616d6f756e74206578636565647320616c6c6f7760008201527f616e636500000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360008201527f7300000000000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a207472616e736665722066726f6d20746865207a65726f20616460008201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760008201527f207a65726f000000000000000000000000000000000000000000000000000000602082015250565b61202281611b0b565b811461202d57600080fd5b50565b61203981611b49565b811461204457600080fd5b5056fea2646970667358221220936c30015e961c857f30b73921621c3711185400ae8752f7a84166fd6ff7727264736f6c63430008070033",
}

// LumerintokenABI is the input ABI used to generate the binding from.
// Deprecated: Use LumerintokenMetaData.ABI instead.
var LumerintokenABI = LumerintokenMetaData.ABI

// LumerintokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LumerintokenMetaData.Bin instead.
var LumerintokenBin = LumerintokenMetaData.Bin

// DeployLumerintoken deploys a new Ethereum contract, binding an instance of Lumerintoken to it.
func DeployLumerintoken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Lumerintoken, error) {
	parsed, err := LumerintokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LumerintokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Lumerintoken{LumerintokenCaller: LumerintokenCaller{contract: contract}, LumerintokenTransactor: LumerintokenTransactor{contract: contract}, LumerintokenFilterer: LumerintokenFilterer{contract: contract}}, nil
}

// Lumerintoken is an auto generated Go binding around an Ethereum contract.
type Lumerintoken struct {
	LumerintokenCaller     // Read-only binding to the contract
	LumerintokenTransactor // Write-only binding to the contract
	LumerintokenFilterer   // Log filterer for contract events
}

// LumerintokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type LumerintokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LumerintokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LumerintokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LumerintokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LumerintokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LumerintokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LumerintokenSession struct {
	Contract     *Lumerintoken     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LumerintokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LumerintokenCallerSession struct {
	Contract *LumerintokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// LumerintokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LumerintokenTransactorSession struct {
	Contract     *LumerintokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// LumerintokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type LumerintokenRaw struct {
	Contract *Lumerintoken // Generic contract binding to access the raw methods on
}

// LumerintokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LumerintokenCallerRaw struct {
	Contract *LumerintokenCaller // Generic read-only contract binding to access the raw methods on
}

// LumerintokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LumerintokenTransactorRaw struct {
	Contract *LumerintokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLumerintoken creates a new instance of Lumerintoken, bound to a specific deployed contract.
func NewLumerintoken(address common.Address, backend bind.ContractBackend) (*Lumerintoken, error) {
	contract, err := bindLumerintoken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lumerintoken{LumerintokenCaller: LumerintokenCaller{contract: contract}, LumerintokenTransactor: LumerintokenTransactor{contract: contract}, LumerintokenFilterer: LumerintokenFilterer{contract: contract}}, nil
}

// NewLumerintokenCaller creates a new read-only instance of Lumerintoken, bound to a specific deployed contract.
func NewLumerintokenCaller(address common.Address, caller bind.ContractCaller) (*LumerintokenCaller, error) {
	contract, err := bindLumerintoken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LumerintokenCaller{contract: contract}, nil
}

// NewLumerintokenTransactor creates a new write-only instance of Lumerintoken, bound to a specific deployed contract.
func NewLumerintokenTransactor(address common.Address, transactor bind.ContractTransactor) (*LumerintokenTransactor, error) {
	contract, err := bindLumerintoken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LumerintokenTransactor{contract: contract}, nil
}

// NewLumerintokenFilterer creates a new log filterer instance of Lumerintoken, bound to a specific deployed contract.
func NewLumerintokenFilterer(address common.Address, filterer bind.ContractFilterer) (*LumerintokenFilterer, error) {
	contract, err := bindLumerintoken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LumerintokenFilterer{contract: contract}, nil
}

// bindLumerintoken binds a generic wrapper to an already deployed contract.
func bindLumerintoken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LumerintokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lumerintoken *LumerintokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lumerintoken.Contract.LumerintokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lumerintoken *LumerintokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lumerintoken.Contract.LumerintokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lumerintoken *LumerintokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lumerintoken.Contract.LumerintokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lumerintoken *LumerintokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lumerintoken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lumerintoken *LumerintokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lumerintoken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lumerintoken *LumerintokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lumerintoken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Lumerintoken *LumerintokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Lumerintoken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Lumerintoken *LumerintokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Lumerintoken.Contract.Allowance(&_Lumerintoken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Lumerintoken *LumerintokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Lumerintoken.Contract.Allowance(&_Lumerintoken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Lumerintoken *LumerintokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Lumerintoken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Lumerintoken *LumerintokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Lumerintoken.Contract.BalanceOf(&_Lumerintoken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Lumerintoken *LumerintokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Lumerintoken.Contract.BalanceOf(&_Lumerintoken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Lumerintoken *LumerintokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Lumerintoken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Lumerintoken *LumerintokenSession) Decimals() (uint8, error) {
	return _Lumerintoken.Contract.Decimals(&_Lumerintoken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Lumerintoken *LumerintokenCallerSession) Decimals() (uint8, error) {
	return _Lumerintoken.Contract.Decimals(&_Lumerintoken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Lumerintoken *LumerintokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Lumerintoken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Lumerintoken *LumerintokenSession) Name() (string, error) {
	return _Lumerintoken.Contract.Name(&_Lumerintoken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Lumerintoken *LumerintokenCallerSession) Name() (string, error) {
	return _Lumerintoken.Contract.Name(&_Lumerintoken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Lumerintoken *LumerintokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lumerintoken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Lumerintoken *LumerintokenSession) Owner() (common.Address, error) {
	return _Lumerintoken.Contract.Owner(&_Lumerintoken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Lumerintoken *LumerintokenCallerSession) Owner() (common.Address, error) {
	return _Lumerintoken.Contract.Owner(&_Lumerintoken.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Lumerintoken *LumerintokenCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Lumerintoken.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Lumerintoken *LumerintokenSession) Paused() (bool, error) {
	return _Lumerintoken.Contract.Paused(&_Lumerintoken.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Lumerintoken *LumerintokenCallerSession) Paused() (bool, error) {
	return _Lumerintoken.Contract.Paused(&_Lumerintoken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Lumerintoken *LumerintokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Lumerintoken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Lumerintoken *LumerintokenSession) Symbol() (string, error) {
	return _Lumerintoken.Contract.Symbol(&_Lumerintoken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Lumerintoken *LumerintokenCallerSession) Symbol() (string, error) {
	return _Lumerintoken.Contract.Symbol(&_Lumerintoken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Lumerintoken *LumerintokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lumerintoken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Lumerintoken *LumerintokenSession) TotalSupply() (*big.Int, error) {
	return _Lumerintoken.Contract.TotalSupply(&_Lumerintoken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Lumerintoken *LumerintokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Lumerintoken.Contract.TotalSupply(&_Lumerintoken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Lumerintoken *LumerintokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Lumerintoken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Lumerintoken *LumerintokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Lumerintoken.Contract.Approve(&_Lumerintoken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Lumerintoken *LumerintokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Lumerintoken.Contract.Approve(&_Lumerintoken.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Lumerintoken *LumerintokenTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Lumerintoken.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Lumerintoken *LumerintokenSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Lumerintoken.Contract.Burn(&_Lumerintoken.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Lumerintoken *LumerintokenTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Lumerintoken.Contract.Burn(&_Lumerintoken.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Lumerintoken *LumerintokenTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Lumerintoken.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Lumerintoken *LumerintokenSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Lumerintoken.Contract.BurnFrom(&_Lumerintoken.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Lumerintoken *LumerintokenTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Lumerintoken.Contract.BurnFrom(&_Lumerintoken.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Lumerintoken *LumerintokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Lumerintoken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Lumerintoken *LumerintokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Lumerintoken.Contract.DecreaseAllowance(&_Lumerintoken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Lumerintoken *LumerintokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Lumerintoken.Contract.DecreaseAllowance(&_Lumerintoken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Lumerintoken *LumerintokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Lumerintoken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Lumerintoken *LumerintokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Lumerintoken.Contract.IncreaseAllowance(&_Lumerintoken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Lumerintoken *LumerintokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Lumerintoken.Contract.IncreaseAllowance(&_Lumerintoken.TransactOpts, spender, addedValue)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Lumerintoken *LumerintokenTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lumerintoken.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Lumerintoken *LumerintokenSession) Pause() (*types.Transaction, error) {
	return _Lumerintoken.Contract.Pause(&_Lumerintoken.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Lumerintoken *LumerintokenTransactorSession) Pause() (*types.Transaction, error) {
	return _Lumerintoken.Contract.Pause(&_Lumerintoken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Lumerintoken *LumerintokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lumerintoken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Lumerintoken *LumerintokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _Lumerintoken.Contract.RenounceOwnership(&_Lumerintoken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Lumerintoken *LumerintokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Lumerintoken.Contract.RenounceOwnership(&_Lumerintoken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Lumerintoken *LumerintokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Lumerintoken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Lumerintoken *LumerintokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Lumerintoken.Contract.Transfer(&_Lumerintoken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Lumerintoken *LumerintokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Lumerintoken.Contract.Transfer(&_Lumerintoken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Lumerintoken *LumerintokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Lumerintoken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Lumerintoken *LumerintokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Lumerintoken.Contract.TransferFrom(&_Lumerintoken.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Lumerintoken *LumerintokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Lumerintoken.Contract.TransferFrom(&_Lumerintoken.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Lumerintoken *LumerintokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Lumerintoken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Lumerintoken *LumerintokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Lumerintoken.Contract.TransferOwnership(&_Lumerintoken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Lumerintoken *LumerintokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Lumerintoken.Contract.TransferOwnership(&_Lumerintoken.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Lumerintoken *LumerintokenTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lumerintoken.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Lumerintoken *LumerintokenSession) Unpause() (*types.Transaction, error) {
	return _Lumerintoken.Contract.Unpause(&_Lumerintoken.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Lumerintoken *LumerintokenTransactorSession) Unpause() (*types.Transaction, error) {
	return _Lumerintoken.Contract.Unpause(&_Lumerintoken.TransactOpts)
}

// LumerintokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Lumerintoken contract.
type LumerintokenApprovalIterator struct {
	Event *LumerintokenApproval // Event containing the contract specifics and raw log

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
func (it *LumerintokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LumerintokenApproval)
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
		it.Event = new(LumerintokenApproval)
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
func (it *LumerintokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LumerintokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LumerintokenApproval represents a Approval event raised by the Lumerintoken contract.
type LumerintokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Lumerintoken *LumerintokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*LumerintokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Lumerintoken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &LumerintokenApprovalIterator{contract: _Lumerintoken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Lumerintoken *LumerintokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LumerintokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Lumerintoken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LumerintokenApproval)
				if err := _Lumerintoken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Lumerintoken *LumerintokenFilterer) ParseApproval(log types.Log) (*LumerintokenApproval, error) {
	event := new(LumerintokenApproval)
	if err := _Lumerintoken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LumerintokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Lumerintoken contract.
type LumerintokenOwnershipTransferredIterator struct {
	Event *LumerintokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LumerintokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LumerintokenOwnershipTransferred)
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
		it.Event = new(LumerintokenOwnershipTransferred)
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
func (it *LumerintokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LumerintokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LumerintokenOwnershipTransferred represents a OwnershipTransferred event raised by the Lumerintoken contract.
type LumerintokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Lumerintoken *LumerintokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LumerintokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Lumerintoken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LumerintokenOwnershipTransferredIterator{contract: _Lumerintoken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Lumerintoken *LumerintokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LumerintokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Lumerintoken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LumerintokenOwnershipTransferred)
				if err := _Lumerintoken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Lumerintoken *LumerintokenFilterer) ParseOwnershipTransferred(log types.Log) (*LumerintokenOwnershipTransferred, error) {
	event := new(LumerintokenOwnershipTransferred)
	if err := _Lumerintoken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LumerintokenPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Lumerintoken contract.
type LumerintokenPausedIterator struct {
	Event *LumerintokenPaused // Event containing the contract specifics and raw log

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
func (it *LumerintokenPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LumerintokenPaused)
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
		it.Event = new(LumerintokenPaused)
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
func (it *LumerintokenPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LumerintokenPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LumerintokenPaused represents a Paused event raised by the Lumerintoken contract.
type LumerintokenPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Lumerintoken *LumerintokenFilterer) FilterPaused(opts *bind.FilterOpts) (*LumerintokenPausedIterator, error) {

	logs, sub, err := _Lumerintoken.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &LumerintokenPausedIterator{contract: _Lumerintoken.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Lumerintoken *LumerintokenFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *LumerintokenPaused) (event.Subscription, error) {

	logs, sub, err := _Lumerintoken.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LumerintokenPaused)
				if err := _Lumerintoken.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Lumerintoken *LumerintokenFilterer) ParsePaused(log types.Log) (*LumerintokenPaused, error) {
	event := new(LumerintokenPaused)
	if err := _Lumerintoken.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LumerintokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Lumerintoken contract.
type LumerintokenTransferIterator struct {
	Event *LumerintokenTransfer // Event containing the contract specifics and raw log

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
func (it *LumerintokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LumerintokenTransfer)
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
		it.Event = new(LumerintokenTransfer)
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
func (it *LumerintokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LumerintokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LumerintokenTransfer represents a Transfer event raised by the Lumerintoken contract.
type LumerintokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Lumerintoken *LumerintokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LumerintokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Lumerintoken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LumerintokenTransferIterator{contract: _Lumerintoken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Lumerintoken *LumerintokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LumerintokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Lumerintoken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LumerintokenTransfer)
				if err := _Lumerintoken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Lumerintoken *LumerintokenFilterer) ParseTransfer(log types.Log) (*LumerintokenTransfer, error) {
	event := new(LumerintokenTransfer)
	if err := _Lumerintoken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LumerintokenUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Lumerintoken contract.
type LumerintokenUnpausedIterator struct {
	Event *LumerintokenUnpaused // Event containing the contract specifics and raw log

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
func (it *LumerintokenUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LumerintokenUnpaused)
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
		it.Event = new(LumerintokenUnpaused)
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
func (it *LumerintokenUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LumerintokenUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LumerintokenUnpaused represents a Unpaused event raised by the Lumerintoken contract.
type LumerintokenUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Lumerintoken *LumerintokenFilterer) FilterUnpaused(opts *bind.FilterOpts) (*LumerintokenUnpausedIterator, error) {

	logs, sub, err := _Lumerintoken.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &LumerintokenUnpausedIterator{contract: _Lumerintoken.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Lumerintoken *LumerintokenFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *LumerintokenUnpaused) (event.Subscription, error) {

	logs, sub, err := _Lumerintoken.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LumerintokenUnpaused)
				if err := _Lumerintoken.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Lumerintoken *LumerintokenFilterer) ParseUnpaused(log types.Log) (*LumerintokenUnpaused, error) {
	event := new(LumerintokenUnpaused)
	if err := _Lumerintoken.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
