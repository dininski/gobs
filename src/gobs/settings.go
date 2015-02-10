package gobs

type settings struct {
    APIKey string
    ApiUrl string
    Scheme string
    MasterKey string
    Version string
}

func getSettings(APIKey string) settings{
    settingsInstance := settings{}
    settingsInstance.APIKey = APIKey
    settingsInstance.ApiUrl = "//api.everlive.com"
    settingsInstance.Scheme = "http"
    settingsInstance.Version = "v1"
    return settingsInstance
}