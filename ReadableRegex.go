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

import "regexp"

var (
	regexRemoveComments = regexp.MustCompile("(?m:#.*$)")
	regexRemoveSpaces   = regexp.MustCompile("(?m:[[:space:]]+)")
)

// MustCompileReadableRegex is similar to regexp.MustCompile
// It accepts a "readable regex" which is a regular RE2 regexp where
// whitespace and comments (a # followed by everything until the end of line)
// are removed prior to compilation.
func MustCompileReadableRegex(r string) *regexp.Regexp {
	r = regexRemoveComments.ReplaceAllString(r, "")
	r = regexRemoveSpaces.ReplaceAllString(r, "")
	return regexp.MustCompile(r)
}
