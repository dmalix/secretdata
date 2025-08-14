package secretdata

type Secretdata struct {
	SecretKey string
}

func NewSecretData(
	SecretKey string) *Secretdata {
	return &Secretdata{
		SecretKey: SecretKey,
	}
}
