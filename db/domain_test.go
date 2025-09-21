package db

import (
	"context"
	"slices"
	"testing"
)

func TestGetAppDomains(t *testing.T) {
	_, err := Q().CreateDomain(context.Background(), CreateDomainParams{App: testAppId, Name: "local.test"})
	if err != nil {
		t.Fatalf("Domain creation failed. %v", err)
	}

	domains, err := Q().GetAppDomains(context.Background(), testAppId)
	if err != nil {
		t.Fatalf("Could not get app domains. %v", err)
	}

	newDomainFound := slices.ContainsFunc(domains, func(domain Domain) bool {
		return domain.App == testAppId
	})

	if !newDomainFound {
		t.Fatal("Could not find created app Domain.")
	}
}
