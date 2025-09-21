package db

import (
	"context"
	"slices"
	"testing"
)

var testAppId string

func TestCreateApp(t *testing.T) {
	var err error
	testApp, err := Q().CreateApp(context.Background(), CreateAppParams{Name: "TestApp"})
	if err != nil {
		t.Fatal(err)
	}
	testAppId = testApp.Id

	_, err = Q().GetApp(context.Background(), testAppId)
	if err != nil {
		t.Fatal(err)
	}

	apps, err := Q().GetApps(context.Background())
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
