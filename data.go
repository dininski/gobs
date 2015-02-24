package gobs

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "reflect"
    "errors"
)

func (d data) GetById(id string, dataObject interface{}) (err error) {
    contentTypeName, error := getContentTypeName(dataObject)
    if error != nil {
        return error
    }

    return d.readOne(contentTypeName, id, dataObject)
}

func (d data) Get(dataObject interface{}) error {
    contentTypeName, err := getContentTypeName(dataObject)

    if err != nil {
        return err
    }

    d.readMany(contentTypeName, dataObject)
    return nil
}


type data struct {
    settings *settings
}

func (d data) getDataUrl(contentType string) string {
    url := fmt.Sprintf(
    "%s:%s/%s/%s/%s",
    d.settings.Scheme,
    d.settings.ApiUrl,
    d.settings.Version,
    d.settings.APIKey,
    contentType)

    return url
}

func (d data) getDataUrlWithId(contentType string, id string) string {
    url := fmt.Sprintf("%s/%s", d.getDataUrl(contentType), id)
    fmt.Println(url)
    return url
}

func getContentTypeName(item interface{}) (string, error) {
    itemType := reflect.TypeOf(item).Elem()
    itemTypeKind := itemType.Kind()
    switch itemTypeKind {
        case reflect.Slice:
        structField, ok := itemType.Elem().FieldByName("DataItem")
        if !ok {
            return "", errors.New("Could not obtain content type from name")
        }
        
        return structField.Tag.Get("contentType"), nil
        default:
        structField, ok := itemType.FieldByName("DataItem")

        if !ok {
            return "", errors.New("Could not obtain content type from name")
        }

        return structField.Tag.Get("contentType"), nil
    }
}

type singleResult struct {
    Result interface{}
}

func (d data) readOne(contentType string, id string, dataObject interface{}) error {
    dataUrl := d.getDataUrlWithId(contentType, id)
    byteData, err := readRequest(dataUrl, http.StatusOK)
    if err != nil {
        return err
    }

    parseForSingle(byteData, dataObject)
    return err
}

func (d data) readMany(contentType string, dataObject interface{}) error {
    dataUrl := d.getDataUrl(contentType)
    byteData, err := readRequest(dataUrl, http.StatusOK)
    if err != nil {
        return err
    }

    err = parseForMultiple(byteData, dataObject)
    return err
}

func parseForMultiple(bytedata []byte, dataObject interface{}) error {
    multi := multipleResult{Result: dataObject}
    error := json.Unmarshal(bytedata, &multi)
    return error
}


type multipleResult struct {
    Result interface{}
    Count int
}

func readRequest(url string, expectedCode int) (byteData []byte, err error) {
    response, error := http.Get(url)
    if error != nil {
        return nil, error
    }

    if response.StatusCode != expectedCode {
        gobsError := gobsErrorFromHttpRequest(response)
        return nil, gobsError
    }

    body, requestError := getResponseBody(response)
    if requestError != nil {
        return nil, requestError
    }

    return body, nil
}

func parseForSingle(bytedata []byte, dataObject interface{}) error {
    single := singleResult{Result: dataObject}
    error := json.Unmarshal(bytedata, &single)
    res := string(bytedata)
    fmt.Println(res)
    return error
}

func getResponseBody(response *http.Response) (body []byte, err error) {
    byteBody, err := ioutil.ReadAll(response.Body)
    response.Body.Close()
    if err != nil {
        return nil, err
    }

    return byteBody, nil
}

//
//func (d data) GetByFilter(filter string) (dataItemResult []DataItem, err error) {
//    return d.readMany(filter)
//}
//
//
//
//
//
//