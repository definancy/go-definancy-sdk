package types

import (
	"crypto/sha512"
	"encoding/base32"
)

const (
	identityPrefix   = "mId"
	maxSigners       = 255
	checksumLenBytes = 4
	hashLenBytes     = sha512.Size256
)

type Identity []Signer

var ZeroIdentity = []Signer{}

func MakeIdentity(signers ...Signer) (Identity, error) {
	if len(signers) > maxSigners {
		return nil, errTooManySigners
	}

	identity := make(Identity, 0)
	for _, signer := range signers {
		identity = append(
			identity,
			signer,
		)
	}

	return identity, nil
}

func (identity Identity) AddSigners(signers ...Signer) (Identity, error) {
	if len(identity)+len(signers) > maxSigners {
		return identity, errTooManySigners
	}

	for _, signer := range signers {
		identity = append(
			identity,
			signer,
		)
	}

	return identity, nil
}

func (identity Identity) Equals(other Identity) bool {
	str1, err := identity.Encode()
	if err != nil {
		return false
	}

	str2, err := other.Encode()
	if err != nil {
		return false
	}

	return str1 == str2
}

func (identity Identity) Encode() (string, error) {
	buffer := make([]byte, 0)
	if len(identity) != 1 {
		buffer = append(buffer, []byte(identityPrefix)...)
		buffer = append(buffer, byte(len(identity)))
	}

	for _, signer := range identity {
		b, err := signerBytes(signer)
		if err != nil {
			return "", err
		}

		buffer = append(buffer, b...)
	}

	checksumHash := sha512.Sum512_256(buffer[:])
	checksumLenBytes := checksumHash[hashLenBytes-checksumLenBytes:]

	checksumIdentity := append(buffer[:], checksumLenBytes...)
	str := base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(checksumIdentity)

	return str, nil
}
