package crypto

import (
	"errors"
	"testing"
)

func TestEncrypt(t *testing.T) {
	passphrase := "passphrase"
	plaintext := "plaintext"
	first, err := Encrypt([]byte(passphrase), []byte(plaintext))
	if err != nil {
		t.Fatalf("expected no error got: %v", err)
	}

	second, err := Encrypt([]byte(passphrase), []byte(plaintext))
	if err != nil {
		t.Fatalf("expected no error got: %v", err)
	}

	if string(first) == string(second) {
		t.Error("expected different ciphertexts for same input")
	}
}

func TestDecrypt(t *testing.T) {
	passphrase := "passphrase"
	wrongPassphrase := "asdasdfdsf"
	plaintext := "plaintext"
	encrypted, err := Encrypt([]byte(passphrase), []byte(plaintext))
	if err != nil {
		t.Fatalf("expected no error got:%v", err)
	}
	_, err = Decrypt([]byte(wrongPassphrase), encrypted)

	if err == nil {
		t.Fatalf("expected error")
	}

	if !errors.Is(err, ErrDecryptionFailed) {
		t.Errorf("expected %v, got %v", ErrDecryptionFailed, err)
	}
}

func TestRoundTrip(t *testing.T) {
	passphrase := "passphrase"
	plaintext := "plaintext"
	encrypted, err := Encrypt([]byte(passphrase), []byte(plaintext))
	if err != nil {
		t.Fatalf("expected no error got: %v", err)
	}

	decrypted, err := Decrypt([]byte(passphrase), encrypted)
	if err != nil {
		t.Fatalf("expected no error got:%v", err)
	}
	if string(decrypted) != plaintext {
		t.Errorf("want: %q, got: %q", plaintext, string(decrypted))
	}
}
