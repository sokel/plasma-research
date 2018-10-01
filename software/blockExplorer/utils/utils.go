package utils

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
)

// StructureToBytes allows to convert structure into bytes array
func StructureToBytes(structure interface{}) []byte {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	enc.Encode(structure)
	return network.Bytes()
}

// PrivateKeyStringToBytes allows to convert hexadecimal private key string into bytes
func PrivateKeyStringToBytes(pvK string) []byte {
	p, err := hex.DecodeString(pvK)
	if err != nil {
		panic(err.Error())
	}
	return p
}

// BytesToHexString allows to convert byte array to hexadecimal string
func BytesToHexString(s []byte) string {
	return hex.EncodeToString(s)
}