package auth

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"os"
	"path/filepath"
	"testing"
)

func TestRSAService_Decrypt(t *testing.T) {
	dir := t.TempDir()
	privPath := filepath.Join(dir, "priv.pem")
	pubPath := filepath.Join(dir, "pub.pem")
	priv, _ := rsa.GenerateKey(rand.Reader, 2048)

	privBytes := x509.MarshalPKCS1PrivateKey(priv)
	if err := os.WriteFile(privPath, pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: privBytes}), 0o600); err != nil {
		t.Fatal(err)
	}
	pubBytes, _ := x509.MarshalPKIXPublicKey(&priv.PublicKey)
	if err := os.WriteFile(pubPath, pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: pubBytes}), 0o644); err != nil {
		t.Fatal(err)
	}

	svc, err := NewRSAService(privPath, pubPath)
	if err != nil {
		t.Fatalf("new service: %v", err)
	}

	plaintext := "hunter2"
	cipher, err := rsa.EncryptPKCS1v15(rand.Reader, &priv.PublicKey, []byte(plaintext))
	if err != nil {
		t.Fatal(err)
	}
	got, err := svc.Decrypt(base64.StdEncoding.EncodeToString(cipher))
	if err != nil {
		t.Fatalf("decrypt: %v", err)
	}
	if got != plaintext {
		t.Fatalf("expected %q got %q", plaintext, got)
	}

	if _, err := svc.Decrypt("@@@invalid base64@@@"); err == nil {
		t.Fatal("expected error on invalid base64")
	}
	if svc.PublicKeyPEM() == "" {
		t.Fatal("public key empty")
	}
}
