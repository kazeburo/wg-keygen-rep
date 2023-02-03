package main

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"

	"github.com/google/uuid"
	"github.com/jessevdk/go-flags"
	"github.com/kazeburo/wg-keygen-rep/internal/keypair"
)

var version string

const StatusCodeOK = 0
const StatusCodeUnknown = 3

type Opt struct {
	Salt    string `short:"s" long:"salt" description:"salt string for generating private key. not specified uuid is used by default"`
	JSON    bool   `long:"json" description:"output with JSON format"`
	Version bool   `short:"v" long:"version" description:"Show version"`
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

	r := keypair.NewPair(opt.Salt)
	if opt.JSON {
		j, _ := json.Marshal(r)
		fmt.Println(string(j))
	} else {
		fmt.Printf("priv: %s\n", r.Priv)
		fmt.Printf("pub: %s\n", r.Pub)
	}
	return StatusCodeOK
}
