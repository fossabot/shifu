# Shifu

[![Join the chat at https://gitter.im/shifu-on-gitter/Lobby](https://badges.gitter.im/shifu-on-gitter/Lobby.svg)](https://gitter.im/shifu-on-gitter/Lobby?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

A simple YAML based application that will help you to test your application with HTTP requests. Configure your http requests ,expected responses and test your application.

# Installation

## Dependencies Installation
```
go get gopkg.in/yaml.v2
go get github.com/docopt/docopt-go
go get net/http
```


## Install locally:
```
go install
```


# Run
```
shifu run --testSeriesfile=<full_path_of_test_describer_yaml_file>
shifu --version
```
