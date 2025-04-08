package types

import (
	"crypto/sha512"
	"fmt"
)

const (
	Sha512_256Size = sha512.Size256
)

var errTooManySigners = fmt.Errorf("too many signers, max signers is %d", maxSigners)

// var errInvalidPublicKey = errors.New("invalid public key")
//
// var errWrongIdentityByteLen = fmt.Errorf("encoding identity is the wrong length, should be %d bytes", hashLenBytes)
// var errWrongIdentityLen = fmt.Errorf("decoded identity is the wrong length, should be %d bytes", hashLenBytes+checksumLenBytes)
// var errWrongChecksum = fmt.Errorf("identity checksum is incorrect, did you copy the identity correctly?")
// var errWrongBlockHashLen = fmt.Errorf("decoded block hash is the wrong length, should be %d bytes", Sha512_256Size)
