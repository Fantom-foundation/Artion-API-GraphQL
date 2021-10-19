// SPDX-License-Identifier: MIT
pragma solidity ^0.6.0;

// RandomNumberConsumer represents a contract capable of receiving random numbers
// provided by the oracle on requests.
interface IRandomNumberConsumer {
    function consumeRandomNumber(bytes32 _seed, uint256 _rnd) external;
}

// RandomTrade provides a trader contract capable of selling random NFT from a given set.
// The randomness of a sale is obtained from an external off-chain Oracle on request-callback basis.
// Tokens are sold for a price obtained from an external price oracle contract.
contract RandomTrade is IRandomNumberConsumer {
    // NFT represents a structure of sell-able NFT token.
    struct NFT {
        bytes32 id;             // joined id of this NFT
        address nftContract;    // contract managing this NFT
        uint256 tokenID;        // ID of the NFT within the contract
        bytes32 next;           // hash of the next NFT in the trade
        bytes32 previous;       // hash of the previous NFT in the trade
        address buyer;          // address of the token buyer
    }

    // getCurrent represents a current NFT pointer.
    bytes32 public getCurrent;

    // getPriceOracle is the address of the price oracle used to get the token price.
    address public getPriceOracle;

    // getRNGOracle is the oracle used to provide random numbers.
    address public getRNGOracle;

    // getOwner if the address of the contract owner.
    address public getOwner;

    // getNFT is the mapping between internal token hash and NFT structure.
    mapping(bytes32 => NFT) public getNFT;

    // getNFTCount provides the total number of tokens traded by the contract.
    uint256 public getNFTCount;

    // getNFTAvailable provides the number of tokens available for purchase.
    uint256 public getNFTAvailable;

    // constructor initializes the contract on deployment.
    constructor (address _priceOracle, address _rngOracle) public {
        getOwner = msg.sender;
        getPriceOracle = _priceOracle;
        getRNGOracle = _rngOracle;
    }

    // transferOwnership changes the contract owner.
    function transferOwnership(address _owner) external {
        require(msg.sender == getOwner, "owner only");
        getOwner = _owner;
    }

    // setPriceOracle changes the price oracle used by the contract.
    function setPriceOracle(address _oracle) external {
        require(msg.sender == getOwner, "owner only");
        getPriceOracle = _oracle;
    }

    // setRNGOracle changes the random number oracle used by the contract.
    function setRNGOracle(address _oracle) external {
        require(msg.sender == getOwner, "owner only");
        getRNGOracle = _oracle;
    }
}
