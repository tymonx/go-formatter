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

/*
Package formatter implements “replacement fields” surrounded by curly braces {} format strings.

Built-in functions

Simple example:

	formatted, err := formatter.Format("{italic}{red}{blink}blinky :){blink | off} no blinky :({default}")

Built-in text functions

List of built-in functions:

	reset      - All text attributes off
	normal 	   - All text attributes off, alias to reset
	default    - All text attributes off, alias to reset
	bold       - Bold text
	faint      - Faint text
	italic     - Italic text
	underline  - Underline text
	overline   - Overline text
	blink      - Blink text
	invert     - Swap foreground and background colors
	hide       - Hide text
	strike     - Strike text
	off        - Disable specific text attribute. Example: blink | off

Built-in string functions

List of built-in functions:

	upper      - Transform provided string to upper case. Example: upper "text"
	lower      - Transform provided string to lower case. Example: lower "TEXT"
	capitalize - Capitalize provided string. Example: capitalize "text"

Built-in color functions

List of built-in functions:

	black      - Black color
	red        - Red color
	green      - Green color
	yellow     - Yellow color
	blue       - Blue color
	magenta    - Magenta color
	cyan       - Cyan color
	white      - White color
	gray       - Gray color
	rgb        - 24-bit color, 3 arguments (red, green, blue), integer values between 0-255
	color      - Set color, 1 argument, color name like "red" or RGB HEX value in "0xXXXXXX" format
	bright     - Make color bright, used with standard color function. Example: green | bright
	foreground - Set as foreground color (default). Example: blue | foreground
	background - Set as background color. Example: cyan | background

Built-in OS functions

List of built-in functions:

	ip         - Get outbound (local) IP address
	user       - Get current user name
	executable - Get current executable path
	cwd        - Get current working directory path
	hostname   - Get hostname
	env        - Get environment variable
	expand     - Get and expand environment variable
	uid        - Get user ID
	gid        - Get group ID
	euid       - Get effective user ID
	egid       - Get effective group ID
	pid        - Get process ID
	ppid       - Get parent process ID
	bell       - Make a sound

Built-in time functions

List of built-in functions:

	now        - Get current time
	rfc3339    - Format time to RFC 3339. Example: now | rfc3339
	iso8601    - Format time to ISO 8601. Example: now | iso8601

Built-in path functions

List of built-in functions:

	absolute   - Returns an absolute representation of path
	base       - Returns the last element of path. Example: base "/dir/dir/file"
	clean      - Returns the shortest path name equivalent to path by purely lexical processing
	directory  - Returns all but the last element of path, typically the path's directory
	extension  - Returns the file name extension used by path. Example: extension "/dir/dir/file.ext"

Built-in object functions

List of built-in functions:

	fields    - Print also struct field names for given object. Example: p | fields
	json      - Marshal object to JSON. Example: p | json
	indent    - Indent marshaled JSON. Example: p | json | indent
*/
package formatter
