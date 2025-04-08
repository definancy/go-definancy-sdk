package types

type Sig struct {
	Signer    Signer    `json:"signer"`
	Signature Signature `json:"signature"`
}

type Signature []byte
