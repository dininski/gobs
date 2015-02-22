package gobs

import "testing"

func TestGetById(t *testing.T) {
    item := DataItem{}
    id := "someId"
    instance := New("")
    instance.Data.GetById(id, &item)
    if item.Id != id {
        t.Errorf("Id not set correctly when getting item")
    }
}