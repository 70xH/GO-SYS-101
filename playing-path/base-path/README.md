# Base() - Get base path

Returns the base (last) element of path. If the path is completely with separators, it returns a single separator.

## Function

```go
func Base(dir path) string
```

* [Code](https://golang.org/src/path/filepath/path.go?s=13035:13064#L424)

## Examples

```
go run abs.go
```

![noargs](img/noargs.png)

```
go run abs.go ~/70xH-Git/
```

![withargs](img/withargs.png)
