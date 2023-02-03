package keypair

import (
	"crypto/sha256"
	"encoding/base64"

	"golang.org/x/crypto/curve25519"
)

const keyLen = 32

type keyByte [keyLen]byte

type keyPair struct {
	Priv string `json:"priv"`
	Pub  string `json:"pub"`
}

func NewPair(salt string) keyPair {
	priv := genPrivateKey(salt)
	pub := genPublicKey(priv)

	return keyPair{
		Priv: encodeBase64(priv),
		Pub:  encodeBase64(pub),
	}
}

func genPrivateKey(salt string) keyByte {
	h := sha256.New()
	h.Write([]byte(salt))
	s := h.Sum(nil)
	var key keyByte
	copy(key[:], s)
	// Modify random bytes using algorithm described at:
	// https://cr.yp.to/ecdh.html.
	key[0] &= 248
	key[31] &= 127
	key[31] |= 64

	return key
}

func genPublicKey(pk keyByte) keyByte {
	var pub [keyLen]byte
	priv := [keyLen]byte(pk)
	curve25519.ScalarBaseMult(&pub, &priv)
	return keyByte(pub)
}

func encodeBase64(key keyByte) string {
	return base64.StdEncoding.EncodeToString(key[:])
}
