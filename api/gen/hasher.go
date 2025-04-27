package gen

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
)

const (
	saltSize   = 16
	iterations = 100_000
	keyLen     = 32
)

type PasswordHasher struct{}

func NewPasswordHasher() *PasswordHasher {
	return &PasswordHasher{}
}

func (p *PasswordHasher) generateSalt() []byte {
	salt := make([]byte, saltSize)
	_, _ = rand.Read(salt)
	return salt
}

func (p *PasswordHasher) pbkdf2(password, salt []byte, iter, length int) []byte {
	hLen := sha256.Size
	numBlocks := (length + hLen - 1) / hLen
	var out []byte

	for block := 1; block <= numBlocks; block++ {
		T := make([]byte, hLen)
		U := make([]byte, hLen)

		blockBytes := []byte{0, 0, 0, byte(block)}
		mac := hmac.New(sha256.New, password)
		mac.Write(salt)
		mac.Write(blockBytes)
		U = mac.Sum(nil)
		copy(T, U)

		for i := 1; i < iter; i++ {
			mac = hmac.New(sha256.New, password)
			mac.Write(U)
			U = mac.Sum(nil)
			for x := range T {
				T[x] ^= U[x]
			}
		}

		out = append(out, T...)
	}

	return out[:length]
}

func (p *PasswordHasher) Hash(password string) string {
	salt := p.generateSalt()
	hash := p.pbkdf2([]byte(password), salt, iterations, keyLen)

	return fmt.Sprintf("%s$%s",
		base64.StdEncoding.EncodeToString(salt),
		base64.StdEncoding.EncodeToString(hash),
	)
}

func (p *PasswordHasher) Verify(password, stored string) bool {
	parts := strings.Split(stored, "$")
	if len(parts) != 2 {
		return false
	}
	salt, _ := base64.StdEncoding.DecodeString(parts[0])
	expected := p.HashWithSalt(password, salt)
	return stored == expected
}

func (p *PasswordHasher) HashWithSalt(password string, salt []byte) string {
	hash := p.pbkdf2([]byte(password), salt, iterations, keyLen)
	return fmt.Sprintf("%s$%s",
		base64.StdEncoding.EncodeToString(salt),
		base64.StdEncoding.EncodeToString(hash),
	)
}
