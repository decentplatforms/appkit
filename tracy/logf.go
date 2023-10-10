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

package logf

import (
	"io"
)

type Logger interface {
	io.Writer
	Configure(conf Config) error
	Log(level LogLevel, msg string, props ...Prop) error
}

type logger struct {
	Config
}

func NewLogger(conf Config) (Logger, error) {
	log := &logger{}
	err := log.Configure(conf)
	return log, err
}

func (log *logger) Configure(conf Config) error {
	log.Config = conf
	if log.Output == nil {
		return NilOutputError
	}
	if log.Format == nil {
		return NilFormatError
	}
	return nil
}

func (log *logger) Log(level LogLevel, msg string, props ...Prop) error {
	if level > log.MaxLevel {
		return nil
	}
	logProps := NewProps(props...)
	out := log.Format.FormatAndNormalize(level, msg, logProps)
	_, err := log.Output.Write([]byte(out))
	logProps.Return()
	return err
}

func (log *logger) Write(msg []byte) (n int, err error) {
	out := log.Format.FormatAndNormalize(log.DefaultLevel, string(msg), NewProps())
	return log.Output.Write([]byte(out))
}
