package gobs

import (
    "net/http"
    "fmt"
    "io/ioutil"
    "encoding/json"
)

type gobsError struct {
    ErrorCode int `json:"errorCode"`
    ErrorMessage string `json:"message"`
    OriginalResponse http.Response
}

func (e *gobsError) Error() string {
    return fmt.Sprintf("Telerik Backend Services request error: %s - error code: %d.", e.ErrorMessage, e.ErrorCode)
}

func gobsErrorFromHttpRequest(response *http.Response) *gobsError {
    byteBody, _ := ioutil.ReadAll(response.Body)
    response.Body.Close()
    err := gobsError{}
    json.Unmarshal(byteBody, &err)
    err.OriginalResponse = *response
    return &err
}