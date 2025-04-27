package gen

import (
	"math/rand"
	"strconv"
)

type (
	PublicKeyGenerator struct{}
	PublicKey          struct {
		Key string
	}
)

func (p PublicKeyGenerator) GeneratePublicKey() PublicKey {
	return PublicKey{
		Key: formKey(rand.Uint32(), rand.Uint32()),
	}
}

func formKey(first, second uint32) string {
	return strconv.FormatUint(uint64(first), 10) + "#" + strconv.FormatUint(uint64(second), 10)
}

func NewPublicKeyGenerator() *PublicKeyGenerator {
	return &PublicKeyGenerator{}
}
