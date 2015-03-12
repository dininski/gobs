package gobs

const (
    MasterKeyAuthentication = "masterKey"
    BearerAuthentication = "bearer"
    None = "none"
)

type gobs struct {
    Settings *settings
    Data data
    Users userData
    AccessToken string
    TokenType string
}

func New(APIKey string) *gobs {
    gobsSettings := getSettings(APIKey)
    filterQuery := filterQuery{settings: &gobsSettings}
    data := data{filterQuery: filterQuery}
    instance := gobs{Settings: &gobsSettings, Data: data}
    return &instance
}