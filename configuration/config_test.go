package configuration

import "testing"

func TestMerge(t *testing.T) {
	// Create a basic config.
	c1 := new(Config)
	c1.Debug = true
	c1.ConfigFile = "default.conf"

	// Another config with different values.
	c2 := new(Config)
	c2.Debug = false
	c2.ConfigFile = "replace.conf"

	c1.mergeConfig(*c2)
	if c1.ConfigFile != "replace.conf" {
		t.Fatal("Expected config to be replace.conf after merge.")
	}
	if c1.Debug != false {
		t.Fatal("Expected debug to be false after merge.")
	}

	// Third config to ensure multiple merges continue to function.
	c3 := new(Config)
	c3.Debug = true
	c3.ConfigFile = "third.conf"

	c1.mergeConfig(*c3)
	if c1.ConfigFile != "third.conf" {
		t.Fatal("Expected config to be third.conf after merge.")
	}
	if c1.Debug != true {
		t.Fatal("Expected debug to be false after merge.")
	}
}
