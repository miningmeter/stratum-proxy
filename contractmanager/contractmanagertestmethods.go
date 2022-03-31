package contractmanager

import (
	"context"
	// "crypto/ecdsa"
	// "crypto/rand"
	// "errors"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/ethclient"

	"gitlab.com/TitanInd/hashrouter/lumerinlib/clonefactory"
	// "gitlab.com/TitanInd/lumerin/lumerinlib/implementation"
	"gitlab.com/TitanInd/hashrouter/lumerinlib/lumerintoken"
)

// type TestSetup struct {
// 	ethClient              *ethclient.Client
// 	nodeEthereumPrivateKey string
// 	nodeEthereumAccount    common.Address
// 	validatorAddress       common.Address
// 	proxyAddress           common.Address
// 	lumerinAddress         common.Address
// 	cloneFactoryAddress    common.Address
// }

func CreateHashrateContract(client *ethclient.Client,
	fromAddress common.Address,
	privateKeyString string,
	contractAddress common.Address,
	_price int,
	_limit int,
	_speed int,
	_length int,
	_validator common.Address) {
	privateKey, err := crypto.HexToECDSA(privateKeyString)
	if err != nil {
		log.Fatalf("Error::%v", err)
	}

	time.Sleep(time.Millisecond * 700)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("Error::%v", err)
	}
	fmt.Println("Nonce: ", nonce)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Error::%v", err)
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("Error::%v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatalf("Error::%v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	instance, err := clonefactory.NewClonefactory(contractAddress, client)
	if err != nil {
		log.Fatalf("Error::%v", err)
	}

	price := big.NewInt(int64(_price))
	limit := big.NewInt(int64(_limit))
	speed := big.NewInt(int64(_speed))
	length := big.NewInt(int64(_length))
	tx, err := instance.SetCreateNewRentalContract(auth, price, limit, speed, length, _validator, "")
	if err != nil {
		log.Fatalf("Error::%v", err)
	}

	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
}

func PurchaseHashrateContract(client *ethclient.Client,
	fromAddress common.Address,
	privateKeyString string,
	contractAddress common.Address,
	_hashrateContract common.Address,
	_buyer common.Address,
	poolData string) {
	privateKey, err := crypto.HexToECDSA(privateKeyString)
	if err != nil {
		log.Fatalf("Error::%v", err)
	}

	time.Sleep(time.Millisecond * 700)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("Error::%v", err)
	}
	fmt.Println("Nonce: ", nonce)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Error::%v", err)
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("Error::%v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatalf("Error::%v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	instance, err := clonefactory.NewClonefactory(contractAddress, client)
	if err != nil {
		log.Fatalf("Error::%v", err)
	}

	tx, err := instance.SetPurchaseRentalContract(auth, _hashrateContract, poolData)
	if err != nil {
		log.Fatalf("Error::%v", err)
	}
	fmt.Printf("tx sent: %s\n\n", tx.Hash().Hex())
	fmt.Printf("Hashrate Contract %s, was purchased by %s\n\n", _hashrateContract, _buyer)
}

func DeployContracts(client *ethclient.Client,
	fromAddress common.Address,
	privateKeyString string) common.Address {
	privateKey, err := crypto.HexToECDSA(privateKeyString)
	if err != nil {
		log.Fatalf("Error::%v", err)
	}

	time.Sleep(time.Millisecond * 700)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("Error::%v", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Error::%v", err)
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("Error::%v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatalf("Error::%v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(6000000) // in units
	auth.GasPrice = gasPrice

	fmt.Println("Deploying Lumerin Token contract")
	laddress, ltransaction, _, err := lumerintoken.DeployLumerintoken(auth, client)
	if err != nil {
		log.Fatalf("Error::%v", err)
	}

	// wait until lumerin token was deployed
	_, lerr := client.TransactionReceipt(context.Background(), ltransaction.Hash())
	for lerr != nil {
		_, lerr = client.TransactionReceipt(context.Background(), ltransaction.Hash())
		time.Sleep(time.Second * 1)
	}

	nonce, err = client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("Error::%v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))

	fmt.Println("Deploying Clone Factory contract")
	cfaddress, ctransaction, _, err := clonefactory.DeployClonefactory(auth, client, laddress, common.HexToAddress(""))
	if err != nil {
		log.Fatalf("Error::%v", err)
	}

	// wait until clonefactory was deployed
	_, cferr := client.TransactionReceipt(context.Background(), ctransaction.Hash())
	for cferr != nil {
		_, cferr = client.TransactionReceipt(context.Background(), ctransaction.Hash())
		time.Sleep(time.Second * 1)
	}

	return cfaddress
}
