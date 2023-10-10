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
	"encoding/json"
	"time"

	"github.com/decentplatforms/appkit/logf"
)

type jsonLog struct {
	Level     logf.LogLevel  `json:"level"`
	LevelStr  string         `json:"level_str"`
	Timestamp string         `json:"timestamp"`
	Message   string         `json:"message"`
	Props     map[string]any `json:"props,omitempty"`
}

type JSONConfig struct {
	TimeFormat string
	Prefix     string
	Indent     string
}

func (conf JSONConfig) withDefaults() JSONConfig {
	if conf.TimeFormat == "" {
		conf.TimeFormat = time.RFC3339
	}
	if conf.Indent == "" {
		conf.Indent = "\t"
	}
	return conf
}

func JSONFormat(conf JSONConfig) logf.Formatter {
	return func(level logf.LogLevel, msg string, props *logf.Props) string {
		raw, _ := json.Marshal(jsonLog{
			Level:     level,
			LevelStr:  level.String(),
			Timestamp: time.Now().UTC().Format(conf.TimeFormat),
			Message:   msg,
			Props:     props.Map(),
		})
		return string(raw)
	}
}

func JSONPrettyFormat(conf JSONConfig) logf.Formatter {
	return func(level logf.LogLevel, msg string, props *logf.Props) string {
		raw, _ := json.MarshalIndent(jsonLog{
			Level:     level,
			LevelStr:  level.String(),
			Timestamp: time.Now().UTC().Format(conf.TimeFormat),
			Message:   msg,
			Props:     props.Map(),
		}, conf.Prefix, conf.Indent)
		return string(raw)
	}
}
