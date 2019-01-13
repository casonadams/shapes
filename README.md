# simpleshapes

This is a simple package to show how to use interfaces in go and packages.

## requirements
### dep
[Install](https://golang.github.io/dep/docs/installation.html)

### Gopkg.toml
```toml
# Gopkg.toml example
#
# Refer to https://golang.github.io/dep/docs/Gopkg.toml.html
# for detailed Gopkg.toml documentation.

required = ["github.com/casonadams/simpleshapes"]
```

### pulling down package with dep
```bash
dep ensure
```

### usage example
```go
package main

import (
	"fmt"

	simpleshape "github.com/casonadams/simpleshapes"
)

func main() {
	r := simpleshape.NewRectangle(10, 5)
	t := simpleshape.NewTriangle(2, 4)

	fmt.Println("Area of Rectangle: ", simpleshape.ShapeArea(r))
	fmt.Println("Area of Triangle: ", simpleshape.ShapeArea(t))

}
```
