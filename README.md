[![Build Status](https://travis-ci.org/hreese/goodregex.svg?branch=master)](https://travis-ci.org/hreese/goodregex)

A collection of regular expressions for Go.

The name does not make any claims about the quality of the regular expressions
included. The author needed a short intuitive name. Feedback, tests and
comments are always welcome.

# Installation

```
go get github.com/hreese/goodregex
```

# Usage

[![godoc](https://img.shields.io/badge/docs-GoDoc-blue.svg)](https://godoc.org/github.com/hreese/goodregex)

This package currently exports two regular expressions:

* ```MatchIPv4``` (matches IPv4 addresses)
* ```MatchIPv6``` (matches IPv6 addresses)

Use them like normal regular expressions from [regexp](https://golang.org/pkg/regexp/).
