# tracy

Tracy is a basic application logging tool. It covers common log formats, such as syslog and JSON, and can be used through a custom tracy.Log API or as a raw io.Writer.

## Installation

`go get github.com/decentplatforms/appkit/tracy`

## Usage

All examples below log messages using the RFC 3164 syslog format.

### Basic

```go
log := loggers.Syslog3164("tag", true, tracy.Warning, tracy.Informational, os.Stdout)
tracy.Use(log)
...
err := someOperation()
if err != nil {
    tracy.Log(tracy.Error, "something bad is happening!", tracy.Prop{Name: "detail", Value: err})
}
```

### Advanced

```go
format := formats.Syslog3164Format(SyslogConfig{
    AppName: "app",
    Tag: "tag",
})
output := os.Stdout
conf := tracy.Config{
    MaxLevel: tracy.Warning,
    Format: format,
    Output: output,
}
log, err := tracy.NewLogger(conf)
if err != nil {
    panic(err)
}
tracy.Use(log)
...
```

### As Writer

Here, someOperation takes an `errorLogger io.Writer` as input to write errors to. We'll make a writer that defaults to `tracy.Error`.

```go
errorLogger := loggers.Syslog3164("log-test", true, tracy.Warning, tracy.Error, os.Stdout)
err := someOperation(errorLogger)
```

This provides a little less control over log content but makes it easier to integrate with systems that expect an io.Writer instead of a tracy.Logger.

## Contributing

See the root CONTRIBUTING.md file in `github.com/decentplatforms/appkit`.
