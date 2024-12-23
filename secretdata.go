package secretdata

type SecretData interface {
	Encrypt(data []byte) ([]byte, error)
	Decrypt(data []byte) ([]byte, error)
}
