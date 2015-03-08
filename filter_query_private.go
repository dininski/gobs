package gobs

import (
    "fmt"
    "net/http"
)

func (fq filterQuery) count(contentType string) (count int, err error) {
    countUrl := fmt.Sprintf("%s/_count", getDataUrl(*fq.settings, contentType))
    byteData, err := readRequest(countUrl, fq, http.StatusOK)
    if err != nil {
        return 0, err
    }

    result := countResult{}
    if parseForCount(byteData, &result); err != nil {
        return 0, err
    }

    return result.Result, nil
}

func (d data) readOne(contentType string, id string, dataObject interface{}) error {
    dataUrl := getDataUrlWithId(*d.settings, contentType, id)
    byteData, err := readRequest(dataUrl, d, http.StatusOK)
    if err != nil {
        return err
    }

    return parseForSingle(byteData, dataObject)
}

func (fq filterQuery) readMany(contentType string, dataObject interface{}) error {
    dataUrl := getDataUrl(*fq.settings, contentType)
    byteData, err := readRequest(dataUrl, fq, http.StatusOK)
    if err != nil {
        return err
    }

    return parseForMultiple(byteData, dataObject)
}

func (fq filterQuery) removeMany(contentType string, dataitem interface{}) error {
    dataUrl := getDataUrl(*fq.settings, contentType)
    _, err := removeRequest(dataUrl, fq, http.StatusOK)

    return err
}

func (fq filterQuery) getFilter() map[string]interface{} {
    filterMap := make(map[string]interface{})
    for _, filter := range fq.filters {
        filterMap[filter.field] = filter.value
    }

    return filterMap
}