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

package main

import (
	"fmt"

	"gitlab.com/tymonx/go-formatter/formatter"
)

func main() {
	fmt.Println(formatter.MustFormat("{black}{white | background}Black{reset}"))
	fmt.Println(formatter.MustFormat("{red}Red{reset}"))
	fmt.Println(formatter.MustFormat("{green}Green{reset}"))
	fmt.Println(formatter.MustFormat("{yellow}Yellow{reset}"))
	fmt.Println(formatter.MustFormat("{blue}Blue{reset}"))
	fmt.Println(formatter.MustFormat("{magenta}Magenta{reset}"))
	fmt.Println(formatter.MustFormat("{cyan}Cyan{reset}"))
	fmt.Println(formatter.MustFormat("{white}{black | background}White{reset}"))
	fmt.Println(formatter.MustFormat("{gray}Gray{reset}"))
	fmt.Println(formatter.MustFormat("{rgb 255 99 71}Tomato{reset}"))
	fmt.Println(formatter.MustFormat(`{color "0xADFF2F"}Greeny yellow{reset}`))

	fmt.Println()
	fmt.Println(formatter.MustFormat("{black | bright}Bright black{reset}"))
	fmt.Println(formatter.MustFormat("{red | bright}Bright red{reset}"))
	fmt.Println(formatter.MustFormat("{green | bright}Bright green{reset}"))
	fmt.Println(formatter.MustFormat("{yellow | bright}Bright yellow{reset}"))
	fmt.Println(formatter.MustFormat("{blue | bright}Bright blue{reset}"))
	fmt.Println(formatter.MustFormat("{magenta | bright}Bright magenta{reset}"))
	fmt.Println(formatter.MustFormat("{cyan | bright}Bright cyan{reset}"))
	fmt.Println(formatter.MustFormat("{white | bright}{black | background}Bright white{reset}"))
	fmt.Println(formatter.MustFormat("{gray | bright}Bright gray{reset}"))

	fmt.Println()
	fmt.Println(formatter.MustFormat("{black | background}{white}Background black{reset}"))
	fmt.Println(formatter.MustFormat("{red | background}Background red{reset}"))
	fmt.Println(formatter.MustFormat("{green | background}Background green{reset}"))
	fmt.Println(formatter.MustFormat("{yellow | background}Background yellow{reset}"))
	fmt.Println(formatter.MustFormat("{blue | background}Background blue{reset}"))
	fmt.Println(formatter.MustFormat("{magenta | background}Background magenta{reset}"))
	fmt.Println(formatter.MustFormat("{cyan | background}Background cyan{reset}"))
	fmt.Println(formatter.MustFormat("{white | background}{black}Background white{reset}"))
	fmt.Println(formatter.MustFormat("{gray | background}Background gray{reset}"))
	fmt.Println(formatter.MustFormat("{rgb 255 99 71 | background}Background tomato{reset}"))
	fmt.Println(formatter.MustFormat(`{color "0xADFF2F" | background}Background greeny yellow{reset}`))

	fmt.Println()
	fmt.Println(formatter.MustFormat("{bold}Bold{reset}"))
	fmt.Println(formatter.MustFormat("{faint}Faint{reset}"))
	fmt.Println(formatter.MustFormat("{italic}Italic{reset}"))
	fmt.Println(formatter.MustFormat("{underline}Underline{reset}"))
	fmt.Println(formatter.MustFormat("{blink}Blink{reset}"))
	fmt.Println(formatter.MustFormat("{overline}Overline{reset}"))
	fmt.Println(formatter.MustFormat("{invert}Invert{reset}"))
	fmt.Println(formatter.MustFormat("Hide: '{hide}Hide{reset}'"))
	fmt.Println(formatter.MustFormat("{strike}Strike{reset}"))

	fmt.Println()
	fmt.Println(formatter.MustFormat("IP: {ip}"))
	fmt.Println(formatter.MustFormat("User: {user}"))
	fmt.Println(formatter.MustFormat("Executable: {executable}"))
	fmt.Println(formatter.MustFormat("Current working directory: {cwd}"))
	fmt.Println(formatter.MustFormat("Hostname: {hostname}"))
	fmt.Println(formatter.MustFormat(`Environment: USER={env "USER"}`))
	fmt.Println(formatter.MustFormat("User ID: {uid}"))
	fmt.Println(formatter.MustFormat("Group ID: {gid}"))
	fmt.Println(formatter.MustFormat("Effective user ID: {euid}"))
	fmt.Println(formatter.MustFormat("Effective group ID: {egid}"))
	fmt.Println(formatter.MustFormat("Process ID: {pid}"))
	fmt.Println(formatter.MustFormat("Parent process ID: {ppid}"))
	fmt.Println(formatter.MustFormat("Bell{bell}"))

	fmt.Println()
	fmt.Println(formatter.MustFormat("Now: {now}"))
	fmt.Println(formatter.MustFormat("ISO 8601: {now | iso8601}"))
}
