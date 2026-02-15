package main

import (
    "fmt"
    "time"
    )

func sayHello(a int) {
    fmt.Println("Hello", a)
}

func main() {
    go sayHello(1)
    go sayHello(2)
    go sayHello(3)

    time.Sleep(time.Second)
}
