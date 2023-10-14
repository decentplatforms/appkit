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
	"errors"
	"os"

	"github.com/decentplatforms/appkit/logf"
	"github.com/decentplatforms/appkit/logf/formats"
)

type TestWriter struct {
	Fail bool
	Last string
}

func (writer *TestWriter) Write(msg []byte) (n int, err error) {
	if writer.Fail {
		return 0, errors.New("writer configured to fail")
	}
	writer.Last = string(msg)
	return len(msg), nil
}

var loggers = map[string]logf.Logger{
	"syslog_rfc3164": Syslog3164(formats.SyslogConfig{Tag: "log-test", UseISO8601: true},
		logf.Informational, logf.Informational, os.Stdout),
	"syslog_rfc5424": Syslog5424(formats.SyslogConfig{AppName: "log-test", Tag: "log"},
		logf.Informational, logf.Informational, os.Stdout),
	"json":        JSON(logf.Informational, logf.Informational, os.Stdout),
	"json_pretty": JSONPretty(formats.JSONConfig{Indent: "  "}, logf.Informational, logf.Informational, os.Stdout),
	"kv":          KV(formats.KVConfig{}, logf.Informational, logf.Informational, os.Stdout),
}

// TODO: test loggers
