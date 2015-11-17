# loggo
A simple logging framework for the Go language.

## Usage
To create a new Logger:
```
l := loggo.NewLog(os.Stderr)
```
### Info
```
l.Info("This is an info.")
```
The output will be:
![Example output of an info command](/examples/info.png)

### Debug
```
l.Debug("This is a debug message.")
```
The output will be:
![Example output of an info command](/examples/debug.png)

### Warning
```
l.Warning("This is a warning.")
```
The output will be:
![Example output of an info command](/examples/warning.png)

### Error
```
l.Error("This is an error.")
```
The output will be:
![Example output of an info command](/examples/error.png)
