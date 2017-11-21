# GolangKatas [![Build Status](https://travis-ci.org/mresti/GolangKatas.svg?branch=master)](https://travis-ci.org/mresti/GolangKatas) [![Code Report](http://goreportcard.com/badge/mresti/GolangKatas)](https://goreportcard.com/report/mresti/GolangKatas)

Golang training repository used to practice and learn by solving some common katas.

### List of katas:

* [Fibonacci.](https://medium.com/@chmeese/fibonacci-kata-93773b30dbb2)
* [String Calculator.](http://osherove.com/tdd-kata-1/)

### Executing tests:

This project contains some tests written using [Golang testing](https://golang.org/pkg/testing/). You can easily run the tests by executing:

```
go test -race -v ./...
```

If you are using GoLand IDE or any other jetbrains IDE to run your tests you'll need to configure the IDE options adding the following configuration:

    File > Settings > Go > GOPATH > project GOPATH -> Add path for this repository

### Linter:

This repository uses [golint](https://github.com/golang/lint) that is a linter for Go source code.

### Go VERSION:

This repository requires Go version 1.7 or greater.

Developed By
------------

* Esteban Dorado Roldan

