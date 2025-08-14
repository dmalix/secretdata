package secretdata

import (
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	secretKey := "my-secret-key"
	data := []byte("this is a secret")

	cipher := NewSecretData(secretKey)

	encryptedData, err := cipher.Encrypt(data)
	if err != nil {
		t.Fatalf("Encrypt failed: %v", err)
	}

	decryptedData, err := cipher.Decrypt(encryptedData)
	if err != nil {
		t.Fatalf("Decrypt failed: %v", err)
	}

	if string(decryptedData) != string(data) {
		t.Errorf("Decrypted data does not match original data. Got %s, want %s", string(decryptedData), string(data))
	}
}

func TestDecryptWithWrongKey(t *testing.T) {
	secretKey := "my-secret-key"
	wrongSecretKey := "wrong-secret-key"
	data := []byte("this is a secret")

	cipher := NewSecretData(secretKey)
	wrongCipher := NewSecretData(wrongSecretKey)

	encryptedData, err := cipher.Encrypt(data)
	if err != nil {
		t.Fatalf("Encrypt failed: %v", err)
	}

	_, err = wrongCipher.Decrypt(encryptedData)
	if err == nil {
		t.Error("Decrypt with wrong key should have failed, but it succeeded.")
	}
}

func BenchmarkEncrypt(b *testing.B) {
	secretKey := "my-secret-key"
	data := []byte("this is a secret")
	cipher := NewSecretData(secretKey)

	for i := 0; i < b.N; i++ {
		_, err := cipher.Encrypt(data)
		if err != nil {
			b.Fatalf("Encrypt failed: %v", err)
		}
	}
}

func BenchmarkDecrypt(b *testing.B) {
	secretKey := "my-secret-key"
	data := []byte("this is a secret")
	cipher := NewSecretData(secretKey)

	encryptedData, err := cipher.Encrypt(data)
	if err != nil {
		b.Fatalf("Encrypt failed: %v", err)
	}

	for i := 0; i < b.N; i++ {
		_, err := cipher.Decrypt(encryptedData)
		if err != nil {
			b.Fatalf("Decrypt failed: %v", err)
		}
	}
}
