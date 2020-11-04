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
	"strings"
	"text/template"

	"gitlab.com/tymonx/go-formatter/cerror"
)

var gFunctions = template.FuncMap{ // nolint: gochecknoglobals
	"Reset":     newEscapeSequence(0),
	"Normal":    newEscapeSequence(0),
	"Default":   newEscapeSequence(0),
	"Bold":      newEscapeSequence(1),
	"Faint":     newEscapeSequence(2),  // nolint: gomnd
	"Italic":    newEscapeSequence(3),  // nolint: gomnd
	"Underline": newEscapeSequence(4),  // nolint: gomnd
	"Blink":     newEscapeSequence(5),  // nolint: gomnd
	"Black":     newEscapeSequence(30), // nolint: gomnd
	"Red":       newEscapeSequence(31), // nolint: gomnd
	"Green":     newEscapeSequence(32), // nolint: gomnd
	"Yellow":    newEscapeSequence(33), // nolint: gomnd
	"Blue":      newEscapeSequence(34), // nolint: gomnd
	"Magenta":   newEscapeSequence(35), // nolint: gomnd
	"Cyan":      newEscapeSequence(36), // nolint: gomnd
	"White":     newEscapeSequence(37), // nolint: gomnd
	"Gray":      newEscapeSequence(90), // nolint: gomnd
	"RGB": func(red, green, blue int) string {
		return fmt.Sprintf("\033[38;2;%d;%d;%dm", scaleColor(red), scaleColor(green), scaleColor(blue))
	},
	"Bright": func(in string) (out string, err error) {
		var code int

		if code, err = getEscapeSequenceCode(in); err != nil {
			return "", err
		}

		if isColorRange(code, 0) || isColorRange(code, backgroundOffset) {
			return newEscapeSequence(code + brightOffset)(), nil
		}

		return "", cerror.New("Bright can be used only with colors")
	},
	"Background": func(in string) (out string, err error) {
		var code int

		if code, err = getEscapeSequenceCode(in); err != nil {
			return "", err
		}

		if isColorRange(code, 0) || isColorRange(code, brightOffset) {
			return newEscapeSequence(code + backgroundOffset)(), nil
		}

		if code == foreground {
			return "\033[48" + strings.TrimPrefix(in, "\033[38"), nil
		}

		return "", cerror.New("Background can be used only with colors")
	},
}
