package common

import (
	"io/ioutil"
	"log"
)

// using asymmetric crypto/RSA keys
const (
	// openssl genrsa -out app.rsa 1024
	privKeyPath = "keys/app.rsa"
	// openssl rsa -in app.rsa -pubout > app.rsa.pub
	pubKeyPath = "keys/app.rsa.pub"
)

// private key for signing and public key for verification
var (
	verifyKey, signKey []byte
)

// Read the key files before starting http handlers
func initKeys() {
	var err error

	signKey, err = ioutil.ReadFile(privKeyPath)

	if err != nil {
		log.Fatalf("[initKeys]:%s\n", err)
	}

	verifyKey, err = ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatalf("[initKeys]:%s\n", err)
		panic(errr)
	}
}
