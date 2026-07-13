package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"path/filepath"
	"time"
)

func generateCert() error {
	// Generate ECDSA P-256 private key
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return err
	}

	// Random serial number between 1 and 2^128
	serialNumber, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))
	if err != nil {
		return err
	}

	// Certificate template
	template := x509.Certificate{
		SerialNumber: serialNumber,

		Subject: pkix.Name{
			CommonName:   "localhost",
			Organization: []string{"Agni"},
		},

		NotBefore: time.Now().Add(-time.Hour),
		NotAfter:  time.Now().AddDate(5, 0, 0), // Valid for 5 years

		KeyUsage: x509.KeyUsageDigitalSignature |
			x509.KeyUsageKeyEncipherment,

		ExtKeyUsage: []x509.ExtKeyUsage{
			x509.ExtKeyUsageServerAuth,
		},

		BasicConstraintsValid: true,

		DNSNames: []string{
			"localhost",
		},

		IPAddresses: []net.IP{
			net.ParseIP("127.0.0.1"),
			net.ParseIP("0.0.0.0"),
		},
	}

	// Self-sign the certificate
	certDER, err := x509.CreateCertificate(
		rand.Reader,
		&template,
		&template,
		&priv.PublicKey,
		priv,
	)
	if err != nil {
		return err
	}

	// Ensure cert directory exists
	if err := os.MkdirAll("certs", 0755); err != nil {
		return err
	}

	// Certificate file
	certOut, err := os.Create(filepath.Join("certs", "server.crt"))
	if err != nil {
		return err
	}
	defer certOut.Close()

	if err := pem.Encode(certOut, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certDER,
	}); err != nil {
		return err
	}

	// Private key file
	keyBytes, err := x509.MarshalECPrivateKey(priv)
	if err != nil {
		return err
	}

	keyOut, err := os.OpenFile(
		filepath.Join("certs", "server.key"),
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
		0600,
	)
	if err != nil {
		return err
	}
	defer keyOut.Close()

	if err := pem.Encode(keyOut, &pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: keyBytes,
	}); err != nil {
		return err
	}

	return nil
}
