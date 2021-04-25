pragma solidity >=0.5.2;

contract Registrar {
    mapping ( bytes32 => bool) internal _voterRegistry;
 
    function registerVoter(bytes32 _voter) public returns (bool) {
        return _updateVoterRegistrar(_voter);
    }

    function _updateVoterRegistrar(bytes32 _voter) internal returns (bool) {
        _voterRegistry[_voter] = true;
        return true;
    }

    function checkVoterRegistered(bytes32 _voter) public view returns (bool) {
        return _voterRegistry[_voter];
        //return true;
    }
}