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

// Error type.
type Error string

// Error message.
func (e Error) Error() string {
	return string(e)
}

// StructError type.
type StructError struct {
	value string
}

// Error message.
func (s *StructError) Error() string {
	return s.value
}

// StructValueError type.
type StructValueError struct {
	value string
}

// Error message.
func (s StructValueError) Error() string {
	return s.value
}

// WriterError mocks Writer interface and returns an error.
type WriterError struct {
	Skip int
}

// Write writes.
func (w *WriterError) Write([]byte) (int, error) {
	if w.Skip > 0 {
		w.Skip--
		return 0, nil
	}

	return 0, Error("error")
}
