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
	"fmt"
	"os"
	"time"

	"github.com/decentplatforms/appkit/logf"
)

type SDElement struct {
	Name string
}

const (
	SYSLOG_HOSTNAME = string("log_syslog_hostname")
	SYSLOG_APPNAME  = string("log_syslog_appname")
	SYSLOG_TAG      = string("log_syslog_TAG")
)

// SyslogConfig sets default values for SyslogXFormat loggers.
// Providing SYSLOG_X props overrides these headers; otherwise the values from config are used.
// Usage notes:
//   - Tag is used as MSGID in 5424
//   - UseISO8601 only applies to RFC 3164; rfc5424 specifies RFC3339 time
//   - You may not set Facility to 0
//   - WithProps uses formats.SyslogKV by default.
type SyslogConfig struct {
	Hostname   string
	AppName    string
	Tag        string
	Facility   int
	UseISO8601 bool
	WithProps  func(string, *logf.Props) string
}

// SyslogJSON is an option for SyslogConfig.WithProps.
// It appends spare props as JSON to the syslog message.
func SyslogJSON(msg string, props *logf.Props) string {
	spareProps := props.Map()
	if len(spareProps) > 0 {
		raw, err := json.Marshal(spareProps)
		if err == nil {
			msg = msg + " " + string(raw)
		}
	}
	return msg
}

func SyslogKV(msg string, props *logf.Props) string {
	return msg + " " + formatProps(props, false)
}

// SyslogIgnore is an option for SyslogConfig.WithProps.
// It ignores spare props, leaving the message as-is.
func SyslogIgnore(msg string, props *logf.Props) string {
	return msg
}

func (conf SyslogConfig) withDefaults() SyslogConfig {
	if conf.Hostname == "" {
		oshost, err := os.Hostname()
		if err != nil {
			conf.Hostname = "nilvalue"
		} else {
			conf.Hostname = oshost
		}
	}
	if conf.AppName == "" {
		conf.AppName = "log"
	}
	if conf.Tag == "" {
		conf.Tag = "log"
	}
	if conf.Facility == 0 {
		conf.Facility = 1
	}
	if conf.WithProps == nil {
		conf.WithProps = SyslogKV
	}
	return conf
}

// Syslog5424Format provides the syslog format (RFC5424) with the following conventions:
//   - Timestamps are RFC3339 in UTC
//   - Hostname is the log.SYSLOG_HOSTNAME prop, conf.Hostname, the machine's hostname at process start, or NILVALUE
//   - App name is log.SYSLOG_APPNAME prop, conf.AppName, or log
//   - Process ID is the application's process ID
//   - Message ID is the log.SYSLOG_MSGID prop, conf.MsgId, or log
//   - Facility is conf.Facility or User (1) and may not be 0
//   - Version is 1
//   - Structured Data is -. Support for this field is planned.
func Syslog5424Format(conf SyslogConfig) logf.Formatter {
	conf = conf.withDefaults()
	return func(level logf.LogLevel, msg string, props *logf.Props) string {
		var timestamp, hostname, appname, msgid, structured string
		var facility, pri, version, pid int
		var ok bool

		timestamp = time.Now().UTC().Format(time.RFC3339)

		if hostname, ok = props.Get(SYSLOG_HOSTNAME).(string); ok {
		} else {
			hostname = conf.Hostname
		}

		if appname, ok = props.Get(SYSLOG_APPNAME).(string); ok {
		} else {
			appname = conf.AppName
		}

		if msgid, ok = props.Get(SYSLOG_TAG).(string); ok {
		} else {
			msgid = conf.Tag
		}

		structured = "-"

		facility = conf.Facility

		pri = 8*facility + int(level)
		version = 1
		pid = os.Getpid()

		props.Delete(SYSLOG_HOSTNAME, SYSLOG_APPNAME, SYSLOG_TAG)
		msg = conf.WithProps(msg, props)

		return fmt.Sprintf("<%d>%d %s %s %s %d %s %s %s", pri, version, timestamp, hostname, appname, pid, msgid, structured, msg)
	}
}

// Syslog3164Format provides the syslog format (RFC3164) with the following conventions:
//   - Timestamps are time.Stamp in UTC (Mmm dd hh:mm:ss)
//   - Hostname is the log.SYSLOG_HOSTNAME prop, the machine's hostname at process start, or NILVALUE
//   - Tag is the log.SYSLOG_TAG prop or log
//   - Facility is User (1)
//
// Spare props are appended to MSG as JSON.
func Syslog3164Format(conf SyslogConfig) logf.Formatter {
	conf = conf.withDefaults()
	return func(level logf.LogLevel, msg string, props *logf.Props) string {
		var timestamp, hostname, tag string
		var facility, pri int
		var ok bool

		timestamp = time.Now().UTC().Format(time.Stamp)
		if conf.UseISO8601 {
			timestamp = time.Now().UTC().Format(time.RFC3339)
		}

		if hostname, ok = props.Get(SYSLOG_HOSTNAME).(string); ok {
		} else {
			hostname = conf.Hostname
		}

		if tag, ok = props.Get(SYSLOG_TAG).(string); ok {
		} else {
			tag = conf.Tag
		}

		facility = conf.Facility

		pri = 8*facility + int(level)

		props.Delete(SYSLOG_HOSTNAME, SYSLOG_APPNAME, SYSLOG_TAG)
		msg = conf.WithProps(msg, props)

		return fmt.Sprintf("<%d>%s %s %s: %s", pri, timestamp, hostname, tag, msg)
	}
}
