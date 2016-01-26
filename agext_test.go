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

// Tests for AGExt additional methods

package regexp

import (
	"testing"
)

type FindNamedTest struct {
	pattern string
	input   string
	output  map[string]string
}

var findNamedTests = []FindNamedTest{
	{"[x-z]", "defabcdef", nil},
	{"[a-c]", "defabcdef", map[string]string{"": "a"}},
	{"([a-c])", "defabcdef", map[string]string{"": "a"}},
	{"(?P<test>[a-c])", "defabcdef", map[string]string{"": "a", "test": "a"}},
	{"(?P<test>[a-c]+).*", "defabcdef", map[string]string{"": "abcdef", "test": "abc"}},
}

func TestFindNamed(t *testing.T) {
	for _, tc := range findNamedTests {
		re, err := Compile(tc.pattern)
		if err != nil {
			t.Errorf("Unexpected error compiling %q: %v", tc.pattern, err)
			continue
		}
		actual := re.FindStringNamed(tc.input)
		if len(actual) != len(tc.output) {
			t.Errorf("%q.FindStringNamed(%q) = %q; want %q",
				tc.pattern, tc.input, actual, tc.output)
		} else {
			for key, value := range actual {
				if value != tc.output[key] {
					t.Errorf("%q.FindStringNamed(%q)[%q] = %q; want %q",
						tc.pattern, tc.input, key, value, tc.output[key])
				}
			}
		}
		// now try bytes
		bactual := re.FindNamed([]byte(tc.input))
		if len(bactual) != len(tc.output) {
			t.Errorf("%q.FindNamed(%q) = %q; want %q",
				tc.pattern, tc.input, bactual, tc.output)
		} else {
			for key, value := range bactual {
				if string(value) != tc.output[key] {
					t.Errorf("%q.FindNamed(%q)[%q] = %q; want %q",
						tc.pattern, tc.input, key, value, tc.output[key])
				}
			}
		}
	}
}

type FindAllNamedTest struct {
	pattern string
	input   string
	output  []map[string]string
}

var findAllNamedTests = []FindAllNamedTest{
	{"[x-z]", "defabcdef", nil},
	{"[a-c]", "defabcdef", []map[string]string{{"": "a"}, {"": "b"}, {"": "c"}}},
	{"([a-c])", "defabcdef", []map[string]string{{"": "a"}, {"": "b"}, {"": "c"}}},
	{"(?P<test>[a-c])", "defabcdef", []map[string]string{{"": "a", "test": "a"}, {"": "b", "test": "b"}, {"": "c", "test": "c"}}},
	{"(?P<test>[a-c]+).*", "defabcdef", []map[string]string{{"": "abcdef", "test": "abc"}}},
}

func TestFindAllNamed(t *testing.T) {
	for _, tc := range findAllNamedTests {
		re, err := Compile(tc.pattern)
		if err != nil {
			t.Errorf("Unexpected error compiling %q: %v", tc.pattern, err)
			continue
		}
		actual := re.FindAllStringNamed(tc.input, -1)
		if len(actual) != len(tc.output) {
			t.Errorf("%q.FindAllStringNamed(%q, -1) = %q; want %q",
				tc.pattern, tc.input, actual, tc.output)
		} else {
			for i, matches := range actual {
				if len(matches) != len(tc.output[i]) {
					for key, value := range matches {
						if value != tc.output[i][key] {
							t.Errorf("%q.FindAllStringNamed(%q, -1)[%d][%q] = %q; want %q",
								tc.pattern, tc.input, i, key, value, tc.output[i][key])
						}
					}
				}
			}
		}
		if len(tc.output) > 0 {
			actual = re.FindAllStringNamed(tc.input, 1)
			if len(actual) != 1 {
				t.Errorf("%q.FindAllStringNamed(%q, 1) = %q; want %q",
					tc.pattern, tc.input, actual, tc.output[:1])
			}
		}
		// now try bytes
		bactual := re.FindAllNamed([]byte(tc.input), -1)
		if len(bactual) != len(tc.output) {
			t.Errorf("%q.FindAllNamed(%q, -1) = %q; want %q",
				tc.pattern, tc.input, bactual, tc.output)
		} else {
			for i, matches := range bactual {
				if len(matches) != len(tc.output[i]) {
					for key, value := range matches {
						if string(value) != tc.output[i][key] {
							t.Errorf("%q.FindAllNamed(%q, -1)[%d][%q] = %q; want %q",
								tc.pattern, tc.input, i, key, value, tc.output[i][key])
						}
					}
				}
			}
		}
		if len(tc.output) > 0 {
			bactual = re.FindAllNamed([]byte(tc.input), 1)
			if len(bactual) != 1 {
				t.Errorf("%q.FindAllStringNamed(%q, 1) = %q; want %q",
					tc.pattern, tc.input, bactual, tc.output[:1])
			}
		}
	}
}

type ReplaceAllSubmatchFuncTest struct {
	pattern       string
	replacement   func([]string) string
	input, output string
}

var replaceAllSubmatchFuncTests = []ReplaceAllSubmatchFuncTest{
	{"[a-c]", func(s []string) string { return "x" + s[0] + "y" }, "defabcdef", "defxayxbyxcydef"},
	{"[a-c]+", func(s []string) string { return "x" + s[0] + "y" }, "defabcdef", "defxabcydef"},
	{"[a-c]*", func(s []string) string { return "x" + s[0] + "y" }, "defabcdef", "xydxyexyfxabcydxyexyfxy"},
	{"x([a-c]+)y", func(s []string) string { return "-" + s[1] + "+" }, "xydxyexyfxabcydxyexyfxy", "xydxyexyf-abc+dxyexyfxy"},
}

func TestReplaceAllSubmatchFunc(t *testing.T) {
	for _, tc := range replaceAllSubmatchFuncTests {
		re, err := Compile(tc.pattern)
		if err != nil {
			t.Errorf("Unexpected error compiling %q: %v", tc.pattern, err)
			continue
		}
		actual := re.ReplaceAllStringSubmatchFunc(tc.input, tc.replacement)
		if actual != tc.output {
			t.Errorf("%q.ReplaceAllSubmatchFunc(%q,fn) = %q; want %q",
				tc.pattern, tc.input, actual, tc.output)
		}
		// now try bytes
		actual = string(re.ReplaceAllSubmatchFunc([]byte(tc.input),
			func(b [][]byte) []byte {
				s := make([]string, len(b))
				for i, v := range b {
					s[i] = string(v)
				}
				return []byte(tc.replacement(s))
			}))
		if actual != tc.output {
			t.Errorf("%q.ReplaceAllSubmatchFunc(%q,fn) = %q; want %q",
				tc.pattern, tc.input, actual, tc.output)
		}
	}
}

type ReplaceAllNamedFuncTest struct {
	pattern       string
	replacement   func(map[string]string) string
	input, output string
}

var replaceAllNamedFuncTests = []ReplaceAllNamedFuncTest{
	{"[a-c]", func(s map[string]string) string { return "x" + s[""] + "y" }, "defabcdef", "defxayxbyxcydef"},
	{"[a-c]+", func(s map[string]string) string { return "x" + s[""] + "y" }, "defabcdef", "defxabcydef"},
	{"[a-c]*", func(s map[string]string) string { return "x" + s[""] + "y" }, "defabcdef", "xydxyexyfxabcydxyexyfxy"},
	{"x(?P<test>[a-c]+)y", func(s map[string]string) string { return "-" + s["test"] + "+" }, "xydxyexyfxabcydxyexyfxy", "xydxyexyf-abc+dxyexyfxy"},
}

func TestReplaceAllNamedFunc(t *testing.T) {
	for _, tc := range replaceAllNamedFuncTests {
		re, err := Compile(tc.pattern)
		if err != nil {
			t.Errorf("Unexpected error compiling %q: %v", tc.pattern, err)
			continue
		}
		actual := re.ReplaceAllStringNamedFunc(tc.input, tc.replacement)
		if actual != tc.output {
			t.Errorf("%q.ReplaceAllNamedFunc(%q,fn) = %q; want %q",
				tc.pattern, tc.input, actual, tc.output)
		}
		// now try bytes
		actual = string(re.ReplaceAllNamedFunc([]byte(tc.input),
			func(b map[string][]byte) []byte {
				s := make(map[string]string, len(b))
				for i, v := range b {
					s[i] = string(v)
				}
				return []byte(tc.replacement(s))
			}))
		if actual != tc.output {
			t.Errorf("%q.ReplaceAllNamedFunc(%q,fn) = %q; want %q",
				tc.pattern, tc.input, actual, tc.output)
		}
	}
}
