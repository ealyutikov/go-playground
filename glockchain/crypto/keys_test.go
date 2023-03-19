package crypto

import (
	"crypto/ed25519"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePrivateKey(t *testing.T) {
	privKey := GeneratePrivateKey()
	assert.Equal(t, len(privKey.Bytes()), ed25519.PrivateKeySize)

	pubKey := privKey.Public()
	assert.Equal(t, len(pubKey.Bytes()), ed25519.PublicKeySize)
}

func TestPrivateKeySign(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.Public()
	msg := []byte("test message")

	sig := privKey.Sign(msg)
	assert.True(t, sig.Verify(pubKey, msg))
	assert.False(t, sig.Verify(pubKey, []byte("new message")))

	assert.False(
		t,
		sig.Verify(GeneratePrivateKey().Public(), msg),
	)
}

func TestPublicKeyToAddress(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.Public()
	address := pubKey.Address()

	assert.Equal(t, addressLen, len(address.Bytes()))
}

func TestPrivateKeyFromString(t *testing.T) {
	var (
		seed       = "b32b405d7a3b72b39663e46d657c1b3a9137e5223bf09b8a69170d7a337b094b"
		privKey    = NewPrivateKeyFromString(seed)
		addressStr = "6e541f7a8ceebcdb3ad80082669e42b727c43f48"
	)

	assert.Equal(t, ed25519.PrivateKeySize, len(privKey.Bytes()))

	address := privKey.Public().Address()
	assert.Equal(t, addressStr, address.String())
}
