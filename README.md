# loggo
A simple logging framework for the Go language.

## Usage
To create a new Logger:
```
l := loggo.NewLog(os.Stderr)
```
```
l.Info("This is an info.")
```
```
l.Debug("This is a debug message.")
```
```
l.Warning("This is a warning.")
```
```
l.Error("This is an error.")
```
