# shapes

This is a simple package to show how to use interfaces in go and packages.

## Requirements

### dep

[Install](https://golang.github.io/dep/docs/installation.html)

### Gopkg.toml

```toml
# Gopkg.toml example
#
# Refer to https://golang.github.io/dep/docs/Gopkg.toml.html
# for detailed Gopkg.toml documentation.

required = ["github.com/casonadams/shapes"]
```

### pull down package with dep

```bash
dep ensure
```

## Example

```go
package main

import (
    "fmt"

    shape "github.com/casonadams/shapes"
)

func main() {
    r := shape.NewRectangle(10, 5)
    t := shape.NewTriangle(2, 4)

    fmt.Println("Area of Rectangle: ", shape.ShapeArea(r))
    fmt.Println("Area of Triangle: ", shape.ShapeArea(t))

}
```

## Tests

### Package Tests

```bash
go test -timeout 30s github.com/casonadams/shapes
```

### Package Benchmark Tests

```bash
go test -benchmem -run=^$ github.com/casonadams/shapes -bench .
```
