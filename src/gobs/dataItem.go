package gobs

import ("time"
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
)

type DataItem struct {
    Id string
    Owner string
    CreatedBy string
    ModifiedBy string
    CreatedAt time.Time
    ModifiedAt time.Time
    contentType string
    gobsInstance gobs
}

func (d DataItem) getDataUrl(id string) string {
    url := fmt.Sprintf(
    "%s:%s/%s/%s/%s/%s",
    d.gobsInstance.Settings.Scheme,
    d.gobsInstance.Settings.ApiUrl,
    d.gobsInstance.Settings.Version,
    d.gobsInstance.Settings.APIKey,
    d.contentType,
    id)

    return url
}

func (d DataItem) Get() (dataItemResult []DataItem, err error) {
    return d.readMany("")
}

func (d DataItem) GetByFilter(filter string) (dataItemResult []DataItem, err error) {
    return d.readMany(filter)
}

func (d DataItem) GetById(id string) (dataItemResult DataItem, err error) {
    return d.readOne(id)
}

func (d DataItem) readOne(id string) (data DataItem, err error) {
    var dataResult = DataItem{}
    dataUrl := d.getDataUrl(id)
    byteData, err := readRequest(dataUrl)
    if err != nil {
        return dataResult, err
    }

    dataResult, err = parseForSingle(byteData)
    if err != nil {
        return dataResult, err
    }

    return dataResult, nil
}

func (d DataItem) readMany(filter string) (dataResult []DataItem, err error) {
    dataUrl := d.getDataUrl("")
    byteData, err := readRequest(dataUrl)
    if err != nil {
        return nil, err
    }

    dataItems, err := parseForMultiple(byteData)
    if err != nil {
        return nil, err
    }

    return dataItems, nil
}

func readRequest(url string) (byteData []byte, err error) {
    response, error := http.Get(url)
    if error != nil {
        return nil, error
    }

    body, requestError := getResponseBody(response)
    if requestError != nil {
        return nil, requestError
    }

    return body, nil
}

func getResponseBody(response *http.Response) (body []byte, err error) {
    byteBody, err := ioutil.ReadAll(response.Body)
    response.Body.Close()
    if err != nil {
        return nil, err
    }

    return byteBody, nil
}

func parseForSingle(bytedata []byte) (data DataItem, err error) {
    single := singleResult{}
    error := json.Unmarshal(bytedata, &single)
    fmt.Println("Single:")
    fmt.Println(single)
    fmt.Println("Single end")
    return single.Result, error
}

func parseForMultiple(bytedata []byte) (data []DataItem, err error) {
    multi := multipleItems{}
    error := json.Unmarshal(bytedata, &multi)
    return multi.Result, error
}

type singleResult struct {
    Result DataItem
}

type multipleItems struct {
    Result []DataItem
    Count int
}