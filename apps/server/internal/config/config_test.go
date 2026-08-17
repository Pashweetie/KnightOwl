package config

import "testing"

func TestLoadRejectsMissingFile(t *testing.T) {
	if _, err := Load("does-not-exist.json"); err == nil {
		t.Fatal("Load accepted a missing file")
	}
}
