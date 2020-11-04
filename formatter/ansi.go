// Copyright 2020 Tymoteusz Blazejczyk
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package formatter

import (
	"fmt"
)

const (
	black                = 30
	white                = 37
	brightOffset         = 60
	backgroundOffset     = 10
	escapeSequenceFormat = "\033[%dm"
	colorMinimum         = 0
	colorMaximum         = 255
	foreground           = 38
)

func newEscapeSequence(code int) func() string {
	return func() string {
		return fmt.Sprintf(escapeSequenceFormat, code)
	}
}

func getEscapeSequenceCode(escapeSequence string) (code int, err error) {
	_, err = fmt.Sscanf(escapeSequence, "\033[%d", &code)
	return code, err
}

func isColorRange(code, offset int) bool {
	code -= offset
	return (code >= black) && (code <= white)
}

func scaleColor(color int) int {
	if color < colorMinimum {
		color = colorMinimum
	}

	if color > colorMaximum {
		color = colorMaximum
	}

	return color
}
