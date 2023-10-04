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
	"strings"
	"sync"
)

type Prop struct {
	Name  string
	Value any
}

type Props struct {
	props []Prop
	hash  map[string]int
}

var propsPool = &sync.Pool{
	New: func() any {
		return &Props{
			props: make([]Prop, 0),
			hash:  make(map[string]int, 0),
		}
	},
}

func getProps(ct int) *Props {
	newProps := propsPool.Get().(*Props)
	diff := ct - len(newProps.props)
	if diff > 0 {
		newProps.props = append(newProps.props, make([]Prop, diff)...)
	} else if diff < 0 {
		newProps.props = newProps.props[:-diff]
	}
	return newProps
}

// *Props returns *Props with all provided Prop objects.
func NewProps(props ...Prop) *Props {
	newProps := getProps(len(props))
	for i, prop := range props {
		newProps.props[i] = prop
		newProps.hash[prop.Name] = i
	}
	return newProps
}

func (props *Props) Get(name string) any {
	if idx, ok := props.hash[name]; ok && idx < len(props.props) {
		return props.props[idx].Value
	}
	return nil
}

func (props *Props) Set(name string, value any) {
	if idx, ok := props.hash[name]; ok {
		props.props[idx].Value = value
		return
	}
	props.hash[name] = len(props.props)
	props.props = append(props.props, Prop{name, value})
}

// Delete removes the specified keys from props.
// This doesn't clear the data from memory, but removes its hash value so that it cannot be accessed
// through Get/Map.
func (props *Props) Delete(propNames ...string) {
	for _, name := range propNames {
		delete(props.hash, name)
	}
}

// Map returns a map of key-value pairs.
func (props *Props) Map() map[string]any {
	propsMap := make(map[string]any, len(props.props))
	for _, idx := range props.hash {
		prop := props.props[idx]
		propsMap[prop.Name] = prop.Value
	}
	return propsMap
}

func (props *Props) Return() {
	clear(props.hash)
	propsPool.Put(props)
}

// Formatter defines how Logger.Log and logger.Write output messages.
// When using Logger.Log, the included props will be passed through, but they are not
// included when using Logger as an io.Writer.
//
// Do not call formatters directly. Use Formatter.FormatAndNormalize; it normalizes whitespace/newlines
// for you so you don't have to worry about it in your formatter.
//
// By default, log uses the RFC5424 syslog format.
type Formatter func(level LogLevel, msg string, props *Props) string

func (formatter Formatter) FormatAndNormalize(level LogLevel, msg string, props *Props) string {
	out := formatter(level, msg, props)
	out = NormalizeWhitespace(out)
	return out
}

// ===== UTILITIES =====

func NormalizeWhitespace(msg string) string {
	return strings.TrimSpace(msg) + "\n"
}
