package gobs

type filterQuery struct {
    filters []filter
    settings *settings
}

type filter struct {
    field string
    value interface{}
}

type getFilter interface{
    getFilter() map[string]interface{}
}

func (fq filterQuery) Find(dataObject interface{}) error {
    contentTypeName, err := getContentTypeName(dataObject)

    if err != nil {
        return err
    }

    return fq.readMany(contentTypeName, dataObject)
}

func (fq filterQuery) RemoveMany(dataObject interface{}) error {
    contentTypeName, err := getContentTypeName(dataObject)

    if err != nil {
        return err
    }

    return fq.removeMany(contentTypeName, dataObject)
}

func (fq filterQuery) Where(field string, value interface{}) filterQuery {
    filter := filter{field: field, value: value}
    fq.filters = append(fq.filters, filter)
    return fq
}

func (fq filterQuery) WhereNot(field string, value interface{}) filterQuery {
    operator := "$ne"
    notFilter := make(map[string]interface{})
    notFilter[operator] = value
    notEqualFilter := filter{field: field, value: notFilter}
    fq.filters = append(fq.filters, notEqualFilter)
    return fq
}

func (fq filterQuery) Count(object interface{}) (count int, err error) {
    contentTypeName, err := getContentTypeName(object)

    if err != nil {
        return 0, err
    }

    return fq.count(contentTypeName)
}