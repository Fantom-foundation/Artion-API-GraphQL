// SPDX-License-Identifier: MIT
pragma solidity ^0.6.0;

// IPriceOracleProxy defines an interface of a price oracle proxy contract
// used to provide up-to-date information about the value of supported tokens.
interface IPriceOracleProxy {
    function getPrice(address _token) external view returns (uint256, uint8);
}

// PriceOracleProxy implements IPriceOracleProxy with an in-place price storage.
// It's intended to be replaced by a proxy with 3rd party price oracle support (i.e. Chainlink, TheBand, etc.)
contract PriceOracleProxy is IPriceOracleProxy {
    // _price represents the actual current price per token.
    // Mapping: (token address -> token price)
    mapping(address => uint256) private _price;

    // _decimals represents the decimals of the price of the token.
    // Mapping: (token address -> price decimals)
    mapping(address => uint8) private _decimals;

    // getOwner is the address of the contract owner.
    address public getOwner;

    // onlyOwner reverts if the caller is not the contract owner.
    modifier onlyOwner() {
        require(getOwner == msg.sender, "PriceOracleProxy: allowed to the owner only");
        _;
    }

    // constructor initializes the contract on deployment.
    constructor () public {
        getOwner = msg.sender;
    }

    // getPrice provides the price of the given token by address.
    function getPrice(address _token) external view override returns (uint256, uint8) {
        require(_price[_token] != 0, "PriceOracleProxy: unknown token");
        return (_price[_token], _decimals[_token]);
    }

    // setPrice sets a new price for the given token.
    function setPrice(address _token, uint256 _newPrice, uint8 _newDecimals) external onlyOwner {
        _price[_token] = _newPrice;
        _decimals[_token] = _newDecimals;
    }

    // transferOwnership changes the owner of the contract.
    function transferOwnership(address _newOwner) public onlyOwner {
        require(_newOwner != address(0), "PriceOracleProxy: zero address not allowed");
        getOwner = _newOwner;
    }
}
