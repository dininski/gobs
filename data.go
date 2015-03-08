package gobs

type data struct {
    filterQuery
}

func (d data) Create(object interface{}) error {
    contentTypeName, err := getContentTypeName(object)
    if err != nil {
        return err
    }

    return d.create(contentTypeName, object)
}

func (d data) Update(object interface{}) error {
    contentTypeName, err := getContentTypeName(object)

    if err != nil {
        return err
    }

    return d.updateSingle(contentTypeName, object)
}


func (d data) FindById(id string, dataObject interface{}) error {
    contentTypeName, error := getContentTypeName(dataObject)
    if error != nil {
        return error
    }

    return d.readOne(contentTypeName, id, dataObject)
}