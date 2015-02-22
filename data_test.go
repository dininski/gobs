package gobs

import (
    "testing"
    "os"
)

type customType struct  {
    DataItem `contentType:"TestType"`
}

func TestGetById(t *testing.T) {
    item := customType{}
    apiKey := os.Getenv("APIKey")
    id := os.Getenv("TypeId")
    instance := New(apiKey)
    error := instance.Data.GetById(id, &item)
    if error != nil {
        t.Errorf(error.Error())
    }
    
    if item.Id != id {
        t.Errorf("Id not set correctly when getting item")
    }
}