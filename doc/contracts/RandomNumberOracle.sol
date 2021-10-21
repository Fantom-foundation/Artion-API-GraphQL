pragma solidity ^0.6.0;

// RandomNumberConsumer represents a contract capable of receiving random numbers
// provided by the oracle on requests.
interface IRandomNumberConsumer {
    function consumeRandomNumber(bytes32 _seed, uint256 _rnd) external;
}

// RandomNumberOracle implements and oracle able to resolve random number requests
// using off-chain random number provider.
contract RandomNumberOracle {
    // Request represents a pending request for random number.
    struct Request {
        bytes32 seed;       // seed provided by the caller
        address consumer;   // address of the consumer waiting for the random number
    }

    // getRequest represents existing requests for random numbers.
    // Mapping: requestID => struct Request
    mapping(bytes32 => Request) public getRequest;

    // getNonce represents a call nonce per caller
    // It's mixed with additional call params to generate call request ID.
    mapping(address => uint256) public getNonce;

    // getProvider represents a map of allowed random number providers.
    mapping(address => bool) public getProvider;

    // getOwner provides address of contract owner.
    address public getOwner;

    // RandomNumberRequested event is emitted to notify off-chain oracle to provide
    // a new random number for the given seed.
    event RandomNumberRequested(bytes32 requestID, bytes32 seed);

    // RequestCanceled event is emitted to notify a random number request was canceled.
    event RequestCanceled(bytes32 requestID);

    // OwnerChanged is emitted on contract owner transfer.
    event OwnerChanged(address newOwner);

    // ProviderAllowed is emitted on a new allowed provider.
    event ProviderAllowed(address provider);

    // ProviderDenied is emitted on a disabled provider.
    event ProviderDenied(address provider);

    /**
     * Reverts if called by any account other than the owner.
     */
    modifier onlyOwner() {
        require(getOwner == msg.sender, "RNG: allowed to owner only");
        _;
    }

    // constructor initializes the oracle contract.
    constructor (address _provider) public {
        getOwner = msg.sender;
        allowProvider(_provider);
    }

    // requestRandomNumber creates a new request for random number
    // and notifies an off-chain provider to feed the random number in.
    function requestRandomNumber(bytes32 _seed) external {
        bytes32 requestID = makeRequestID(msg.sender, _seed, getNonce[msg.sender], block.timestamp);
        require(getRequest[requestID].consumer == address(0), "RNG: existing request ID");

        // store the request and advance the nonce for the caller
        // we should not need to deal with nonce overflow, since we delete fulfilled requests
        getRequest[requestID] = Request({
        seed : _seed,
        consumer : msg.sender
        });
        getNonce[msg.sender] += 1;

        // notify provider
        emit RandomNumberRequested(requestID, _seed);
    }

    // makeRequestID generates a new request ID.
    function makeRequestID(address _sender, bytes32 _seed, uint256 _nonce, uint256 _ts) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(_sender, _seed, _nonce, _ts));
    }

    // fulfillRandomNumber is called by an off-chain provider
    // to satisfy pending random number request.
    function fulfillRandomNumber(bytes32 _requestID, uint256 _rnd) external {
        // check conditions
        require(getProvider[msg.sender], "RNG: only provider can fulfill");
        require(getRequest[_requestID].consumer != address(0), "RNG: unknown request ID");

        // remove the request
        address memory _target = getRequest[_requestID].consumer;
        bytes memory payload = abi.encodeWithSignature("consumeRandomNumber(bytes32,uint256)", getRequest[_requestID].seed, _rnd);
        delete (getRequest[_requestID]);

        // fulfill the request calling the recipient
        (bool success,) = _target.call(payload);
        require(success, "RNG: failed to fulfill");
    }

    // transferOwnership changes the contract owner.
    function transferOwnership(address _newOwner) external onlyOwner {
        require(_newOwner != address(0), "RandomTrade: zero address not allowed");
        getOwner = _newOwner;
        emit OwnerChanged(_newOwner);
    }

    // allowProvider enables a new provider address.
    function allowProvider(address _provider) public onlyOwner {
        getProvider[_provider] = true;
        emit ProviderAllowed(_provider);
    }

    // denyProvider disables a new provider address.
    function denyProvider(address _provider) public onlyOwner {
        getProvider[_provider] = false;
        emit ProviderDenied(_provider);
    }

    // cancelRequest allows owner to remove given pending request.
    function cancelRequest(bytes32 _requestID) external onlyOwner {
        delete (getRequest[_requestID]);
        emit RequestCanceled(_requestID);
    }
}
