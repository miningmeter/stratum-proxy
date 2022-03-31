package contractmanager

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"sync"

	//"encoding/hex"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	//"github.com/ethereum/go-ethereum/crypto/ecies"
	"github.com/ethereum/go-ethereum/ethclient"

	i "gitlab.com/TitanInd/hashrouter/interfaces"
	"gitlab.com/TitanInd/hashrouter/lumerinlib/implementation"
	// "gitlab.com/TitanInd/lumerin/cmd/log"
	// "gitlab.com/TitanInd/lumerin/cmd/msgbus"
	// "gitlab.com/TitanInd/lumerin/lumerinlib"
	// "gitlab.com/TitanInd/lumerin/lumerinlib/implementation"
)

const (
	AvailableState uint8 = 0
	RunningState   uint8 = 1
)

const HASHRATE_TOLERANCE = .10

type hashrateContractValues struct {
	State                  uint8
	Price                  int
	Limit                  int
	Speed                  int
	Length                 int
	StartingBlockTimestamp int
	Buyer                  common.Address
	Seller                 common.Address
}

type nonce struct {
	mutex sync.Mutex
	nonce uint64
}

//TODO: replace logging
type SellerContractManager struct {
	ps                  i.IEventManager
	l                   *log.Logger
	ethClient           *ethclient.Client
	cloneFactoryAddress common.Address
	account             common.Address
	privateKey          string
	claimFunds          bool
	currentNonce        nonce
	nodeOperator        NodeOperator
	ctx                 context.Context
}

func (s *SellerContractManager) SetLogger(l *log.Logger) {
	s.l = l
}

type MinerState string

const (
	ContAvailableState string = "AvailableState"
	ContRunningState   string = "RunningState"
)

// Do we still need this with the config package in place?

type ConfigInfo struct {
	ID           string
	DefaultDest  string
	NodeOperator string
}

type ContractManagerConfig struct {
	ID                  string
	Mnemonic            string
	AccountIndex        int
	EthNodeAddr         string
	ClaimFunds          bool
	TimeThreshold       int
	CloneFactoryAddress string
	LumerinTokenAddress string
	ValidatorAddress    string
	ProxyAddress        string
}

type NodeOperator struct {
	ID                     string
	IsBuyer                bool
	DefaultDest            string
	EthereumAccount        string
	TotalAvailableHashRate int
	UnusedHashRate         int
	Contracts              map[string]string
}

type Contract struct {
	IsSeller               bool
	ID                     string
	State                  string
	Buyer                  string
	Price                  int
	Limit                  int
	Speed                  int
	Length                 int
	StartingBlockTimestamp int
	Dest                   string
}

type Dest struct {
	ID     string
	NetUrl string
}

const (
	NoMsg                    string = "NoMsg"
	ConfigMsg                string = "ConfigMsg"
	ContractManagerConfigMsg string = "ContractManagerConfigMsg"
	DestMsg                  string = "DestMsg"
	NodeOperatorMsg          string = "NodeOperatorMsg"
	ContractMsg              string = "ContractMsg"
	MinerMsg                 string = "MinerMsg"
	ConnectionMsg            string = "ConnectionMsg"
	LogMsg                   string = "LogMsg"
)

type ContractMsgHandler struct{}

func (c *ContractMsgHandler) Update(eventState interface{}) {

}

var contractMsgHandler = &ContractMsgHandler{}

func Run(ctx *context.Context, contractManager *SellerContractManager, eventManager i.IEventManager, contractAddr string, ethNodeAddr string) (err error) {
	contractManagerCtx, _ := context.WithCancel(*ctx)

	err = contractManager.init(&contractManagerCtx, ethNodeAddr, eventManager)
	if err != nil {
		return err
	}
	err = contractManager.start(contractAddr)
	if err != nil {
		return err
	}

	return err
}

func (seller *SellerContractManager) init(ctx *context.Context, ethNodeAddr string, eventManager i.IEventManager) (err error) {
	// func (seller *SellerContractManager) init(ctx *context.Context, contractManagerConfigID IDString, nodeOperatorMsg *NodeOperator) (err error) {
	seller.ctx = *ctx
	// cs := contextlib.GetContextStruct(seller.ctx)
	//seller.ctx.Value(contextlib.ContextKey).(*contextlib.ContextStruct)
	//TODO: REWRITE
	seller.ps = eventManager
	// seller.l = cs.Log

	// event, err := seller.ps.GetWait(ContractManagerConfigMsg, contractManagerConfigID)
	// if err != nil {
	// 	return err
	// }
	// contractManagerConfig := event.Data.(ContractManagerConfig)
	// seller.claimFunds = contractManagerConfig.ClaimFunds
	// ethNodeAddr := contractManagerConfig.EthNodeAddr
	// mnemonic := contractManagerConfig.Mnemonic
	// accountIndex := contractManagerConfig.AccountIndex

	// account, privateKey := hdWalletKeys(mnemonic, accountIndex)
	// seller.account = account.Address
	// seller.privateKey = privateKey

	var client *ethclient.Client
	client, err = SetUpClient(ethNodeAddr, seller.account)
	if err != nil {
		return err
	}
	seller.ethClient = client
	// seller.cloneFactoryAddress = common.HexToAddress(contractManagerConfig.CloneFactoryAddress)

	// seller.nodeOperator = *nodeOperatorMsg
	seller.nodeOperator.EthereumAccount = seller.account.Hex()

	if seller.nodeOperator.Contracts == nil {
		seller.nodeOperator.Contracts = make(map[string]string)
	}

	return err
}

func (seller *SellerContractManager) start(addr string) (err error) {
	// err = seller.setupExistingContracts()
	// if err != nil {
	// 	return err
	// }

	// routine for listensing to contract creation events that will update seller msg with new contracts and load new contract onto msgbus
	// cfLogs, cfSub, err := SubscribeToContractEvents(seller.ethClient, seller.cloneFactoryAddress)
	// if err != nil {
	// 	return err
	// }
	// go seller.watchContractCreation(cfLogs, cfSub)

	// routine starts routines for seller's contracts that monitors contract purchase, close, and cancel events
	// go func() {
	// start routines for existing contracts
	// for addr := range seller.nodeOperator.Contracts {
	//TODO: get addr from josh
	hrLogs, hrSub, err := SubscribeToContractEvents(seller.ethClient, common.HexToAddress(addr))
	if err != nil {
		seller.l.Fatalf("Panic: %v", fmt.Sprintf("Failed to subscribe to events on hashrate contract %v, Error::%v", addr, err))
	}
	go seller.watchHashrateContract(addr, hrLogs, hrSub)
	// }

	// // monitor new contracts getting created and start hashrate conrtract monitor routine when they are created
	// contractEventChan := NewEventChan()
	// _, err = seller.ps.Sub(ContractMsg, "", contractEventChan)
	// if err != nil {
	// 	seller.l.Printf("panic", fmt.Sprintf("Failed to subscribe to contract events on msgbus, Fileline::%v, Error::", lumerinlib.FileLine()), err)
	// }
	// for {
	// 	select {
	// 	case <-seller.ctx.Done():
	// 		seller.l.Printf(log.LevelInfo, "Cancelling current contract manager context: cancelling start routine")
	// 		return
	// 	case event := <-contractEventChan:
	// 		if event.EventType == PublishEvent {
	// 			newContract := event.Data.(Contract)
	// 			if newContract.State == ContAvailableState {
	// 				addr := common.HexToAddress(string(newContract.ID))
	// 				hrLogs, hrSub, err := SubscribeToContractEvents(seller.ethClient, addr)
	// 				if err != nil {
	// 					seller.l.Printf("panic", fmt.Sprintf("Failed to subscribe to events on hashrate contract %v, Fileline::%v, Error::", newContract.ID, lumerinlib.FileLine()), err)
	// 				}
	// 				go seller.watchHashrateContract(ContractID(addr.Hex()), hrLogs, hrSub)
	// 			}
	// 		}
	// 	}
	// }
	// }()
	return err
}

// func (seller *SellerContractManager) setupExistingContracts() (err error) {
// 	var contractValues []hashrateContractValues
// 	var contractMsgs []Contract

// 	sellerContracts, err := seller.readContracts()
// 	if err != nil {
// 		return err
// 	}
// 	seller.l.Printf(log.LevelInfo, "Existing Buyer Contracts: %v", sellerContracts)

// 	for i := range sellerContracts {
// 		id := ContractID(sellerContracts[i].Hex())
// 		if _, ok := seller.nodeOperator.Contracts[id]; !ok {
// 			contract, err := readHashrateContract(seller.ethClient, sellerContracts[i])
// 			if err != nil {
// 				return err
// 			}
// 			contractValues = append(contractValues, contract)
// 			contractMsgs = append(contractMsgs, createContractMsg(sellerContracts[i], contractValues[i], true))

// 			seller.nodeOperator.Contracts[ContractID(sellerContracts[i].Hex())] = ContAvailableState

// 			if contractValues[i].State == RunningState {
// 				seller.nodeOperator.Contracts[ContractID(sellerContracts[i].Hex())] = ContRunningState

// 				// get existing dests in msgbus to see if contract's dest already exists
// 				event, err := seller.ps.GetWait(DestMsg, "")
// 				if err != nil {
// 					seller.l.Printf("panic", "Getting existing dests Failed: %v", err)
// 				}
// 				existingDests := event.Data.(IDIndex)

// 				destUrl, err := readDestUrl(seller.ethClient, sellerContracts[i], seller.privateKey)
// 				if err != nil {
// 					seller.l.Printf("panic", fmt.Sprintf("Reading dest url failed, Fileline::%v, Error::", lumerinlib.FileLine()), err)
// 				}

// 				// if msgbus has dest with same target address, use that as contract msg dest
// 				for _, v := range existingDests {
// 					existingDest, err := seller.ps.DestGetWait(DestID(v))
// 					if err != nil {
// 						seller.l.Printf("panic", "Getting existing dest Failed: %v", err)
// 					}
// 					if existingDest.NetUrl == DestNetUrl(destUrl) {
// 						contractMsgs[i].Dest = DestID(v)
// 					}
// 				}

// 				// msgbus does not have dest with that target address
// 				if contractMsgs[i].Dest == "" {
// 					destMsg := Dest{
// 						ID:     DestID(GetRandomIDString()),
// 						NetUrl: DestNetUrl(destUrl),
// 					}
// 					seller.ps.PubWait(DestMsg, IDString(destMsg.ID), destMsg)

// 					contractMsgs[i].Dest = destMsg.ID
// 				}
// 			}

// 			seller.ps.PubWait(ContractMsg, IDString(contractMsgs[i].ID), contractMsgs[i])
// 		}
// 	}

// 	seller.ps.SetWait(NodeOperatorMsg, IDString(seller.nodeOperator.ID), seller.nodeOperator)

// 	return err
// }

// func (seller *SellerContractManager) readContracts() ([]common.Address, error) {
// 	var sellerContractAddresses []common.Address
// 	var hashrateContractInstance *implementation.Implementation
// 	var hashrateContractSeller common.Address

// 	instance, err := clonefactory.NewClonefactory(seller.cloneFactoryAddress, seller.ethClient)
// 	if err != nil {
// 		seller.l.Printf(log.LevelError, fmt.Sprintf("Funcname::%v, Fileline::%v, Error::", lumerinlib.Funcname(), lumerinlib.FileLine()), err)
// 		return sellerContractAddresses, err
// 	}

// 	hashrateContractAddresses, err := instance.GetContractList(&bind.CallOpts{})
// 	if err != nil {
// 		seller.l.Printf(log.LevelError, fmt.Sprintf("Funcname::%v, Fileline::%v, Error::", lumerinlib.Funcname(), lumerinlib.FileLine()), err)
// 		return sellerContractAddresses, err
// 	}

// 	// parse existing hashrate contracts for ones that belong to seller
// 	for i := range hashrateContractAddresses {
// 		hashrateContractInstance, err = implementation.NewImplementation(hashrateContractAddresses[i], seller.ethClient)
// 		if err != nil {
// 			seller.l.Printf(log.LevelError, fmt.Sprintf("Funcname::%v, Fileline::%v, Error::", lumerinlib.Funcname(), lumerinlib.FileLine()), err)
// 			return sellerContractAddresses, err
// 		}
// 		hashrateContractSeller, err = hashrateContractInstance.Seller(nil)
// 		if err != nil {
// 			seller.l.Printf(log.LevelError, fmt.Sprintf("Funcname::%v, Fileline::%v, Error::", lumerinlib.Funcname(), lumerinlib.FileLine()), err)
// 			return sellerContractAddresses, err
// 		}
// 		if hashrateContractSeller == seller.account {
// 			sellerContractAddresses = append(sellerContractAddresses, hashrateContractAddresses[i])
// 		}
// 	}

// 	return sellerContractAddresses, err
// }

// func (seller *SellerContractManager) watchContractCreation(cfLogs chan types.Log, cfSub ethereum.Subscription) {
// 	defer close(cfLogs)
// 	defer cfSub.Unsubscribe()

// 	// create event signature to parse out creation event
// 	contractCreatedSig := []byte("contractCreated(address)")
// 	contractCreatedSigHash := crypto.Keccak256Hash(contractCreatedSig)
// 	for {
// 		select {
// 		case err := <-cfSub.Err():
// 			seller.l.Printf("panic", fmt.Sprintf("Funcname::%v, Fileline::%v, Error::", lumerinlib.Funcname(), lumerinlib.FileLine()), err)
// 		case <-seller.ctx.Done():
// 			seller.l.Printf(log.LevelInfo, "Cancelling current contract manager context: cancelling watchContractCreation go routine")
// 			return
// 		case cfLog := <-cfLogs:
// 			if cfLog.Topics[0].Hex() == contractCreatedSigHash.Hex() {
// 				address := common.HexToAddress(cfLog.Topics[1].Hex())
// 				// check if contract created belongs to seller
// 				hashrateContractInstance, err := implementation.NewImplementation(address, seller.ethClient)
// 				if err != nil {
// 					seller.l.Printf("panic", fmt.Sprintf("Funcname::%v, Fileline::%v, Error::", lumerinlib.Funcname(), lumerinlib.FileLine()), err)
// 				}
// 				hashrateContractSeller, err := hashrateContractInstance.Seller(nil)
// 				if err != nil {
// 					seller.l.Printf("panic", fmt.Sprintf("Funcname::%v, Fileline::%v, Error::", lumerinlib.Funcname(), lumerinlib.FileLine()), err)
// 				}
// 				if hashrateContractSeller == seller.account {
// 					seller.l.Printf(log.LevelInfo, "Address of created Hashrate Contract: %v\n\n", address.Hex())

// 					createdContractValues, err := readHashrateContract(seller.ethClient, address)
// 					if err != nil {
// 						seller.l.Printf("panic", fmt.Sprintf("Reading hashrate contract failed, Fileline::%v, Error::", lumerinlib.FileLine()), err)
// 					}
// 					createdContractMsg := createContractMsg(address, createdContractValues, true)
// 					seller.ps.PubWait(ContractMsg, IDString(address.Hex()), createdContractMsg)

// 					seller.nodeOperator.Contracts[ContractID(address.Hex())] = ContAvailableState

// 					seller.ps.SetWait(NodeOperatorMsg, IDString(seller.nodeOperator.ID), seller.nodeOperator)
// 				}
// 			}
// 		}
// 	}
// }

func (seller *SellerContractManager) watchHashrateContract(addr string, hrLogs chan types.Log, hrSub ethereum.Subscription) {
	// contractEventChan := NewEventChan()

	// check if contract is already in the running state and needs to be monitored for closeout
	// event, err := seller.ps.GetWait(ContractMsg, IDString(addr))
	// if err != nil {
	// 	seller.l.Printf("panic", "Getting Hashrate Contract Failed: %v", err)
	// }
	// if event.Err != nil {
	// 	seller.l.Printf("panic", "Getting Hashrate Contract Failed: %v", event.Err)
	// }
	// hashrateContractMsg := event.Data.(Contract)
	// if hashrateContractMsg.State == ContRunningState {
	// go seller.closeOutMonitor(hashrateContractMsg)
	// }

	// create event signatures to parse out which event was being emitted from hashrate contract
	contractPurchasedSig := []byte("contractPurchased(address)")
	// contractClosedSig := []byte("contractClosed()")
	// purchaseInfoUpdatedSig := []byte("purchaseInfoUpdated()")
	// cipherTextUpdatedSig := []byte("cipherTextUpdated(string)")
	contractPurchasedSigHash := crypto.Keccak256Hash(contractPurchasedSig)
	// contractClosedSigHash := crypto.Keccak256Hash(contractClosedSig)
	// purchaseInfoUpdatedSigHash := crypto.Keccak256Hash(purchaseInfoUpdatedSig)
	// cipherTextUpdatedSigHash := crypto.Keccak256Hash(cipherTextUpdatedSig)

	// routine monitoring and acting upon events emmited by hashrate contract
	go func() {
		defer close(hrLogs)
		defer hrSub.Unsubscribe()
		for {
			select {
			case err := <-hrSub.Err():
				seller.l.Printf("Panic %v", fmt.Sprintf("Funcname::%v, Error::%v", "watchHashrateContract", err))
			case <-seller.ctx.Done():
				seller.l.Printf("Info %v", "Cancelling current contract manager context: cancelling watchHashrateContract go routine")
				return
			case hLog := <-hrLogs:
				switch hLog.Topics[0].Hex() {
				case contractPurchasedSigHash.Hex():
					buyer := common.HexToAddress(hLog.Topics[1].Hex())
					seller.l.Printf("Info %v purchased Hashrate Contract: %v\n\n", buyer.Hex(), addr)

					destUrl, err := readDestUrl(seller.ethClient, common.HexToAddress(string(addr)), seller.privateKey)
					if err != nil {
						seller.l.Printf("Panic %v Reading dest url failed, Error::%v", "\r\n", err)
					}
					destMsg := Dest{
						ID:     GetRandomIDString(),
						NetUrl: destUrl,
					}

					seller.ps.GoDispatch(DestMsg, destMsg)

					// event, err := seller.ps.GetWait(ContractMsg, string(addr))
					if err != nil {
						seller.l.Printf("panic Getting Purchased Contract Failed: %v", err)
					}
					// if event.Err != nil {
					// 	seller.l.Printf("panic", "Getting Purchased Contract Failed: %v", event.Err)
					// }
					contractValues, err := readHashrateContract(seller.ethClient, common.HexToAddress(string(addr)))
					if err != nil {
						seller.l.Printf("panic Reading hashrate contract failed, Error::%v", err)
					}
					contractMsg := createContractMsg(common.HexToAddress(string(addr)), contractValues, true)
					contractMsg.Dest = destMsg.ID
					contractMsg.State = ContRunningState
					contractMsg.Buyer = string(buyer.Hex())
					// seller.ps.SetWait(ContractMsg, string(addr), contractMsg)

					seller.nodeOperator.Contracts[addr] = ContRunningState
					// seller.ps.SetWait(NodeOperatorMsg, string(seller.nodeOperator.ID), seller.nodeOperator)

					// case cipherTextUpdatedSigHash.Hex():
					// 	seller.l.Printf(log.LevelInfo, "Hashrate Contract %v Cipher Text Updated \n\n", addr)

					// 	event, err := seller.ps.GetWait(ContractMsg, IDString(addr))
					// 	if err != nil {
					// 		seller.l.Printf("panic", "Getting Purchased Contract Failed: %v", err)
					// 	}
					// 	if event.Err != nil {
					// 		seller.l.Printf("panic", "Getting Purchased Contract Failed: %v", event.Err)
					// 	}
					// 	contractMsg := event.Data.(Contract)
					// 	event, err = seller.ps.GetWait(DestMsg, IDString(contractMsg.Dest))
					// 	if err != nil {
					// 		seller.l.Printf("panic", "Getting Dest Failed: %v", err)
					// 	}
					// 	if event.Err != nil {
					// 		seller.l.Printf("panic", "Getting Dest Failed: %v", event.Err)
					// 	}
					// 	destMsg := event.Data.(Dest)

					// 	destUrl, err := readDestUrl(seller.ethClient, common.HexToAddress(string(addr)), seller.privateKey)
					// 	if err != nil {
					// 		seller.l.Printf("panic", fmt.Sprintf("Reading dest url failed, Fileline::%v, Error::", lumerinlib.FileLine()), err)
					// 	}
					// 	destMsg.NetUrl = DestNetUrl(destUrl)
					// 	seller.ps.SetWait(DestMsg, IDString(destMsg.ID), destMsg)

					// case contractClosedSigHash.Hex():
					// 	seller.l.Printf(log.LevelInfo, "Hashrate Contract %v Closed \n\n", addr)

					// 	event, err := seller.ps.GetWait(ContractMsg, IDString(addr))
					// 	if err != nil {
					// 		seller.l.Printf("panic", "Getting Purchased Contract Failed: %v", err)
					// 	}
					// 	if event.Err != nil {
					// 		seller.l.Printf("panic", "Getting Purchased Contract Failed: %v", event.Err)
					// 	}
					// 	contractMsg := event.Data.(Contract)
					// 	contractMsg.State = ContAvailableState
					// 	contractMsg.Buyer = ""
					// 	seller.ps.SetWait(ContractMsg, IDString(contractMsg.ID), contractMsg)

					// 	seller.nodeOperator.Contracts[addr] = ContAvailableState
					// 	seller.ps.SetWait(NodeOperatorMsg, IDString(seller.nodeOperator.ID), seller.nodeOperator)

					// case purchaseInfoUpdatedSigHash.Hex():
					// 	seller.l.Printf(log.LevelInfo, "Hashrate Contract %v Purchase Info Updated \n\n", addr)

					// 	event, err := seller.ps.GetWait(ContractMsg, IDString(addr))
					// 	if err != nil {
					// 		seller.l.Printf("panic", "Getting Purchased Contract Failed: %v", err)
					// 	}
					// 	if event.Err != nil {
					// 		seller.l.Printf("panic", "Getting Purchased Contract Failed: %v", event.Err)
					// 	}
					// 	contractMsg := event.Data.(Contract)

					// 	updatedContractValues, err := readHashrateContract(seller.ethClient, common.HexToAddress(string(addr)))
					// 	if err != nil {
					// 		seller.l.Printf("panic", fmt.Sprintf("Reading hashrate contract failed, Fileline::%v, Error::", lumerinlib.FileLine()), err)
					// 	}
					// 	updateContractMsg(&contractMsg, updatedContractValues)
					// 	seller.ps.SetWait(ContractMsg, IDString(contractMsg.ID), contractMsg)
				}
			}
		}
	}()

	// once contract is running, closeout after length of contract has passed if it was not closed out early
	// for {
	// 	select {
	// 	case <-seller.ctx.Done():
	// 		seller.l.Printf(log.LevelInfo, "Cancelling current contract manager context: cancelling watchHashrateContract go routine")
	// 		return
	// 	case event := <-contractEventChan:
	// 		if event.EventType == UpdateEvent {
	// 			runningContractMsg := event.Data.(Contract)
	// 			if runningContractMsg.State == ContRunningState {
	// 				// run routine for each running contract to check if contract length has passed and contract should be closed out
	// 				go seller.closeOutMonitor(runningContractMsg)
	// 			}
	// 		}
	// 	}
	// }
}

// func (seller *SellerContractManager) closeOutMonitor(contractMsg Contract) {
// 	contractFinishedTimestamp := contractMsg.StartingBlockTimestamp + contractMsg.Length

// 	// subscribe to latest block headers
// 	headers := make(chan *types.Header)
// 	sub, err := seller.ethClient.SubscribeNewHead(context.Background(), headers)
// 	if err != nil {
// 		seller.l.Printf("panic", fmt.Sprintf("Funcname::%v, Fileline::%v, Error::", lumerinlib.Funcname(), lumerinlib.FileLine()), err)
// 	}
// 	defer close(headers)
// 	defer sub.Unsubscribe()

// loop:
// 	for {
// 		select {
// 		case err := <-sub.Err():
// 			seller.l.Printf("panic", fmt.Sprintf("Funcname::%v, Fileline::%v, Error::", lumerinlib.Funcname(), lumerinlib.FileLine()), err)
// 		case <-seller.ctx.Done():
// 			seller.l.Printf(log.LevelInfo, "Cancelling current contract manager context: cancelling closeout monitor go routine")
// 			return
// 		case header := <-headers:
// 			// get latest block from header
// 			block, err := seller.ethClient.BlockByHash(context.Background(), header.Hash())
// 			if err != nil {
// 				seller.l.Printf("panic", fmt.Sprintf("Funcname::%v, Fileline::%v, Error::", lumerinlib.Funcname(), lumerinlib.FileLine()), err)
// 			}

// 			// check if contract length has passed
// 			if block.Time() >= uint64(contractFinishedTimestamp) {
// 				var closeOutType uint

// 				// seller only wants to closeout
// 				closeOutType = 2
// 				// seller wants to claim funds with closeout
// 				if seller.claimFunds {
// 					closeOutType = 3
// 				}

// 				// if contract was not already closed early, close out here
// 				contractValues, err := readHashrateContract(seller.ethClient, common.HexToAddress(string(contractMsg.ID)))
// 				if err != nil {
// 					seller.l.Printf("panic", fmt.Sprintf("Reading hashrate contract failed, Fileline::%v, Error::", lumerinlib.FileLine()), err)
// 				}
// 				if contractValues.State == RunningState {
// 					var wg sync.WaitGroup
// 					wg.Add(1)
// 					err = setContractCloseOut(seller.ethClient, seller.account, seller.privateKey, common.HexToAddress(string(contractMsg.ID)), &wg, &seller.currentNonce, closeOutType)
// 					if err != nil {
// 						seller.l.Printf("panic", fmt.Sprintf("Contract Close Out failed, Fileline::%v, Error::", lumerinlib.FileLine()), err)
// 					}
// 					wg.Wait()
// 				}
// 				break loop
// 			}
// 		}
// 	}
// }

// func hdWalletKeys(mnemonic string, accountIndex int) (accounts.Account, string) {
// 	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
// 	if err != nil {
// 		panic(fmt.Sprintf("Funcname::%v, Fileline::%v, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err))
// 	}
// 	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/" + fmt.Sprint(accountIndex))
// 	account, err := wallet.Derive(path, false)
// 	if err != nil {
// 		panic(fmt.Sprintf("Funcname::%v, Fileline::%v, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err))
// 	}
// 	privateKey, err := wallet.PrivateKeyHex(account)
// 	if err != nil {
// 		panic(fmt.Sprintf("Funcname::%v, Fileline::%v, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err))
// 	}
// 	return account, privateKey
// }

func SetUpClient(clientAddress string, contractManagerAccount common.Address) (client *ethclient.Client, err error) {
	client, err = ethclient.Dial(clientAddress)
	if err != nil {
		//fmt.Printf("Funcname::%v, Fileline::%v, Error::%v\n", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
		return client, err
	}

	//fmt.Printf("Connected to rpc client at %v\n", clientAddress)

	// var balance *big.Int
	// balance, err = client.BalanceAt(context.Background(), contractManagerAccount, nil)
	// if err != nil {
	// 	//fmt.Printf("Funcname::%v, Fileline::%v, Error::%v\n", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
	// 	return client, err
	// }
	// fbalance := new(big.Float)
	// fbalance.SetString(balance.String())
	// ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	// fmt.Println("Balance of contract manager account:", ethValue, "ETH")

	return client, err
}

func SubscribeToContractEvents(client *ethclient.Client, contractAddress common.Address) (chan types.Log, ethereum.Subscription, error) {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		fmt.Printf("Funcname::%v, Error::%v\n", "SubscribeToContractEvents", err)
		return logs, sub, err
	}

	return logs, sub, err
}

func readHashrateContract(client *ethclient.Client, contractAddress common.Address) (hashrateContractValues, error) {
	var contractValues hashrateContractValues

	instance, err := implementation.NewImplementation(contractAddress, client)
	if err != nil {
		//fmt.Printf("Funcname::%v, Fileline::%v, Error::%v\n", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
		return contractValues, err
	}

	state, price, limit, speed, length, startingBlockTimestamp, buyer, seller, _, err := instance.GetPublicVariables(&bind.CallOpts{})
	if err != nil {
		//fmt.Printf("Funcname::%v, Fileline::%v, Error::%v\n", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
		return contractValues, err
	}
	contractValues.State = state
	contractValues.Price = int(price.Int64())
	contractValues.Limit = int(limit.Int64())
	contractValues.Speed = int(speed.Int64())
	contractValues.Length = int(length.Int64())
	contractValues.StartingBlockTimestamp = int(startingBlockTimestamp.Int64())
	contractValues.Buyer = buyer
	contractValues.Seller = seller

	return contractValues, err
}

func readDestUrl(client *ethclient.Client, contractAddress common.Address, privateKeyString string) (string, error) {
	instance, err := implementation.NewImplementation(contractAddress, client)
	if err != nil {
		//fmt.Printf("Funcname::%v, Fileline::%v, Error::%v\n", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
		return "", err
	}

	//fmt.Printf("Getting Dest url from contract %v\n\n", contractAddress)

	encryptedDestUrl, err := instance.EncryptedPoolData(nil)
	if err != nil {
		//fmt.Printf("Funcname::%v, Fileline::%v, Error::%v\n", "", err)
		return "", err
	}

	/*
		// Decryption Logic
		destUrlBytes,_ := hex.DecodeString(encryptedDestUrl)
		privateKey, err := crypto.HexToECDSA(privateKeyString)
		if err != nil {
			log.Printf("Funcname::%v, Fileline::%v, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
			return "", err
		}
		privateKeyECIES := ecies.ImportECDSA(privateKey)
		decryptedDestUrlBytes, err := privateKeyECIES.Decrypt(destUrlBytes, nil, nil)
		if err != nil {
			log.Printf("Funcname::%v, Fileline::%v, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
			return "", err
		}
		decryptedDestUrl := string(decryptedDestUrlBytes)

		return decryptedDestUrl, err
	*/
	return encryptedDestUrl, err
}

//UNUSED
// func setContractCloseOut(client *ethclient.Client, fromAddress common.Address, privateKeyString string, contractAddress common.Address, wg *sync.WaitGroup, currentNonce *nonce, closeOutType uint) error {
// 	defer wg.Done()
// 	defer currentNonce.mutex.Unlock()

// 	currentNonce.mutex.Lock()

// 	instance, err := implementation.NewImplementation(contractAddress, client)
// 	if err != nil {
// 		//fmt.Printf("Funcname::%v, Fileline::%v, Error::%v\n", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
// 		return err
// 	}

// 	privateKey, err := crypto.HexToECDSA(privateKeyString)
// 	if err != nil {
// 		//fmt.Printf("Funcname::%v, Fileline::%v, Error::%v\n", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
// 		return err
// 	}

// 	chainId, err := client.ChainID(context.Background())
// 	if err != nil {
// 		//fmt.Printf("Funcname::%v, Fileline::%v, Error::%v\n", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
// 		return err
// 	}

// 	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
// 	if err != nil {
// 		//fmt.Printf("Funcname::%v, Fileline::%v, Error::%v\n", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
// 		return err
// 	}

// 	gasPrice, err := client.SuggestGasPrice(context.Background())
// 	if err != nil {
// 		//fmt.Printf("Funcname::%v, Fileline::%v, Error::%v\n", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
// 		return err
// 	}
// 	auth.GasPrice = gasPrice
// 	auth.GasLimit = uint64(3000000) // in units
// 	auth.Value = big.NewInt(0)      // in wei

// 	currentNonce.nonce, err = client.PendingNonceAt(context.Background(), fromAddress)
// 	if err != nil {
// 		//fmt.Printf("Funcname::%v, Fileline::%v, Error::%v\n", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
// 		return err
// 	}
// 	auth.Nonce = big.NewInt(int64(currentNonce.nonce))

// 	tx, err := instance.SetContractCloseOut(auth, big.NewInt(int64(closeOutType)))
// 	if err != nil {
// 		fmt.Printf("Funcname::%v,  Error::%v\n", "setContractCloseOut", lumerinlib.FileLine(), err)
// 		return err
// 	}

// 	//fmt.Printf("tx sent: %v\n\n", tx.Hash().Hex())
// 	fmt.Println("Closing Out Contract: ", contractAddress)
// 	return err
// }

//TODO: repurpose for POC
func createContractMsg(contractAddress common.Address, contractValues hashrateContractValues, isSeller bool) Contract {
	convertToMsgBusState := map[uint8]string{
		AvailableState: ContAvailableState,
		RunningState:   ContRunningState,
	}

	var contractMsg Contract
	contractMsg.IsSeller = isSeller
	contractMsg.ID = string(contractAddress.Hex())
	contractMsg.State = convertToMsgBusState[contractValues.State]
	contractMsg.Buyer = string(contractValues.Buyer.Hex())
	contractMsg.Price = contractValues.Price
	contractMsg.Limit = contractValues.Limit
	contractMsg.Speed = contractValues.Speed
	contractMsg.Length = contractValues.Length
	contractMsg.StartingBlockTimestamp = contractValues.StartingBlockTimestamp

	return contractMsg
}

//UNUSED
func updateContractMsg(contractMsg *Contract, contractValues hashrateContractValues) {
	contractMsg.Price = contractValues.Price
	contractMsg.Limit = contractValues.Limit
	contractMsg.Speed = contractValues.Speed
	contractMsg.Length = contractValues.Length
}

func GetRandomIDString() (i string) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		//fmt.Printf("Error reading random file: %v\n", err)
		panic(err)
	}
	str := fmt.Sprintf("%08x-%08x-%08x-%08x", b[0:4], b[4:8], b[8:12], b[12:16])
	i = string(str)
	return i
}
