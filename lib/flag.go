package lib

import (
	"flag"
	"fmt"
	"github.com/binganao/breachkit/config"
	"os"
)

var Target string
var Targets string
var OutPut string
var TimeOut int
var NoIcmp bool

func init() {
	flag.StringVar(&Target, "target", "", "Specify a target [eg. -target https://example.com/file.php]")
	flag.StringVar(&Targets, "targets", "", "Reads all targets from the specified file [eg. -targets filename.txt]")
	flag.StringVar(&OutPut, "output", "output.txt", "Save the scan results to the specified file")
	flag.IntVar(&TimeOut, "timeout", 5, "Response timeout time [eg. -timeout 10]")
	flag.BoolVar(&NoIcmp, "np", false, "Set all hosts to live [eg. -np]")
	flag.Usage = usage
}

func help() {
	flag.Usage()
}

func usage() {
	fmt.Fprintf(os.Stderr, ""+
		"breachkit version: breachkit/"+config.BreachkitVersion+"\n"+
		"Usage: breachkit [-target] [-targets] [-output] [-timeout] [-np] \n"+
		"\n"+
		"Options:\n")
	flag.PrintDefaults()
}
