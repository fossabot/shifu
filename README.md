# Shifu

[![Build Status](https://travis-ci.org/ayonious/shifu.svg?branch=master)](https://travis-ci.org/ayonious/shifu)
[![codecov](https://codecov.io/gh/ayonious/shifu/branch/master/graph/badge.svg)](https://codecov.io/gh/ayonious/shifu)
[![Join the chat at https://gitter.im/shifu-on-gitter/Lobby](https://badges.gitter.im/shifu-on-gitter/Lobby.svg)](https://gitter.im/shifu-on-gitter/Lobby?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fayonious%2Fshifu.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fayonious%2Fshifu?ref=badge_shield)

A simple YAML based application that will help you to test your application with HTTP requests. Configure your http requests ,expected responses and test your application.

# Installation

## Dependencies
```
go get gopkg.in/yaml.v2
go get github.com/docopt/docopt-go
go get net/http
```


## Local Installation:
```
go install
```


# Run
```
shifu run --testSeriesfile=<full_path_of_test_describer_yaml_file>
shifu --version
```


## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fayonious%2Fshifu.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fayonious%2Fshifu?ref=badge_large)