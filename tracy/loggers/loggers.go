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

package loggers

import (
	"io"

	"github.com/decentplatforms/appkit/tracy"
	"github.com/decentplatforms/appkit/tracy/formats"
)

// Syslog3164 returns a logger that logs to output with max level max and default level def.
// Uses tag as the default tag; can be changed on a per-message basis using*Props.
func Syslog3164(tag string, timeDetail bool, max, def tracy.LogLevel, output io.Writer) tracy.Logger {
	logger, _ := tracy.NewLogger(tracy.Config{
		MaxLevel:     max,
		DefaultLevel: def,
		Format: formats.Syslog3164Format(formats.SyslogConfig{
			Tag:        tag,
			UseISO8601: timeDetail,
		}),
		Output: output,
	})
	return logger
}

// Syslog5424 returns a logger that logs to output with max level max and default level def.
// Uses app and msgid as defaults for those values; can be changed on a per-message basis using*Props.
func Syslog5424(app, msgid string, max, def tracy.LogLevel, output io.Writer) tracy.Logger {
	logger, _ := tracy.NewLogger(tracy.Config{
		MaxLevel:     max,
		DefaultLevel: def,
		Format: formats.Syslog5424Format(formats.SyslogConfig{
			AppName: app,
			Tag:     msgid,
		}),
		Output: output,
	})
	return logger
}

// JSON returns a logger that logs to output with max level max and default level def.
// Uses default JSONFormat settings.
func JSON(max, def tracy.LogLevel, output io.Writer) tracy.Logger {
	logger, _ := tracy.NewLogger(tracy.Config{
		MaxLevel:     max,
		DefaultLevel: def,
		Format:       formats.JSONFormat(formats.JSONConfig{}),
		Output:       output,
	})
	return logger
}

// JSON returns a logger that logs to output with max level max and default level def.
// Uses default JSONFormat settings plus the specified indent. If left empty, indents with tabs.
func JSONPretty(indent string, max, def tracy.LogLevel, output io.Writer) tracy.Logger {
	logger, _ := tracy.NewLogger(tracy.Config{
		MaxLevel:     max,
		DefaultLevel: def,
		Format:       formats.JSONPrettyFormat(formats.JSONConfig{Indent: indent}),
		Output:       output,
	})
	return logger
}
