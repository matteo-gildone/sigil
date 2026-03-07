package crypto

import (
	"errors"
	"testing"
)

func TestEncrypt(t *testing.T) {
	passphrase := "passphrase"
	plaintext := "plaintext"
	encrypted, err := Encrypt([]byte(passphrase), []byte(plaintext))
	if err != nil {
		t.Fatalf("expected no error got: %v", err)
	}

	encryptedTwice, err := Encrypt([]byte(passphrase), encrypted)
	if err != nil {
		t.Fatalf("expected no error got: %v", err)
	}

	if string(encrypted) == string(encryptedTwice) {
		t.Error("expected to not be the same")
	}
}

func TestDecrypt(t *testing.T) {
	passphrase := "passphrase"
	wrongPassphrase := "asdasdfdsf"
	plaintext := "plaintext"
	encrypted, err := Encrypt([]byte(passphrase), []byte(plaintext))

	_, err = Decrypt([]byte(wrongPassphrase), encrypted)

	if err == nil {
		t.Fatalf("expected error")
	}

	if !errors.Is(err, ErrDecryptionFailed) {
		t.Errorf("expected %v, got %v", ErrDecryptionFailed, err)
	}
}

func TestRoundtrip(t *testing.T) {
	passphrase := "passphrase"
	plaintext := "plaintext"
	encrypted, err := Encrypt([]byte(passphrase), []byte(plaintext))
	if err != nil {
		t.Fatalf("expected no error got: %v", err)
	}

	decrypted, err := Decrypt([]byte(passphrase), encrypted)

	if string(decrypted) != plaintext {
		t.Errorf("want: %q, got: %q", plaintext, string(decrypted))
	}
}
