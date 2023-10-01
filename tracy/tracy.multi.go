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

import (
	"errors"
)

type multiLogger struct {
	logs []Logger
}

func (log *multiLogger) Configure(conf Config) error {
	return MultiConfigError
}

func (log *multiLogger) Log(level LogLevel, msg string, props ...Prop) error {
	var err error
	for _, log := range log.logs {
		logErr := log.Log(level, msg, props...)
		errors.Join(err, logErr)
	}
	return err
}

func (log *multiLogger) Write(msg []byte) (n int, err error) {
	for _, log := range log.logs {
		_, logErr := log.Write(msg)
		errors.Join(err, logErr)
	}
	return len(msg), err
}
