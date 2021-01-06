# ReadFule() - Reads the file whole at a time

Reads the specified file by loading into memory. On success, returns `nil` and byte array, on failure return `EOF`.

## Function

```go
func ReadFile(string path) ([]byte, error)
```

* [Code](https://golang.org/src/io/ioutil/ioutil.go?s=1503:1549#L42)

## Examples

```
go run readFile.go ./readTest.txt
```

![readfile](img/readfile.png)
