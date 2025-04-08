package account

import (
	"go-definancy-sdk/encoding/msgpack"
	"go-definancy-sdk/types"

	"golang.org/x/crypto/ed25519"
)

const (
	Class = "account.ed25519"
	//checksumLenBytes = 4
	//hashLenBytes     = sha512.Size256
)

type AccountSigner struct {
	types.Signer
}

type Account struct {
	AccountSigner
	PrivateKey ed25519.PrivateKey
	PublicKey  ed25519.PublicKey
}

var zeroAccount = Account{}

func GenerateAccount() (Account, error) {
	publicKey, privateKey, err := ed25519.GenerateKey(nil)
	if err != nil {
		return zeroAccount, err
	}

	types.MakeIdentity()

	return Account{
		AccountSigner: AccountSigner{
			Signer: types.Signer{
				Class: Class,
				Data:  publicKey,
			},
		},
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}, err
}

func (signer AccountSigner) Encode() ([]byte, error) {
	if len(signer.Data) != ed25519.PublicKeySize {
		return nil, errInvalidPublicKey
	}

	return signer.Data[:], nil
}

func rawMessageBytes(message types.Message) ([]byte, error) {
	return msgpack.Encode(message), nil
}

func (signer AccountSigner) Verify(message types.Message, signature types.Signature) bool {
	signed, err := rawMessageBytes(message)
	if err != nil {
		return false
	}

	if len(signer.Data) != ed25519.PublicKeySize {
		return false
	}

	return ed25519.Verify(signer.Data, signed, signature)
}

func (account Account) Sign(message types.Message) (types.Signature, error) {
	signature := make(types.Signature, ed25519.SignatureSize)

	toBeSigned, err := rawMessageBytes(message)
	if err != nil {
		return nil, err
	}

	s := ed25519.Sign(account.PrivateKey, toBeSigned)
	if len(s) != ed25519.SignatureSize {
		return signature, errInvalidSignatureReturned
	}

	n := copy(signature[:], s)
	if n != len(signature) {
		return signature, errInvalidSignatureReturned
	}

	return signature, nil
}

func (account Account) GetSigner() types.Signer {
	return account.Signer
}
