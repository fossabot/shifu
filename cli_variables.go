package main


import (
	"github.com/docopt/docopt-go"
	"fmt"
	"net/http"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

var version = "0.0.5"

const usage = `
  Usage:
    shifu run [--testSeriesfile=<file>]
    shifu -h | --help
    shifu --version
  Options:
    -t, --testSeriesfile file   test file to execute
    -h, --help          		output help information
    -v, --version       		output version
  Examples:
    output tasks
    $ shifu
    run a test series
    $ shifu run -t test1.yaml
`

