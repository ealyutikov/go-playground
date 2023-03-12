package types

import (
	"crypto/ed25519"
	"crypto/sha256"

	"github.com/elyutikov/goblockchain/crypto"

	"github.com/elyutikov/goblockchain/proto"
	pb "google.golang.org/protobuf/proto"
)

// HashBlock returns a SHA256 of the header
func HashBlock(block *proto.Block) []byte {
	return HashHeader(block.Header)
}

func HashHeader(header *proto.Header) []byte {
	b, err := pb.Marshal(header)
	if err != nil {
		panic(err)
	}

	hash := sha256.Sum256(b)
	return hash[:]
}

func SignBlock(pk *crypto.PrivateKey, b *proto.Block) *crypto.Signature {
	hash := HashBlock(b)
	sig := pk.Sign(hash)
	b.PublicKey = pk.Public().Bytes()
	b.Signature = sig.Bytes()

	return sig
}

func VerifyBlock(b *proto.Block) bool {
	if len(b.PublicKey) != ed25519.PublicKeySize {
		return false
	}

	if len(b.Signature) != ed25519.SignatureSize {
		return false
	}

	sig := crypto.SignatureFromBytes(b.Signature)
	pubKey := crypto.PublicKeyFromBytes(b.PublicKey)
	hash := HashBlock(b)
	return sig.Verify(pubKey, hash)
}
