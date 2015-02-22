package gobs
import "fmt"
import "net/http"
import "io/ioutil"
import "encoding/json"

type data struct {
    settings *settings
}

func (d data) getDataUrl(id string) string {
    url := fmt.Sprintf(
    "%s:%s/%s/%s/%s/%s",
    d.settings.Scheme,
    d.settings.ApiUrl,
    d.settings.Version,
    d.settings.APIKey,
    "Type1",
    id)

    fmt.Println(url)

    return url
}

func (d data) GetById(id string, dataItem interface{}) (err error) {
    //    st := reflect.TypeOf(*dataItem)
    //    field, ok := st.FieldByName("DataItem")
    //    fmt.Println(field, ok)
    return nil
    //    return d.readOne(id, dataItem)
}

func (d data) Get() (dataItemResult []DataItem, err error) {
    return d.readMany("")
}

func (d data) GetByFilter(filter string) (dataItemResult []DataItem, err error) {
    return d.readMany(filter)
}

func (d data) readOne(id string, dataItem *DataItem) (err error) {
    dataUrl := d.getDataUrl(id)
    byteData, err := readRequest(dataUrl)
    if err != nil {
        return err
    }

    *dataItem, err = parseForSingle(byteData)
    return err
}

func (d data) readMany(filter string) (dataResult []DataItem, err error) {
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
    res := string(bytedata)
    fmt.Println(res)
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