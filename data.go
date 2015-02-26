package gobs

func (d data) FindById(id string, dataObject interface{}) (err error) {
    contentTypeName, error := getContentTypeName(dataObject)
    if error != nil {
        return error
    }

    return d.readOne(contentTypeName, id, dataObject)
}

func (d data) Find(dataObject interface{}) error {
    contentTypeName, err := getContentTypeName(dataObject)

    if err != nil {
        return err
    }

    d.readMany(contentTypeName, dataObject)
    return nil
}

func (d data) Where(field string, value interface{}) data {
    filter := filter{field: field, value:value}
    d.filters = append(d.filters, filter)
    return d
}

type filter struct {
    field string
    value interface{}
}

type data struct {
    settings *settings
    filters []filter
}