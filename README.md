# logsimple

### Description

Simple log for Go with methods like Info, Warn, Error, Fatal

### Isntructions

Create a new logger, change the DateFormat or LogFormat (optional)

Default values for both configurations:
* DateFormat : "2006-01-02T15:04:05.000Z"
* LogFormat : LogFormatString ([%s] [%s] [%d] [%s:%v] [%s])

Example with text format:

`
var log logsimple.Logger
log.DateFormat = "2006-01-02 15:04:05.000 UTC"
log.LogFormat = LogFormatString
log.Init()
log.Info("Hey I'm logging info to my app!")
`

Output:

`
[2020-02-29T18:41:39.672Z] [INFO] [github.com/fernhtls/logsimple.TestLogSimple] [/home/fernando/go/src/github.com/fernhtls/logsimple/logsimple_test.go:32] [Hey I'm logging info to my app!]
`

Example with JSON format:

`
var log logsimple.Logger
log.DateFormat = "2006-01-02 15:04:05.000 UTC"
log.LogFormat = LogFormatJSON
log.Init()
log.Info("Hey I'm logging info to my app!")
`

Output:

`
{"timestamp":"2020-02-29T18:42:46.390Z","level":"INFO","function":"github.com/fernhtls/logsimple.TestLogSimple","file":"/home/fernando/go/src/github.com/fernhtls/logsimple/logsimple_test.go:22","msg":"Hey I'm logging info to my app!"}
`
