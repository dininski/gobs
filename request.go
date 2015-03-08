package gobs

import (
    "encoding/json"
    "net/http"
    "bytes"
)

type request struct {
    method string
    url string
    filter getFilter
    data interface{}
    expectedCode int
}

var client *http.Client = &http.Client{}

func createRequest(url string, dataObject interface{}, filter getFilter, expectedCode int) (byteResponse []byte, err error) {
    request := request{url: url, expectedCode: expectedCode, method: "POST", data: dataObject, filter: filter}
    return processRequest(request)
}

func readRequest(url string, filter getFilter, expectedCode int) (byteResponse []byte, err error) {
    request := request{url: url, expectedCode: expectedCode, method: "GET", data: nil, filter: filter}
    return processRequest(request)
}

func removeRequest(url string, filter getFilter, expectedCode int) (byteResponse []byte, err error) {
    request := request{url: url, expectedCode: expectedCode, method: "DELETE", data: nil, filter: filter}
    return processRequest(request)
}

func updateRequest(url string, updateExpression interface{}, filter getFilter, expectedCode int) (byteResponse []byte, err error) {
    request := request{url: url, expectedCode: expectedCode, method: "PUT", data: updateExpression, filter: filter}
    return processRequest(request)
}

func processRequest(request request) (body []byte, err error) {
    byteBody, err := json.Marshal(request.data)
    if err != nil {
        return nil, err
    }

    req, err := http.NewRequest(request.method, request.url, bytes.NewReader(byteBody))
    if err != nil {
        return nil, err
    }

    if request.filter != nil {
        filterMap := request.filter.getFilter()
        filterObject, _ := json.Marshal(filterMap)
        req.Header.Add("X-Everlive-Filter", string(filterObject))
    }

    req.Header.Add("Content-type", "application/json")

    response, error := client.Do(req)
    if error != nil {
        return nil, error
    }

    if response.StatusCode != request.expectedCode {
        gobsError := gobsErrorFromHttpRequest(response)
        return nil, gobsError
    }

    body, requestError := getResponseBody(response)
    if requestError != nil {
        return nil, requestError
    }

    return body, nil
}