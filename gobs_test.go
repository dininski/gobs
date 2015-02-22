package gobs

import "testing"

func TestAPIKeyCorrectlySet(t *testing.T) {
    sampleKey := "sampleApiKey"
    gobs := New(sampleKey)
    if gobs.Settings.APIKey != sampleKey {
        t.Errorf("The API Key was not set correctly")
    }
}