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
	"os"
	"strconv"
	"strings"

	"github.com/mattn/go-isatty"
)

const (
	redOffset   = 16
	greenOffset = 8
)

var gBrightMap = map[string]string{ // nolint: gochecknoglobals
	"\033[30m":  "\033[90m",
	"\033[31m":  "\033[91m",
	"\033[32m":  "\033[92m",
	"\033[33m":  "\033[93m",
	"\033[34m":  "\033[94m",
	"\033[35m":  "\033[95m",
	"\033[36m":  "\033[96m",
	"\033[37m":  "\033[97m",
	"\033[40m":  "\033[100m",
	"\033[41m":  "\033[101m",
	"\033[42m":  "\033[102m",
	"\033[43m":  "\033[103m",
	"\033[44m":  "\033[104m",
	"\033[45m":  "\033[105m",
	"\033[46m":  "\033[106m",
	"\033[47m":  "\033[107m",
	"\033[90m":  "\033[90m",
	"\033[91m":  "\033[91m",
	"\033[92m":  "\033[92m",
	"\033[93m":  "\033[93m",
	"\033[94m":  "\033[94m",
	"\033[95m":  "\033[95m",
	"\033[96m":  "\033[96m",
	"\033[97m":  "\033[97m",
	"\033[100m": "\033[100m",
	"\033[101m": "\033[101m",
	"\033[102m": "\033[102m",
	"\033[103m": "\033[103m",
	"\033[104m": "\033[104m",
	"\033[105m": "\033[105m",
	"\033[106m": "\033[106m",
	"\033[107m": "\033[107m",
}

var gOffMap = map[string]string{ // nolint: gochecknoglobals
	"\033[1m":  "\033[21m",
	"\033[2m":  "\033[22m",
	"\033[3m":  "\033[23m",
	"\033[4m":  "\033[24m",
	"\033[5m":  "\033[25m",
	"\033[7m":  "\033[27m",
	"\033[8m":  "\033[28m",
	"\033[9m":  "\033[29m",
	"\033[21m": "\033[21m",
	"\033[22m": "\033[22m",
	"\033[23m": "\033[23m",
	"\033[24m": "\033[24m",
	"\033[25m": "\033[25m",
	"\033[27m": "\033[27m",
	"\033[28m": "\033[28m",
	"\033[29m": "\033[29m",
	"\033[53m": "\033[55m",
	"\033[55m": "\033[55m",
}

var gBackgroundMap = map[string]string{ // nolint: gochecknoglobals,dupl
	"\033[0m":   "\033[49m",
	"\033[30m":  "\033[40m",
	"\033[31m":  "\033[41m",
	"\033[32m":  "\033[42m",
	"\033[33m":  "\033[43m",
	"\033[34m":  "\033[44m",
	"\033[35m":  "\033[45m",
	"\033[36m":  "\033[46m",
	"\033[37m":  "\033[47m",
	"\033[40m":  "\033[40m",
	"\033[41m":  "\033[41m",
	"\033[42m":  "\033[42m",
	"\033[43m":  "\033[43m",
	"\033[44m":  "\033[44m",
	"\033[45m":  "\033[45m",
	"\033[46m":  "\033[46m",
	"\033[47m":  "\033[47m",
	"\033[90m":  "\033[100m",
	"\033[91m":  "\033[101m",
	"\033[92m":  "\033[102m",
	"\033[93m":  "\033[103m",
	"\033[94m":  "\033[104m",
	"\033[95m":  "\033[105m",
	"\033[96m":  "\033[106m",
	"\033[97m":  "\033[107m",
	"\033[100m": "\033[100m",
	"\033[101m": "\033[101m",
	"\033[102m": "\033[102m",
	"\033[103m": "\033[103m",
	"\033[104m": "\033[104m",
	"\033[105m": "\033[105m",
	"\033[106m": "\033[106m",
	"\033[107m": "\033[107m",
}

var gForegroundMap = map[string]string{ // nolint: gochecknoglobals,dupl
	"\033[0m":   "\033[39m",
	"\033[30m":  "\033[30m",
	"\033[31m":  "\033[31m",
	"\033[32m":  "\033[32m",
	"\033[33m":  "\033[33m",
	"\033[34m":  "\033[34m",
	"\033[35m":  "\033[35m",
	"\033[36m":  "\033[36m",
	"\033[37m":  "\033[37m",
	"\033[40m":  "\033[30m",
	"\033[41m":  "\033[31m",
	"\033[42m":  "\033[32m",
	"\033[43m":  "\033[33m",
	"\033[44m":  "\033[34m",
	"\033[45m":  "\033[35m",
	"\033[46m":  "\033[36m",
	"\033[47m":  "\033[37m",
	"\033[90m":  "\033[90m",
	"\033[91m":  "\033[91m",
	"\033[92m":  "\033[92m",
	"\033[93m":  "\033[93m",
	"\033[94m":  "\033[94m",
	"\033[95m":  "\033[95m",
	"\033[96m":  "\033[96m",
	"\033[97m":  "\033[97m",
	"\033[100m": "\033[90m",
	"\033[101m": "\033[91m",
	"\033[102m": "\033[92m",
	"\033[103m": "\033[93m",
	"\033[104m": "\033[94m",
	"\033[105m": "\033[95m",
	"\033[106m": "\033[96m",
	"\033[107m": "\033[97m",
}

var gColorMap = map[string]string{ // nolint: gochecknoglobals
	"default": "\033[0m",
	"normal":  "\033[0m",
	"reset":   "\033[0m",
	"black":   "\033[30m",
	"red":     "\033[31m",
	"green":   "\033[32m",
	"yellow":  "\033[33m",
	"blue":    "\033[34m",
	"magenta": "\033[35m",
	"cyan":    "\033[36m",
	"white":   "\033[37m",
	"gray":    "\033[90m",
}

func isTerminal() bool {
	return isatty.IsTerminal(os.Stdout.Fd()) || isatty.IsCygwinTerminal(os.Stdout.Fd())
}

func setDummy() string {
	return ""
}

func setDummyTransform(string) (string, error) {
	return "", nil
}

func setDummyRGB(int, int, int) string {
	return ""
}

func setNormal() string {
	return "\033[0m"
}

func setBold() string {
	return "\033[1m"
}

func setFaint() string {
	return "\033[2m"
}

func setItalic() string {
	return "\033[3m"
}

func setUnderline() string {
	return "\033[4m"
}

func setOverline() string {
	return "\033[53m"
}

func setBlink() string {
	return "\033[5m"
}

func setInvert() string {
	return "\033[7m"
}

func setHide() string {
	return "\033[8m"
}

func setStrike() string {
	return "\033[9m"
}

func setBlack() string {
	return "\033[30m"
}

func setRed() string {
	return "\033[31m"
}

func setGreen() string {
	return "\033[32m"
}

func setYellow() string {
	return "\033[33m"
}

func setBlue() string {
	return "\033[34m"
}

func setMagenta() string {
	return "\033[35m"
}

func setCyan() string {
	return "\033[36m"
}

func setWhite() string {
	return "\033[37m"
}

func setGray() string {
	return "\033[90m"
}

func setBell() string {
	return "\a"
}

func setColor(in string) (out string, err error) {
	var ok bool

	in = strings.TrimSpace(strings.ToLower(in))

	if out, ok = gColorMap[in]; ok {
		return out, nil
	}

	if strings.HasPrefix(in, "0x") {
		var value uint64

		if value, err = strconv.ParseUint(strings.TrimPrefix(in, "0x"), 16, 24); err != nil {
			return "", err
		}

		return setRGB(uint8(value>>redOffset), uint8(value>>greenOffset), uint8(value)), nil
	}

	return "", fError("color is not supported")
}

func setOff(in string) (string, error) {
	if out, ok := gOffMap[in]; ok {
		return out, nil
	}

	return "", fError("off can be used with that function")
}

func setRGB(red, green, blue uint8) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", red, green, blue)
}

func setBright(in string) (string, error) {
	if out, ok := gBrightMap[in]; ok {
		return out, nil
	}

	return "", fError("bright can be used only with colors")
}

func setBackground(in string) (string, error) {
	if out, ok := gBackgroundMap[in]; ok {
		return out, nil
	}

	switch {
	case strings.HasPrefix(in, "\033[38"):
		return "\033[48" + strings.TrimPrefix(in, "\033[38"), nil
	case strings.HasPrefix(in, "\033[48"):
		return in, nil
	default:
		return "", fError("background can be used only with colors")
	}
}

func setForeground(in string) (out string, err error) {
	if out, ok := gForegroundMap[in]; ok {
		return out, nil
	}

	switch {
	case strings.HasPrefix(in, "\033[48"):
		return "\033[38" + strings.TrimPrefix(in, "\033[48"), nil
	case strings.HasPrefix(in, "\033[38"):
		return in, nil
	default:
		return "", fError("foreground can be used only with colors")
	}
}
