package gobs

import (
    "fmt"
    "reflect"
    "errors"
)

func getDataUrl(settings settings, contentType string) string {
    url := fmt.Sprintf(
    "%s:%s/%s/%s/%s",
    settings.Scheme,
    settings.ApiUrl,
    settings.Version,
    settings.APIKey,
    contentType)

    return url
}

func getDataUrlWithId(settings settings, contentType string, id string) string {
    url := fmt.Sprintf("%s/%s", getDataUrl(settings, contentType), id)
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