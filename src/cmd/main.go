package main

import (
    "gobs"
    "fmt"
)

func main() {
    gobsInstance := gobs.New("dcqW03sJbOs8nn8A")
    fmt.Println(gobsInstance.Settings.APIKey)
    type1 := CustomType{}
    error := gobsInstance.Data.GetById("0d097020-ea5c-11e3-aa8b-d70b1ee81b73", &type1)
    if error != nil {
        panic(error)
    }

    fmt.Println(type1.Id)
}

type CustomType struct {
    gobs.DataItem `type:"Type1"`
    Ivan string
}