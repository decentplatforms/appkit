package logf

import (
	"fmt"
	"time"

	"github.com/decentplatforms/appkit/tracy"
)

type KVConfig struct {
	TimeFormat      string
	UseSingleQuotes bool
}

func (conf KVConfig) withDefaults() KVConfig {
	if conf.TimeFormat == "" {
		conf.TimeFormat = time.RFC3339
	}

	return conf
}

func formatProps(props *tracy.Props, useSingleQuotes bool) string {
	if props == nil {
		return ""
	}

	propsIter := props.Slice()

	raw := ""

	for _, prop := range propsIter {
		// reflect check for int,floats,uint,bool
		switch prop.Value.(type) {
		case int, float64, uint, bool:
			raw += fmt.Sprintf("%s=%v ", prop.Name, prop.Value)
		default:
			if useSingleQuotes {
				raw += fmt.Sprintf("%s='%v' ", prop.Name, prop.Value)
			} else {
				raw += fmt.Sprintf("%s=\"%v\" ", prop.Name, prop.Value)
			}
		}
	}

	if raw == "" {
		return ""
	}

	return raw
}

func KVFormat(conf KVConfig) tracy.Formatter {
	return func(level tracy.LogLevel, msg string, props *tracy.Props) string {

		timestamp := time.Now().UTC().Format(conf.TimeFormat)

		formattedProps := formatProps(props, conf.UseSingleQuotes)

		if conf.UseSingleQuotes {
			return fmt.Sprintf("level=%d timestamp=%s message='%s' %s", level, timestamp, msg, formattedProps)
		} else {
			return fmt.Sprintf("level=%d timestamp=%s message=\"%s\" %s", level, timestamp, msg, formattedProps)
		}
	}
}
