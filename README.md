[![CI](https://github.com/hreese/goodregex/actions/workflows/ci.yml/badge.svg)](https://github.com/hreese/goodregex/actions/workflows/ci.yml)

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

[![Go Reference](https://pkg.go.dev/badge/github.com/hreese/goodregex.svg)](https://pkg.go.dev/github.com/hreese/goodregex)

This package currently exports three regular expressions:

* `MatchIPv4` (matches IPv4 addresses)
* `MatchIPv6` (matches IPv6 addresses)
* `MatchHostname` (matches hostnames)

Experimental bounded versions (`MatchBoundedIPv4`, `MatchBoundedHostname`) are
also available. These only match when the pattern is surrounded by word
boundaries, preventing partial matches inside longer strings.

Use them like normal regular expressions from [regexp](https://pkg.go.dev/regexp).

This package also exports `CompileReadableRegex` and `MustCompileReadableRegex`,
which compile "readable regexes" -- RE2 expressions with embedded whitespace and
`#`-style comments stripped before compilation.
