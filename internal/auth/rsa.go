package auth

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

type RSAService struct {
	privateKey *rsa.PrivateKey
	publicPEM  string
}

func NewRSAService(privatePath, publicPath string) (*RSAService, error) {
	privPEM, err := os.ReadFile(privatePath)
	if err != nil {
		return nil, fmt.Errorf("read private key: %w", err)
	}
	priv, err := parseRSAPrivateKey(privPEM)
	if err != nil {
		return nil, fmt.Errorf("parse private key: %w", err)
	}

	var pubPEM []byte
	if publicPath != "" {
		pubPEM, err = os.ReadFile(publicPath)
		if err != nil {
			return nil, fmt.Errorf("read public key: %w", err)
		}
	} else {
		pubBytes, err := x509.MarshalPKIXPublicKey(&priv.PublicKey)
		if err != nil {
			return nil, fmt.Errorf("marshal public: %w", err)
		}
		pubPEM = pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: pubBytes})
	}

	return &RSAService{
		privateKey: priv,
		publicPEM:  string(pubPEM),
	}, nil
}

func parseRSAPrivateKey(pemBytes []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, errors.New("invalid PEM data")
	}
	if key, err := x509.ParsePKCS1PrivateKey(block.Bytes); err == nil {
		return key, nil
	}
	if anyKey, err := x509.ParsePKCS8PrivateKey(block.Bytes); err == nil {
		if rk, ok := anyKey.(*rsa.PrivateKey); ok {
			return rk, nil
		}
		return nil, errors.New("PKCS8 key is not RSA")
	}
	return nil, errors.New("unrecognized RSA private key format")
}

// Decrypt mirrors PHP openssl_private_decrypt + base64_decode(strict). Default
// PHP padding is PKCS1 v1.5, which matches the jsencrypt frontend lib.
func (s *RSAService) Decrypt(base64Cipher string) (string, error) {
	cipherBytes, err := base64.StdEncoding.DecodeString(base64Cipher)
	if err != nil {
		return "", fmt.Errorf("invalid base64: %w", err)
	}
	plain, err := rsa.DecryptPKCS1v15(rand.Reader, s.privateKey, cipherBytes)
	if err != nil {
		return "", fmt.Errorf("rsa decrypt: %w", err)
	}
	return string(plain), nil
}

func (s *RSAService) PublicKeyPEM() string {
	return s.publicPEM
}
