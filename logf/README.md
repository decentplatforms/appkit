# logf

`appkit/logf` is a tool for formatted, deterministic application logging.

- Formatted: logf provides extensible formatting tools, and supports structured (kvp, json) and non-structured (both standard syslog RFCs) out-of-the-box
- Deterministic: The same input properties always result in the same log message output

## Installation

`go get github.com/decentplatforms/appkit/logf`

## Usage

All examples below log messages using the RFC 3164 syslog format.

Create a logger using `package loggers` (arguments vary by format):

```go
syslogTag := "application"
useTimeDetail := true
maxLevel := logf.Warning
defaultLevel := logf.Informational
output := os.Stdout
log := loggers.Syslog3164(syslogTag, useTimeDetail, maxLevel, defaultLevel, output)
```

or by using a custom configuration:

```go
format := formats.Syslog3164Format(SyslogConfig{
    AppName: "app",
    Tag: "tag",
})
output := os.Stdout
conf := logf.Config{
    MaxLevel: logf.Warning,
    Format: format,
    Output: output,
}
log, err := logf.NewLogger(conf)
if err != nil {
    panic(err)
}
```

Log messages using `Logger.Log`:

```go
// Log with only a message...
log.Log(logf.Informational, "test log message!")
// ...or with additional properties.
log.Log(logf.Informational, "test log message!", tracy.String("detail", "extra detail here"))
```

Treatment of additional properties depends on the format.

## As Writer

Loggers are an `io.Writer`, so you can `Logger.Write(msg []byte)` to write the message with the logger's default level.

```go
log.Write([]byte("test log message!"))
```

## Contributing

See the root CONTRIBUTING.md file in `github.com/decentplatforms/appkit`.
