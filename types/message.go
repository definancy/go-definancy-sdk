package types

type Message struct {
	Version string `json:"version"`
	Type    string `json:"type"`
	Data    any    `json:"data"`
}

type SignedMessage struct {
	Message Message  `json:"message"`
	Signers Identity `json:"signers"`
	Sigs    []Sig    `json:"sigs"`
}

func (message Message) PrepareToSign() (SignedMessage, error) {
	return SignedMessage{
		Message: message,
		Signers: ZeroIdentity,
		Sigs:    []Sig{},
	}, nil
}

func (message Message) Sign(account Account) (SignedMessage, error) {
	signedMessage, err := message.PrepareToSign()
	if err != nil {
		return signedMessage, err
	}

	return signedMessage.Sign(account)
}

func (signedMessage SignedMessage) VerifySigners() bool {
	signers := make([]Signer, 0)
	for _, sig := range signedMessage.Sigs {
		signers = append(signers, sig.Signer)
	}

	identityFromSigners, err := MakeIdentity(signers...)
	if err != nil {
		return false
	}

	if !signedMessage.Signers.Equals(identityFromSigners) {
		return false
	}

	return true
}

func (signedMessage SignedMessage) Sign(account Account) (SignedMessage, error) {
	signature, err := account.Sign(signedMessage.Message)
	if err != nil {
		return signedMessage, err
	}

	signer := account.GetSigner()

	sigs := append(
		signedMessage.Sigs,
		Sig{
			Signer:    signer,
			Signature: signature,
		},
	)

	identity, err := signedMessage.Signers.AddSigners(signer)
	if err != nil {
		return signedMessage, err
	}
	signedMessage.Signers = identity
	signedMessage.Sigs = sigs

	return signedMessage, nil
}
