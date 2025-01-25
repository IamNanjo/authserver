package db

import (
	"slices"
	"testing"
)

func TestGetAppDomains(t *testing.T) {
	err := CreateDomain(testAppId, "local.test")
	if err != nil {
		t.Fatalf("Domain creation failed. %v", err)
	}

	domains, err := GetAppDomains(testAppId)
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
