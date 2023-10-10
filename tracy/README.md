# logf

Tracy is a basic application logging tool. It covers common log formats, such as syslog and JSON, and can be used through a custom logf.Log API or as a raw io.Writer.

## Installation

`go get github.com/decentplatforms/appkit/logf`

## Usage

All examples below log messages using the RFC 3164 syslog format.

### Basic

```go
log := loggers.Syslog3164("tag", true, logf.Warning, logf.Informational, os.Stdout)
logf.Use(log)
...
err := someOperation()
if err != nil {
    logf.Log(logf.Error, "something bad is happening!", logf.Prop{Name: "detail", Value: err})
}
```

### Advanced

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
logf.Use(log)
...
```

### As Writer

Here, someOperation takes an `errorLogger io.Writer` as input to write errors to. We'll make a writer that defaults to `logf.Error`.

```go
errorLogger := loggers.Syslog3164("log-test", true, logf.Warning, logf.Error, os.Stdout)
err := someOperation(errorLogger)
```

This provides a little less control over log content but makes it easier to integrate with systems that expect an io.Writer instead of a logf.Logger.

## Contributing

See the root CONTRIBUTING.md file in `github.com/decentplatforms/appkit`.
