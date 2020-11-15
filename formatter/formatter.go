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
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"text/template"
)

var gEscapeSequences = AreEscapeSequencesSupported() // nolint: gochecknoglobals

// These constants define default values used by formatter.
const (
	DefaultPlaceholder      = "p"
	DefaultLeftDelimiter    = "{"
	DefaultRightDelimiter   = "}"
	ForceEscapeSequencesEnv = "FORCE_ESCAPE_SEQUENCES"
)

// Named defines named arguments.
type Named map[string]interface{}

// Functions defines a map of template functions.
type Functions map[string]interface{}

// Formatter defines a formatter object that formats string using
// “replacement fields” surrounded by curly braces {}.
type Formatter struct {
	placeholder     string
	leftDelimiter   string
	rightDelimiter  string
	escapeSequences bool
	functions       Functions
}

// New creates a new formatter object.
func New() *Formatter {
	return &Formatter{
		placeholder:     DefaultPlaceholder,
		leftDelimiter:   DefaultLeftDelimiter,
		rightDelimiter:  DefaultRightDelimiter,
		escapeSequences: gEscapeSequences,
		functions:       Functions{},
	}
}

// Format formats string.
func Format(message string, arguments ...interface{}) (string, error) {
	return New().Format(message, arguments...)
}

// MustFormat is like Format but panics if provided message cannot be formatted.
// It simplifies safe initialization of global variables holding formatted strings.
func MustFormat(message string, arguments ...interface{}) string {
	return New().MustFormat(message, arguments...)
}

// FormatWriter formats string to writer.
func FormatWriter(writer io.Writer, message string, arguments ...interface{}) error {
	return New().FormatWriter(writer, message, arguments...)
}

// AreEscapeSequencesSupported returns true if environment supports ANSI escape sequences.
// Otherwise, it returns false.
func AreEscapeSequencesSupported() bool {
	switch strings.TrimSpace(strings.ToLower(os.Getenv(ForceEscapeSequencesEnv))) {
	case "1", "true", "on", "yes", "enable", "y":
		return true
	case "0", "false", "off", "no", "disable", "n":
		return false
	default:
		return (os.Getenv("TERM") != "dumb") && isTerminal()
	}
}

// Format formats string.
func (f *Formatter) Format(message string, arguments ...interface{}) (string, error) {
	var buffer bytes.Buffer

	if err := f.FormatWriter(&buffer, message, arguments...); err != nil {
		return "", err
	}

	return buffer.String(), nil
}

// MustFormat is like Format but panics if provided message cannot be formatted.
// It simplifies safe initialization of global variables holding formatted strings.
func (f *Formatter) MustFormat(message string, arguments ...interface{}) string {
	formatted, err := f.Format(message, arguments...)

	if err != nil {
		panic(err)
	}

	return formatted
}

// Reset resets formatter to default state.
func (f *Formatter) Reset() *Formatter {
	*f = *New()

	return f
}

// SetFunctions sets template functions used by formatter.
func (f *Formatter) SetFunctions(functions Functions) *Formatter {
	f.functions = functions
	return f
}

// GetFunction returns template function used by formatter.
func (f *Formatter) GetFunction(name string) interface{} {
	return f.functions[name]
}

// GetFunctions returns template functions used by formatter.
func (f *Formatter) GetFunctions() Functions {
	return f.functions
}

// AddFunction adds template function used by formatter.
func (f *Formatter) AddFunction(name string, function interface{}) *Formatter {
	f.functions[name] = function
	return f
}

// AddFunctions adds template functions used by formatter.
func (f *Formatter) AddFunctions(functions Functions) *Formatter {
	for name, function := range functions {
		f.functions[name] = function
	}

	return f
}

// RemoveFunction removes template function used by formatter.
func (f *Formatter) RemoveFunction(name string) *Formatter {
	if _, ok := f.functions[name]; !ok {
		return f
	}

	delete(f.functions, name)

	return f
}

// RemoveFunctions removes template functions used by formatter.
func (f *Formatter) RemoveFunctions(names []string) *Formatter {
	for _, name := range names {
		f.RemoveFunction(name)
	}

	return f
}

// ResetFunctions resets template functions used by formatter.
func (f *Formatter) ResetFunctions() *Formatter {
	f.functions = Functions{}
	return f
}

// SetPlaceholder sets placeholder string prefix used for automatic and
// positional placeholders to format string. Default is p.
func (f *Formatter) SetPlaceholder(placeholder string) *Formatter {
	f.placeholder = placeholder
	return f
}

// GetPlaceholder returns placeholder string prefix used for automatic and
// positional placeholders to format string. Default is p.
func (f *Formatter) GetPlaceholder() string {
	return f.placeholder
}

// ResetPlaceholder resets placeholder to default value.
func (f *Formatter) ResetPlaceholder() *Formatter {
	f.placeholder = DefaultPlaceholder
	return f
}

// SetDelimiters sets delimiters used by formatter. Default is {}.
func (f *Formatter) SetDelimiters(left, right string) *Formatter {
	return f.SetLeftDelimiter(left).SetRightDelimiter(right)
}

// SetLeftDelimiter sets left delimiter used by formatter. Default is {.
func (f *Formatter) SetLeftDelimiter(delimiter string) *Formatter {
	f.leftDelimiter = delimiter
	return f
}

// SetRightDelimiter sets right delimiter used by formatter. Default is }.
func (f *Formatter) SetRightDelimiter(delimiter string) *Formatter {
	f.rightDelimiter = delimiter
	return f
}

// GetDelimiters returns delimiters used by formatter. Default is {}.
func (f *Formatter) GetDelimiters() (left, right string) {
	return f.GetLeftDelimiter(), f.GetRightDelimiter()
}

// GetLeftDelimiter returns left delimiter used by formatter. Default is {.
func (f *Formatter) GetLeftDelimiter() string {
	return f.leftDelimiter
}

// GetRightDelimiter returns right delimiter used by formatter. Default is }.
func (f *Formatter) GetRightDelimiter() string {
	return f.rightDelimiter
}

// ResetDelimiters resets delimiters used by formatter to default values.
func (f *Formatter) ResetDelimiters() *Formatter {
	return f.ResetLeftDelimiter().ResetRightDelimiter()
}

// ResetLeftDelimiter resets left delimiter used by formatter to default value.
func (f *Formatter) ResetLeftDelimiter() *Formatter {
	f.leftDelimiter = DefaultLeftDelimiter
	return f
}

// ResetRightDelimiter resets right delimiter used by formatter to default value.
func (f *Formatter) ResetRightDelimiter() *Formatter {
	f.rightDelimiter = DefaultRightDelimiter
	return f
}

// SetEscapeSequences enables or disables ANSI escape sequences in formatted messages.
func (f *Formatter) SetEscapeSequences(escapeSequences bool) *Formatter {
	f.escapeSequences = escapeSequences
	return f
}

// EnableEscapeSequences allows ANSI escape sequences in formatted messages.
func (f *Formatter) EnableEscapeSequences() *Formatter {
	return f.SetEscapeSequences(true)
}

// DisableEscapeSequences removes ANSI escape sequences in formatted messages.
func (f *Formatter) DisableEscapeSequences() *Formatter {
	return f.SetEscapeSequences(false)
}

// AreEscapeSequencesEnabled returns true if escape sequences are allowed in formatted messages.
// Otherwise, it returns false.
func (f *Formatter) AreEscapeSequencesEnabled() bool {
	return f.escapeSequences
}

// FormatWriter formats string to writer.
func (f *Formatter) FormatWriter(writer io.Writer, message string, arguments ...interface{}) error {
	var object interface{}

	var objectPosition int

	used := make(map[int]bool)
	placeholders := make(template.FuncMap)
	placeholders[f.placeholder] = argumentAutomatic(used, arguments)

	for position, argument := range arguments {
		placeholder := f.placeholder + strconv.Itoa(position)
		placeholders[placeholder] = argumentValue(used, position, argument)
		valueOf := reflect.ValueOf(argument)

		switch valueOf.Kind() {
		case reflect.Map:
			if reflect.TypeOf(argument).Key().Kind() == reflect.String {
				for _, key := range valueOf.MapKeys() {
					placeholders[key.String()] = argumentValue(used, position, valueOf.MapIndex(key).Interface())
				}
			}
		case reflect.Struct:
			object = argument
			objectPosition = position
		case reflect.Ptr:
			if isObjectPointer(valueOf) {
				object = argument
				objectPosition = position
			}
		}
	}

	t := template.New("").Delims(f.leftDelimiter, f.rightDelimiter).Funcs(f.getEscapeFunctions()).
		Funcs(gFunctions).Funcs(placeholders).Funcs(template.FuncMap(f.functions))

	if _, err := t.Parse(message); err != nil {
		return err
	}

	if err := t.Execute(writer, object); err != nil {
		return err
	}

	if len(used) >= len(arguments) {
		return nil
	}

	if err := f.objectUsed(used, objectPosition, object, message); err != nil {
		return err
	}

	separator := getSeparator(message)
	message = ""

	for position, argument := range arguments {
		if !used[position] {
			message += separator + fmt.Sprint(argument)
			separator = " "
		}
	}

	return write(writer, message)
}

func (f *Formatter) getEscapeFunctions() template.FuncMap {
	if f.escapeSequences {
		return gEscapeFunctions
	}

	return gDummyFunctions
}

func (f *Formatter) objectUsed(used map[int]bool, position int, object interface{}, message string) (err error) {
	if (object == nil) || used[position] {
		return nil
	}

	var r *regexp.Regexp

	if r, err = regexp.Compile(f.leftDelimiter + `\s*(\.|[^\.].* \.).+` + f.rightDelimiter); err != nil {
		return err
	}

	used[position] = r.MatchString(message)

	return nil
}

func getSeparator(message string) string {
	if (message == "") || (message[len(message)-1] == ' ') {
		return ""
	}

	return " "
}

func isObjectPointer(value reflect.Value) bool {
	return !value.IsNil() && (value.Elem().Kind() == reflect.Struct)
}

func argumentValue(used map[int]bool, position int, argument interface{}) func() interface{} {
	return func() interface{} {
		used[position] = true
		return argument
	}
}

func argumentAutomatic(used map[int]bool, arguments []interface{}) func() interface{} {
	length := len(arguments)
	position := 0

	return func() interface{} {
		var argument interface{}

		if position < length {
			used[position] = true
			argument = arguments[position]
			position++
		}

		return argument
	}
}

func write(writer io.Writer, message string) error {
	if _, err := writer.Write([]byte(message)); err != nil {
		return err
	}

	return nil
}
