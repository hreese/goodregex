[![Build Status](https://travis-ci.org/hreese/goodregex.svg?branch=master)](https://travis-ci.org/hreese/goodregex)

A collection of regular expressions for Go. This is the core of
[resolveip](https://github.com/hreese/resolveip).

The name does not make any claims about the quality of the regular expressions
included. The author just needed a short intuitive name :-)

Feedback, tests and comments are always welcome.

# Installation

```
go get github.com/hreese/goodregex
```

# Usage

[![godoc](https://img.shields.io/badge/docs-GoDoc-blue.svg)](https://godoc.org/github.com/hreese/goodregex)

This package currently exports three regular expressions:

* ```MatchIPv4``` (matches IPv4 addresses)
* ```MatchIPv6``` (matches IPv6 addresses)
* ```MatchHostname``` (matches hostnames)

I would like to add versions that do not match parts inside faulty inputs (named `*Bounded*`), but it's currently WIP…

Use them like normal regular expressions from [regexp](https://golang.org/pkg/regexp/).
