package logf

import (
	"errors"
	"fmt"
	"regexp"
	"testing"

	"github.com/decentplatforms/appkit/tracy"
	"github.com/decentplatforms/appkit/tracy/testhelp"
)

var syslog_formats = map[string]tracy.Formatter{
	"syslog_rfc3164": Syslog3164Format(SyslogConfig{
		Hostname: "test-host",
		Tag:      "test-log",
	}),
	"syslog_rfc3164.json": Syslog3164Format(SyslogConfig{
		Hostname:  "test-host",
		Tag:       "test-log",
		WithProps: SyslogJSON,
	}),
	"syslog_rfc3164.timedetail": Syslog3164Format(SyslogConfig{
		Hostname:   "test-host",
		Tag:        "test-log",
		UseISO8601: true,
	}),
	"syslog_rfc3164.ignore": Syslog3164Format(SyslogConfig{
		Hostname:  "test-host",
		Tag:       "test-log",
		WithProps: SyslogIgnore,
	}),
	"syslog_rfc5424": Syslog5424Format(SyslogConfig{
		Hostname: "test-host",
		AppName:  "test-app",
		Tag:      "test-log",
	}),
}

var syslog_regexes = map[string]*regexp.Regexp{
	"syslog_rfc3164":            regexp.MustCompile(`(?P<pri><\d+>)(?P<timestamp>[A-Z][a-z]{2} [ \d]\d \d{2}:\d{2}:\d{2}) (?P<hostname>\S+) (?P<tag>[^:\s]+): (?P<message>.+)`),
	"syslog_rfc3164.timedetail": regexp.MustCompile(`(?P<pri><\d+>)(?P<timestamp>\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}Z) (?P<hostname>\S+) (?P<tag>[^:\s]+): (?P<message>.+)`),
	"syslog_rfc5424":            regexp.MustCompile(`(?P<pri><\d+>)(?P<version>\d+) (?P<timestamp>\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(?:Z|[\+-]\d{2}:\d{2})) (?P<hostname>\S+) (?P<appname>\S+) (?P<pid>\d+) (?P<msgid>\S+) (?P<structured>\S+) (?P<message>.+)`),
}

var syslog_expects = map[string]testhelp.ResultsMap{
	"syslog_rfc3164": {
		"pri":      "<14>",
		"hostname": "test-host",
		"tag":      "test-log",
		"message":  `test log detail="testing format using regex"`,
	},
	"syslog_rfc3164.json": {
		"pri":      "<14>",
		"hostname": "test-host",
		"tag":      "test-log",
		"message":  `test log {"detail":"testing format using regex"}`,
	},
	"syslog_rfc3164.ignore": {
		"pri":      "<14>",
		"hostname": "test-host",
		"tag":      "test-log",
		"message":  `test log`,
	},
	"syslog_rfc5424": {
		"pri":        "<14>",
		"version":    "1",
		"hostname":   "test-host",
		"appname":    "test-app",
		"msgid":      "test-log",
		"structured": "-",
		"message":    `test log detail="testing format using regex"`,
	},
}

var syslog_custom_props = map[string][]tracy.Prop{
	"syslog_rfc3164": {tracy.String(SYSLOG_HOSTNAME, "custom-host"), tracy.String(SYSLOG_TAG, "custom-log")},
	"syslog_rfc5424": {tracy.String(SYSLOG_HOSTNAME, "custom-host"), tracy.String(SYSLOG_APPNAME, "custom-app"), tracy.String(SYSLOG_TAG, "custom-log")},
}

var syslog_custom_expects = map[string]testhelp.ResultsMap{
	"syslog_rfc3164": {
		"pri":      "<14>",
		"hostname": "custom-host",
		"tag":      "custom-log",
		"message":  `test log detail="testing with custom props"`,
	},
	"syslog_rfc3164.json": {
		"pri":      "<14>",
		"hostname": "custom-host",
		"tag":      "custom-log",
		"message":  `test log {"detail":"testing with custom props"}`,
	},
	"syslog_rfc3164.ignore": {
		"pri":      "<14>",
		"hostname": "custom-host",
		"tag":      "custom-log",
		"message":  `test log`,
	},
	"syslog_rfc5424": {
		"pri":        "<14>",
		"version":    "1",
		"hostname":   "custom-host",
		"appname":    "custom-app",
		"msgid":      "custom-log",
		"structured": "-",
		"message":    `test log detail="testing with custom props"`,
	},
}

func results(msg, format string) (map[string]string, error) {
	regex := testhelp.GetTestOption(syslog_regexes, format, nil)
	matches := regex.FindStringSubmatch(msg)
	if matches == nil {
		fmt.Println(msg)
		return nil, errors.New("message didn't match format " + format)
	}
	result := make(map[string]string)
	for i, name := range regex.SubexpNames() {
		result[name] = matches[i]
	}
	return result, nil
}

func TestSyslog(t *testing.T) {
	for name, format := range syslog_formats {
		t.Run(name, func(t *testing.T) {
			tw := &TestWriter{}
			conf := tracy.Config{
				MaxLevel: tracy.Informational,
				Format:   format,
				Output:   tw,
			}
			log, err := tracy.NewLogger(conf)
			if err != nil {
				t.Fatal(err)
			}
			err = log.Log(tracy.Informational, "test log", tracy.String("detail", "testing format using regex"))
			if err != nil {
				t.Fatal(err)
			}
			res, err := results(tw.Last, name)
			if err != nil {
				t.Fatal(err)
			}
			wanted := testhelp.GetTestOption(syslog_expects, name, nil)
			err = testhelp.ValidateResults(res, wanted)
			if err != nil {
				t.Fatal(err)
			}
		})
		t.Run(name+"-custom", func(t *testing.T) {
			tw := &TestWriter{}
			conf := tracy.Config{
				MaxLevel: tracy.Informational,
				Format:   format,
				Output:   tw,
			}
			log, err := tracy.NewLogger(conf)
			if err != nil {
				t.Fatal(err)
			}
			props := append(testhelp.GetTestOption(syslog_custom_props, name, nil), tracy.String("detail", "testing with custom props"))
			err = log.Log(tracy.Informational, "test log", props...)
			if err != nil {
				t.Fatal(err)
			}
			res, err := results(tw.Last, name)
			if err != nil {
				t.Fatal(err)
			}
			wanted := testhelp.GetTestOption(syslog_custom_expects, name, nil)
			err = testhelp.ValidateResults(res, wanted)
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}
