package gobs

import (
    "testing"
    "os"
    "encoding/json"
)

type testConfig struct {
    APIKey string
}

type customType struct {
    DataItem `contentType:"TestType"`
    SomeValue int
}

var testInstance *gobs
var configuration *testConfig

func TestMain(m *testing.M) {
    configuration = &testConfig{}
    if err := populateConfigurationFromFile(); err != nil {
        populateConfigurationFromEnvironment()
    }

    testInstance = New(configuration.APIKey)
    testInstance.Data.Remove(&customType{})
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
}