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
	"io"
	"os"
)

type File struct {
	writer io.Writer
	queue  chan []byte
	close  chan struct{}
}

func Open(path string, buffer uint8) (*File, error) {
	f := &File{}
	var err error
	f.writer, err = os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	f.queue = make(chan []byte, buffer)
	f.close = make(chan struct{})
	go f.work()
	return f, nil
}

func (f *File) Write(msg []byte) (n int, err error) {
	f.queue <- msg
	return
}

func (f *File) Close() {
	f.close <- struct{}{}
}

func (f *File) work() {
	for {
		select {
		case msg := <-f.queue:
			f.writer.Write(msg)
		case <-f.close:
			return
		default:
		}
	}
}
