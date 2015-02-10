package gobs

type gobs struct {
    Settings settings
}

func New(APIKey string) *gobs {
    gobsSettings := getSettings(APIKey)
    instance := gobs{Settings: gobsSettings}
    return &instance
}

func (g *gobs) Data(ContentType string, data *DataItem) {
    data.gobsInstance = *g
}