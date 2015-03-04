package gobs

import (
    "testing"
//    "fmt"
//    "strconv"
)

type customType struct {
    DataItem `contentType:"TestType"`
    SomeValue int
}

func TestGetById(t *testing.T) {
    item := customType{}
    item.Id = "CustomId"

    if err := testInstance.Data.Create(&item); err != nil {
        t.Errorf(err.Error())
        return
    }

    newItem := customType{}
    if err := testInstance.Data.FindById(item.Id, &newItem); err != nil {
        t.Errorf(err.Error())
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
    if err := testInstance.Data.Create(&customItem); err != nil {
        t.Errorf(err.Error())
        return
    }

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
    targetItem := customType{SomeValue: 2}
    additionalItem := customType{}
    items = append(items, targetItem, additionalItem)
    if err := testInstance.Data.Create(&items); err != nil {
        t.Errorf(err.Error())
        return
    }
    
    // should be additionally expanded as data.where.and.where.or.where.find....
    if error := testInstance.Data.Where("Id", configuration.TypeId).Where("SomeValue", 2).Find(&items); error != nil {
        t.Errorf(error.Error())
        return
    }

    if len(items) != 1 {
        t.Errorf("Incorrect number of elements")
    }
}

//func TestDeleteAllItems(t *testing.T) {
//    simpleType := customType{}
//    error := testInstance.Data.Remove(&simpleType)
//    if error != nil {
//        t.Errorf(error.Error())
//        return
//    }
//
//    count, err := testInstance.Data.Count(&simpleType)
//
//    if err != nil {
//        t.Errorf(err.Error())
//        return
//    }
//
//    if count != 0 {
//        t.Errorf("Not all items were deleted")
//    }
//}