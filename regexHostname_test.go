/*
Copyright 2017 Heiko Reese <mail@example.org>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package goodregex

import (
	"reflect"
	"regexp"
	"testing"
)

type HostnameFindAllString struct {
	re     regexp.Regexp
	input  string
	result []string
}

var HostnameFindAllStringTests = map[string][]HostnameFindAllString{
	"MatchHostname": {
		HostnameFindAllString{*MatchHostname, `goodregex.example.org`, []string{`goodregex.example.org`}},
		HostnameFindAllString{*MatchHostname, `-goodregex.example.org`, []string{`goodregex.example.org`}},
		HostnameFindAllString{*MatchHostname, `goodregex example.org`, []string{"goodregex", "example.org"}},
		HostnameFindAllString{*MatchHostname, `goodregex@example.org`, []string{"goodregex", "example.org"}},
	},
	"MatchBoundedHostname": {
		HostnameFindAllString{*MatchBoundedHostname, `goodregex.example.org`, []string{`goodregex.example.org`}},
		//HostnameFindAllString{*MatchBoundedHostname, `-goodregex.example.org`, []string{``}}, // TODO: not correct
		HostnameFindAllString{*MatchBoundedHostname, `goodregex example.org`, []string{"goodregex", "example.org"}},
		HostnameFindAllString{*MatchBoundedHostname, `goodregex@example.org`, []string{"goodregex", "example.org"}},
	},
}

func TestHostnameFindAllString(t *testing.T) {
	for comment, tests := range HostnameFindAllStringTests {
		for _, test := range tests {
			res := test.re.FindAllString(test.input, -1)
			if !reflect.DeepEqual(res, test.result) {
				t.Errorf("Testset %s tested\n  %s\n and expected\n  %#v\ngot\n  %#v\n", comment, test.input, test.result, res)
			}
		}
	}
}
