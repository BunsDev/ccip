// SPDX-License-Identifier: MIT
// solhint-disable not-rely-on-time
pragma solidity ^0.8.15;
pragma abicoder v2;

import "../vendor/ECDSA.sol";
import "../vendor/ERC165.sol";

import "./IForwarder.sol";
import {OwnerIsCreator} from "../ccip/OwnerIsCreator.sol";

/// @title The Forwarder Implementation
/// @notice This implementation of the `IForwarder` interface uses ERC-712 signatures and stored nonces for verification.
/// @dev This implementation has been ported from OpenGSN's Forwarder.sol and modified in following ways:
/// @dev 1. execute() does not accept "gas" parameter which allows caller to specify max gas limit for the forwarded call
/// @dev 2. execute() does not accept "value" parameter which allows caller to pass native token to the forwarded call
/// @dev 3. renamed field: "address to" => "address target"
contract Forwarder is IForwarder, ERC165, OwnerIsCreator {
  using ECDSA for bytes32;

  address private constant DRY_RUN_ADDRESS = 0x0000000000000000000000000000000000000000;

  string public constant GENERIC_PARAMS = "address from,address target,uint256 nonce,bytes data,uint256 validUntilTime";

  string public constant EIP712_DOMAIN_TYPE =
    "EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)";

  /// @dev mapping of EIP712 request type to boolean.
  /// @dev request type is validated during EIP712 signature verification
  mapping(bytes32 => bool) public s_typeHashes;
  /// @dev mapping of EIP712 domain separator to boolean.
  /// @dev domain separator must be registered for every target contract
  /// @dev domain separator is validated during EIP712 signature verification
  mapping(bytes32 => bool) public s_domains;

  /// @notice Nonces of senders, used to prevent replay attacks
  mapping(address => uint256) private s_nonces;

  // solhint-disable-next-line no-empty-blocks
  receive() external payable {}

  /// @inheritdoc IForwarder
  function getNonce(address from) public view override returns (uint256) {
    return s_nonces[from];
  }

  constructor() {
    string memory requestType = string(abi.encodePacked("ForwardRequest(", GENERIC_PARAMS, ")"));
    registerRequestTypeInternal(requestType);
  }

  /// @inheritdoc IERC165
  function supportsInterface(bytes4 interfaceId) public view virtual override(IERC165, ERC165) returns (bool) {
    return interfaceId == type(IForwarder).interfaceId || super.supportsInterface(interfaceId);
  }

  /// @inheritdoc IForwarder
  function verify(
    ForwardRequest calldata req,
    bytes32 domainSeparator,
    bytes32 requestTypeHash,
    bytes calldata suffixData,
    bytes calldata sig
  ) external view override {
    _verifyNonce(req);
    _verifySig(req, domainSeparator, requestTypeHash, suffixData, sig);
  }

  error ForwardFailed(bytes reason);
  event ForwardSucceeded(
    address indexed from,
    address indexed target,
    bytes32 indexed domainSeparator,
    uint256 nonce,
    bytes data,
    bytes returnValue
  );

  error RequestExpired(uint256 expected, uint256 actual);

  /// @inheritdoc IForwarder
  function execute(
    ForwardRequest calldata req,
    bytes32 domainSeparator,
    bytes32 requestTypeHash,
    bytes calldata suffixData,
    bytes calldata sig
  ) external payable override returns (bool success, bytes memory ret) {
    _verifySig(req, domainSeparator, requestTypeHash, suffixData, sig);
    _verifyAndUpdateNonce(req);

    if (req.validUntilTime != 0 && req.validUntilTime <= block.timestamp) {
      revert RequestExpired(block.timestamp, req.validUntilTime);
    }

    bytes memory callData = abi.encodePacked(req.data, req.from);
    // solhint-disable-next-line avoid-low-level-calls
    (success, ret) = req.target.call(callData);

    if (!success) {
      if (ret.length == 0) revert("Forwarded call reverted without reason");
      // assembly below extracts revert reason from the low-level call
      assembly {
        revert(add(32, ret), mload(ret))
      }
    }

    emit ForwardSucceeded(req.from, req.target, domainSeparator, req.nonce, req.data, ret);

    return (success, ret);
  }

  error NonceMismatch(uint256 expected, uint256 actual);

  function _verifyNonce(ForwardRequest calldata req) internal view {
    if (s_nonces[req.from] != req.nonce) {
      revert NonceMismatch(s_nonces[req.from], req.nonce);
    }
  }

  function _verifyAndUpdateNonce(ForwardRequest calldata req) internal {
    if (s_nonces[req.from]++ != req.nonce) {
      revert NonceMismatch(s_nonces[req.from], req.nonce);
    }
  }

  error InvalidTypeName(string typeName);

  /// @inheritdoc IForwarder
  function registerRequestType(string calldata typeName, string calldata typeSuffix) external override {
    for (uint256 i = 0; i < bytes(typeName).length; i++) {
      bytes1 c = bytes(typeName)[i];
      if (c == "(" || c == ")") {
        revert InvalidTypeName(typeName);
      }
    }

    string memory requestType = string(abi.encodePacked(typeName, "(", GENERIC_PARAMS, ",", typeSuffix));
    registerRequestTypeInternal(requestType);
  }

  function getDomainSeparator(string calldata name, string calldata version) public view returns (bytes memory) {
    return
      abi.encode(
        keccak256(bytes(EIP712_DOMAIN_TYPE)),
        keccak256(bytes(name)),
        keccak256(bytes(version)),
        block.chainid,
        address(this)
      );
  }

  /// @inheritdoc IForwarder
  function registerDomainSeparator(string calldata name, string calldata version) external override onlyOwner {
    bytes memory domainSeparator = getDomainSeparator(name, version);
    bytes32 domainHash = keccak256(domainSeparator);
    s_domains[domainHash] = true;

    emit DomainRegistered(domainHash, domainSeparator);
  }

  function registerRequestTypeInternal(string memory requestType) internal {
    bytes32 requestTypehash = keccak256(bytes(requestType));
    s_typeHashes[requestTypehash] = true;
    emit RequestTypeRegistered(requestTypehash, requestType);
  }

  error UnregisteredDomainSeparator();
  error UnregisteredTypeHash();
  error SignatureMismatch();

  function _verifySig(
    ForwardRequest calldata req,
    bytes32 domainSeparator,
    bytes32 requestTypeHash,
    bytes calldata suffixData,
    bytes calldata sig
  ) internal view virtual {
    if (!s_domains[domainSeparator]) {
      revert UnregisteredDomainSeparator();
    }
    if (!s_typeHashes[requestTypeHash]) {
      revert UnregisteredTypeHash();
    }
    bytes32 digest = keccak256(
      abi.encodePacked("\x19\x01", domainSeparator, keccak256(_getEncoded(req, requestTypeHash, suffixData)))
    );
    // solhint-disable-next-line avoid-tx-origin

    if (tx.origin != DRY_RUN_ADDRESS && digest.recover(sig) != req.from) {
      revert SignatureMismatch();
    }
  }

  /// @notice Creates a byte array that is a valid ABI encoding of a request of a `RequestType` type. See `execute()`.
  function _getEncoded(
    ForwardRequest calldata req,
    bytes32 requestTypeHash,
    bytes calldata suffixData
  ) public pure returns (bytes memory) {
    // we use encodePacked since we append suffixData as-is, not as dynamic param.
    // still, we must make sure all first params are encoded as abi.encode()
    // would encode them - as 256-bit-wide params.
    return
      abi.encodePacked(
        requestTypeHash,
        uint256(uint160(req.from)),
        uint256(uint160(req.target)),
        req.nonce,
        keccak256(req.data),
        req.validUntilTime,
        suffixData
      );
  }
}
