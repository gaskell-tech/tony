package helm

import (
	"encoding/base64"
	"testing"
)

func TestEncryptAES(t *testing.T) {
	ciphertext, _ := EncryptAES("password", "plaintext")
	_, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		t.Error(err)
	}
}

func TestDecryptAES(t *testing.T) {
	plaintext, _ := DecryptAES("password", "+Cujolsjujn9Nj9QoqZLTOTiU4s8ENortswY43IFV54=")
	if plaintext != "test" {
		t.Error("failed to decrypt ciphertext")
	}
}
