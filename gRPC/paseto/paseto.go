package paseto

import (
	"crypto/rand"
	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
	"high-traffic-practice/config"
	auth "high-traffic-practice/gRPC/proto"
	"log"
)

type PasetoMaker struct {
	Pt  *paseto.V2
	Key []byte
}

func NewPasetoMaker(cfg *config.Config) *PasetoMaker {
	key := make([]byte, chacha20poly1305.KeySize)
	if _, err := rand.Read(key); err != nil {
		log.Fatalf("failed to generate key: %v", err)
	}
	return &PasetoMaker{
		Pt:  paseto.NewV2(),
		Key: key,
	}
}

func (m *PasetoMaker) CreateNewToken(auth *auth.AuthData) (string, error) {
	randomBytes := make([]byte, 16)
	if _, err := rand.Read(randomBytes); err != nil {
		return "", err
	}

	return m.Pt.Encrypt(m.Key, auth, randomBytes)
}

func (m *PasetoMaker) VerifyToken(token string) error {
	var authData *auth.AuthData
	return m.Pt.Decrypt(token, m.Key, authData, nil)
}
