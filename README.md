# ptrace-instruction-counter

It's all in the name.

This program runs an executable file in another process and uses `ptrace` to
count how many instruction were executed.

# Build

```
$ go build -o ptrace-instruction-counter cmd/ptrace-instruction-counter/main.go
```

# Usage

```
$ ptrace-instruction-counter elffile [args...]
```
