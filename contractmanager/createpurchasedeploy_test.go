package contractmanager

import (
	"log"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestHashrateContractCreation(t *testing.T) {
	clonefactoryAddress := common.HexToAddress("")
	accountAddress := common.HexToAddress("")
	accountPrivateKey := ""
	gethNodeAddress := ""

	// hashrate contract params
	price := 0
	limit := 10
	speed := 100
	length := 100

	client, err := setUpClient(gethNodeAddress, accountAddress)
	if err != nil {
		log.Fatalf("Error::%v", err)
	}
	
	CreateHashrateContract(client, accountAddress, accountPrivateKey, clonefactoryAddress, price, limit, speed, length, clonefactoryAddress)

	// subcribe to creation events emitted by clonefactory contract
	cfLogs, cfSub, _ := subscribeToContractEvents(client, clonefactoryAddress)
	// create event signature to parse out creation event
	contractCreatedSig := []byte("contractCreated(address,string)")
	contractCreatedSigHash := crypto.Keccak256Hash(contractCreatedSig)
	for {
		select {
		case err := <-cfSub.Err():
			log.Fatalf("Error::%v", err)		
		case cfLog := <-cfLogs:
			
			if cfLog.Topics[0].Hex() == contractCreatedSigHash.Hex() {
				hashrateContractAddress := common.HexToAddress(cfLog.Topics[1].Hex())
				fmt.Printf("Address of created Hashrate Contract: %v\n\n", hashrateContractAddress.Hex())
			}
		}
	}
}


func TestHashrateContractPurchase(t *testing.T) {
	clonefactoryAddress := common.HexToAddress("")
	accountAddress := common.HexToAddress("")
	accountPrivateKey := ""
	gethNodeAddress := ""

	hashrateContractAddress := common.HexToAddress("")
	poolUrl := ""

	client, err := setUpClient(gethNodeAddress, accountAddress)
	if err != nil {
		log.Fatalf("Error::%v", err)
	}

	PurchaseHashrateContract(client, accountAddress, accountPrivateKey, clonefactoryAddress, hashrateContractAddress, accountAddress, poolUrl)
	
	// subcribe to purchase events emitted by clonefactory contract
	cfLogs, cfSub, _ := subscribeToContractEvents(client, clonefactoryAddress)
	// create event signature to parse out purchase event
	clonefactoryContractPurchasedSig := []byte("clonefactoryContractPurchased(address)")
	clonefactoryContractPurchasedSigHash := crypto.Keccak256Hash(clonefactoryContractPurchasedSig)
	for {
		select {
		case err := <-cfSub.Err():
			log.Fatalf("Error::%v", err)		
		case cfLog := <-cfLogs:
			
			if cfLog.Topics[0].Hex() == clonefactoryContractPurchasedSigHash.Hex() {
				hashrateContractAddress := common.HexToAddress(cfLog.Topics[1].Hex())
				fmt.Printf("Address of purchased Hashrate Contract: %v\n\n", hashrateContractAddress.Hex())
			}
		}
	}
}

func TestDeployContracts(t *testing.T) {
	accountAddress := common.HexToAddress("")
	accountPrivateKey := ""
	gethNodeAddress := ""

	client, err := setUpClient(gethNodeAddress, accountAddress)
	if err != nil {
		log.Fatalf("Error::%v", err)
	}

	cloneFactoryAddress := DeployContracts(client, accountAddress, accountPrivateKey)

	fmt.Printf("Address of CloneFactory contract: %v\n\n", cloneFactoryAddress.Hex())
}