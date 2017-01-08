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
		IPFindAllString{*MatchIPv4, "1.1.1.1", []string{"1.1.1.1"}},
		IPFindAllString{*MatchIPv4, "127.0.0.1", []string{"127.0.0.1"}},
		IPFindAllString{*MatchIPv4, "8.9.10.333", []string{"8.9.10.33"}},
		IPFindAllString{*MatchIPv4, `"8.8.8.8"`, []string{"8.8.8.8"}},
		IPFindAllString{*MatchIPv4, `999.0.0.0`, []string{"99.0.0.0"}},
	},
	"MatchBoundedIPv4": []IPFindAllString{
		IPFindAllString{*MatchBoundedIPv4, "_1.1.1.1_", []string{"1.1.1.1"}},
		IPFindAllString{*MatchBoundedIPv4, "1127.0.0.1", []string{"127.0.0.1"}},
		IPFindAllString{*MatchBoundedIPv4, "8.9.10.333", []string{"8.9.10.33"}},
		IPFindAllString{*MatchBoundedIPv4, `"8.8.8.8"`, []string{"8.8.8.8"}},
		IPFindAllString{*MatchBoundedIPv4, `xxx9.1.0.0xxx`, []string{"9.1.0.0"}},
	},
}

func TestFindIPv4(t *testing.T) {
	for comment, tests := range IPFindAllStringTests {
		for _, test := range tests {
			res := MatchIPv4.FindAllString(test.input, -1)
			if !reflect.DeepEqual(res, test.result) {
				t.Errorf("Testset %s tested\n  %s\n and expected\n  %#v\ngot\n  %#v\n", comment, test.input, test.result, res)
			}
		}
	}
}
