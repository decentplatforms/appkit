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
	"fmt"
	"strings"
	"sync"
)

type Prop struct {
	Name  string
	Value any
}

// String returns a prop whose value is a string.
// T may be any type that has an underlying type of string.
//
// If you need to use a fmt.Stringer, use tracy.Stringer.
func String[T ~string](name string, value T) Prop {
	return Prop{
		Name:  name,
		Value: value,
	}
}

// Stringer returns a prop whose value is a string.
// T may be any type that implements fmt.Stringer.
func Stringer[T fmt.Stringer](name string, value T) Prop {
	return Prop{
		Name:  name,
		Value: value.String(),
	}
}

// Int returns a prop whose value is an int.
// T may be any type that has an underlying type of int, int64, uint, or uint64.
// This converts the value to an int64.
func Int[T ~int | ~int64](name string, value T) Prop {
	return Prop{
		Name:  name,
		Value: value,
	}
}

// UInt returns a prop whose value is an unsigned int.
// T may be any type that has an underlying type of uint, or uint64.
// This converts the value to a uint64.
func UInt[T ~uint | ~uint64](name string, value T) Prop {
	return Prop{
		Name:  name,
		Value: value,
	}
}

// Float returns a prop whose value is a float.
// T may be any type that has an underlying type of float64.
func Float[T ~float64](name string, value T) Prop {
	return Prop{
		Name:  name,
		Value: value,
	}
}

// Bool returns a prop whose value is a bool.
// T may be any type that has an underlying type of bool.
func Bool[T ~bool](name string, value T) Prop {
	return Prop{
		Name:  name,
		Value: value,
	}
}

// Props are an ordered collection of log properties.
type Props struct {
	props []Prop
	hash  map[string]int
}

// The propsPool is used to manage log properties in async contexts.
// It keeps memory usage lower when there's a high volume of log messages.
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

// Get returns a named log property in the calling props.
// If there's no matching property, returns nil instead.
func (props *Props) Get(name string) any {
	if idx, ok := props.hash[name]; ok && idx < len(props.props) {
		return props.props[idx].Value
	}
	return nil
}

// GetString gets a named string prop with default value def.
func GetString[T ~string](props *Props, name string, def T) T {
	if v, ok := props.Get(name).(T); ok {
		return v
	}
	return def
}

// GetInt gets a named int prop with default value def.
func GetInt[T ~int | ~int64](props *Props, name string, def T) T {
	if v, ok := props.Get(name).(T); ok {
		return v
	}
	return def
}

// GetUInt gets a named uint prop with default value def.
func GetUInt[T ~uint | ~uint64](props *Props, name string, def T) T {
	if v, ok := props.Get(name).(T); ok {
		return v
	}
	return def
}

// GetFloat gets a named float prop with default value def.
func GetFloat[T ~float64](props *Props, name string, def T) T {
	if v, ok := props.Get(name).(T); ok {
		return v
	}
	return def
}

// GetBool gets a named bool prop with default value def.
func GetBool[T ~bool](props *Props, name string, def T) T {
	if v, ok := props.Get(name).(T); ok {
		return v
	}
	return def
}

func (props *Props) Set(prop Prop) {
	if idx, ok := props.hash[prop.Name]; ok {
		props.props[idx].Value = prop.Value
		return
	}
	props.hash[prop.Name] = len(props.props)
	props.props = append(props.props, prop)
}

// Delete removes the specified keys from props.
// This doesn't clear the data from memory, but removes its hash value so that it cannot be accessed
// through Get/Map.
func (props *Props) Delete(propnames ...string) {
	for _, name := range propnames {
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
