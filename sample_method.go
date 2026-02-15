package main

import "fmt"

type Drone struct {
    ID string
    Latitude float64
    Longitude float64
}

func (d *Drone) Move(lat, lng float64) {
    d.Latitude = lat
    d.Longitude = lng
}

func (d Drone) Print() {
    fmt.Println("Drone:", d.ID, "Latitude:", d.Latitude, "Longitude: ", d.Longitude)
}

func main() {
    d := Drone{ID:"01"}
    d.Move(31, 32)
    d.Print()
}
