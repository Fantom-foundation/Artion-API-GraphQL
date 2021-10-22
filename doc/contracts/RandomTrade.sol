// SPDX-License-Identifier: MIT
pragma solidity ^0.6.0;

// IRandomNumberOracle represents a contract capable of requesting random numbers from off-chain source.
interface IRandomNumberOracle {
    function requestRandomNumber(bytes32 _seed) external;
}

// RandomNumberConsumer represents a contract capable of receiving random numbers
// provided by the oracle on requests.
interface IRandomNumberConsumer {
    function consumeRandomNumber(bytes32 _seed, uint256 _rnd) external;
}

// IERC721Receiver defines an interface of a contract capable of receiving ERC721 tokens safely.
interface IERC721Receiver {
    function onERC721Received(address operator, address from, uint256 tokenId, bytes calldata data) external returns (bytes4);
}

// IPriceOracleProxy defines an interface of a price oracle proxy contract
// used to provide up-to-date information about the value of supported tokens.
interface IPriceOracleProxy {
    function getPrice(address _token) external view returns (uint256, uint8);
}

// IERC165 defines ERC-165 interface used to check contract capability
// https://eips.ethereum.org/EIPS/eip-165
interface IERC165 {
    function supportsInterface(bytes4 interfaceId) external view returns (bool);
}

// IERC721 defines reduced ERC-721 token interface.
// We don't really need all the implemented functionality, so we list only functions we require.
interface IERC721 is IERC165 {
    // ownerOf provides information about the owner of the given token.
    function ownerOf(uint256 tokenId) external view returns (address owner);

    // safeTransferFrom executes safe transfer of the given token to designated recipient.
    function safeTransferFrom(address from, address to, uint256 tokenId) external;
}

// IERC20 defines reduced interface of an ERC-20 token.
// We don't need all the implemented functionality, so we list only functions we require.
interface IERC20 {
    // decimals provides the n umber of decimals used by the token.
    function decimals() external view returns (uint8);

    // transfer moves given amount of tokens from caller's account to recipient.
    function transfer(address recipient, uint256 amount) external returns (bool);

    // allowance returns the remaining number of tokens that spender will be
    // allowed to spend on behalf of owner through transferFrom() call.
    function allowance(address owner, address spender) external view returns (uint256);

    // transferFrom moves given amount of tokens from sender to recipient .
    function transferFrom(address sender, address recipient, uint256 amount) external returns (bool);
}

// RandomTrade provides a trader contract capable of selling random NFT from a given set.
// The randomness of a sale is obtained from an external off-chain Oracle on request-callback basis.
// Tokens are sold for a price obtained from an external price oracle contract.
contract RandomTrade is IRandomNumberConsumer, IERC721Receiver {
    // Equals to `bytes4(keccak256("onERC721Received(address,address,uint256,bytes)"))`
    bytes4 private constant _ERC721_RECEIVED = 0x150b7a02;

    // ID of the ERC721 interface for EIP-165 check.
    bytes4 private constant _INTERFACE_ID_ERC721 = 0x80ac58cd;

    // NFT represents a structure of NFT token offered on the trade.
    struct NFT {
        // token reference
        address nftContract;    // contract managing this NFT
        uint256 tokenID;        // ID of the NFT within the contract

        // pool linkage
        bytes32 next;           // hash of the next NFT
        bytes32 previous;       // hash of the previous NFT
    }

    // Purchase represents a structure of a purchase request.
    struct Purchase {
        address buyer;          // address of the buyer
        address payToken;       // address of the pay token deposited to cover the purchase
        uint256 deposit;        // amount of tokens deposited on a purchase
    }

    // getOwner is the address of the contract owner/administrator.
    address public getOwner;

    // getPriceOracle is the address of the price oracle proxy used to get the token price.
    address public getPriceOracle;

    // getRNGOracle is the oracle used to provide random numbers.
    address public getRNGOracle;

    // getTradeStart is the time stamp of the trade opening.
    uint256 public getTradeStarts;

    // getTradeEnds is the time stamp of the trade closing.
    uint256 public getTradeEnds;

    // getName provides the name of the random trade.
    string public getName;

    // isPayTokenAllowed is a map of allowed ERC20 pay tokens.
    // mapping: (ERC20 address -> allowed true/false)
    mapping(address => bool) public isPayTokenAllowed;

    // getNFT is the mapping between internal token hash and NFT structure.
    // mapping: (token hash -> NFT structure)
    mapping(bytes32 => NFT) public getNFT;

    // getPurchase is the mapping between internal purchase ID hash and Purchase structure.
    // mapping: (purchase hash -> Purchase structure)
    mapping(bytes32 => Purchase) public getPurchase;

    // getCurrent represents the pointer to current NFT in the pool.
    bytes32 public getCurrentNFT;

    // getNFTCount represents the total number of tokens traded by the contract.
    uint256 public getNFTCount;

    // getNFTAvailable represents the number of tokens available for purchase.
    uint256 public getNFTAvailable;

    // getUnitPrice represents the unit price all the tokens in the pool are traded for.
    // The price denomination depends on the price oracle in use.
    uint256 public getUnitPrice;

    // getUnitPriceDecimals represents the number of decimals used on unit price.
    uint8 public getUnitPriceDecimals;

    // _nonce stores purchase request counter.
    uint256 private _nonce;

    // OwnerChanged is emitted on contract owner transfer.
    event OwnerChanged(address newOwner);

    // PayTokenAdded event is emitted on a new pay token.
    event PayTokenAdded(address token);

    //  PayTokenRemoved event is emitted on a pay token being disabled.
    event PayTokenRemoved(address token);

    // PriceChanged is emitted on a unit price update.
    event PriceChanged(uint256 newPrice, uint8 newDecimals);

    // TokenAdded is emitted with a new token added to the trade pool.
    event TokenAdded(address collection, uint256 tokenID, bytes32 tradeTokenID);

    // PurchaseCreated is emitted on a purchase request creation.
    event PurchaseCreated(address indexed buyer, bytes32 purchaseID, address payToken, uint256 price);

    // PurchaseFinished is emitted on a purchase finalization.
    event PurchaseFinished(address indexed buyer, bytes32 purchaseID, address collection, uint256 tokenID);

    // PurchaseCanceled is emitted on a purchase being canceled.
    event PurchaseCanceled(address indexed buyer, bytes32 purchaseID);

    // onlyOwner reverts if called by any account other than the owner.
    modifier onlyOwner() {
        require(msg.sender == getOwner, "RandomTrade: allowed to owner only");
        _;
    }

    // constructor initializes the contract on deployment.
    constructor (string memory _name, address _priceOracle, address _rngOracle, uint256 _price, uint8 _priceDecimals, uint256 _start, uint256 _end) public {
        getOwner = msg.sender;

        getName = _name;
        getPriceOracle = _priceOracle;
        getRNGOracle = _rngOracle;
        getUnitPrice = _price;
        getUnitPriceDecimals = _priceDecimals;
        getTradeStarts = _start;
        getTradeEnds = _end;
    }

    // onERC721Received is called upon ERC-721 token delivery.
    // Implements IERC721Receiver for safe token transfer.
    function onERC721Received(address operator, address /* from */, uint256 tokenId, bytes calldata /* data */) external override returns (bytes4) {
        // we accept only ERC-721 tokens sent by the current trade owner
        require(isContract(msg.sender), "RandomTrade: invalid caller");
        require(IERC165(msg.sender).supportsInterface(_INTERFACE_ID_ERC721), "RandomTrade: invalid NFT received");

        // only owner can send tokens
        require(operator == getOwner, "RandomTrade: tokens accepted from the owner only");

        // no new tokens after the trade starts
        require(block.timestamp < getTradeStarts, "RandomTrade: the trade already started");

        // add the token to the pool and return expected selector
        _addToken(msg.sender, tokenId);
        return _ERC721_RECEIVED;
    }

    // _addToken adds the given token to the trading pool.
    function _addToken(address _contract, uint256 _tokenID) internal {
        // gen the id of the token from ERC-721 contract and token ID
        bytes32 id = keccak256(abi.encodePacked(_contract, _tokenID));

        // store the token base
        getNFT[id].nftContract = _contract;
        getNFT[id].tokenID = _tokenID;

        // adding the first token? loop the circle
        if (getCurrentNFT == bytes32(0)) {
            getNFT[id].next = id;
            getNFT[id].previous = id;
        } else {
            // add the new one in front of the cursor
            getNFT[id].previous = getCurrentNFT;
            getNFT[id].next = getNFT[getCurrentNFT].next;
            getNFT[getNFT[getCurrentNFT].next].previous = id;
            getNFT[getCurrentNFT].next = id;
        }

        // move the cursor; advance counters
        getCurrentNFT = id;
        getNFTCount += 1;
        getNFTAvailable += 1;

        emit TokenAdded(_contract, _tokenID, id);
    }

    // purchase random token from the pool using the given ERC20 token.
    function purchase(address _token) external {
        // no trade before opening, or after the end
        require(block.timestamp >= getTradeStarts, "RandomTrade: the trade not open");
        require(block.timestamp < getTradeEnds, "RandomTrade: the trade is closed");

        // any available tokens?
        require(0 < getNFTAvailable, "RandomTrade: no NFT available for purchase");

        // the pay token must be whitelisted
        require(isPayTokenAllowed[_token], "RandomTrade: the pay token not allowed");

        // get the amount of pay token required to satisfy the NFT price
        uint256 price = _getPrice(_token);
        require(0 < price, "RandomTrade: the price must be non-zero");
        require(price <= IERC20(_token).allowance(msg.sender, address(this)), "RandomTrade: the pay token allowance too low");

        // make the purchase record ID and check it's unique
        _nonce += 1;
        bytes32 id = keccak256(abi.encodePacked(msg.sender, _token, _nonce));
        require(getPurchase[id].deposit == 0, "RandomTrade: pending purchase overwrite rejected");

        // one less token is available from now
        getNFTAvailable -= 1;

        // transfer pay tokens
        safeTransferFrom(IERC20(_token), msg.sender, address(this), price);

        // create the purchase record
        getPurchase[id].buyer = msg.sender;
        getPurchase[id].deposit = price;
        getPurchase[id].payToken = _token;

        // request random number from RNG oracle
        IRandomNumberOracle(getRNGOracle).requestRandomNumber(id);

        // inform we initiated the purchase
        emit PurchaseCreated(msg.sender, id, _token, price);
    }

    // cancelPurchase cancels the given purchase.
    function cancelPurchase(bytes32 _id) external {
        // make sure the purchase exists and is made by the caller
        require(msg.sender == getPurchase[_id].buyer, "RandomTrade: not your purchase request");
        require(block.timestamp > getTradeEnds, "RandomTrade: the trade is still open");

        _cancelPurchase(_id);
    }

    // clearPurchase allows trade owner to force purchase request removal.
    function clearPurchase(bytes32 _id) external onlyOwner {
        require(0 < getPurchase[_id].deposit, "RandomTrade: closed purchase request");
        _cancelPurchase(_id);
    }

    // _cancelPurchase executes the given purchase removal.
    function _cancelPurchase(bytes32 _id) internal {
        address token = getPurchase[_id].payToken;
        uint256 amount = getPurchase[_id].deposit;

        // clear the purchase storage and
        delete (getPurchase[_id]);

        // send the deposit back to owner and inform about cancel
        safeTransfer(IERC20(token), msg.sender, amount);

        // the token is now available again
        getNFTAvailable += 1;
        emit PurchaseCanceled(msg.sender, _id);
    }

    // consumeRandomNumber processes callback with random number.
    function consumeRandomNumber(bytes32 _seed, uint256 _rnd) external override {
        // make sure it comes from trusted source
        require(msg.sender == getRNGOracle, "RandomTrade: allowed to RNG oracle only");

        // make sure the request still exists;
        require(0 != getPurchase[_seed].deposit, "RandomTrade: unknown purchase");

        // finish the purchase
        _finishPurchase(_seed, _rnd);
    }

    // _finishPurchase finishes the purchase with the given random number.
    function _finishPurchase(bytes32 _id, uint256 _rnd) internal {
        // move the internal cursor using random number of steps
        require(bytes32(0) != getCurrentNFT, "RandomTrade: no NFT left to trade");
        _moveCursor(_rnd);

        // get the current token from the pool; this deletes the NFT from pool
        (address collection, uint256 tokenID) = _removeToken();

        // purchase details
        address buyer = getPurchase[_id].buyer;
        address payToken = getPurchase[_id].payToken;
        uint256 price = getPurchase[_id].deposit;

        delete (getPurchase[_id]);

        // send the NFT to buyer
        IERC721(collection).safeTransferFrom(address(this), buyer, tokenID);

        // send purchase price to the trade owner
        safeTransfer(IERC20(payToken), getOwner, price);
    }

    // removeToken allows trade owner to remove un-booked token from the contract.
    function removeToken(uint256 _shift) external onlyOwner {
        // any available tokens?
        require(0 < getNFTAvailable, "RandomTrade: no available NFT left");

        // move the internal cursor using random number of steps
        require(bytes32(0) != getCurrentNFT, "RandomTrade: trade pool empty");
        _moveCursor(_shift);

        // get the current token from the pool; this deletes the NFT from pool
        (address collection, uint256 tokenID) = _removeToken();

        // send the NFT to buyer
        IERC721(collection).safeTransferFrom(address(this), getOwner, tokenID);
    }

    // _moveCursor traverses the internal active token cursor
    // by the given amount of steps normalized to the active loop elements.
    function _moveCursor(uint256 _shift) internal {
        // nothing to do?
        if (_shift == 0) return;

        require(0 < getNFTCount, "RandomTrade: trade pool empty");
        uint256 steps = _shift % getNFTCount;
        for (uint256 i = 0; i < steps; i++) {
            getCurrentNFT = getNFT[getCurrentNFT].next;
        }
    }

    // _removeToken removes the current token from the pool and returns the token details.
    function _removeToken() internal returns (address collection, uint256 tokenID) {
        // what is going to be removed
        collection = getNFT[getCurrentNFT].nftContract;
        tokenID = getNFT[getCurrentNFT].tokenID;

        // remove the NFT from collection
        bytes32 next = getNFT[getCurrentNFT].next;
        bytes32 previous = getNFT[getCurrentNFT].previous;
        delete (getNFT[getCurrentNFT]);

        // rewire the loop if it's not the last NFT in the pool
        if (next != previous) {
            getNFT[next].previous = previous;
            getNFT[previous].next = next;
            getCurrentNFT = next;
        } else {
            getCurrentNFT = bytes32(0);
        }

        getNFTCount -= 1;
        return (collection, tokenID);
    }

    // getPrice provides the price of the pool NFT denominated in given token.
    function getPrice(address _token) public view returns (uint256) {
        require(isPayTokenAllowed[_token], "RandomTrade: the pay token not allowed");
        return _getPrice(_token);
    }

    // _getPrice provides the price of the pool NFT denominated in given token.
    function _getPrice(address _token) internal view returns (uint256) {
        // get the token price from oracle
        (uint256 value, uint8 valDecimals) = IPriceOracleProxy(getPriceOracle).getPrice(_token);
        require(0 < value, "RandomTrade: the pay token has unknown value");

        // how many decimals are used by the token
        uint8 tokDecimals = IERC20(_token).decimals();
        require(0 < tokDecimals, "RandomTrade: the pay token has unknown decimals");

        // calculate the expected price denominated in pay token
        return (getUnitPrice * (uint256(10) ** (valDecimals + tokDecimals))) / (value * (uint256(10) ** getUnitPriceDecimals));
    }

    // transferOwnership changes the contract owner.
    function transferOwnership(address _newOwner) external onlyOwner {
        require(_newOwner != address(0), "RandomTrade: zero address not allowed");
        getOwner = _newOwner;
    }

    // setPriceOracle changes the price oracle used by the contract.
    function setPriceOracle(address _oracle) external onlyOwner {
        require(_oracle != address(0), "RandomTrade: zero address not allowed");
        getPriceOracle = _oracle;
    }

    // setRNGOracle changes the random number oracle used by the contract.
    function setRNGOracle(address _oracle) external onlyOwner {
        require(_oracle != address(0), "RandomTrade: zero address not allowed");
        getRNGOracle = _oracle;
    }

    // setPrice updates the unit price of the tokens in the pool.
    function setPrice(uint256 _price, uint8 _priceDecimals) external onlyOwner {
        getUnitPrice = _price;
        getUnitPriceDecimals = _priceDecimals;

        emit PriceChanged(_priceDecimals, _priceDecimals);
    }

    // allowPayToken enables the given pay token.
    function allowPayToken(address _token) external onlyOwner {
        uint256 price = getPrice(_token);
        require(0 < price, "RandomTrade: the pay token not supported by price oracle");

        // how many decimals are used by the token
        uint8 tokDecimals = IERC20(_token).decimals();
        require(0 < tokDecimals, "RandomTrade: the pay token has unknown decimals");

        isPayTokenAllowed[_token] = true;
        emit PayTokenAdded(_token);
    }

    // denyPayToken disables the given pay token.
    function denyPayToken(address _token) external onlyOwner {
        isPayTokenAllowed[_token] = false;
        emit PayTokenRemoved(_token);
    }

    // isContract checks if the given account is a contract.
    // We use OpenZeppelin v3.4 implementation.
    function isContract(address account) internal view returns (bool) {
        uint256 size;
        // solhint-disable-next-line no-inline-assembly
        assembly {size := extcodesize(account)}
        return size > 0;
    }

    // safeTransfer moves the given amount of pay tokens from this contract to recipient safely.
    function safeTransfer(IERC20 _token, address _to, uint256 _amount) internal {
        bytes memory payload = abi.encodeWithSelector(_token.transfer.selector, _to, _amount);

        (bool success, bytes memory res) = address(_token).call(payload);
        require(success, "RandomTrade: pay token transfer failed");

        if (res.length > 0) {
            // solhint-disable-next-line max-line-length
            require(abi.decode(res, (bool)), "RandomTrade: pay token transfer did not succeed");
        }
    }

    // safeTransferFrom moves the given amount of pay tokens from sender to recipient safely.
    function safeTransferFrom(IERC20 _token, address _from, address _to, uint256 _amount) internal {
        bytes memory payload = abi.encodeWithSelector(_token.transferFrom.selector, _from, _to, _amount);

        (bool success, bytes memory res) = address(_token).call(payload);
        require(success, "RandomTrade: pay token transfer failed");

        if (res.length > 0) {
            // solhint-disable-next-line max-line-length
            require(abi.decode(res, (bool)), "RandomTrade: pay token transfer did not succeed");
        }
    }
}
