package goodregex

import (
	"reflect"
	"testing"
)

type findIPv4 struct {
	input         string
	result        []string
	resultBounded []string
}

var findIPv4Tests = []findIPv4{
	findIPv4{"1.1.1.1", []string{"1.1.1.1"}, []string{"1.1.1.1"}},
	findIPv4{"127.0.0.1", []string{"127.0.0.1"}, []string{"127.0.0.1"}},
	findIPv4{"8.9.10.333", []string{"8.9.10.33"}, []string(nil)},
	findIPv4{`"8.8.8.8"`, []string{"8.8.8.8"}, []string{"8.8.8.8"}},
	findIPv4{`999.0.0.0`, []string{"99.0.0.0"}, []string(nil)},
}

func TestFindIPv4(t *testing.T) {
	for _, test := range findIPv4Tests {
		res := MatchV4.FindAllString(test.input, -1)
		if !reflect.DeepEqual(res, test.result) {
			t.Errorf("Extracted IPv4 addresses (unbounded) from\n  %s\nexpected\n  %#v\ngot\n  %#v\n", test.input, test.result, res)
		}
		res = MatchBoundedV4.FindAllString(test.input, -1)
		if !reflect.DeepEqual(res, test.resultBounded) {
			t.Errorf("Extracted IPv4 addresses (bounded) from\n  %s\nexpected\n  %#v\ngot\n  %#v\n", test.input, test.resultBounded, res)
		}
	}
}
