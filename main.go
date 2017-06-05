package main

import (
	"github.com/tj/docopt"
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
		cli.Fatalf("error parsing arguments: %s", err)
	}

	file := args["--testSeriesfile"].(string)
	c, err := config.New(file)
	if err != nil {
		cli.Fatalf("error loading Testseries File: %s", err)
	}
}