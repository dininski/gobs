package gobs

import (
    "testing"
    "fmt"
)

const singleItemId string = "7c80bc3e-07c7-4145-8e0b-c21c127c2c94"

func TestGetById(t *testing.T) {
    item := customType{}
    item.Id = singleItemId

    if err := testInstance.Data.Create(&item); err != nil {
        t.Errorf(err.Error())
        return
    }

    newItem := customType{}
    if err := testInstance.Data.FindById(item.Id, &newItem); err != nil {
        t.Errorf(err.Error())
        return
    }

    if item.Id != item.Id {
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
    if error := testInstance.Data.Where("Id", singleItemId).Find(&items); error != nil {
        t.Errorf(error.Error())
        return
    }

    if len(items) != 1 {
        t.Errorf("Incorrect number of elements")
    }
}

func TestGetWithNotFilter(t *testing.T) {
    count, err := testInstance.Data.Count(&customType{})
    if err != nil {
        t.Errorf(err.Error())
        return
    }

    itemsWhereNotCount, err := testInstance.Data.WhereNot("Id", singleItemId).Count(&customType{})
    if err != nil {
        t.Errorf(err.Error())
        return
    }

    if itemsWhereNotCount != count - 1 {
        t.Errorf(fmt.Sprintf("Expected where not to exclude the filter. Count result is: %d", itemsWhereNotCount))
        return
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
    if error := testInstance.Data.Where("SomeValue", 2).Find(&items); error != nil {
        t.Errorf(error.Error())
        return
    }

    elementCount := len(items)
    if elementCount != 1 {
        t.Errorf(fmt.Sprintf("Incorrect number of elements, expected %d to be %d", elementCount, 1))
    }
}

func TestDeleteAllItems(t *testing.T) {
    simpleType := customType{}
    error := testInstance.Data.RemoveMany(&simpleType)
    if error != nil {
        t.Errorf(error.Error())
        return
    }

    count, err := testInstance.Data.Count(&simpleType)

    if err != nil {
        t.Errorf(err.Error())
        return
    }

    if count != 0 {
        t.Errorf("Not all items were deleted")
    }
}

func TestUpdateItem(t *testing.T) {
    simpleType := customType{}
    simpleType.Id = singleItemId
    if err := testInstance.Data.RemoveMany(&simpleType); err != nil {
        t.Errorf(err.Error())
        return
    }

    error := testInstance.Data.Create(&simpleType)
    if error != nil {
        t.Errorf(error.Error())
        return
    }

    randomValue := 100
    simpleType.SomeValue = randomValue
    err := testInstance.Data.Update(&simpleType)
    if err != nil {
        t.Errorf(err.Error())
        return
    }

    err = testInstance.Data.FindById(singleItemId, &simpleType)
    if err != nil {
        t.Errorf(err.Error())
        return
    }

    if simpleType.SomeValue != randomValue {
        t.Errorf("Object not updated correctly on the server")
        return
    }
}