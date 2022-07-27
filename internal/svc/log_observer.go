// Package svc implements monitoring and scanning services of the API server.
package svc

import (
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	eth "github.com/ethereum/go-ethereum/core/types"
)

const (
	// observedBlocksCapacity represents the capacity of channel for observed block IDs.
	observedBlocksCapacity = 100

	// nftQueueCapacity is the capacity of NFT updater queue.
	nftQueueCapacity = 1000
)

// EventHandler represents a function used to process event log record.
type EventHandler func(*eth.Log, *logObserver)

// logObserver represents the service responsible for processing event logs of interest.
type logObserver struct {
	// mgr represents the Manager instance
	mgr *Manager

	// sigStop represents the signal for closing the observer
	sigStop chan bool

	// outEvent represents an output channel being fed
	// with recognized block events for processing
	inEvents chan eth.Log

	// outNftTokens represents an output channel receiving new NFT tokens
	// for processing and metadata update
	outNftTokens chan *types.Token

	// outObservedBlock is fed with numbers of processed blocks.
	outObservedBlocks chan uint64

	// currentBlock contains the number of the currently processed block
	currentBlock uint64

	// topics represents a map of topics to their respective event handlers.
	topics map[common.Hash]EventHandler
}

// newLogObserver creates a new instance of the event logs observer service.
func newLogObserver(mgr *Manager) *logObserver {
	return &logObserver{
		mgr:               mgr,
		sigStop:           make(chan bool, 1),
		outNftTokens:      make(chan *types.Token, nftQueueCapacity),
		outObservedBlocks: make(chan uint64, observedBlocksCapacity),
		topics: map[common.Hash]EventHandler{
			/* Factory::event ContractCreated(address creator, address nft) */
			common.HexToHash("0x2d49c67975aadd2d389580b368cfff5b49965b0bd5da33c144922ce01e7a4d7b"): newNFTContract,

			/* erc721::event Minted(uint256 tokenId, address beneficiary, string tokenUri, address minter) */
			common.HexToHash("0x997115af5924f5e38964c6d65c804d4cb85129b65e62eb20a8ca6329dbe57e18"): erc721TokenMinted,

			/* erc721::event Transfer(address indexed from, address indexed to, uint256 indexed tokenId) */
			common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"): erc721TokenTransfer,

			/* erc1155::event TransferSingle(address indexed _operator, address indexed _from, address indexed _to, uint256 _id, uint256 _amount) */
			common.HexToHash("0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62"): erc1155TokenTransfer,

			/* erc1155::event TransferBatch(address indexed _operator, address indexed _from, address indexed _to, uint256[] _ids, uint256[] _amounts) */
			// common.HexToHash("0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb"): erc1155BatchTransfer,

			/* erc1155::event URI(string _uri, uint256 indexed _id) */
			// common.HexToHash("0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b"): erc1155UriChanged,

			/* Marketplace::event ItemListed(address indexed owner, address indexed nft, uint256 tokenId, uint256 quantity, address payToken, uint256 pricePerItem, uint256 startingTime) */
			common.HexToHash("0xa0294f02f8ad82fe4744717b0f953a105547196cd3c67056200c1a4ae3aa2629"): marketNFTListed,

			/* Marketplace::event ItemUpdated(address indexed owner, address indexed nft, uint256 tokenId, address payToken, uint256 newPrice) */
			common.HexToHash("0x60a11f1619b1716bc2857bf610d4bc631336e14d197025fd5875c1aca1ac7cbd"): marketNFTUpdated,

			/* Marketplace::event ItemCanceled(address indexed owner, address indexed nft, uint256 tokenId) */
			common.HexToHash("0x9ba1a3cb55ce8d63d072a886f94d2a744f50cddf82128e897d0661f5ec623158"): marketNFTUnlisted,

			/* Marketplace::event ItemSold(address indexed seller, address indexed buyer, address indexed nft, uint256 tokenId, uint256 quantity, address payToken, int256 unitPrice, uint256 pricePerItem) */
			common.HexToHash("0x949d1413baca5c0e4ab96b0198d536cac8cdcc17cb909b9ea24594f42ed9fa0d"): marketItemSold,

			/* Marketplace::event OfferCreated(address indexed creator, address indexed nft, uint256 tokenId, uint256 quantity, address payToken, uint256 pricePerItem, uint256 deadline) */
			common.HexToHash("0x89f255157c655b5155655107b77c620998e5ad4e7485d749e4e6d7ddb63e70f6"): marketOfferCreated,

			/* Marketplace::event OfferCanceled(address indexed creator, address indexed nft, uint256 tokenId) */
			common.HexToHash("0xc6e24dcedb16cc237925b586889d0a38102c719734d6cc56acb89b013099b3a7"): marketOfferCanceled,

			/* Auction::event AuctionCreated(address indexed nftAddress, uint256 indexed tokenId, address payToken) */
			common.HexToHash("0xca437d90ed6373b827a01275bd2fdfe7e6406d7ecd400662ee0533d3546ab17a"): auctionCreated,

			/* AuctionV1::event UpdateAuctionStartTime(address indexed nftAddress, uint256 indexed tokenId, uint256 startTime) */
			common.HexToHash("0xf123182d1bfce603eb20a2f74cce94af6542f353a4c2a8d5cfadf41822b32b38"): auctionStartTimeUpdated,

			/* AuctionV1::event UpdateAuctionEndTime(address indexed nftAddress, uint256 indexed tokenId, uint256 endTime) */
			common.HexToHash("0xd8aaad775c4c9b21112b5cb22ab4618c74936d873850fbd6f2e691701119ffe1"): auctionEndTimeUpdated,

			/* Auction::event UpdateAuctionReservePrice(address indexed nftAddress, uint256 indexed tokenId, address payToken, uint256 reservePrice) */
			common.HexToHash("0x1492815b4a41d34cc0015f8ba8013fa45717ea6d8591db6036bdca706981fe83"): auctionReserveUpdated,

			/* Auction::event AuctionCancelled(address indexed nftAddress, uint256 indexed tokenId) */
			common.HexToHash("0x018b64b6242d32aa550e95d78985b938d71af5b3f10827b0683f55da16393048"): auctionCanceled,

			/* AuctionV1::event AuctionResulted(address indexed nftAddress, uint256 indexed tokenId, address indexed winner, address payToken, int256 unitPrice, uint256 winningBid) */
			common.HexToHash("0x0427247b7170429ed72451c5779575834ddc75e9788eb907e40025c62ed7a258"): auctionResolved,

			/* AuctionV2::event AuctionResulted(address oldOwner, address indexed nftAddress, uint256 indexed tokenId, address indexed winner, address payToken, int256 unitPrice, uint256 winningBid) */
			common.HexToHash("0xa717395bf4be1915f0bc8e6cf9b0f526c5ad40a0c3750b709a29a834daf0fd9b"): auctionResolvedV2,

			/* Auction::event BidPlaced(address indexed nftAddress, uint256 indexed tokenId, address indexed bidder, uint256 bid) */
			common.HexToHash("0x0158f5674dc243762459b88cfc91b10d2d1ef9d40821cca978c2b680aa444682"): auctionBidPlaced,

			/* Auction::event BidWithdrawn(address indexed nftAddress, uint256 indexed tokenId, address indexed bidder, uint256 bid) */
			common.HexToHash("0x867b8ea96dd803063f905a19f8117cbb1866ec7a594dfede75ab4a5235f61d7c"): auctionBidWithdrawn,

			/* RandomNumberOracle::event RandomNumberRequested(bytes32 requestID, bytes32 seed) */
			common.HexToHash("0xac2e43d9741627d0f2e7a61dba4f97dfa56414d39e787163b0e6dbde34e3a6b2"): requestedRandomNumber,
		},
	}
}

// name provides the name fo the log observer service.
func (lo *logObserver) name() string {
	return "logs observer"
}

// init configures the log observer and subscribes it with the manager.
func (lo *logObserver) init() {
	// link channels
	lo.inEvents = lo.mgr.blkObserver.outEvents

	// add and run
	lo.mgr.add(lo)
}

// close signals the log observer to terminate.
func (lo *logObserver) close() {
	lo.sigStop <- true
}

// run collects incoming event logs from the channel and processes them using
// pre-configured callbacks.
func (lo *logObserver) run() {
	defer func() {
		close(lo.outNftTokens)
		close(lo.outObservedBlocks)

		lo.mgr.closed(lo)
	}()

	for {
		select {
		case <-lo.sigStop:
			return
		case evt, ok := <-lo.inEvents:
			if !ok {
				return
			}
			lo.process(&evt)
			lo.processed(&evt)
		}
	}
}

// process an incoming event
func (lo *logObserver) process(evt *eth.Log) {
	// get the handler
	handler, ok := lo.topics[evt.Topics[0]]
	if !ok {
		log.Criticalf("event log handler not found for event #%d / %d, topic %s", evt.BlockNumber, evt.Index, evt.Topics[0].String())
		return
	}

	// do the handler job
	log.Debugf("processing event #%d/#%d at %s -> topic %s", evt.BlockNumber, evt.Index, evt.Address.String(), evt.Topics[0].String())
	handler(evt, lo)
}

// processed manages information about block being evaluated and notifies about blocks done.
func (lo *logObserver) processed(evt *eth.Log) {
	if lo.currentBlock != evt.BlockNumber {
		// notify we did process the last block
		select {
		case lo.outObservedBlocks <- lo.currentBlock:
			log.Debugf("block #%d done", lo.currentBlock)
		default:
		}

		lo.currentBlock = evt.BlockNumber
	}
}

// topicsList provides a list of observed topics for blocks event filtering
func (lo *logObserver) topicsList() [][]common.Hash {
	list := make([][]common.Hash, 1)
	list[0] = make([]common.Hash, len(lo.topics))

	var i int
	for h := range lo.topics {
		list[0][i] = h
		i++
	}

	return list
}
