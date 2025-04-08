package types

type Signer struct {
	Class string `json:"class"`
	Data  []byte `json:"data"`
}

func signerBytes(signer Signer) ([]byte, error) {
	return append([]byte(signer.Class), signer.Data...), nil
}
