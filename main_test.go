package gobs

import (
    "testing"
    "os"
    "encoding/json"
    "fmt")

type testConfig struct {
    APIKey string
    TypeId string
}

var testInstance *gobs
var configuration *testConfig

func TestMain(m *testing.M) {
    configuration = &testConfig{}
    if err := populateConfigurationFromFile(); err != nil {
        populateConfigurationFromEnvironment()
    }

    fmt.Println(configuration.APIKey)
    fmt.Println(configuration.TypeId)
    testInstance = New(configuration.APIKey)
    os.Exit(m.Run())
}

func populateConfigurationFromFile() error {
    configFile, err := os.Open("tests-config.json")
    if err != nil {
        return err
    }
    
    jsonParser := json.NewDecoder(configFile)
    if err = jsonParser.Decode(&configuration); err != nil {
        return err
    }

    return nil
}

func populateConfigurationFromEnvironment() {
    configuration.APIKey = os.Getenv("APIKey")
    configuration.TypeId = os.Getenv("TypeId")
}