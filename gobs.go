package gobs

type gobs struct {
    Settings *settings
    Data data
}

func New(APIKey string) *gobs {
    gobsSettings := getSettings(APIKey)
    filterQuery := filterQuery{settings: &gobsSettings}
    data := data{filterQuery: filterQuery}
    instance := gobs{Settings: &gobsSettings, Data: data}
    return &instance
}