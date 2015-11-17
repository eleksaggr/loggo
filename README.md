# loggo
A simple logging framework for the Go language.

## Usage
To create a new Logger:
```go
l := loggo.NewLog(os.Stderr)
```
### Info
```go
l.Info("This is an info.")
```
The output will be:
![Example output of an info command](/examples/info.png)

### Debug
```go
l.Debug("This is a debug message.")
```
The output will be:
![Example output of an debug command](/examples/debug.png)

### Warning
```go
l.Warning("This is a warning.")
```
The output will be:
![Example output of an warning command](/examples/warning.png)

### Error
```go
l.Error("This is an error.")
```
The output will be:
![Example output of an error command](/examples/error.png)
