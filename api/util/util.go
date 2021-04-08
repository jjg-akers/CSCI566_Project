package util

import "github.com/ethereum/go-ethereum/crypto"

//stringToKeccak256 converts a string to a keccak256 hash of type [32]byte
func StringToKeccak256(s string) [32]byte {
	var toReturn [32]byte
	copy(toReturn[:], crypto.Keccak256([]byte(s))[:])
	return toReturn
}

func KeccakToString(b []byte) string{
	h := crypto.Keccak256Hash(b)
	return h.String()
}