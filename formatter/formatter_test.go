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
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/tymonx/go-formatter/formatter"
)

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
