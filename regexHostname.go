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
	// source: https://stackoverflow.com/questions/106179/regular-expression-to-match-dns-hostname-or-ip-address
	regexMatchHostname string = `(
        ([a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9]|[a-zA-Z0-9]) # may not start with a dash
        \.)*                                               # fqhn parts separated by dots
        ([a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9]|[a-zA-Z0-9])`
)

// MatchHostname matches a hostname according to RfC 1123 2.1
var MatchHostname = MustCompileReadableRegex(regexMatchHostname)

// MatchBoundedHostname is an experimental version of MatchHostname that
// only matches if the hostname is surrounded by a word border.
var MatchBoundedHostname = MustCompileReadableRegex(`(?:\b` + regexMatchHostname + `\b)`)
