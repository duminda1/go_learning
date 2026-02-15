package main

import (
    "fmt"
    "encoding/json"
    )

type DroneUpdate struct {
    ID string `json:"id"`
    Latitude float64 `json:"latitude"`
}

func main() {
    du := DroneUpdate{ID:"D1", Latitude:55.2}
    
    data,err := json.Marshal(du)
    if err != nil {
        fmt.Println("Error in marshalling")
    } else {
        fmt.Println(string(data))
    }

    var droneup DroneUpdate
    json.Unmarshal(data, &droneup)

    fmt.Println(droneup.ID, droneup.Latitude)
}


