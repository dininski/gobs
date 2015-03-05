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