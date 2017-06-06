package main


import (
	"github.com/docopt/docopt-go"
	"fmt"
	//"net/http"
	"log"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"os"
	//"time"
)

var version = "0.0.1"

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


type TestDescriber struct {
    test_name string
    comment string
    command_sequence []Command
}

type Command struct {
    order_id int
    type_ string
    request_type string
    url string
    data string
    expect Expect
    headers map[string]string
    repeat_times int
    waiting_time int
}

type Expect struct {
	file string
	respones_code int
	type_ string
}



func main() {
	cmd_args := os.Args[1:]
	fmt.Println(cmd_args)

	args, err := docopt.Parse(usage, cmd_args, true, version, false)
	if err != nil {
		log.Fatal("error parsing arguments",err)
	}

	for k := range args {
	  fmt.Println(k)
	}

	file := args["--testSeriesfile"].(string)

	fmt.Println("file name", file)


	yamlFile, err := ioutil.ReadFile(file)
    if err != nil {
        panic(err)
    }


    var testDescriber TestDescriber

    err = yaml.Unmarshal(yamlFile, &testDescriber)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Value: %#v\n", testDescriber.test_name)

	//c, err := config.New(file)
	//if err != nil {
	//	log.Fatal("error loading Testseries File",err)
	//}
}