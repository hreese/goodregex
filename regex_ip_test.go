package goodregex

import (
	"reflect"
	"regexp"
	"testing"
)

type IPFindAllString struct {
	re     regexp.Regexp
	input  string
	result []string
}

var IPFindAllStringTests = map[string][]IPFindAllString{
	"MatchIPv4": []IPFindAllString{
		IPFindAllString{*MatchIPv4, `1.1.1.1`, []string{`1.1.1.1`}},
		IPFindAllString{*MatchIPv4, `127.0.0.1`, []string{`127.0.0.1`}},
		IPFindAllString{*MatchIPv4, `8.9.10.333`, []string{`8.9.10.33`}},
		IPFindAllString{*MatchIPv4, `"8.8.8.8"`, []string{`8.8.8.8`}},
		IPFindAllString{*MatchIPv4, `999.0.0.888`, []string{`99.0.0.88`}},
		IPFindAllString{*MatchIPv4, `1.2.3.4 5.6.7.8 9.10.200.1`, []string{`1.2.3.4`, `5.6.7.8`, `9.10.200.1`}},
	},
	"MatchBoundedIPv4": []IPFindAllString{
		IPFindAllString{*MatchBoundedIPv4, `1.1.1.1`, []string{`1.1.1.1`}},
		IPFindAllString{*MatchBoundedIPv4, `127.0.0.1`, []string{`127.0.0.1`}},
		IPFindAllString{*MatchBoundedIPv4, `8.9.10.333`, []string(nil)},
		IPFindAllString{*MatchBoundedIPv4, `"8.8.8.8"`, []string{`8.8.8.8`}},
		IPFindAllString{*MatchBoundedIPv4, `999.0.0.888`, []string(nil)},
		IPFindAllString{*MatchBoundedIPv4, `1.2.3.4 5.6.7.8 9.10.200.1`, []string{`1.2.3.4`, `5.6.7.8`, `9.10.200.1`}},
		IPFindAllString{*MatchBoundedIPv4, `::1.2.3.4:5.6.7.8:9.10.200.1:::`, []string{`1.2.3.4`, `5.6.7.8`, `9.10.200.1`}},
	},
	"MatchIPv6": []IPFindAllString{
		IPFindAllString{*MatchIPv6, `::0`, []string{`::0`}},
		IPFindAllString{*MatchIPv6, `::1`, []string{`::1`}},
		IPFindAllString{*MatchIPv6, `::ff`, []string{`::ff`}},
		IPFindAllString{*MatchIPv6, `1::`, []string{`1::`}},
		IPFindAllString{*MatchIPv6, `111111111::aaaaaaaa`, []string{`1111::aaaa`}},
		IPFindAllString{*MatchIPv6, `a::b`, []string{`a::b`}},
		IPFindAllString{*MatchIPv6, `1234:5678::90ab`, []string{`1234:5678::90ab`}},
		IPFindAllString{*MatchIPv6, `2001:db8::1`, []string{`2001:db8::1`}},
		IPFindAllString{*MatchIPv6, `1:2:3:4:5:6:7:8`, []string{`1:2:3:4:5:6:7:8`}},
		IPFindAllString{*MatchIPv6, `12:34:56:78:90:ab:cd:ef`, []string{`12:34:56:78:90:ab:cd:ef`}},
		IPFindAllString{*MatchIPv6, `12::56:78:90:ab:cd:ef`, []string{`12::56:78:90:ab:cd:ef`}},
		IPFindAllString{*MatchIPv6, `12:34::78:90:ab:cd:ef`, []string{`12:34::78:90:ab:cd:ef`}},
		IPFindAllString{*MatchIPv6, `12:34:56::90:ab:cd:ef`, []string{`12:34:56::90:ab:cd:ef`}},
		IPFindAllString{*MatchIPv6, `12:34:56:78::ab:cd:ef`, []string{`12:34:56:78::ab:cd:ef`}},
		IPFindAllString{*MatchIPv6, `12:34:56:78:90::cd:ef`, []string{`12:34:56:78:90::cd:ef`}},
		IPFindAllString{*MatchIPv6, `12:34:56:78:90:ab::ef`, []string{`12:34:56:78:90:ab::ef`}},
		IPFindAllString{*MatchIPv6, `::ffff:255.255.255.255`, []string{`::ffff:255.255.255.255`}},
		IPFindAllString{*MatchIPv6, `2001:db8:3:4::192.0.2.33`, []string{`2001:db8:3:4::192.0.2.33`}},
		IPFindAllString{*MatchIPv6, `64:ff9b::129.13.64.5`, []string{`64:ff9b::129.13.64.5`}},
	},
}

func TestIPFindAllString(t *testing.T) {
	for comment, tests := range IPFindAllStringTests {
		for _, test := range tests {
			res := test.re.FindAllString(test.input, -1)
			if !reflect.DeepEqual(res, test.result) {
				t.Errorf("Testset %s tested\n  %s\n and expected\n  %#v\ngot\n  %#v\n", comment, test.input, test.result, res)
			}
		}
	}
}
