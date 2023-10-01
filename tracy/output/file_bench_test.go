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

package output

import (
	"fmt"
	"testing"

	"github.com/decentplatforms/appkit/tracy"
	"github.com/decentplatforms/appkit/tracy/logf"
)

func BenchmarkFile(b *testing.B) {
	writer, err := Open("../../dat/test.log", 100)
	if err != nil {
		b.Fatal(err)
	}
	conf := tracy.Config{
		MaxLevel:     tracy.Debug,
		DefaultLevel: tracy.Informational,
		Format:       logf.Syslog5424Format(logf.SyslogConfig{Tag: "file-bench"}),
		Output:       writer,
	}
	log, err := tracy.NewLogger(conf)
	if err != nil {
		b.Fatal(err)
	}

	for iter := 0; iter < b.N; iter++ {
		if iter%10000 == 0 {
			writer.Close()
			fn := fmt.Sprintf("../../dat/bench-%d.log", iter)
			writer, _ = Open(fn, 100)
			conf.Output = writer
			log.Configure(conf)
		}
		msg := fmt.Sprintf("test log %d", iter)
		log.Log(tracy.Informational, msg)
	}
}
