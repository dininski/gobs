package gobs

type gobs struct {
    Settings *settings
    Data data
}

func New(APIKey string) *gobs {
    gobsSettings := getSettings(APIKey)
    data := data{settings: &gobsSettings}
    instance := gobs{Settings: &gobsSettings, Data: data}
    return &instance
}