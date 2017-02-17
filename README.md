# Extended `regexp` package for Go (golang)

[![Release](https://img.shields.io/github/release/agext/regexp.svg?style=flat)](https://github.com/agext/regexp/releases/latest)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/agext/regexp) 
[![Build Status](https://travis-ci.org/agext/regexp.svg?branch=master&style=flat)](https://travis-ci.org/agext/regexp)
[![Coverage Status](https://coveralls.io/repos/github/agext/regexp/badge.svg?style=flat)](https://coveralls.io/github/agext/regexp)
[![Go Report Card](https://goreportcard.com/badge/github.com/agext/regexp?style=flat)](https://goreportcard.com/report/github.com/agext/regexp)


This is an extension of the standard [Go](http://golang.org) package with the same name. It is a modified copy that can be used as a simple drop-in replacement, being 100% backwards-compatible with the standard package. It adds new methods to the `Regexp` type to allow for more efficient processing of submatches.

## Maturity

v1.2 Stable: Guaranteed no breaking changes to the API in future v1.x releases. No known bugs or performance issues introduced by the added code. Probably safe to use in production, though provided on "AS IS" basis.

This package is being actively maintained. If you encounter any problems or have any suggestions for improvement, please [open an issue](https://github.com/agext/regexp/issues). Pull requests are welcome.

**Note on failing test for older versions (and possibly tip)**

The `TestFoldConstants` in syntax/parse_test.go depends on the standard `unicode/utf8` package. This test fails when this package is used with a Go version in which the unicode/utf8 package handles different folding ranges (new ranges are added from time to time, as unicode/utf8 is refined). The behavior of agext/regexp will be the one you expect from your Go version, based on its unicode/utf8 package, so it is safe to ignore this failing test.

## Overview

This package provides the following additional methods on the `Regexp` type:

### `FindNamed` and `FindStringNamed`

These methods work like the standard `FindSubmatch` and `FindStringSubmatch`, except they return a map `{subexpName: subMatch...}`, with elements only for the named subexpressions, plus the whole match with an empty string key. This ensures the returned map has at least one element when there was a match, even if the pattern has no named subexpressions, and is consistent with the standard 'Submatch' methods returning the whole match as the zeroth element.

### `FindAllNamed` and `FindAllStringNamed`

These are the 'All' version of `FindNamed` and `FindStringNamed`. They return a slice of all successive matches of the expression, as defined by the 'All' description in the package comment. The slice elements are maps with the same semantics as the returns of `FindNamed` and `FindStringNamed`, respectively.

### `ReplaceAllSubmatchFunc` and `ReplaceAllStringSubmatchFunc`

These methods work like the standard `ReplaceAllFunc` and `ReplaceAllStringFunc`, but the replace function receives a slice containing the match and all submatches (like the return of 'FindSubmatch' methods), instead of just the match.

### `ReplaceAllNamedFunc` and `ReplaceAllStringNamedFunc`

These methods work like the standard `ReplaceAllFunc` and `ReplaceAllStringFunc`, but the replace function receives a map containing the match and all named submatches (like the return of 'FindNamed' methods), instead of just the match.

## Installation

```
go get github.com/agext/regexp
```

## License

Package regexp is released under the Apache 2.0 license. See the [LICENSE](LICENSE) file for details.
