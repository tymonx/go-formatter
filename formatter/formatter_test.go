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

package formatter_test

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"os/user"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mattn/go-isatty"
	"github.com/stretchr/testify/assert"
	"gitlab.com/tymonx/go-formatter/formatter"
	"gitlab.com/tymonx/go-formatter/mocks"
)

func ExampleMustFormat() {
	fmt.Println(formatter.MustFormat("With arguments", 3, nil, false, 4.5, "text", []byte{}, Error("error")))
	// Output: With arguments 3 <nil> false 4.5 text [] error
}

func ExampleFormat_withoutArguments() {
	formatted, err := formatter.Format("Without arguments")

	if err != nil {
		panic(err)
	}

	fmt.Println(formatted)
	// Output: Without arguments
}

func ExampleFormat_withArguments() {
	formatted, err := formatter.Format("With arguments", 3, nil, 4.5, true, "arg1", []byte{}, Error("error"))

	if err != nil {
		panic(err)
	}

	fmt.Println(formatted)
	// Output: With arguments 3 <nil> 4.5 true arg1 [] error
}

func ExampleFormat_automaticPlaceholder() {
	formatted, err := formatter.Format("Automatic placeholder {p}:{p}:{p}():", "dir/file", 1, "func1")

	if err != nil {
		panic(err)
	}

	fmt.Println(formatted)
	// Output: Automatic placeholder dir/file:1:func1():
}

func ExampleFormat_positionalPlaceholders() {
	formatted, err := formatter.Format("Positional placeholders {p1}:{p0}:{p2}():", 2, "dir/file", "func1")

	if err != nil {
		panic(err)
	}

	fmt.Println(formatted)
	// Output: Positional placeholders dir/file:2:func1():
}

func ExampleFormat_namedPlaceholders() {
	formatted, err := formatter.Format("Named placeholders {file}:{line}:{function}():", formatter.Named{
		"line":     3,
		"function": "func1",
		"file":     "dir/file",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(formatted)
	// Output: Named placeholders dir/file:3:func1():
}

func ExampleFormat_objectPlaceholders() {
	object := struct {
		Line     int
		Function string
		File     string
	}{
		Line:     4,
		Function: "func1",
		File:     "dir/file",
	}

	formatted, err := formatter.Format("Object placeholders {.File}:{.Line}:{.Function}():", object)

	if err != nil {
		panic(err)
	}

	fmt.Println(formatted)
	// Output: Object placeholders dir/file:4:func1():
}

func ExampleFormat_objectAutomaticPlaceholders() {
	object1 := struct {
		X int
	}{
		X: 1,
	}

	object2 := struct {
		Message string
	}{
		Message: "msg",
	}

	formatted, err := formatter.Format("{p.X} {p.Message}", object1, object2)

	if err != nil {
		panic(err)
	}

	fmt.Println(formatted)
	// Output: 1 msg
}

func ExampleFormat_objectPositionalPlaceholders() {
	object1 := struct {
		X int
	}{
		X: 1,
	}

	object2 := struct {
		Y int
	}{
		Y: 2,
	}

	formatted, err := formatter.Format("{p1.Y} {p0.X}", object1, object2)

	if err != nil {
		panic(err)
	}

	fmt.Println(formatted)
	// Output: 2 1
}

func ExampleFormat_objectPointerPlaceholders() {
	objectPointer := &struct {
		X int
		Y int
		Z int
	}{
		X: 4,
		Z: 3,
		Y: 1,
	}

	formatted, err := formatter.Format("Object placeholders {.X}.{.Y}.{.Z}", objectPointer)

	if err != nil {
		panic(err)
	}

	fmt.Println(formatted)
	// Output: Object placeholders 4.1.3
}

func ExampleFormat_mixedPlaceholders() {
	objectPointer := &struct {
		X int
		Y int
		Z int
	}{
		X: 2,
		Z: 6,
		Y: 3,
	}

	formatted, err := formatter.Format("Mixed placeholders {.X}.{p}.{.Y}.{.Z} {p1} {p0}", objectPointer, "b", "c", nil)

	if err != nil {
		panic(err)
	}

	fmt.Println(formatted)
	// Output: Mixed placeholders 2.{2 3 6}.3.6 b {2 3 6} c <nil>
}

func ExampleFormatWriter() {
	buffer := new(bytes.Buffer)

	if err := formatter.FormatWriter(buffer, "Writer {p2}", 3, "foo", "bar"); err != nil {
		panic(err)
	}

	fmt.Println(buffer.String())
	// Output: Writer bar 3 foo
}

func ExampleFormatter_SetFunctions() {
	functions := formatter.Functions{
		"str": func() string {
			return "text"
		},
		"number": func() int {
			return 3
		},
		"boolean": func() bool {
			return true
		},
		"floating": func() float64 {
			return 4.5
		},
	}

	formatted, err := formatter.New().SetFunctions(functions).Format("Custom functions {str} {p} {number} {boolean} {floating}", 5)

	if err != nil {
		panic(err)
	}

	fmt.Println(formatted)
	// Output: Custom functions text 5 3 true 4.5
}

func ExampleFormatter_SetPlaceholder() {
	formatted, err := formatter.New().SetPlaceholder("arg").Format("Custom placeholder {arg1} {arg0}", "2", 3)

	if err != nil {
		panic(err)
	}

	fmt.Println(formatted)
	// Output: Custom placeholder 3 2
}

func ExampleFormatter_SetDelimiters() {
	formatted, err := formatter.New().SetDelimiters("<", ">").Format("Custom delimiters <p1> <p0>", "4", 3)

	if err != nil {
		panic(err)
	}

	fmt.Println(formatted)
	// Output: Custom delimiters 3 4
}

func ExampleFormat_colors() {
	formatted, err := formatter.Format("With colors {red}red{normal} {green}green{normal} {blue}blue{normal}")

	if err != nil {
		panic(err)
	}

	fmt.Println(formatted)
}

func ExampleFormat_rgb() {
	formatted, err := formatter.Format("With RGB {rgb 255 165 0}funky{normal}")

	if err != nil {
		panic(err)
	}

	fmt.Println(formatted)
}

func ExampleFormat_backgroundRGB() {
	formatted, err := formatter.Format("With background RGB {rgb 255 165 0 | background}funky{normal}")

	if err != nil {
		panic(err)
	}

	fmt.Println(formatted)
}

func ExampleFormat_brightColors() {
	formatted, err := formatter.Format("With bright colors {magenta | bright}magenta{normal}")

	if err != nil {
		panic(err)
	}

	fmt.Println(formatted)
}

func ExampleFormat_backgroundColors() {
	formatted, err := formatter.Format("With background colors {yellow | background}yellow{normal}")

	if err != nil {
		panic(err)
	}

	fmt.Println(formatted)
}

func ExampleFormat_backgroundBrightColors() {
	formatted, err := formatter.Format("With background bright colors {cyan | bright | background}cyan{normal}")

	if err != nil {
		panic(err)
	}

	fmt.Println(formatted)
}

func ExampleFormatter_SetEscapeSequences() {
	f := formatter.New()

	fmt.Println(f.SetEscapeSequences(false).AreEscapeSequencesEnabled())
	fmt.Println(f.SetEscapeSequences(true).AreEscapeSequencesEnabled())
	// Output:
	// false
	// true
}

func ExampleFormatter_EnableEscapeSequences() {
	fmt.Println(formatter.New().EnableEscapeSequences().AreEscapeSequencesEnabled())
	// Output: true
}

func ExampleFormatter_DisableEscapeSequences() {
	formatted, err := formatter.New().DisableEscapeSequences().Format("{rgb 255 134 5 | background}Escape sequences disabled{normal}")

	if err != nil {
		panic(err)
	}

	fmt.Println(formatted)
	// Output: Escape sequences disabled
}

func ExampleFormatter_AreEscapeSequencesEnabled() {
	f := formatter.New()

	fmt.Println(f.DisableEscapeSequences().AreEscapeSequencesEnabled())
	fmt.Println(f.EnableEscapeSequences().AreEscapeSequencesEnabled())
	// Output:
	// false
	// true
}

func TestFormatterNew(test *testing.T) {
	assert.NotNil(test, formatter.New())
}

func TestFormatterNoArguments(test *testing.T) {
	formatted, err := formatter.New().Format("My message")

	assert.NoError(test, err)
	assert.Equal(test, "My message", formatted)
}

func TestFormatterArguments(test *testing.T) {
	formatted, err := formatter.New().Format("My test", 13, "foo", false)

	assert.NoError(test, err)
	assert.Equal(test, "My test 13 foo false", formatted)
}

func TestFormatterAutomaticArguments(test *testing.T) {
	formatted, err := formatter.New().Format("{p} {p}-{p}", 4, "test", true)

	assert.NoError(test, err)
	assert.Equal(test, "4 test-true", formatted)
}

func TestFormatterPositionalArguments(test *testing.T) {
	formatted, err := formatter.New().Format("{p1} {p0}", 1, 2)

	assert.NoError(test, err)
	assert.Equal(test, "2 1", formatted)
}

func TestFormatterNamedArguments(test *testing.T) {
	formatted, err := formatter.New().Format("{z} {y} {x} {z}", formatter.Named{
		"x": 1,
		"y": 2,
		"z": 3,
	}, "c")

	assert.NoError(test, err)
	assert.Equal(test, "3 2 1 3 c", formatted)
}

func TestFormatterObjectArguments(test *testing.T) {
	formatted, err := formatter.New().Format("{.Z} {.Y} {.X} {.Z}", struct {
		X, Y, Z int
	}{
		X: 4,
		Y: 5,
		Z: 6,
	}, "b")

	assert.NoError(test, err)
	assert.Equal(test, "6 5 4 6 b", formatted)
}

func TestFormatterObjectArgumentsError(test *testing.T) {
	formatted, err := formatter.New().SetLeftDelimiter("[").Format("[.Z} [.Y} [.X} [.Z}", struct {
		X, Y, Z int
	}{
		X: 7,
		Y: 8,
		Z: 9,
	}, "c")

	assert.Error(test, err)
	assert.Empty(test, formatted)
}

func TestFormatterReset(test *testing.T) {
	f := formatter.New()

	assert.Equal(test, "z", f.SetPlaceholder("z").GetPlaceholder())
	assert.Equal(test, "[", f.SetDelimiters("[", "]").GetLeftDelimiter())
	assert.NotEmpty(test, f.AddFunction("f", func() {}).GetFunctions())

	assert.Equal(test, f, f.Reset())
	assert.Equal(test, formatter.DefaultPlaceholder, f.GetPlaceholder())
	assert.Equal(test, formatter.DefaultLeftDelimiter, f.GetLeftDelimiter())
	assert.Equal(test, formatter.DefaultRightDelimiter, f.GetRightDelimiter())
	assert.Empty(test, f.GetFunctions())
}

func TestFormatterDelimiters(test *testing.T) {
	f := formatter.New().SetDelimiters("<", ">")

	left, right := f.GetDelimiters()

	assert.Equal(test, "<", left)
	assert.Equal(test, ">", right)

	formatted, err := f.Format("<p1> <p0>", "c", 3)

	assert.NoError(test, err)
	assert.Equal(test, "3 c", formatted)

	assert.Equal(test, f, f.ResetDelimiters())
	assert.Equal(test, formatter.DefaultLeftDelimiter, f.GetLeftDelimiter())
	assert.Equal(test, formatter.DefaultRightDelimiter, f.GetRightDelimiter())
}

func TestFormatterPlaceholder(test *testing.T) {
	f := formatter.New().SetPlaceholder("c")

	formatted, err := f.Format("{c1} {c0}", "d", 4)

	assert.NoError(test, err)
	assert.Equal(test, "4 d", formatted)

	assert.Equal(test, f, f.ResetPlaceholder())
	assert.Equal(test, formatter.DefaultPlaceholder, f.GetPlaceholder())
}

func TestFormatterFunctions(test *testing.T) {
	f := formatter.New().SetFunctions(formatter.Functions{
		"Nn": func() int { return 8 },
	})

	assert.NotNil(test, f.AddFunctions(formatter.Functions{
		"cc": func() string { return "C" },
		"xx": func() error { return nil },
	}))

	assert.Len(test, f.GetFunctions(), 3)
	assert.NotNil(test, f.GetFunction("cc"))

	formatted, err := f.Format("{cc} {Nn}")

	assert.NoError(test, err)
	assert.Equal(test, "C 8", formatted)
	assert.Len(test, f.RemoveFunction("cc").GetFunctions(), 2)
	assert.Len(test, f.RemoveFunction("cc").GetFunctions(), 2)
	assert.Len(test, f.RemoveFunctions([]string{"xx"}).GetFunctions(), 1)
	assert.Empty(test, f.ResetFunctions().GetFunctions())
}

func TestFormatterFormatError(test *testing.T) {
	formatted, err := formatter.New().Format("{c}", 3)

	assert.Error(test, err)
	assert.Empty(test, formatted)
}

func TestFormatterWriterError(test *testing.T) {
	writer := new(WriterError)

	assert.Error(test, formatter.New().FormatWriter(writer, "{p}", 2))
}

func TestFormatterUnusedError(test *testing.T) {
	writer := &WriterError{
		Skip: 1,
	}

	assert.Error(test, formatter.New().FormatWriter(writer, "{p}", 2, 5))
}

func TestFormatterObjectNil(test *testing.T) {
	var object *struct {
		Foo int
	}

	formatted, err := formatter.New().Format("Message {p} {p} {p2}", 3, object, "object")

	assert.Nil(test, object)
	assert.NoError(test, err)
	assert.Equal(test, "Message 3 <nil> object", formatted)
}

func TestFormatterMustFormatPanics(test *testing.T) {
	assert.Panics(test, func() {
		formatter.MustFormat("{invalid}")
	})
}

func TestFormatterColors(test *testing.T) {
	formatted, err := formatter.Format("{red}red{normal} {green}green{normal} {blue}blue{normal}")

	assert.NoError(test, err)
	assert.Equal(test, "\x1b[31mred\x1b[0m \x1b[32mgreen\x1b[0m \x1b[34mblue\x1b[0m", formatted)
}

func TestFormatterRGB(test *testing.T) {
	formatted, err := formatter.Format("{rgb 255 165 0}funky{normal}")

	assert.NoError(test, err)
	assert.Equal(test, "\x1b[38;2;255;165;0mfunky\x1b[0m", formatted)
}

func TestFormatterRGBBackground(test *testing.T) {
	formatted, err := formatter.Format("{rgb 255 165 0 | background}funky{normal}")

	assert.NoError(test, err)
	assert.Equal(test, "\x1b[48;2;255;165;0mfunky\x1b[0m", formatted)
}

func TestFormatterRGBBackgroundForeground(test *testing.T) {
	formatted, err := formatter.Format("{rgb 0 165 7 | background | foreground}funky{normal}")

	assert.NoError(test, err)
	assert.Equal(test, "\x1b[38;2;0;165;7mfunky\x1b[0m", formatted)
}

func TestFormatterRGBForeground(test *testing.T) {
	formatted, err := formatter.Format("{rgb 255 165 0 | foreground}funky{normal}")

	assert.NoError(test, err)
	assert.Equal(test, "\x1b[38;2;255;165;0mfunky\x1b[0m", formatted)
}

func TestFormatterRGBBackgroundBackground(test *testing.T) {
	formatted, err := formatter.Format("{rgb 255 165 3 | background | background}funky{normal}")

	assert.NoError(test, err)
	assert.Equal(test, "\x1b[48;2;255;165;3mfunky\x1b[0m", formatted)
}

func TestFormatterBrightColors(test *testing.T) {
	formatted, err := formatter.Format("{magenta | bright}magenta{normal}")

	assert.NoError(test, err)
	assert.Equal(test, "\x1b[95mmagenta\x1b[0m", formatted)
}

func TestFormatterBrightBright(test *testing.T) {
	formatted, err := formatter.Format("{red | bright | bright}red{normal}")

	assert.NoError(test, err)
	assert.Equal(test, "\x1b[91mred\x1b[0m", formatted)
}

func TestFormatterBrightInvalid(test *testing.T) {
	formatted, err := formatter.Format("{print 5 | bright}magenta{normal}")

	assert.Error(test, err)
	assert.Empty(test, formatted)
}

func TestFormatterBrightError(test *testing.T) {
	formatted, err := formatter.Format("{normal | bright}magenta{normal}")

	assert.Error(test, err)
	assert.Empty(test, formatted)
}

func TestFormatterBackgroundColors(test *testing.T) {
	formatted, err := formatter.Format("{yellow | background}yellow{normal}")

	assert.NoError(test, err)
	assert.Equal(test, "\x1b[43myellow\x1b[0m", formatted)
}

func TestFormatterBackgroundForeground(test *testing.T) {
	formatted, err := formatter.Format("{white | background | foreground}white{normal}")

	assert.NoError(test, err)
	assert.Equal(test, "\x1b[37mwhite\x1b[0m", formatted)
}

func TestFormatterBackgroundDefault(test *testing.T) {
	formatted, err := formatter.Format("{gray | background}gray{default | background}")

	assert.NoError(test, err)
	assert.Equal(test, "\x1b[100mgray\x1b[49m", formatted)
}

func TestFormatterBackgroundBackground(test *testing.T) {
	formatted, err := formatter.Format("{yellow | background | background}yellow{normal}")

	assert.NoError(test, err)
	assert.Equal(test, "\x1b[43myellow\x1b[0m", formatted)
}

func TestFormatterBackgroundInvalid(test *testing.T) {
	formatted, err := formatter.Format("{print 6 | background}yellow{normal}")

	assert.Error(test, err)
	assert.Empty(test, formatted)
}

func TestFormatterBackgroundError(test *testing.T) {
	formatted, err := formatter.Format("{blink | background}yellow{normal}")

	assert.Error(test, err)
	assert.Empty(test, formatted)
}

func TestFormatterBackgroundBrightColors(test *testing.T) {
	formatted, err := formatter.Format("{cyan | bright | background}cyan{normal}")

	assert.NoError(test, err)
	assert.Equal(test, "\x1b[106mcyan\x1b[0m", formatted)
}

func TestFormatterForegroundColors(test *testing.T) {
	formatted, err := formatter.Format("{yellow | foreground}yellow{normal}")

	assert.NoError(test, err)
	assert.Equal(test, "\x1b[33myellow\x1b[0m", formatted)
}

func TestFormatterForegroundDefault(test *testing.T) {
	formatted, err := formatter.Format("{green | foreground}green{default | foreground}")

	assert.NoError(test, err)
	assert.Equal(test, "\x1b[32mgreen\x1b[39m", formatted)
}

func TestFormatterForegroundForeground(test *testing.T) {
	formatted, err := formatter.Format("{black | foreground | foreground}black{normal}")

	assert.NoError(test, err)
	assert.Equal(test, "\x1b[30mblack\x1b[0m", formatted)
}

func TestFormatterForegroundInvalid(test *testing.T) {
	formatted, err := formatter.Format("{print 6 | foreground}yellow{normal}")

	assert.Error(test, err)
	assert.Empty(test, formatted)
}

func TestFormatterForegroundError(test *testing.T) {
	formatted, err := formatter.Format("{blink | foreground}yellow{normal}")

	assert.Error(test, err)
	assert.Empty(test, formatted)
}

func TestFormatterForegroundBrightColors(test *testing.T) {
	formatted, err := formatter.Format("{cyan | bright | foreground}cyan{normal}")

	assert.NoError(test, err)
	assert.Equal(test, "\x1b[96mcyan\x1b[0m", formatted)
}

func TestFormatterError(test *testing.T) {
	formatted, err := formatter.Format("Error", error(&StructError{"message"}))

	assert.NoError(test, err)
	assert.Equal(test, "Error message", formatted)
}

func TestFormatterValueError(test *testing.T) {
	formatted, err := formatter.Format("Error", error(StructValueError{"message"}))

	assert.NoError(test, err)
	assert.Equal(test, "Error message", formatted)
}

func TestFormatterBold(test *testing.T) {
	formatted, err := formatter.Format("{bold}text{bold | off}")

	assert.NoError(test, err)
	assert.Equal(test, "\x1b[1mtext\x1b[21m", formatted)
}

func TestFormatterFaint(test *testing.T) {
	formatted, err := formatter.Format("{faint}text{faint | off}")

	assert.NoError(test, err)
	assert.Equal(test, "\x1b[2mtext\x1b[22m", formatted)
}

func TestFormatterItalic(test *testing.T) {
	formatted, err := formatter.Format("{italic}text{italic | off}")

	assert.NoError(test, err)
	assert.Equal(test, "\x1b[3mtext\x1b[23m", formatted)
}

func TestFormatterUnderline(test *testing.T) {
	formatted, err := formatter.Format("{underline}text{underline | off}")

	assert.NoError(test, err)
	assert.Equal(test, "\x1b[4mtext\x1b[24m", formatted)
}

func TestFormatterBlink(test *testing.T) {
	formatted, err := formatter.Format("{blink}text{blink | off}")

	assert.NoError(test, err)
	assert.Equal(test, "\x1b[5mtext\x1b[25m", formatted)
}

func TestFormatterInvert(test *testing.T) {
	formatted, err := formatter.Format("{invert}text{invert | off}")

	assert.NoError(test, err)
	assert.Equal(test, "\x1b[7mtext\x1b[27m", formatted)
}

func TestFormatterHide(test *testing.T) {
	formatted, err := formatter.Format("{hide}text{hide | off}")

	assert.NoError(test, err)
	assert.Equal(test, "\x1b[8mtext\x1b[28m", formatted)
}

func TestFormatterStrike(test *testing.T) {
	formatted, err := formatter.Format("{strike}text{strike | off}")

	assert.NoError(test, err)
	assert.Equal(test, "\x1b[9mtext\x1b[29m", formatted)
}

func TestFormatterOverline(test *testing.T) {
	formatted, err := formatter.Format("{overline}text{overline | off}")

	assert.NoError(test, err)
	assert.Equal(test, "\x1b[53mtext\x1b[55m", formatted)
}

func TestFormatterOffOff(test *testing.T) {
	formatted, err := formatter.Format("{blink | off | off}")

	assert.NoError(test, err)
	assert.Equal(test, "\x1b[25m", formatted)
}

func TestFormatterOffError(test *testing.T) {
	formatted, err := formatter.Format("{print 5 | off}")

	assert.Error(test, err)
	assert.Empty(test, formatted)
}

func TestFormatterOffInvalid(test *testing.T) {
	formatted, err := formatter.Format("{red | off}")

	assert.Error(test, err)
	assert.Empty(test, formatted)
}

func TestFormatterIPAddress(test *testing.T) {
	formatted, err := formatter.Format("{ip}")

	assert.NoError(test, err)
	assert.NotEmpty(test, formatted)
}

func TestFormatterIPAddressDialError(test *testing.T) {
	defer func() {
		formatter.Dial = net.Dial
	}()

	controller := gomock.NewController(test)
	defer controller.Finish()

	connection := mocks.NewMockConn(controller)

	connection.EXPECT().Close().Times(1)
	connection.EXPECT().LocalAddr().Times(1).Return(new(net.UDPAddr))

	formatter.Dial = func(string, string) (net.Conn, error) {
		return connection, Error("error")
	}

	formatted, err := formatter.Format("{ip}")

	assert.NoError(test, err)
	assert.NotEmpty(test, formatted)
}

func TestFormatterIPAddressDialNil(test *testing.T) {
	defer func() {
		formatter.Dial = net.Dial
	}()

	formatter.Dial = func(string, string) (net.Conn, error) {
		return nil, Error("error")
	}

	formatted, err := formatter.Format("{ip}")

	assert.NoError(test, err)
	assert.NotEmpty(test, formatted)
	assert.Equal(test, "127.0.0.1", formatted)
}

func TestFormatterIPAddressCloseError(test *testing.T) {
	defer func() {
		formatter.Dial = net.Dial
	}()

	controller := gomock.NewController(test)
	defer controller.Finish()

	connection := mocks.NewMockConn(controller)

	connection.EXPECT().Close().Times(1).Return(Error("error"))
	connection.EXPECT().LocalAddr().Times(1).Return(new(net.UDPAddr))

	formatter.Dial = func(string, string) (net.Conn, error) {
		return connection, nil
	}

	formatted, err := formatter.Format("{ip}")

	assert.NoError(test, err)
	assert.NotEmpty(test, formatted)
}

func TestFormatterUser(test *testing.T) {
	formatted, err := formatter.Format("{user}")

	assert.NoError(test, err)
	assert.NotEmpty(test, formatted)
}

func TestFormatterUserError(test *testing.T) {
	defer func() {
		formatter.Current = user.Current
	}()

	formatter.Current = func() (*user.User, error) {
		return nil, Error("error")
	}

	formatted, err := formatter.Format("{user}")

	assert.NoError(test, err)
	assert.Empty(test, formatted)
}

func TestFormatterNow(test *testing.T) {
	formatted, err := formatter.Format("{now}")

	assert.NoError(test, err)
	assert.NotEmpty(test, formatted)
}

func TestFormatterISO8601(test *testing.T) {
	formatted, err := formatter.Format("{now | iso8601}")

	assert.NoError(test, err)
	assert.NotEmpty(test, formatted)
}

func TestFormatterUpper(test *testing.T) {
	formatted, err := formatter.Format(`{"text" | upper}`)

	assert.NoError(test, err)
	assert.Equal(test, "TEXT", formatted)
}

func TestFormatterLower(test *testing.T) {
	formatted, err := formatter.Format(`{"teXt" | lower}`)

	assert.NoError(test, err)
	assert.Equal(test, "text", formatted)
}

func TestFormatterCapitalize(test *testing.T) {
	formatted, err := formatter.Format(`{"text" | capitalize}`)

	assert.NoError(test, err)
	assert.Equal(test, "Text", formatted)
}

func TestFormatterColor(test *testing.T) {
	formatted, err := formatter.Format(`{color "red"}red{normal}`)

	assert.NoError(test, err)
	assert.Equal(test, "\x1b[31mred\x1b[0m", formatted)
}

func TestFormatterColorHex(test *testing.T) {
	formatted, err := formatter.Format(`{color "0xF3AC67"}funky{normal}`)

	assert.NoError(test, err)
	assert.Equal(test, "\x1b[38;2;243;172;103mfunky\x1b[0m", formatted)
}

func TestFormatterColorHexError(test *testing.T) {
	formatted, err := formatter.Format(`{color "0xFFF3AC67"}funky{normal}`)

	assert.Error(test, err)
	assert.Empty(test, formatted)
}

func TestFormatterColorInvalid(test *testing.T) {
	formatted, err := formatter.Format(`{color "foo"}`)

	assert.Error(test, err)
	assert.Empty(test, formatted)
}

func TestFormatterBell(test *testing.T) {
	formatted, err := formatter.Format("{bell}")

	assert.NoError(test, err)
	assert.Equal(test, "\a", formatted)
}

func TestFormatterAreEscapeSequencesSupported(test *testing.T) {
	defer func(value string) {
		assert.NoError(test, os.Setenv(formatter.ForceEscapeSequencesEnv, value))
	}(os.Getenv(formatter.ForceEscapeSequencesEnv))

	assert.NoError(test, os.Setenv(formatter.ForceEscapeSequencesEnv, "true"))
	assert.True(test, formatter.AreEscapeSequencesSupported())

	assert.NoError(test, os.Setenv(formatter.ForceEscapeSequencesEnv, "false"))
	assert.False(test, formatter.AreEscapeSequencesSupported())

	supported := (os.Getenv("TERM") != "dumb") && (isatty.IsTerminal(os.Stdout.Fd()) || isatty.IsCygwinTerminal(os.Stdout.Fd()))

	assert.NoError(test, os.Setenv(formatter.ForceEscapeSequencesEnv, ""))
	assert.Equal(test, supported, formatter.AreEscapeSequencesSupported())
}

func TestFormatterObject(test *testing.T) {
	object := struct {
		Value   int
		Message string
	}{
		Value:   2,
		Message: "text",
	}

	formatted, err := formatter.Format("", object)

	assert.NoError(test, err)
	assert.Equal(test, "{2 text}", formatted)

	formatted, err = formatter.Format(" ", object)

	assert.NoError(test, err)
	assert.Equal(test, " {2 text}", formatted)
}

func TestFormatterFields(test *testing.T) {
	object := struct {
		Value   int
		Message string
	}{
		Value:   3,
		Message: "text",
	}

	formatted, err := formatter.Format("{p | fields}", object)

	assert.NoError(test, err)
	assert.Equal(test, "{Value:3 Message:text}", formatted)
}

func TestFormatterJSON(test *testing.T) {
	object := struct {
		Value   int
		Message string
	}{
		Value:   4,
		Message: "text",
	}

	formatted, err := formatter.Format("{p | json}", object)

	assert.NoError(test, err)
	assert.Equal(test, `{"Value":4,"Message":"text"}`, formatted)
}

func TestFormatterJSONIndent(test *testing.T) {
	object := struct {
		Value   int
		Message string
	}{
		Value:   5,
		Message: "text",
	}

	formatted, err := formatter.Format("{p | json | indent}", object)

	assert.NoError(test, err)
	assert.Equal(test, "{\n\t\"Value\": 5,\n\t\"Message\": \"text\"\n}", formatted)
}

func TestFormatterJSONIndentError(test *testing.T) {
	formatted, err := formatter.Format("{print \"[\" | indent}")

	assert.Error(test, err)
	assert.Empty(test, formatted)
}

func TestFormatterJSONError(test *testing.T) {
	object := struct {
		Invalid chan struct{}
	}{}

	formatted, err := formatter.Format("{p | json}", object)

	assert.Error(test, err)
	assert.Empty(test, formatted)
}
