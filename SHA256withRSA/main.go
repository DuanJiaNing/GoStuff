package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
)

// https://cloud.google.com/storage/docs/xml-api/post-object#policydocument
func main() {
	secretKey, err := loadSecretKey()
	if err != nil {
		panic(err)
	}

	policy := `
{"expiration": "2006-01-02T15:04:05Z",
  "conditions": []
}`
	// sign policy document using RSA with SHA-256 using a secret key
	d := sha256SumMessage(base64.StdEncoding.EncodeToString([]byte(policy)))
	messageDigest, err := rsa.SignPKCS1v15(rand.Reader, secretKey, crypto.SHA256, d)
	if err != nil {
		panic(err)
	}

	signature := base64.StdEncoding.EncodeToString(messageDigest)
	fmt.Printf("Encoded signature: %v", signature)
}

func sha256SumMessage(msg string) []byte {
	h := sha256.New()
	h.Write([]byte(msg))
	d := h.Sum(nil)
	return d
}

func loadSecretKey() (*rsa.PrivateKey, error) {
	// RSA private key, extract from service account key json file, private_key filed
	// Service account key email: *@appspot.gserviceaccount.com
	// Key ID: *
	blockPri, _ := pem.Decode([]byte(`-----BEGIN PRIVATE KEY----- 
*****
-----END PRIVATE KEY-----
`))

	// may returns a *rsa.PrivateKey, a *ecdsa.PrivateKey, or a ed25519.PrivateKey
	// see doc here: https://golang.org/src/crypto/x509/pkcs8.go
	prkI, err := x509.ParsePKCS8PrivateKey(blockPri.Bytes)
	if err != nil {
		return nil, err
	}

	return prkI.(*rsa.PrivateKey), err
}
