package main

import (
	"encoding/base64"
	"fmt"
	"go-definancy-sdk/crypto/account"
	"go-definancy-sdk/encoding/json"
	"go-definancy-sdk/types"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("Please write a message to be verified")
	}

	var signedMessage types.SignedMessage
	signedDocBytes := []byte(os.Args[1])

	ok := true
	err := json.Decode(signedDocBytes, &signedMessage)
	if err != nil {
		ok = false
	}

	ok = ok && signedMessage.VerifySigners()
	fmt.Printf("Signers: %v\n", ok)

	for _, sig := range signedMessage.Sigs {
		ok = ok && sig.Signer.Class == account.Class

		accountSigner := account.AccountSigner{
			Signer: sig.Signer,
		}

		signerData := base64.StdEncoding.EncodeToString(sig.Signer.Data)
		ok = ok && accountSigner.Verify(signedMessage.Message, sig.Signature)
		fmt.Printf("Signer[%s{%s}]: %v\n", sig.Signer.Class, signerData, ok)
	}

	fmt.Print("Result: ")
	if ok {
		fmt.Print("Message ok\n")
	} else {
		fmt.Print("Message tampered\n")
	}
}
