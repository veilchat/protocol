package gen

import "math/rand"

type (
	PublicKeyGenerator struct{}
	PublicKey          struct {
		Key uint64
	}
)

func (p PublicKeyGenerator) GeneratePublicKey() PublicKey {
	return PublicKey{
		Key: rand.Uint64(),
	}
}
