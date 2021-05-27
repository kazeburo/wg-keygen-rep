package main

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"runtime"

	"github.com/jessevdk/go-flags"
	"golang.org/x/crypto/curve25519"
)

var version string

const UNKNOWN = 3
const OK = 0

const keyLen = 32

type keyByte [keyLen]byte

type keyPair struct {
	Priv string `json:"priv"`
	Pub  string `json:"pub"`
}

type commandOpts struct {
	Salt    string `short:"s" long:"salt" description:"salt string for generating private key"`
	JSON    bool   `long:"json" description:"output with JSON format"`
	Version bool   `short:"v" long:"version" description:"Show version"`
}

func printVersion() {
	fmt.Printf(`%s Compiler: %s %s`,
		version,
		runtime.Compiler,
		runtime.Version())
}

func GenerateKeyPair(salt string) keyPair {
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

func main() {
	os.Exit(_main())
}

func _main() int {
	opts := commandOpts{}
	psr := flags.NewParser(&opts, flags.Default)
	_, err := psr.Parse()
	if err != nil {
		os.Exit(UNKNOWN)
	}

	if opts.Version {
		printVersion()
		return OK
	}

	r := GenerateKeyPair(opts.Salt)
	if opts.JSON {
		j, _ := json.Marshal(r)
		fmt.Println(string(j))
	} else {
		fmt.Printf("priv: %s\n", r.Priv)
		fmt.Printf("pub: %s\n", r.Pub)
	}
	return OK
}
