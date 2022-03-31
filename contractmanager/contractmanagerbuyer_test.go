package contractmanager

// import (
// 	//"crypto/ecdsa"
// 	//"crypto/rand"
// 	//"errors"
// 	"context"
// 	"fmt"
// 	"testing"
// 	"time"

// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/crypto"

// 	//"github.com/ethereum/go-ethereum/crypto/ecies"

// 	"gitlab.com/TitanInd/lumerin/cmd/connectionscheduler"
// 	"gitlab.com/TitanInd/lumerin/cmd/log"
// 	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
// 	"gitlab.com/TitanInd/lumerin/lumerinlib"
// 	contextlib "gitlab.com/TitanInd/lumerin/lumerinlib/context"
// )

// func TestBuyerRoutine(t *testing.T) {
// 	configPath := "../../ganacheconfig.json"
// 	// ps := msgbus.New(10, nil)
// 	l := log.New()
// 	ts := BeforeEach(configPath)
// 	var hashrateContractAddress [3]common.Address
// 	var purchasedHashrateContractAddress [3]common.Address

// 	ctxStruct := contextlib.NewContextStruct(nil, ps, l, nil, nil)
// 	mainCtx := context.WithValue(context.Background(), contextlib.ContextKey, ctxStruct)

// 	contractManagerCtx, contractManagerCancel := context.WithCancel(mainCtx)

// 	var contractManagerConfig msgbus.ContractManagerConfig
// 	contractManagerConfigID := msgbus.GetRandomIDString()

// 	contractLength := 10000

// 	defaultpooladdr := "stratum+tcp://127.0.0.1:33334/"
// 	defaultDest := msgbus.Dest{
// 		ID:     msgbus.DestID(msgbus.DEFAULT_DEST_ID),
// 		NetUrl: msgbus.DestNetUrl(defaultpooladdr),
// 	}
// 	event, err := ps.PubWait(msgbus.DestMsg, msgbus.IDString(msgbus.DEFAULT_DEST_ID), defaultDest)
// 	if err != nil {
// 		panic(fmt.Sprintf("Adding Default Dest Failed: %s", err))
// 	}
// 	if event.Err != nil {
// 		panic(fmt.Sprintf("Adding Default Dest Failed: %s", event.Err))
// 	}

// 	contractManagerConfigFile, err := LoadTestConfiguration("contract", configPath)
// 	if err != nil {
// 		panic(fmt.Sprintf("failed to load contract manager configuration:%s", err))
// 	}

// 	contractManagerConfig.Mnemonic = contractManagerConfigFile["mnemonic"].(string)
// 	contractManagerConfig.AccountIndex = int(contractManagerConfigFile["accountIndex"].(float64))
// 	contractManagerConfig.TimeThreshold = int(contractManagerConfigFile["timeThreshold"].(float64))
// 	contractManagerConfig.EthNodeAddr = contractManagerConfigFile["ethNodeAddr"].(string)
// 	contractManagerConfig.CloneFactoryAddress = ts.cloneFactoryAddress.Hex()

// 	sleepTime := 5000 // 5000 ms sleeptime in ganache
// 	if contractManagerConfig.EthNodeAddr != "ws://127.0.0.1:7545" {
// 		sleepTime = 30000 // 20000 ms on testnet
// 	}

// 	account, privateKey := hdWalletKeys(contractManagerConfig.Mnemonic, contractManagerConfig.AccountIndex+1)
// 	sellerAddress := account.Address
// 	sellerPrivateKey := privateKey
// 	fmt.Println("Seller account", sellerAddress)
// 	fmt.Println("Seller private key", sellerPrivateKey)

// 	ps.PubWait(msgbus.ContractManagerConfigMsg, contractManagerConfigID, contractManagerConfig)

// 	nodeOperator := msgbus.NodeOperator{
// 		ID:          msgbus.NodeOperatorID(msgbus.GetRandomIDString()),
// 		DefaultDest: defaultDest.ID,
// 		IsBuyer:     true,
// 	}
// 	event, err = ps.PubWait(msgbus.NodeOperatorMsg, msgbus.IDString(nodeOperator.ID), nodeOperator)
// 	if err != nil {
// 		panic(fmt.Sprintf("Adding Node Operator Failed: %s", err))
// 	}
// 	if event.Err != nil {
// 		panic(fmt.Sprintf("Adding Node Operator Failed: %s", event.Err))
// 	}

// 	// start connection scheduler look at miners
// 	cs, err := connectionscheduler.New(&mainCtx, &nodeOperator, false)
// 	if err != nil {
// 		panic(fmt.Sprintf("schedule manager failed:%s", err))
// 	}
// 	err = cs.Start()
// 	if err != nil {
// 		panic(fmt.Sprintf("schedule manager failed to start:%s", err))
// 	}

// 	var cman BuyerContractManager
// 	go newConfigMonitor(&mainCtx, contractManagerCtx, contractManagerCancel, &cman, contractManagerConfigID, &nodeOperator)
// 	err = cman.init(&contractManagerCtx, contractManagerConfigID, &nodeOperator)
// 	if err != nil {
// 		panic(fmt.Sprintf("contract manager init failed:%s", err))
// 	}

// 	// subcribe to creation events emitted by clonefactory contract
// 	cfLogs, cfSub, _ := subscribeToContractEvents(ts.ethClient, ts.cloneFactoryAddress)
// 	// create event signature to parse out creation event
// 	contractCreatedSig := []byte("contractCreated(address)")
// 	contractCreatedSigHash := crypto.Keccak256Hash(contractCreatedSig)
// 	clonefactoryContractPurchasedSig := []byte("clonefactoryContractPurchased(address)")
// 	clonefactoryContractPurchasedSigHash := crypto.Keccak256Hash(clonefactoryContractPurchasedSig)
// 	go func() {
// 		i := 0
// 		j := 0
// 		for {
// 			select {
// 			case err := <-cfSub.Err():
// 				panic(fmt.Sprintf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err))
// 			case cfLog := <-cfLogs:
// 				switch {
// 				case cfLog.Topics[0].Hex() == contractCreatedSigHash.Hex():
// 					hashrateContractAddress[i] = common.HexToAddress(cfLog.Topics[1].Hex())
// 					fmt.Printf("Address of created Hashrate Contract %d: %s\n\n", i+1, hashrateContractAddress[i].Hex())
// 					i++

// 				case cfLog.Topics[0].Hex() == clonefactoryContractPurchasedSigHash.Hex():
// 					purchasedHashrateContractAddress[j] = common.HexToAddress(cfLog.Topics[1].Hex())
// 					fmt.Printf("Address of purchased Hashrate Contract %d: %s\n\n", j+1, purchasedHashrateContractAddress[j].Hex())
// 					j++
// 				}
// 			}
// 		}
// 	}()

// 	//
// 	// test startup with 1 running contract and 1 availabe contract
// 	//
// 	CreateHashrateContract(cman.ethClient, sellerAddress, sellerPrivateKey, ts.cloneFactoryAddress, int(0), int(10), int(31), int(contractLength), cman.account)
// 	CreateHashrateContract(cman.ethClient, sellerAddress, sellerPrivateKey, ts.cloneFactoryAddress, int(0), int(10), int(41), int(contractLength), cman.account)

// 	// wait until created hashrate contract was found before continuing
// loop1:
// 	for {
// 		if hashrateContractAddress[0] != common.HexToAddress("0x0000000000000000000000000000000000000000") {
// 			break loop1
// 		}
// 	}
// 	time.Sleep(time.Millisecond * time.Duration(sleepTime/5))
// 	PurchaseHashrateContract(cman.ethClient, cman.account, cman.privateKey, ts.cloneFactoryAddress, hashrateContractAddress[0], cman.account, "stratum+tcp://127.0.0.1:3333/testrig")

// 	// wait until hashrate contract was purchased before continuing
// loop2:
// 	for {
// 		if purchasedHashrateContractAddress[0] != common.HexToAddress("0x0000000000000000000000000000000000000000") {
// 			break loop2
// 		}
// 	}
// 	// publish miners sent from seller to fulfill hashrate promised by contract
// 	miner1 := msgbus.Miner{
// 		ID:              msgbus.MinerID("MinerID01"),
// 		IP:              "IpAddress1",
// 		CurrentHashRate: 20,
// 		State:           msgbus.OnlineState,
// 	}
// 	miner2 := msgbus.Miner{
// 		ID:              msgbus.MinerID("MinerID02"),
// 		IP:              "IpAddress2",
// 		CurrentHashRate: 10,
// 		State:           msgbus.OnlineState,
// 	}
// 	ps.Pub(msgbus.MinerMsg, msgbus.IDString(miner1.ID), miner1)
// 	ps.Pub(msgbus.MinerMsg, msgbus.IDString(miner2.ID), miner2)

// 	err = cman.start()
// 	if err != nil {
// 		panic(fmt.Sprintf("contract manager failed to start:%s", err))
// 	}
// 	if err != nil {
// 		panic(fmt.Sprintf("contract manager failed to start:%s", err))
// 	}

// 	// contract manager sees existing contracts and states are correct
// 	if cman.nodeOperator.Contracts[msgbus.ContractID(hashrateContractAddress[0].Hex())] != msgbus.ContRunningState {
// 		t.Errorf("Contract 1 was not found or is not in correct state")
// 	}
// 	if _, ok := cman.nodeOperator.Contracts[msgbus.ContractID(hashrateContractAddress[1].Hex())]; ok {
// 		t.Errorf("Contract 2 was found by buyer node while in the available state")
// 	}

// 	// connection scheduler sets contract to correct miners
// 	m1, _ := ps.MinerGetWait(miner1.ID)
// 	m2, _ := ps.MinerGetWait(miner2.ID)
// 	if m1.Contract != msgbus.ContractID(hashrateContractAddress[0].Hex()) || m2.Contract != msgbus.ContractID(hashrateContractAddress[0].Hex()) {
// 		t.Errorf("Miner contracts not set correctly")
// 	}

// 	// contract manager should updated states
// 	// wait until created hashrate contract was found before continuing
// loop3:
// 	for {
// 		if hashrateContractAddress[1] != common.HexToAddress("0x0000000000000000000000000000000000000000") {
// 			break loop3
// 		}
// 	}
// 	time.Sleep(time.Millisecond * time.Duration(sleepTime/5))
// 	PurchaseHashrateContract(cman.ethClient, cman.account, cman.privateKey, ts.cloneFactoryAddress, hashrateContractAddress[1], cman.account, "stratum+tcp://127.0.0.1:3333/testrig")

// 	// wait until hashrate contract was purchased before continuing
// loop4:
// 	for {
// 		if purchasedHashrateContractAddress[1] != common.HexToAddress("0x0000000000000000000000000000000000000000") {
// 			break loop4
// 		}
// 	}
// 	time.Sleep(time.Millisecond * time.Duration(sleepTime/5))
// 	miner3 := msgbus.Miner{
// 		ID:              msgbus.MinerID("MinerID03"),
// 		IP:              "IpAddress3",
// 		CurrentHashRate: 40,
// 		State:           msgbus.OnlineState,
// 	}
// 	ps.Pub(msgbus.MinerMsg, msgbus.IDString(miner3.ID), miner3)
// 	time.Sleep(time.Millisecond * time.Duration(sleepTime))

// 	if cman.nodeOperator.Contracts[msgbus.ContractID(hashrateContractAddress[1].Hex())] != msgbus.ContRunningState {
// 		t.Errorf("Contract 2 is not in correct state")
// 	}

// 	// connection scheduler sets contracts to correct miners
// 	m1, _ = ps.MinerGetWait(miner1.ID)
// 	m2, _ = ps.MinerGetWait(miner2.ID)
// 	m3, _ := ps.MinerGetWait(miner3.ID)
// 	time.Sleep(time.Millisecond * time.Duration(sleepTime/5))
// 	if m1.Contract != msgbus.ContractID(hashrateContractAddress[0].Hex()) || m2.Contract != msgbus.ContractID(hashrateContractAddress[0].Hex()) || m3.Contract != msgbus.ContractID(hashrateContractAddress[1].Hex()) {
// 		t.Errorf("Miner contracts not set correctly")
// 	}

// 	/*
// 		//
// 		// Test early closeout from seller
// 		//
// 		CreateHashrateContract(cman.ethClient, sellerAddress, sellerPrivateKey, ts.cloneFactoryAddress, int(0), int(0), int(30), int(contractLength*10), cman.account)
// 		time.Sleep(time.Millisecond * time.Duration(sleepTime))
// 		if _,ok := cman.msg.Contracts[msgbus.ContractID(hashrateContractAddress[2].Hex())] ; ok {
// 			t.Errorf("Contract 3 was found by buyer node while in the available state")
// 		}

// 		PurchaseHashrateContract(cman.ethClient, cman.account, cman.privateKey, ts.cloneFactoryAddress, hashrateContractAddress[2], cman.account, "stratum+tcp://127.0.0.1:3333/testrig")
// 		time.Sleep(time.Millisecond * time.Duration(sleepTime))
// 		if cman.msg.Contracts[msgbus.ContractID(hashrateContractAddress[2].Hex())] != msgbus.ContRunningState {
// 			t.Errorf("Contract 3 is not in correct state")
// 		}

// 		var wg sync.WaitGroup
// 		wg.Add(1)
// 		setContractCloseOut(cman.ethClient, sellerAddress, sellerPrivateKey, hashrateContractAddress[2], &wg, &cman.currentNonce, 0)
// 		wg.Wait()
// 		time.Sleep(time.Millisecond * time.Duration(sleepTime))
// 		if _,ok := cman.msg.Contracts[msgbus.ContractID(hashrateContractAddress[2].Hex())]; ok {
// 			t.Errorf("Contract 3 did not close out correctly")
// 		}
// 	*/

// 	//
// 	// Test contract creation, purchasing, and target dest being updated while node is running
// 	//
// 	CreateHashrateContract(cman.ethClient, sellerAddress, sellerPrivateKey, ts.cloneFactoryAddress, int(0), int(10), int(100), int(contractLength), cman.account)

// loop5:
// 	for {
// 		if hashrateContractAddress[2] != common.HexToAddress("0x0000000000000000000000000000000000000000") {
// 			break loop5
// 		}
// 	}
// 	time.Sleep(time.Millisecond * time.Duration(sleepTime/5))
// 	if _, ok := cman.nodeOperator.Contracts[msgbus.ContractID(hashrateContractAddress[2].Hex())]; ok {
// 		t.Errorf("Contract 3 was found by buyer node while in the available state")
// 	}
// 	PurchaseHashrateContract(cman.ethClient, cman.account, cman.privateKey, ts.cloneFactoryAddress, hashrateContractAddress[2], cman.account, "stratum+tcp://127.0.0.1:3333/testrig")

// 	// wait until hashrate contract was purchased before continuing
// loop6:
// 	for {
// 		if purchasedHashrateContractAddress[2] != common.HexToAddress("0x0000000000000000000000000000000000000000") {
// 			break loop6
// 		}
// 	}
// 	time.Sleep(time.Millisecond * time.Duration(sleepTime/5))
// 	miner4 := msgbus.Miner{
// 		ID:              msgbus.MinerID("MinerID04"),
// 		IP:              "IpAddress4",
// 		CurrentHashRate: 100,
// 		State:           msgbus.OnlineState,
// 	}
// 	ps.Pub(msgbus.MinerMsg, msgbus.IDString(miner4.ID), miner4)
// 	time.Sleep(time.Millisecond * time.Duration(sleepTime/5))

// 	if cman.nodeOperator.Contracts[msgbus.ContractID(hashrateContractAddress[2].Hex())] != msgbus.ContRunningState {
// 		t.Errorf("Contract 3 is not in correct state")
// 	}

// 	// connection scheduler sets contracts to correct miners
// 	m1, _ = ps.MinerGetWait(miner1.ID)
// 	m2, _ = ps.MinerGetWait(miner2.ID)
// 	m3, _ = ps.MinerGetWait(miner3.ID)
// 	m4, _ := ps.MinerGetWait(miner4.ID)
// 	time.Sleep(time.Millisecond * time.Duration(sleepTime/5))
// 	if m1.Contract != msgbus.ContractID(hashrateContractAddress[0].Hex()) || m2.Contract != msgbus.ContractID(hashrateContractAddress[0].Hex()) || m3.Contract != msgbus.ContractID(hashrateContractAddress[1].Hex()) || m4.Contract != msgbus.ContractID(hashrateContractAddress[2].Hex()) {
// 		t.Errorf("Miner contracts not set correctly")
// 	}

// 	UpdateCipherText(cman.ethClient, cman.account, cman.privateKey, hashrateContractAddress[2], "stratum+tcp://127.0.0.1:3333/updated")
// 	time.Sleep(time.Millisecond * time.Duration(sleepTime*2))
// 	// check dest msg with associated contract was updated in msgbus
// 	event, err = cman.ps.GetWait(msgbus.ContractMsg, msgbus.IDString(hashrateContractAddress[2].Hex()))
// 	if err != nil {
// 		panic(fmt.Sprintf("Getting Purchased Contract Failed: %s", err))
// 	}
// 	if event.Err != nil {
// 		panic(fmt.Sprintf("Getting Purchased Contract Failed: %s", event.Err))
// 	}
// 	contractMsg := event.Data.(msgbus.Contract)
// 	event, err = cman.ps.GetWait(msgbus.DestMsg, msgbus.IDString(contractMsg.Dest))
// 	if err != nil {
// 		panic(fmt.Sprintf("Getting Dest Failed: %s", err))
// 	}
// 	if event.Err != nil {
// 		panic(fmt.Sprintf("Getting Dest Failed: %s", event.Err))
// 	}
// 	destMsg := event.Data.(msgbus.Dest)
// 	if destMsg.NetUrl != "stratum+tcp://127.0.0.1:3333/updated" {
// 		t.Errorf("Contract 3's target dest was not updated")
// 	}

// 	//
// 	// Test miners being updated below min, deleted, and set to offline
// 	//
// 	// miner 4's hashrate is updated to below min
// 	miner4.CurrentHashRate = 5
// 	ps.Set(msgbus.MinerMsg, msgbus.IDString(miner1.ID), miner1)
// 	time.Sleep(time.Millisecond * time.Duration(sleepTime*3))

// 	// miner 2 deleted
// 	ps.UnpubWait(msgbus.MinerMsg, msgbus.IDString(miner2.ID))
// 	time.Sleep(time.Millisecond * time.Duration(sleepTime*3))

// 	//
// 	// Test miners are set to offline state so running contracts should close out
// 	//
// 	miner1.State = msgbus.OfflineState
// 	ps.Set(msgbus.MinerMsg, msgbus.IDString(miner1.ID), miner1)
// 	miner3.State = msgbus.OfflineState
// 	ps.Set(msgbus.MinerMsg, msgbus.IDString(miner3.ID), miner3)
// 	miner4.State = msgbus.OfflineState
// 	ps.Set(msgbus.MinerMsg, msgbus.IDString(miner4.ID), miner4)
// 	time.Sleep(time.Millisecond * time.Duration(sleepTime*4))

// 	// check contracts map is empty now
// 	if len(cman.nodeOperator.Contracts) != 0 {
// 		t.Errorf("Contracts did not closeout after all miners were set to offline")
// 	}

// 	// connection scheduler removes contracts from miners
// 	m1, _ = ps.MinerGetWait(miner1.ID)
// 	m3, _ = ps.MinerGetWait(miner3.ID)
// 	m4, _ = ps.MinerGetWait(miner4.ID)
// 	if m1.Contract != "" || m3.Contract != "" || m4.Contract != "" {
// 		t.Errorf("Miner contracts not removed after being closed out")
// 	}

// 	//
// 	// test contract manager config updated
// 	//
// 	contractManagerConfig.AccountIndex = 2
// 	ps.SetWait(msgbus.ContractManagerConfigMsg, contractManagerConfigID, contractManagerConfig)
// 	time.Sleep(time.Second * 3)
// 	newAccount, _ := hdWalletKeys(contractManagerConfig.Mnemonic, contractManagerConfig.AccountIndex)
// 	if cman.account != newAccount.Address {
// 		t.Errorf("Contract manager's configuration was not updated after msgbus update")
// 	}
// }
