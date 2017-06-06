package main


import (
	"github.com/docopt/docopt-go"
	"fmt"
	//"sort"
	//"net/http"
	"log"
	//"reflect"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	//"os"
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
    Test_name string
    Comment string
    Command_sequence []Command
}


type Command struct {
    Order_id int
    Type string
    Request_type string
    Url string
    Data string
    Expect Expect
    Headers map[string]string
    Repeat_times int
    Waiting_time int
}

type Expect struct {
	File string
	Respones_code int
	Type string
}




func main() {

	args, err := docopt.Parse(usage, nil, false, version, false)
	if err != nil {
		log.Fatal("error parsing arguments",err)
	}
	files, _ := args["--testSeriesfile"].([]string)
	for _,file := range files {
		fmt.Println("the file with instruction: " , file)

		
		yamlFile, err := ioutil.ReadFile(file)
		if err != nil {
		    panic(err)
		}
		
		testDescriber := TestDescriber{}

	    err = yaml.Unmarshal([]byte(yamlFile), &testDescriber)
	    if err != nil {
	        panic(err)
	    }

	    fmt.Printf("--- t:\n%v\n\n", testDescriber.Command_sequence[0].Type)
	}

	/*
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
	
	*/
	//c, err := config.New(file)
	//if err != nil {
	//	log.Fatal("error loading Testseries File",err)
	//}
}

