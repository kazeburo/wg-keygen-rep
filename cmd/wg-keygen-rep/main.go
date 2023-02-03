package main

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"runtime"

	"github.com/google/uuid"
	"github.com/jessevdk/go-flags"
	"golang.org/x/crypto/curve25519"
)

var version string

const StatusCodeOK = 0
const StatusCodeUnknown = 3

const keyLen = 32

type keyByte [keyLen]byte

type keyPair struct {
	Priv string `json:"priv"`
	Pub  string `json:"pub"`
}

type Opt struct {
	Salt    string `short:"s" long:"salt" description:"salt string for generating private key. not specified uuid is used by default"`
	JSON    bool   `long:"json" description:"output with JSON format"`
	Version bool   `short:"v" long:"version" description:"Show version"`
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
	opt := &Opt{}
	psr := flags.NewParser(opt, flags.HelpFlag|flags.PassDoubleDash)
	_, err := psr.Parse()
	if opt.Version {
		fmt.Printf(`%s %s
Compiler: %s %s
`,
			os.Args[0],
			version,
			runtime.Compiler,
			runtime.Version())
		os.Exit(StatusCodeOK)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(StatusCodeUnknown)
	}

	if opt.Salt == "" {
		opt.Salt = uuid.NewString()
	}

	r := GenerateKeyPair(opt.Salt)
	if opt.JSON {
		j, _ := json.Marshal(r)
		fmt.Println(string(j))
	} else {
		fmt.Printf("priv: %s\n", r.Priv)
		fmt.Printf("pub: %s\n", r.Pub)
	}
	return StatusCodeOK
}
