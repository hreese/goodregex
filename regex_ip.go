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
	"regexp"
)

const (
	// this will not match the unspecified address "::" (RfC 4291, Section 2.5.2.)
	regexMatchV6 = `(
        # IPv4-embedded IPv6 addresses
        (?:64:ff9b::                                                                                  # Well-Known Prefix
            (?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?) # embedded IPv4 address
        ) |
        (?:(?:[[:alnum:]]{1,4}:){0,3} (?:[[:alnum:]]{1,4}) ::                                         # /96 prefix
            (?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?) # embedded IPv4 address
        ) |
        # IPv4-Mapped IPv6 Address
        (?: :: [fF]{4}:                                                                               # ::ffff:
            (?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?) # embedded IPv4 address
        ) |
        # addresses containing a single ::
                                                      :: (?:(?:[[:alnum:]]{1,4}:){0,6}(?:[[:alnum:]]{1,4})){0,1} |
                                 (?:[[:alnum:]]{1,4}) :: (?:(?:[[:alnum:]]{1,4}:){0,5}(?:[[:alnum:]]{1,4})){0,1} |
        (?:[[:alnum:]]{1,4}:){1} (?:[[:alnum:]]{1,4}) :: (?:(?:[[:alnum:]]{1,4}:){0,4}(?:[[:alnum:]]{1,4})){0,1} |
        (?:[[:alnum:]]{1,4}:){2} (?:[[:alnum:]]{1,4}) :: (?:(?:[[:alnum:]]{1,4}:){0,3}(?:[[:alnum:]]{1,4})){0,1} |
        (?:[[:alnum:]]{1,4}:){3} (?:[[:alnum:]]{1,4}) :: (?:(?:[[:alnum:]]{1,4}:){0,2}(?:[[:alnum:]]{1,4})){0,1} |
        (?:[[:alnum:]]{1,4}:){4} (?:[[:alnum:]]{1,4}) :: (?:(?:[[:alnum:]]{1,4}:){0,1}(?:[[:alnum:]]{1,4})){0,1} |
        (?:[[:alnum:]]{1,4}:){5} (?:[[:alnum:]]{1,4}) :: (?:[[:alnum:]]{1,4}){0,1}                               |
        (?:[[:alnum:]]{1,4}:){6} (?:[[:alnum:]]{1,4}) ::                                                         |
        # plain IPv6 address
        (?:[[:alnum:]]{1,4}:){7} (?:[[:alnum:]]{1,4})
        )`
	regexMatchV4 = `
        (?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}
           (?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)`
)

var (
	regexRemoveComments = regexp.MustCompile("(?m:#.*$)")
	regexRemoveSpaces   = regexp.MustCompile("(?m:[[:space:]]+)")
	MatchIPv4           = MustCompileReadableRegex(regexMatchV4)
	MatchBoundedIPv4    = MustCompileReadableRegex(`(?:\b` + regexMatchV4 + `\b)`)
	MatchIPv6           = MustCompileReadableRegex(regexMatchV6)
)

// MustCompileReadableRegex is similar to regexp.MustCompile
// It accepts a "readable regex" which is a regular RE2 regexp where
// * all whitespace and
// * comments (# follow by everythin until EOL)
// are removed prior to compilation
func MustCompileReadableRegex(r string) *regexp.Regexp {
	r = regexRemoveComments.ReplaceAllString(r, "")
	r = regexRemoveSpaces.ReplaceAllString(r, "")
	return regexp.MustCompile(r)
}
