package main

import (
	"fmt"
	"go-definancy-sdk/crypto/account"
	"go-definancy-sdk/encoding/json"
	"go-definancy-sdk/types"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("Please write a message to sign")
	}

	var err error
	msg := os.Args[1]

	var data any
	var msgDecoded map[string]any

	err = json.Decode([]byte(msg), &msgDecoded)
	if err == nil {
		data = msgDecoded
	} else {
		data = msg
	}

	accounts := make([]account.Account, 0)

	for range 2 {
		acc, err := account.GenerateAccount()
		if err != nil {
			panic(err)
		}

		accounts = append(accounts, acc)

		pretty := json.Encode(acc)
		fmt.Printf("Account:\n%s\n", pretty)
	}

	message := types.Message{
		//Version: "v1",
		Type: "msg",
		Data: data,
	}

	signedMessage, err := message.PrepareToSign()
	if err != nil {
		panic(err)
	}

	for _, acc := range accounts {
		signedMessage, err = signedMessage.Sign(acc)
		if err != nil {
			panic(err)
		}
	}

	pretty := json.Encode(signedMessage)
	fmt.Printf("Signed Doc:\n%s\n", pretty)
}
