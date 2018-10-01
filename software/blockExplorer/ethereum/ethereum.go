package ethereum

import (
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto"
	"../utils"
)

type wallet struct {
	key *ecdsa.PrivateKey // ecdsa private key object
}

// Wallet allows to create new ecdsa private key or from existing string private key in []byte format
func Wallet(pvK []byte) (*wallet, error) {
	var key *ecdsa.PrivateKey
	var err error

	if pvK != nil {
		key, err = crypto.ToECDSA(pvK)
	} else {
		key, err = crypto.GenerateKey()
	}

	if err != nil {
		return nil, err
	}

	return &wallet{
		key: key,
	}, nil
}

// PrivateKey allows to get string private key
func (w *wallet) PrivateKey() string  {
	return hex.EncodeToString(w.key.D.Bytes())
}

// Allows to get string address
func (w *wallet) Address() string  {
	return crypto.PubkeyToAddress(w.key.PublicKey).Hex()
}

type signature struct {
	s []byte
}

// Sign allows to sign any byte array via user private key
func (w *wallet) Sign(data []byte) (*signature, error) {
	hash := crypto.Keccak256(data)
	s, err := crypto.Sign(hash, w.key)
	if err != nil {
		return nil, err
	}
	return &signature{
		s,
	}, nil
}

// ToBytes allows to get byte array of signature
func (s *signature) ToBytes() []byte {
	return s.s
}

// ToHexString allows to get hexadecimal string of signature
func (s *signature) ToHexString() string {
	return utils.BytesToHexString(s.s)
}

