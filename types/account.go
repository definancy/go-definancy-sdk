package types

type Account interface {
	GetSigner() Signer
	Sign(Message) (Signature, error)
}
