# Bazinga

some missing tool for golang


## Examples

### Compare

```go
package main

import (
    "fmt"
    "github.com/stornado/bazinga/pkg/compare"
)

func main() {
    max := compare.Max(3,2,5,1,3)
    fmt.Println(max) // 5

    min := compare.Min(3,2,5,1,3)
    fmt.Println(min) // 1
}
```

### URLJoin

```go
package main

import (
    "fmt"
    "github.com/stornado/bazinga/pkg/urlparse"
)

func main() {
    fullURL, _ := urlparse.URLJoin("http://github.com/", "stornado", "bazinga")
	fmt.Println(fullURL) // http://github.com/stornado/bazinga
}
```

## License

[MIT](https://github.com/stornado/bazinga/blob/develop/LICENSE)