package gobs

import (
    "testing"
//    "fmt"
//    "strconv"
)

type customType struct  {
    DataItem `contentType:"TestType"`
}

func TestGetById(t *testing.T) {
    item := customType{}
    error := testInstance.Data.FindById(configuration.TypeId, &item)
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
    if error := testInstance.Data.Find(&items); error != nil {
        t.Errorf(error.Error())
    }

    if len(items) == 0 {
        t.Errorf("Incorrect number of elements")
    }
}

func TestGetWithSimpleFilter(t *testing.T) {
    items := []customType{}
    // should be additionally expanded as data.where.and.where.or.where.find....
    if error := testInstance.Data.Where("Id", configuration.TypeId).Find(&items); error != nil {
        t.Errorf(error.Error())
        return
    }
    
    if len(items) != 1 {
        t.Errorf("Incorrect number of elements")
    }
}

func TestCreateSingleItem(t *testing.T) {
    customItem := customType{}
    itemCountBeforeCreate, err := testInstance.Data.Count(&customItem);
    if err!= nil {
        t.Errorf(err.Error())
        return
    }

    if err := testInstance.Data.Create(&customItem); err != nil {
        t.Errorf(err.Error())
        return
    }
    
    itemsAfterCreate := []customType{}
    countAfterCreate, err := testInstance.Data.Count(&itemsAfterCreate);
    if err != nil {
        t.Errorf(err.Error())
        return
    }
    
    if itemCountBeforeCreate == countAfterCreate {
        t.Errorf("Item was not created succcessfully")
    }
}

func TestItemCount(t *testing.T) {
    customItem := customType{}
    count, err := testInstance.Data.Count(&customItem);
    if err != nil {
        t.Errorf(err.Error())
        return
    }
    
    if count == 0 {
        t.Errorf("Count did not return correct number of items")
    }
}

func TestGetWithMoreComplexFilter(t *testing.T) {
    items := []customType{}
    // should be additionally expanded as data.where.and.where.or.where.find....
    if error := testInstance.Data.Where("Id", configuration.TypeId).Where("SomeValue", 2).Find(&items); error != nil {
        t.Errorf(error.Error())
        return
    }

    if len(items) != 1 {
        t.Errorf("Incorrect number of elements")
    }
}