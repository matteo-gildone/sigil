package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/pbkdf2"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
)

func generateSalt(n int) ([]byte, error) {
	salt := make([]byte, n)

	_, err := rand.Read(salt)
	if err != nil {
		return nil, fmt.Errorf("failed to generate salt: %w", err)
	}

	return salt, nil
}

func Encrypt(passphrase, plaintext []byte) ([]byte, error) {
	salt, err := generateSalt(16)
	if err != nil {
		return nil, err
	}

	key, err := pbkdf2.Key(sha256.New, string(passphrase), salt, 260000, 32)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aesgcm.NonceSize())

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	ciphertext := aesgcm.Seal(nonce, nonce, plaintext, nil)

	return append(salt, ciphertext...), nil
}

//func Decrypt(passphrase, data []byte) ([]byte, error) {}
