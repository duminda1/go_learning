package main

import "fmt"

func main() {
    m := make(map[string]int)

    m["a"] = 10
    m["b"] = 20

    value, ok := m["x"]
    if ok {
        fmt.Println(value)
    } else {
        fmt.Println("not found")
    }

    fmt.Println(m["b"])
}
