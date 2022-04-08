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
	accountAddress := common.HexToAddress("0xE5DE0c1fF501c7c31C4CC895a4e9f8D047E8A85b")
	accountPrivateKey := "03d2afa85bb362a6b89b01daa9f811a99f57b4f79db77ac6ed543ab926de108d"
	gethNodeAddress := "ws://127.0.0.1:7545"

	client, err := contractmanager.SetUpClient(gethNodeAddress, accountAddress)
	if err != nil {
		log.Fatalf("Error::%v", err)
	}

	//
	// Deploy new CloneFactory Contract
	//
	cloneFactoryAddress := contractmanager.DeployContracts(client, accountAddress, accountPrivateKey)

	t.Logf("Address of CloneFactory contract: %v\n\n", cloneFactoryAddress.Hex())

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

	time.Sleep(time.Second * 30)
	//
	// Purchase hashrate contract
	//
	poolUrl := "stratum.slushpool.com:3333"
	// poolUrl := "stratum+tcp://stratum.slushpool.com:3333"
	contractmanager.PurchaseHashrateContract(client, accountAddress, accountPrivateKey, cloneFactoryAddress, hashrateContractAddress, accountAddress, poolUrl)

	// hang until signal interrupt
	<-sigInt
}
