package gen

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

type TokenGenerator struct {
	secret []byte
}

type TokenScope struct {
	Read  bool
	Write bool
}

func NewTokenGenerator(secret []byte) *TokenGenerator {
	return &TokenGenerator{secret: secret}
}

func (tg *TokenGenerator) generateRandomHex(n int) string {
	b := make([]byte, n)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func (tg *TokenGenerator) scopeString(scope TokenScope) string {
	var s []string
	if scope.Read {
		s = append(s, "read")
	}
	if scope.Write {
		s = append(s, "write")
	}
	if len(s) == 0 {
		s = append(s, "none")
	}
	return strings.Join(s, ",")
}

func (tg *TokenGenerator) sign(data string) string {
	h := hmac.New(sha256.New, tg.secret)
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func (tg *TokenGenerator) Generate(scope TokenScope) string {
	id := tg.generateRandomHex(16)
	scopeStr := tg.scopeString(scope)
	payload := fmt.Sprintf("%s:%s", id, scopeStr)
	sig := tg.sign(payload)
	return fmt.Sprintf("%s:%s", payload, sig)
}

func (tg *TokenGenerator) Validate(token string) (valid bool, id string, scopes []string) {
	parts := strings.Split(token, ":")
	if len(parts) != 3 {
		return false, "", nil
	}

	id, scopeStr, sig := parts[0], parts[1], parts[2]
	payload := fmt.Sprintf("%s:%s", id, scopeStr)
	expectedSig := tg.sign(payload)

	if !hmac.Equal([]byte(sig), []byte(expectedSig)) {
		return false, "", nil
	}

	return true, id, strings.Split(scopeStr, ",")
}
