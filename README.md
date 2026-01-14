# gopsutil-user-test

Quick test for https://github.com/shirou/gopsutil/pull/1990

## Setup:

```
git clone -b fix/darwin-utmpx https://github.com/uubulb/gopsutil.git /tmp/gopsutil-fork
go mod tidy

go run main.go

```


### Results


Linux (for comparison)

```
User:  herbert
Terminal:  pts/1
Host:
Started:  1768374935
```

macOS 26.2 - arm64
```
User:  macbert
Terminal:  ttys000
Host:
Started:  1768375108
```
