package ethereum

import (
	"../utils"
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto"
)

type wallet struct {
	privateKey *ecdsa.PrivateKey `json:"privateKey"`
	publicKey *ecdsa.PublicKey   `json:"publicKey"`
}

// Wallet allows to create new ecdsa private key or from existing string private key in []byte format
func Wallet(pvK ...[]byte) (*wallet, error) {
	var privateKey *ecdsa.PrivateKey
	var err error

	if pvK != nil {
		privateKey, err = crypto.ToECDSA(pvK[0])
	} else {
		privateKey, err = crypto.GenerateKey()
	}

	publicKey := &ecdsa.PublicKey{Curve: privateKey.Curve, X: privateKey.X, Y: privateKey.Y}

	if err != nil {
		return nil, err
	}

	return &wallet{
		privateKey: privateKey,
		publicKey: publicKey,
	}, nil
}

// PrivateKey allows to get string private key
func (w *wallet) PrivateKey() string  {
	return hex.EncodeToString(w.privateKey.D.Bytes())
}

// Allows to get public key associated with private key
func (w *wallet) PublicKey() string {
	return hex.EncodeToString(crypto.CompressPubkey(w.publicKey))
}

// Allows to get string address
func (w *wallet) Address() string  {
	return crypto.PubkeyToAddress(w.privateKey.PublicKey).Hex()
}

type signature struct {
	signature []byte			`json:"signature"`
	hash []byte			`json:"hash"`
}

// Sign allows to sign any byte array via user private key
func (w *wallet) Sign(data []byte) (*signature, error) {
	hash := crypto.Keccak256(data)
	s, err := crypto.Sign(hash, w.privateKey)
	if err != nil {
		return nil, err
	}
	return &signature{
		s,
		hash,
	}, nil
}

// ToBytes allows to get byte array of signature
func (s *signature) ToBytes() []byte {
	return s.signature
}

// ToHexString allows to get hexadecimal string of signature
func (s *signature) ToHexString() string {
	return utils.BytesToHexString(s.signature)
}

// GetHash allows to get signed data hash
func (s *signature) GetHash() []byte {
	return s.hash
}

// GetPublicKeyFromSignature allows to get Public Key from signature
func GetPublicKeyFromSignature(s *signature) *ecdsa.PublicKey {
	pubK, err := crypto.SigToPub(s.hash, s.signature)
	if err != nil {
		panic(err.Error())
	}

	return pubK
}

