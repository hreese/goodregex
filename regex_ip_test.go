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
	"regexp"
	"testing"
)

type findAllStringTest struct {
	re     *regexp.Regexp
	input  string
	result []string
}

var ipFindAllStringTests = map[string][]findAllStringTest{
	"MatchIPv4": {
		{MatchIPv4, `1.1.1.1`, []string{`1.1.1.1`}},
		{MatchIPv4, `127.0.0.1`, []string{`127.0.0.1`}},
		{MatchIPv4, `8.9.10.333`, []string{`8.9.10.33`}},
		{MatchIPv4, `"8.8.8.8"`, []string{`8.8.8.8`}},
		{MatchIPv4, `999.0.0.888`, []string{`99.0.0.88`}},
		{MatchIPv4, `1.2.3.4 5.6.7.8 9.10.200.1`, []string{`1.2.3.4`, `5.6.7.8`, `9.10.200.1`}},
	},
	"MatchBoundedIPv4": {
		{MatchBoundedIPv4, `1.1.1.1`, []string{`1.1.1.1`}},
		{MatchBoundedIPv4, `127.0.0.1`, []string{`127.0.0.1`}},
		{MatchBoundedIPv4, `8.9.10.333`, nil},
		{MatchBoundedIPv4, `"8.8.8.8"`, []string{`8.8.8.8`}},
		{MatchBoundedIPv4, `999.0.0.888`, nil},
		{MatchBoundedIPv4, `1.2.3.4 5.6.7.8 9.10.200.1`, []string{`1.2.3.4`, `5.6.7.8`, `9.10.200.1`}},
		{MatchBoundedIPv4, `::1.2.3.4:5.6.7.8:9.10.200.1:::`, []string{`1.2.3.4`, `5.6.7.8`, `9.10.200.1`}},
	},
	"MatchIPv6": {
		{MatchIPv6, `::0`, []string{`::0`}},
		{MatchIPv6, `::1`, []string{`::1`}},
		{MatchIPv6, `::ff`, []string{`::ff`}},
		{MatchIPv6, `1::`, []string{`1::`}},
		{MatchIPv6, `a::b`, []string{`a::b`}},
		{MatchIPv6, `1234:5678::90ab`, []string{`1234:5678::90ab`}},
		{MatchIPv6, `2001:db8::1`, []string{`2001:db8::1`}},
		{MatchIPv6, `1:2:3:4:5:6:7:8`, []string{`1:2:3:4:5:6:7:8`}},
		{MatchIPv6, `12:34:56:78:90:ab:cd:ef`, []string{`12:34:56:78:90:ab:cd:ef`}},
		{MatchIPv6, `12::56:78:90:ab:cd:ef`, []string{`12::56:78:90:ab:cd:ef`}},
		{MatchIPv6, `12:34::78:90:ab:cd:ef`, []string{`12:34::78:90:ab:cd:ef`}},
		{MatchIPv6, `12:34:56::90:ab:cd:ef`, []string{`12:34:56::90:ab:cd:ef`}},
		{MatchIPv6, `12:34:56:78::ab:cd:ef`, []string{`12:34:56:78::ab:cd:ef`}},
		{MatchIPv6, `12:34:56:78:90::cd:ef`, []string{`12:34:56:78:90::cd:ef`}},
		{MatchIPv6, `12:34:56:78:90:ab::ef`, []string{`12:34:56:78:90:ab::ef`}},
		{MatchIPv6, `::ffff:255.255.255.255`, []string{`::ffff:255.255.255.255`}},
		{MatchIPv6, `2001:db8:3:4::192.0.2.33`, []string{`2001:db8:3:4::192.0.2.33`}},
		{MatchIPv6, `64:ff9b::129.13.64.5`, []string{`64:ff9b::129.13.64.5`}},
	},
}

func TestIPFindAllString(t *testing.T) {
	for name, tests := range ipFindAllStringTests {
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
