package gobs

import (
    "encoding/json"
    "net/http"
    "io/ioutil"
    "errors"
    "reflect"
    "fmt"
    "bytes"
)

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

type countResult struct {
    Result int
}

func (d data) count(contentType string) (count int, err error) {
    countUrl := fmt.Sprintf("%s/_count", d.getDataUrl(contentType))
    byteData, err := d.readRequest(countUrl, http.StatusOK)
    if err != nil {
        return 0, err
    }

    result := countResult{}
    if parseForCount(byteData, &result); err != nil {
        return 0, err
    }

    return result.Result, nil
}

func (d data) create(contentType string, dataObject interface{}) error {
    dataUrl := d.getDataUrl(contentType)
    return d.createRequest(dataUrl, dataObject, http.StatusCreated)
}

func (d data) createRequest(createUrl string, dataObject interface{}, expectedStatusCode int) error {
    byteBody, err := json.Marshal(dataObject)
    if err != nil {
        return err
    }

    req, err := http.NewRequest("POST", createUrl, bytes.NewReader(byteBody))
    if err != nil {
        return err
    }

    req.Header.Add("Content-type", "application/json")

    response, error := client.Do(req)
    if error != nil {
        return error
    }

    if response.StatusCode != expectedStatusCode {
        gobsError := gobsErrorFromHttpRequest(response)
        return gobsError
    }

    _, requestError := getResponseBody(response)
    if requestError != nil {
        return requestError
    }

    return nil
}

func (d data) readOne(contentType string, id string, dataObject interface{}) error {
    dataUrl := d.getDataUrlWithId(contentType, id)
    byteData, err := d.readRequest(dataUrl, http.StatusOK)
    if err != nil {
        return err
    }

    parseForSingle(byteData, dataObject)
    return nil
}

func (d data) readMany(contentType string, dataObject interface{}) error {
    dataUrl := d.getDataUrl(contentType)
    byteData, err := d.readRequest(dataUrl, http.StatusOK)
    if err != nil {
        return err
    }

    err = parseForMultiple(byteData, dataObject)
    return nil
}

func (d data) removeMany(contentType string, dataitem interface{}) error {
    dataUrl := d.getDataUrl(contentType)
    _, err := d.removeRequest(dataUrl, http.StatusOK)
    if err != nil {
        return err
    }

    return nil
}

func (d data) removeRequest(url string, expectedCode int) (body []byte, err error) {
    req, err := http.NewRequest("DELETE", url, nil)
    if err != nil {
        return nil, err
    }

    d.applyFiltering(req)
    response, error := client.Do(req)
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

func parseForMultiple(bytedata []byte, dataObject interface{}) error {
    multi := multipleResult{Result: dataObject}
    return json.Unmarshal(bytedata, &multi)
}

type multipleResult struct {
    Result interface{}
    Count int
}

var client *http.Client = &http.Client{}

func (d data) readRequest(url string, expectedCode int) (byteData []byte, err error) {
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }

    d.applyFiltering(req)
    response, error := client.Do(req)
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

type jsonObject struct {
    key string
    value string
}

func (d data) applyFiltering(req *http.Request) {
    filterMap := make(map[string]interface{})
    for _, filter := range d.filters {
        filterMap[filter.field] = filter.value
    }

    filterObject, _ := json.Marshal(filterMap)
    req.Header.Add("X-Everlive-Filter", string(filterObject))
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