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

package formats

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/decentplatforms/appkit/logf"
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

var formats = map[string]logf.Formatter{
	"syslog_rfc3164": Syslog3164Format(SyslogConfig{}),
	"syslog_rfc5424": Syslog5424Format(SyslogConfig{}),
	"json": JSONFormat(JSONConfig{
		TimeFormat: time.RFC3339,
	}),
	"kv": KVFormat(KVConfig{
		TimeFormat: time.RFC3339,
	}),
	"kv_single_quote": KVFormat(KVConfig{
		TimeFormat:      time.RFC3339,
		UseSingleQuotes: true,
	}),
	"json_pretty": JSONPrettyFormat(JSONConfig{Indent: "  ", TimeFormat: time.RFC3339}),
}

func testProps() []logf.Prop {
	return []logf.Prop{
		logf.String("property", "value"),
		logf.Int("num1", 1),
		logf.Float("num2", 2.0),
	}
}

func TestLogger(t *testing.T) {
	t.Run("default logger", func(t *testing.T) {
		for name, format := range formats {
			tw := &TestWriter{}
			conf := logf.Config{
				MaxLevel:     logf.Warning,
				DefaultLevel: logf.Informational,
				Format:       format,
				Output:       tw,
			}
			t.Run(name+" format", func(t *testing.T) {
				conf.Format = format
				log, err := logf.NewLogger(conf)
				if err != nil {
					t.Fatal(err)
				}
				for i := logf.MOST_SEVERE; i < logf.LEAST_SEVERE; i++ {
					lvl := logf.LogLevel(i)
					msg := fmt.Sprintf("test log at level %s", lvl)
					log.Log(logf.LogLevel(i), msg)
					if i <= conf.MaxLevel {
						if expected := format.FormatAndNormalize(lvl, msg, logf.NewProps()); tw.Last != expected {
							t.Error("wrong log at", lvl, tw.Last, expected)
						}
					} else {
						if tw.Last != "" {
							t.Error("logger shouldn't have logged at", lvl)
						}
					}
					tw.Last = ""
				}
			})
		}
	})
	t.Run("with props", func(t *testing.T) {
		for name, format := range formats {
			tw := &TestWriter{}
			conf := logf.Config{
				MaxLevel:     logf.Warning,
				DefaultLevel: logf.Informational,
				Format:       format,
				Output:       tw,
			}
			t.Run(name+" format", func(t *testing.T) {
				conf.Format = format
				log, err := logf.NewLogger(conf)
				if err != nil {
					t.Fatal(err)
				}
				for i := logf.MOST_SEVERE; i < logf.LEAST_SEVERE; i++ {
					lvl := logf.LogLevel(i)
					msg := fmt.Sprintf("test log at level %s", lvl)
					log.Log(logf.LogLevel(i), msg, testProps()...)
					if i <= conf.MaxLevel {
						if expected := format.FormatAndNormalize(lvl, msg, logf.NewProps(testProps()...)); tw.Last != expected {
							t.Error("wrong log at", lvl, tw.Last, expected)
						}
					} else {
						if tw.Last != "" {
							t.Error("logger shouldn't have logged at", lvl)
						}
					}
					tw.Last = ""
				}
			})
		}
	})
}
