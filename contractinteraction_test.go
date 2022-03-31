package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"gitlab.com/TitanInd/hashrouter/contractmanager"
)

func TestContractInteraction(t *testing.T) {
	sigInt := make(chan os.Signal, 1)
	signal.Notify(sigInt, os.Interrupt)
	
	//
	// Create connection to geth node
	//
	accountAddress := common.HexToAddress("0x3f573AAAb5cf414ECc3A763cb90a8233870769A1")
	accountPrivateKey := "13bf255f212342e388d9ff7fcd27ff56823a72a45c30eceabdfcb2da44842c1a"
	gethNodeAddress := "ws://127.0.0.1:7545"

	client, err := contractmanager.SetUpClient(gethNodeAddress, accountAddress)
	if err != nil {
		log.Fatalf("Error::%v", err)
	}

	//
	// Deploy new CloneFactory Contract
	//
	cloneFactoryAddress := contractmanager.DeployContracts(client, accountAddress, accountPrivateKey)

	fmt.Printf("Address of CloneFactory contract: %v\n\n", cloneFactoryAddress.Hex())

	//
	// Create hashrate contract
	//
	var hashrateContractAddress common.Address
	price := 0
	limit := 10
	speed := 100
	length := 100

	contractmanager.CreateHashrateContract(client, accountAddress, accountPrivateKey, cloneFactoryAddress, price, limit, speed, length, cloneFactoryAddress)

	// subcribe to creation events emitted by clonefactory contract
	cfLogs, cfSub, _ := contractmanager.SubscribeToContractEvents(client, cloneFactoryAddress)
	// create event signature to parse out creation event
	contractCreatedSig := []byte("contractCreated(address,string)")
	contractCreatedSigHash := crypto.Keccak256Hash(contractCreatedSig)
	loop1:
	for {
		select {
		case err := <-cfSub.Err():
			log.Fatalf("Error::%v", err)
		case cfLog := <-cfLogs:
			if cfLog.Topics[0].Hex() == contractCreatedSigHash.Hex() {
				hashrateContractAddress = common.HexToAddress(cfLog.Topics[1].Hex())
				fmt.Printf("Address of created Hashrate Contract: %v\n\n", hashrateContractAddress.Hex())
				break loop1
			}
		}
	}

	//
	// Run proxy node
	//	
	os.Args[0] = "Test Contract Interaction"
	os.Args[1] = "-contract.addr=" + hashrateContractAddress.Hex()
	os.Args[2] = "-ethNode.addr=" + gethNodeAddress
	os.Args[3] = "-stratum.addr=" + "127.0.0.1:9332"
	os.Args[4] = "-pool.addr=" + "mining.dev.pool.titan.io:4242"

	go main()

	time.Sleep(time.Second*5)
	//
	// Purchase hashrate contract
	//
	poolUrl := "mining.dev.pool.titan.io:4242?test=test"
	contractmanager.PurchaseHashrateContract(client, accountAddress, accountPrivateKey, cloneFactoryAddress, hashrateContractAddress, accountAddress, poolUrl)
	
	// hang until signal interrupt
	<-sigInt
}