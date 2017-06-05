package main


import (
	"github.com/tj/docopt"
	"request-runner"
	"fmt"
	//"net/http"
	"log"
	//"io"
	//"os"
	//"time"
)

var version = "0.0.1"

const usage = `
  Usage:
    shifu run [--testSeriesfile file]
    shifu -h | --help
    shifu --version
  Options:
    -t, --testSeriesfile file   config file to load [default: robo.yml]
    -h, --help          		output help information
    -v, --version       		output version
  Examples:
    output tasks
    $ shifu
    run a test series
    $ shifu run -t test1.yaml
`

func main() {
	args, err := docopt.Parse(usage, nil, true, version, true)
	if err != nil {
		log.Fatal("error parsing arguments",err)
	}

	fmt.Println("asd", args)

	file := args["--testSeriesfile"].(string)
	
	//c, err := config.New(file)
	//if err != nil {
	//	log.Fatal("error loading Testseries File",err)
	//}
}

