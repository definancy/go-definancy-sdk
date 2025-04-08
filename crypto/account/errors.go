package account

import (
	"errors"
)

// var errInvalidPrivateKey = errors.New("invalid private key")
var errInvalidPublicKey = errors.New("invalid public key")
var errInvalidSignatureReturned = errors.New("ed25519 library returned an invalid signature")
