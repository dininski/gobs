package gobs

import (
    "testing"
)

type customType struct  {
    DataItem `contentType:"TestType"`
}

func TestGetById(t *testing.T) {
    item := customType{}
    error := testInstance.Data.GetById(configuration.TypeId, &item)
    if error != nil {
        t.Errorf(error.Error())
        return
    }
    
    if item.Id != configuration.TypeId {
        t.Errorf("Id not set correctly when getting item")
    }
}

func TestGetMultipleNoFilter(t *testing.T) {
    items := []customType{}
    if error := testInstance.Data.Get(&items); error != nil {
        t.Errorf(error.Error())
    }

    if len(items) == 0 {
        t.Errorf("Incorrect number of elements")
    }
}