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

	"github.com/decentplatforms/appkit/logf"
	"github.com/decentplatforms/appkit/logf/formats"
)

// Syslog3164 returns a logger that logs to output with max level max and default level def.
// Uses the given Syslog3164Format settings.
func Syslog3164(conf formats.SyslogConfig, max, def logf.LogLevel, output io.Writer) logf.Logger {
	logger, _ := logf.NewLogger(logf.Config{
		MaxLevel:     max,
		DefaultLevel: def,
		Format:       formats.Syslog3164Format(conf),
		Output:       output,
	})
	return logger
}

// Syslog5424 returns a logger that logs to output with max level max and default level def.
// Uses the given Syslog5424Format settings.
func Syslog5424(conf formats.SyslogConfig, max, def logf.LogLevel, output io.Writer) logf.Logger {
	logger, _ := logf.NewLogger(logf.Config{
		MaxLevel:     max,
		DefaultLevel: def,
		Format:       formats.Syslog5424Format(conf),
		Output:       output,
	})
	return logger
}

// JSON returns a logger that logs to output with max level max and default level def.
// Uses default JSONFormat settings.
func JSON(max, def logf.LogLevel, output io.Writer) logf.Logger {
	logger, _ := logf.NewLogger(logf.Config{
		MaxLevel:     max,
		DefaultLevel: def,
		Format:       formats.JSONFormat(formats.JSONConfig{}),
		Output:       output,
	})
	return logger
}

// JSON returns a logger that logs to output with max level max and default level def.
// Uses the given JSONFormat settings.
func JSONPretty(conf formats.JSONConfig, max, def logf.LogLevel, output io.Writer) logf.Logger {
	logger, _ := logf.NewLogger(logf.Config{
		MaxLevel:     max,
		DefaultLevel: def,
		Format:       formats.JSONPrettyFormat(conf),
		Output:       output,
	})
	return logger
}

// KV returns a logger that logs to output with max level max and default level def.
// Uses the given KVFormat settings.
func KV(conf formats.KVConfig, max, def logf.LogLevel, output io.Writer) logf.Logger {
	logger, _ := logf.NewLogger(logf.Config{
		MaxLevel:     max,
		DefaultLevel: def,
		Format:       formats.KVFormat(conf),
		Output:       output,
	})
	return logger
}
