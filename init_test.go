package secretdata

import (
	"testing"
)

func Test_SUCCESS(t *testing.T) {

	dataset := []byte("dataset")

	data := NewSecretData("secret")

	encryptedData, err := data.Encrypt(dataset)
	if err != nil {
		t.Errorf("failed to encrypt the dataset: %s", err)
	}

	decryptedData, err := data.Decrypt(encryptedData)
	if err != nil {
		t.Errorf("failed to decrypt the dataset: %s", err)
	}

	if string(decryptedData) != string(dataset) {
		t.Errorf("function returned wrong status code: got '%v' want '%v'", string(decryptedData), string(dataset))
	}
}

func Test_FAIL(t *testing.T) {

	dataset := []byte("dataset")

	data1 := NewSecretData("secret1")
	data2 := NewSecretData("secret2")

	encryptedData, err := data1.Encrypt(dataset)
	if err != nil {
		t.Errorf("failed to encrypt the dataset: %s", err)
	}

	_, err = data2.Decrypt(encryptedData)
	if err == nil {
		t.Errorf("function returned wrong status code: got 'nil' want 'error'")
	}
}
