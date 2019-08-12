// staticcheck analyses Go code and makes it better.
package main // import "github.com/meshifyiot/go-tools/cmd/staticcheck"

import (
	"log"
	"os"

	"golang.org/x/tools/go/analysis"
	"github.com/meshifyiot/go-tools/lint"
	"github.com/meshifyiot/go-tools/lint/lintutil"
	"github.com/meshifyiot/go-tools/simple"
	"github.com/meshifyiot/go-tools/staticcheck"
	"github.com/meshifyiot/go-tools/stylecheck"
	"github.com/meshifyiot/go-tools/unused"
)

func main() {
	fs := lintutil.FlagSet("staticcheck")
	wholeProgram := fs.Bool("unused.whole-program", false, "Run unused in whole program mode")
	debug := fs.String("debug.unused-graph", "", "Write unused's object graph to `file`")
	fs.Parse(os.Args[1:])

	var cs []*analysis.Analyzer
	for _, v := range simple.Analyzers {
		cs = append(cs, v)
	}
	for _, v := range staticcheck.Analyzers {
		cs = append(cs, v)
	}
	for _, v := range stylecheck.Analyzers {
		cs = append(cs, v)
	}

	u := unused.NewChecker(*wholeProgram)
	if *debug != "" {
		f, err := os.OpenFile(*debug, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			log.Fatal(err)
		}
		u.Debug = f
	}
	cums := []lint.CumulativeChecker{u}
	lintutil.ProcessFlagSet(cs, cums, fs)
}
