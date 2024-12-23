package secretdata

type Cipher struct {
	SecretKey string
}

func NewSecretData(
	SecretKey string) *Cipher {
	return &Cipher{
		SecretKey: SecretKey,
	}
}
