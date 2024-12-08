```go
package main

import (

    "fmt"
)

func main() {
    var m map[string]int
    if m == nil {
        fmt.Println("Map is nil")
    } else {
        fmt.Println(m["a"])
    }
}
```