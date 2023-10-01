// Copyright 2023 appkit Authors
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

package tracy

// LogLevels denote log severity, with lower values being more severe
// The native set of LogLevels follows syslog severity, from EMERGENCY/0 to DEBUG/7.
// You can define extra constants using const LEVEL = LogLevel(<value>) but most logging
// systems use a subset of what's provided here. Suggested usage:
//
//   - Error: Something failed, and there's no way to automatically recover.
//   - Warning: Something failed, but the program can recover/continue.
//   - Informational: Messages that should be seen during normal, successful runs.
//   - Debug: Messages that should only be seen when debugging the application.
//
// Emergency, Alert, and Critical are usually reserved for OS errors, and Notice isn't used frequently.
//
// LogLevel implements fmt.Stringer, which by default maps LogLevels to their respective syslog keywords.
// You can call SetKeyword with a LogLevel (including self-defined levels) to set or change its string representation:
//
//	log.Error.SetKeyword("ERROR")
//	fmt.Println(log.Error) // ERROR
//
// Some presets are provided through log.Keywords_X() functions.
type LogLevel int8

const (
	Emergency = LogLevel(iota)
	Alert
	Critical
	Error
	Warning
	Notice
	Informational
	Debug
)

const MOST_SEVERE = Emergency
const LEAST_SEVERE = Debug

var keywords map[LogLevel]string = map[LogLevel]string{
	Emergency:     "emerg",
	Alert:         "alert",
	Critical:      "crit",
	Error:         "err",
	Warning:       "warn",
	Notice:        "notice",
	Informational: "info",
	Debug:         "debug",
}

func (level LogLevel) SetKeyword(keyword string) {
	keywords[level] = keyword
}

func (level LogLevel) String() string {
	return keywords[level]
}

// ===== KEYWORD SETS =====

func Keywords_Syslog() {
	keywords = map[LogLevel]string{
		Emergency:     "emerg",
		Alert:         "alert",
		Critical:      "crit",
		Error:         "err",
		Warning:       "warn",
		Notice:        "notice",
		Informational: "info",
		Debug:         "debug",
	}
}

func Keywords_AllCaps() {
	keywords = map[LogLevel]string{
		Emergency:     "EMERGENCY",
		Alert:         "ALERT",
		Critical:      "CRITICAL",
		Error:         "ERROR",
		Warning:       "WARN",
		Notice:        "NOTICE",
		Informational: "INFORMATIONAL",
		Debug:         "DEBUG",
	}
}
