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

const (
	// this will not match the unspecified address "::" (RfC 4291, Section 2.5.2.)
	regexMatchV6 = `(
        # IPv4-embedded IPv6 addresses
        (?:64:ff9b::                                                                                  # Well-Known Prefix
            (?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?) # embedded IPv4 address
        ) |
        (?:(?:[0-9a-fA-F]{1,4}:){0,3} (?:[0-9a-fA-F]{1,4}) ::                                         # /96 prefix
            (?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?) # embedded IPv4 address
        ) |
        # IPv4-Mapped IPv6 Address
        (?: :: [fF]{4}:                                                                               # ::ffff:
            (?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?) # embedded IPv4 address
        ) |
        # addresses containing a single ::
                                                      :: (?:(?:[0-9a-fA-F]{1,4}:){0,6}(?:[0-9a-fA-F]{1,4})){0,1} |
                                 (?:[0-9a-fA-F]{1,4}) :: (?:(?:[0-9a-fA-F]{1,4}:){0,5}(?:[0-9a-fA-F]{1,4})){0,1} |
        (?:[0-9a-fA-F]{1,4}:){1} (?:[0-9a-fA-F]{1,4}) :: (?:(?:[0-9a-fA-F]{1,4}:){0,4}(?:[0-9a-fA-F]{1,4})){0,1} |
        (?:[0-9a-fA-F]{1,4}:){2} (?:[0-9a-fA-F]{1,4}) :: (?:(?:[0-9a-fA-F]{1,4}:){0,3}(?:[0-9a-fA-F]{1,4})){0,1} |
        (?:[0-9a-fA-F]{1,4}:){3} (?:[0-9a-fA-F]{1,4}) :: (?:(?:[0-9a-fA-F]{1,4}:){0,2}(?:[0-9a-fA-F]{1,4})){0,1} |
        (?:[0-9a-fA-F]{1,4}:){4} (?:[0-9a-fA-F]{1,4}) :: (?:(?:[0-9a-fA-F]{1,4}:){0,1}(?:[0-9a-fA-F]{1,4})){0,1} |
        (?:[0-9a-fA-F]{1,4}:){5} (?:[0-9a-fA-F]{1,4}) :: (?:[0-9a-fA-F]{1,4}){0,1}                               |
        (?:[0-9a-fA-F]{1,4}:){6} (?:[0-9a-fA-F]{1,4}) ::                                                         |
        # plain IPv6 address
        (?:[0-9a-fA-F]{1,4}:){7} (?:[0-9a-fA-F]{1,4})
        )`
	regexMatchV4 = `
        (?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}
           (?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)`
)

// MatchIPv4 matches IPv4 addresses in dotted decimal notation
var MatchIPv4 = MustCompileReadableRegex(regexMatchV4)

// MatchBoundedIPv4 is an experimental version of MatchIPv4 that
// only matches if the IPv4 address is surrounded by a word border.
var MatchBoundedIPv4 = MustCompileReadableRegex(`(?:\b` + regexMatchV4 + `\b)`)

// MatchIPv6 matches IPv6 addresses. It supports colon-separated hexadecimal notation
// including ::-abbreviation, IPv4-embedded IPv6 addresses and IPv4-Mapped IPv6 Address.
// The unspecified address "::" is not matched as it cannot be resolved or connected to.
var MatchIPv6 = MustCompileReadableRegex(regexMatchV6)
