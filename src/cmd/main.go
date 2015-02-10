package main

import (
    "gobs"
    "fmt"
)

func main() {
    gobsInstance := gobs.New("dcqW03sJbOs8nn8A")
    fmt.Println(gobsInstance.Settings.APIKey)
    type1 := gobs.DataItem{}
    gobsInstance.Data("Type1", &type1)
    types, error := type1.GetById("0d097020-ea5c-11e3-aa8b-d70b1ee81b73")
    if error != nil {
        panic(error)
    }

    fmt.Println(types.Id)
}

type CustomType struct {
    gobs.DataItem
    Ivan string
}