pragma solidity >=0.5.2 <0.7.4;

contract Registrar {
    mapping ( bytes32 => bool) internal _voterRegistry;
 
    function registerVoter(bytes32 _voter) public returns (bool) {
        _voterRegistry[_voter] = true;
        return true;
    }

    function checkVoterRegistered(bytes32 _voter) public view returns (bool) {
        return _voterRegistry[_voter];
    }
}