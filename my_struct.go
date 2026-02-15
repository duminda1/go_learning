package main

import "fmt"

type Drone struct {
    ID string
    Latitude float64
    Longitude float64
}

func main() {
    drone := Drone {ID : "drone1", Latitude: -31.04, Longitude: 151.1}

    fmt.Println("Drone", drone.ID, "at", drone.Latitude, drone.Longitude)
}
