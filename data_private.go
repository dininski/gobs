package gobs

import (
    "encoding/json"
    "net/http"
    "io/ioutil"
)

type singleResult struct {
    Result interface{}
}

type countResult struct {
    Result int
}

func (d data) create(contentType string, dataObject interface{}) error {
    dataUrl := getDataUrl(*d.settings, contentType)
    _, err := createRequest(dataUrl, dataObject, d, http.StatusCreated)
    return err
}

func parseForMultiple(bytedata []byte, dataObject interface{}) error {
    multi := multipleResult{Result: dataObject}
    return json.Unmarshal(bytedata, &multi)
}

type multipleResult struct {
    Result interface{}
    Count int
}

type jsonObject struct {
    key string
    value string
}

func parseForSingle(bytedata []byte, dataObject interface{}) error {
    single := singleResult{Result: dataObject}
    return json.Unmarshal(bytedata, &single)
}

func parseForCount(bytedata []byte, countRes *countResult) error {
    return json.Unmarshal(bytedata, countRes)
}

func getResponseBody(response *http.Response) (body []byte, err error) {
    byteBody, err := ioutil.ReadAll(response.Body)
    response.Body.Close()
    if err != nil {
        return nil, err
    }

    return byteBody, nil
}