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
	"net"
)

// Dial is used only in testing and mocking.
var Dial = net.Dial // nolint: gochecknoglobals

func getIPAddress() string {
	connection, err := Dial("udp", "8.8.8.8:80")

	if connection == nil {
		return "127.0.0.1"
	}

	defer func() {
		err = connection.Close()
		_ = err
	}()

	return connection.LocalAddr().(*net.UDPAddr).IP.String()
}
