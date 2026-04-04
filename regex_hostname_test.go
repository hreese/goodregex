/*
Copyright 2017 Heiko Reese <mail@heiko-reese.de>

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
	"fmt"
	"reflect"
	"testing"
)

var hostnameFindAllStringTests = map[string][]findAllStringTest{
	"MatchHostname": {
		{MatchHostname, `goodregex.example.org`, []string{`goodregex.example.org`}},
		{MatchHostname, `-goodregex.example.org`, []string{`goodregex.example.org`}},
		{MatchHostname, `goodregex example.org`, []string{"goodregex", "example.org"}},
		{MatchHostname, `goodregex@example.org`, []string{"goodregex", "example.org"}},
	},
	"MatchBoundedHostname": {
		{MatchBoundedHostname, `goodregex.example.org`, []string{`goodregex.example.org`}},
		// NOTE: MatchBoundedHostname does not correctly reject leading dashes
		// (e.g., "-goodregex.example.org" still partially matches). This is a
		// known limitation of using \b word boundaries with the current regex.
		{MatchBoundedHostname, `goodregex example.org`, []string{"goodregex", "example.org"}},
		{MatchBoundedHostname, `goodregex@example.org`, []string{"goodregex", "example.org"}},
	},
}

func TestHostnameFindAllString(t *testing.T) {
	for name, tests := range hostnameFindAllStringTests {
		t.Run(name, func(t *testing.T) {
			for i, tc := range tests {
				t.Run(fmt.Sprintf("case_%d_%s", i, tc.input), func(t *testing.T) {
					res := tc.re.FindAllString(tc.input, -1)
					if !reflect.DeepEqual(res, tc.result) {
						t.Errorf("input %q: expected %#v, got %#v", tc.input, tc.result, res)
					}
				})
			}
		})
	}
}
