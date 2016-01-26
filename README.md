# Extended `regexp` package for Go (golang)

This is an extension of the standard [Go](http://golang.org) package with the same name. It is a modified copy that can be used as a simple drop-in replacement, being 100% backwards-compatible with the standard package. It adds new methods to the `Regexp` type to allow for more efficient processing of submatches.

## Maturity

[![Build Status](https://travis-ci.org/agext/regexp.svg?branch=master)](https://travis-ci.org/agext/regexp)

Stable: no known bugs or performance issues introduced by the added code.

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
