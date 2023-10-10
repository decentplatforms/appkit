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

import "errors"

var NilOutputError = errors.New("loggers must have a non-nil output")
var NilFormatError = errors.New("loggers must have a format")
var MultiConfigError = errors.New("can't configure MultiLogger; configure subloggers instead")

var NoActiveLoggerError = errors.New("no active logger")
