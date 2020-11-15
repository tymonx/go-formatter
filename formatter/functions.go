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
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

var gDummyFunctions = template.FuncMap{ // nolint: gochecknoglobals
	"reset":      setDummy,
	"normal":     setDummy,
	"default":    setDummy,
	"bold":       setDummy,
	"faint":      setDummy,
	"italic":     setDummy,
	"underline":  setDummy,
	"overline":   setDummy,
	"blink":      setDummy,
	"invert":     setDummy,
	"hide":       setDummy,
	"strike":     setDummy,
	"off":        setDummyTransform,
	"bell":       setDummy,
	"black":      setDummy,
	"red":        setDummy,
	"green":      setDummy,
	"yellow":     setDummy,
	"blue":       setDummy,
	"magenta":    setDummy,
	"cyan":       setDummy,
	"white":      setDummy,
	"gray":       setDummy,
	"rgb":        setDummyRGB,
	"bright":     setDummyTransform,
	"background": setDummyTransform,
	"foreground": setDummyTransform,
	"color":      setDummyTransform,
}

var gEscapeFunctions = template.FuncMap{ // nolint: gochecknoglobals
	"reset":      setNormal,
	"normal":     setNormal,
	"default":    setNormal,
	"bold":       setBold,
	"faint":      setFaint,
	"italic":     setItalic,
	"underline":  setUnderline,
	"overline":   setOverline,
	"blink":      setBlink,
	"invert":     setInvert,
	"hide":       setHide,
	"strike":     setStrike,
	"off":        setOff,
	"bell":       setBell,
	"black":      setBlack,
	"red":        setRed,
	"green":      setGreen,
	"yellow":     setYellow,
	"blue":       setBlue,
	"magenta":    setMagenta,
	"cyan":       setCyan,
	"white":      setWhite,
	"gray":       setGray,
	"rgb":        setRGB,
	"bright":     setBright,
	"background": setBackground,
	"foreground": setForeground,
	"color":      setColor,
}

var gFunctions = template.FuncMap{ // nolint: gochecknoglobals
	"ip":         getIPAddress,
	"user":       getUser,
	"executable": os.Executable,
	"cwd":        os.Getwd,
	"hostname":   os.Hostname,
	"env":        os.Getenv,
	"expand":     os.ExpandEnv,
	"uid":        os.Getuid,
	"gid":        os.Getgid,
	"euid":       os.Geteuid,
	"egid":       os.Getegid,
	"pid":        os.Getpid,
	"ppid":       os.Getppid,
	"upper":      strings.ToUpper,
	"lower":      strings.ToLower,
	"capitalize": strings.Title,
	"now":        time.Now,
	"rfc3339":    setISO8601,
	"iso8601":    setISO8601,
	"absolute":   filepath.Abs,
	"base":       filepath.Base,
	"clean":      filepath.Clean,
	"directory":  filepath.Dir,
	"extension":  filepath.Ext,
	"json":       setJSON,
	"indent":     setIndent,
	"fields":     setFields,
}
