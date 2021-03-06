pragma solidity ^0.4.24;

import "openzeppelin-solidity/contracts/cryptography/ECDSA.sol";
import "openzeppelin-solidity/contracts/ownership/Ownable.sol";
import "openzeppelin-solidity/contracts/math/SafeMath.sol";


contract PlasmaBlocks is Ownable {
    using SafeMath for uint256;

    uint256[] private _blocks;

    event BlocksSubmitted(uint256 indexed length, uint256 time);

    function blocksLength() public view returns(uint) {
        return _blocks.length;
    }

    function blocks(uint i) public view returns(uint256) {
        return _blocks[i];
    }

    function allBlocks() public view returns(uint256[]) {
        return _blocks;
    }

    function submitBlocks(uint256 fromIndex, uint256[] newBlocks) public onlyOwner returns(uint) {
        return _submitBlocks(fromIndex, newBlocks);
    }

    function submitBlocksSigned(uint256 fromIndex, uint256[] newBlocks, bytes rsv) public returns(uint) {
        bytes32 messageHash = keccak256(abi.encodePacked(fromIndex, newBlocks));
        bytes32 signedHash = ECDSA.toEthSignedMessageHash(messageHash);
        require(owner() == ECDSA.recover(signedHash, rsv), "Invalid signature");
        return _submitBlocks(fromIndex, newBlocks);
    }

    function _submitBlocks(uint256 fromIndex, uint256[] newBlocks) internal returns(uint) {
        uint256 begin = _blocks.length.sub(fromIndex);
        _blocks.length = fromIndex.add(newBlocks.length);
        for (uint i = begin; i < newBlocks.length; i++) {
            _blocks[fromIndex + i] = newBlocks[i];
        }

        if (begin < newBlocks.length) {
            // solium-disable-next-line security/no-block-members
            emit BlocksSubmitted(_blocks.length, block.timestamp);
        }
        return newBlocks.length - begin;
    }
}
