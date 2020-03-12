// Copyright 2015 ALRUX Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// AGExt additional methods

package regexp

// FindNamed returns a map of byte slices holding the text of the leftmost match in src
// of the regular expression as the value for the empty-string key (""),
// plus all named submatches (keys are sub-expression names, values are the
// corresponding matched texts).
// A return value of nil indicates no match.
func (re *Regexp) FindNamed(src []byte) map[string][]byte {
	match := re.doExecute(nil, src, "", 0, re.prog.NumCap, nil)
	if len(match) == 0 {
		return nil
	}
	result := make(map[string][]byte, len(re.subexpNames))
	for i, name := range re.subexpNames {
		if (i == 0 || name != "") && 2*i+1 < len(match) && match[2*i] >= 0 {
			result[name] = src[match[2*i]:match[2*i+1]]
		}
	}
	return result
}

// FindStringNamed returns a map of strings holding the text of the leftmost match in src
// of the regular expression as the value for the empty-string key (""),
// plus all named submatches (keys are sub-expression names, values are the
// corresponding matched texts).
// If there is no match, the return value is nil.
func (re *Regexp) FindStringNamed(src string) map[string]string {
	match := re.doExecute(nil, nil, src, 0, re.prog.NumCap, nil)
	if len(match) == 0 {
		return nil
	}
	result := make(map[string]string, len(re.subexpNames))
	for i, name := range re.subexpNames {
		if (i == 0 || name != "") && 2*i+1 < len(match) && match[2*i] >= 0 {
			result[name] = src[match[2*i]:match[2*i+1]]
		}
	}
	return result
}

// FindAllNamed is the 'All' version of FindNamed; it returns a slice of all successive
// matches of the expression, as defined by the 'All' description in the
// package comment, with slice elements as described for FindNamed.
// A return value of nil indicates no match.
func (re *Regexp) FindAllNamed(src []byte, n int) []map[string][]byte {
	if n < 0 {
		n = len(src) + 1
	}
	matches := make([]map[string][]byte, 0, startSize)
	re.allMatches("", src, n, func(match []int) {
		result := make(map[string][]byte, len(re.subexpNames))
		for i, name := range re.subexpNames {
			if (i == 0 || name != "") && 2*i+1 < len(match) && match[2*i] >= 0 {
				result[name] = src[match[2*i]:match[2*i+1]]
			}
		}
		matches = append(matches, result)
	})
	if len(matches) == 0 {
		return nil
	}
	return matches
}

// FindAllStringNamed is the 'All' version of FindStringNamed; it returns a slice of all
// successive matches of the expression, as defined by the 'All' description
// in the package comment, with slice elements as described for FindStringNamed.
// A return value of nil indicates no match.
func (re *Regexp) FindAllStringNamed(src string, n int) []map[string]string {
	if n < 0 {
		n = len(src) + 1
	}
	matches := make([]map[string]string, 0, startSize)
	re.allMatches(src, nil, n, func(match []int) {
		result := make(map[string]string, len(re.subexpNames))
		for i, name := range re.subexpNames {
			if (i == 0 || name != "") && 2*i+1 < len(match) && match[2*i] >= 0 {
				result[name] = src[match[2*i]:match[2*i+1]]
			}
		}
		matches = append(matches, result)
	})
	if len(matches) == 0 {
		return nil
	}
	return matches
}

// ReplaceAllStringSubmatchFunc returns a copy of src in which all matches of the
// Regexp have been replaced by the return value of function repl applied
// to the submatches.  The replacement returned by repl is substituted
// directly, without using Expand.
func (re *Regexp) ReplaceAllStringSubmatchFunc(src string, repl func([]string) string) string {
	b := re.replaceAll(nil, src, 2*(re.numSubexp+1), func(dst []byte, match []int) []byte {
		matches := make([]string, len(match)/2)
		for i := range matches {
			if match[2*i] >= 0 {
				matches[i] = src[match[2*i]:match[2*i+1]]
			}
		}
		return append(dst, repl(matches)...)
	})
	return string(b)
}

// ReplaceAllSubmatchFunc returns a copy of src in which all matches of the
// Regexp have been replaced by the return value of function repl applied
// to the submatches.  The replacement returned by repl is substituted
// directly, without using Expand.
func (re *Regexp) ReplaceAllSubmatchFunc(src []byte, repl func([][]byte) []byte) []byte {
	return re.replaceAll(src, "", 2*(re.numSubexp+1), func(dst []byte, match []int) []byte {
		matches := make([][]byte, len(match)/2)
		for i := range matches {
			if match[2*i] >= 0 {
				matches[i] = src[match[2*i]:match[2*i+1]]
			}
		}
		return append(dst, repl(matches)...)
	})
}

// ReplaceAllStringNamedFunc returns a copy of src in which all matches of the
// Regexp have been replaced by the return value of function repl applied
// to the submatches.  The replacement returned by repl is substituted
// directly, without using Expand.
func (re *Regexp) ReplaceAllStringNamedFunc(src string, repl func(map[string]string) string) string {
	b := re.replaceAll(nil, src, 2*(re.numSubexp+1), func(dst []byte, match []int) []byte {
		matches := make(map[string]string, len(re.subexpNames))
		for i, name := range re.subexpNames {
			if (i == 0 || name != "") && 2*i+1 < len(match) && match[2*i] >= 0 {
				matches[name] = src[match[2*i]:match[2*i+1]]
			}
		}
		return append(dst, repl(matches)...)
	})
	return string(b)
}

// ReplaceAllNamedFunc returns a copy of src in which all matches of the
// Regexp have been replaced by the return value of function repl applied
// to the submatches.  The replacement returned by repl is substituted
// directly, without using Expand.
func (re *Regexp) ReplaceAllNamedFunc(src []byte, repl func(map[string][]byte) []byte) []byte {
	return re.replaceAll(src, "", 2*(re.numSubexp+1), func(dst []byte, match []int) []byte {
		matches := make(map[string][]byte, len(re.subexpNames))
		for i, name := range re.subexpNames {
			if (i == 0 || name != "") && 2*i+1 < len(match) && match[2*i] >= 0 {
				matches[name] = src[match[2*i]:match[2*i+1]]
			}
		}
		return append(dst, repl(matches)...)
	})
}
