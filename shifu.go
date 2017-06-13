package main


import (
	"github.com/docopt/docopt-go"
	"fmt"
	"net/http"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

func ProcessTestDescriber(testDescriber TestDescriber) {
	fmt.Println("🐼 Running Test Describer " + testDescriber.Test_name)
    fmt.Println(testDescriber.Comment)
    tr := &http.Transport{
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}

	for _, command := range testDescriber.Command_sequence {
		fmt.Println("🐯 " + command.Comment)
		req, _ := http.NewRequest(command.Method, command.Url, nil)
		for key, val := range command.Headers {
			req.Header.Add(key,val)
		}

		resp, _ := client.Do(req)
	    defer resp.Body.Close()
		byteArray, _ := ioutil.ReadAll(resp.Body)
		AssertEquealInt(resp.StatusCode, command.Expect.Respones_code)
		AssertEquealString(string(byteArray[:]), command.Expect.Value)
	}
}

func ProcessTestSeriesFile(file string) {
	fmt.Println("the file with instruction: " , file)
	yamlFile, _ := ioutil.ReadFile(file)
	testDescriber := TestDescriber{}
    yaml.Unmarshal([]byte(yamlFile), &testDescriber)
    ProcessTestDescriber(testDescriber)
}

func main() {

	args, _ := docopt.Parse(usage, nil, false, version, false)
	files, _ := args["--testSeriesfile"].([]string)
	
	for _,file := range files {
		ProcessTestSeriesFile(file)
	}
}

