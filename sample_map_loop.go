package main

import "fmt"

func main() {
    drones := make(map[string]float64)
    drones["drone1"] = -33.9
    drones["drone2"] = -34.1

    for k,v := range drones {
        fmt.Println(k,v)
    }
}
