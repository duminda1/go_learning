package main

import "fmt"

func main() {
    s := []float64{}

    s = append(s, 10.5)
    s = append(s, 20.2)
    s = append(s, 30.7)

    sum := 0.0

    for _, v := range s {
        sum += v
    }

    fmt.Println("Sum:", sum)
}
