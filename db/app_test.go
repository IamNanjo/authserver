package db

import (
	"slices"
	"testing"
)

var testAppId string

func TestCreateApp(t *testing.T) {
	var err error
	testAppId, err = CreateApp("TestApp", "")

	_, err = GetAppById(testAppId)
	if err != nil {
		t.Fatal(err)
	}

	apps, err := GetApps()
	if err != nil {
		t.Fatal(err)
	}

	newAppFound := slices.ContainsFunc(apps, func(app App) bool {
		return app.Id == testAppId
	})

	if !newAppFound {
		t.Fatal("New app not found with GetApps")
	}
}
