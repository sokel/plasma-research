package main

import (
	"fmt"
	"../ethereum"
	"../utils"
)

func main()  {
	type Structure struct {
		a string
		b int
	}

	// Create New Wallet
	wallet, _ := ethereum.Wallet()

	// Get private key in Bytes or Hex String
	wallet.PrivateKey().Bytes()
	privateKeyHex := wallet.PrivateKey().Hex()

	// Get public key in Bytes or Hex String
	publicKeyBytes := wallet.PublicKey().Bytes()
	publicKeyHex := wallet.PublicKey().Hex()

	// Get address in Bytes or Hex String
	wallet.Address()
	addressHex := utils.BytesToHexString(wallet.Address())

	// Get signature from String or Structure
	wallet.Sign([]byte("Hello"))
	structureSig, _ := wallet.Sign(utils.StructureToBytes(Structure{"Hello", 1}))

	// Get signature Bytes or Hex String
	sigBytes := structureSig.Bytes()
	sigHex := structureSig.HexString()

	// Get signed message Hash Bytes or Hex string
	sigMsgHashBytes := structureSig.GetHash()
	sigMsgHashHex := utils.BytesToHexString(structureSig.GetHash())

	// Verify signature
	verify := ethereum.VerifySignature(publicKeyBytes, sigMsgHashBytes, sigBytes)




	// Create wallet from existing private key
	walletE, _ := ethereum.Wallet(utils.PrivateKeyStringToBytes("3d345c1036f325e046c8a013707def71a8854f563f7b17f7c61d81975c7de479"))

	walletE.PrivateKey().Bytes()
	privateKeyHexE := walletE.PrivateKey().Hex()

	publicKeyBytesE := walletE.PublicKey().Bytes()
	publicKeyHexE := walletE.PublicKey().Hex()

	walletE.Address()
	addressHexE := utils.BytesToHexString(walletE.Address())

	walletE.Sign([]byte("Hello"))
	structureSigE, _ := walletE.Sign(utils.StructureToBytes(Structure{"Hello", 1}))

	sigBytesE := structureSigE.Bytes()
	sigHexE := structureSigE.HexString()

	sigMsgHashBytesE := structureSigE.GetHash()
	sigMsgHashHexE := utils.BytesToHexString(structureSigE.GetHash())

	verifyE := ethereum.VerifySignature(publicKeyBytesE, sigMsgHashBytesE, sigBytesE)

	fmt.Print("New Wallet\n")
	fmt.Print("Private Key Hex: " + privateKeyHex + "\n\n")
	fmt.Print("Public Key Hex: " + publicKeyHex + "\n\n")
	fmt.Print("Get Address Hex: " + addressHex + "\n\n")
	fmt.Print("Get signature hex: " + sigHex + "\n\n")
	fmt.Print("Get message hash from signature: " + sigMsgHashHex + "\n\n")
	fmt.Print("Verify signature\n")
	fmt.Print(verify)

	fmt.Print("\n\nWallet from existing Private Key\n")
	fmt.Print("Private Key Hex: " + privateKeyHexE + "\n\n")
	fmt.Print("Public Key Hex: " + publicKeyHexE + "\n\n")
	fmt.Print("Get Address Hex: " + addressHexE + "\n\n")
	fmt.Print("Get signature hex: " + sigHexE + "\n\n")
	fmt.Print("Get message hash from signature: " + sigMsgHashHexE + "\n\n")
	fmt.Print("Verify signature\n")
	fmt.Print(verifyE)
}