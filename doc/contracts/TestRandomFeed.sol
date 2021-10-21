// SPDX-License-Identifier: MIT
pragma solidity ^0.6.0;

interface IRandomNumberOracle {
    function requestRandomNumber(bytes32 _seed) external;
}

interface IRandomNumberConsumer {
    function consumeRandomNumber(bytes32 _seed, uint256 _rnd) external;
}

contract TestRandomFeed is IRandomNumberConsumer {
    // Equals to `bytes4(keccak256("onERC721Received(address,address,uint256,bytes)"))`
    bytes4 private constant _ERC721_RECEIVED = 0x150b7a02;

    address public owner;

    address public oracle;

    uint256 public nonce;

    uint256 public lastRandom;

    constructor (address _oracle) public {
        owner = msg.sender;
        oracle = _oracle;
    }

    function request() external {
        require(msg.sender == owner, "owner only");

        nonce += 1;
        bytes32 id = keccak256(abi.encodePacked(msg.sender, nonce));

        IRandomNumberOracle(oracle).requestRandomNumber(id);
    }

    function consumeRandomNumber(bytes32 _seed, uint256 _rnd) external override {
        require(msg.sender == oracle, "oracle only");
        lastRandom = _rnd;
    }
}
