package configuration

import (
	"reflect"
	"testing"
	"time"
)

func TestLoadMixedConfig(t *testing.T) {
	setOSFlagsForTesting([]string{
		"splendid",
		"-c=../test.conf",
		"--debug",
	})

	// Mixture of Default values, File Values and Flag Values
	config := loadConfig()

	// Grab defaults as base values.
	expect := getConfigDefaults()
	// Expected File Values
	expect.DefaultUser = "splendid"
	expect.DefaultPass = "splendid"
	expect.Timeout = 1337 * time.Second
	expect.Concurrency = 4
	// Expected Flag Values
	expect.Debug = true
	expect.ConfigFile = "../test.conf"
	// Device
	expect.Devices = []DeviceConfig{
		{"localhost", "pfsense", "pfuser", "pfpass", 22, 0, 0},
	}
	if !reflect.DeepEqual(config, expect) {
		t.Fatalf("Loaded config not as expected.\nFound: %v\nExpected: %v", config, expect)
	}
	if config.ConfigFile != "../test.conf" {
		t.Fatal("ConfigFile value not as expected.")
	}
}

// TestLoadFlagRevertToDefault sets an OS Flag to set concurrency to a value of 30.
// This is testing the default value of 30, being overwritten by the test.conf
// config file value of 4, and then again overwritten by flags back to 30 which is the default value.
func TestLoadFlagRevertToDefault(t *testing.T) {
	setOSFlagsForTesting([]string{
		"splendid",
		"-c=../test.conf",
		"-p=30",
	})

	config := loadConfig()
	if config.Concurrency != 30 {
		t.Fatalf("Expected [%v] concurrency, but found [%v]", 30, config.Concurrency)
	}
}

// Test simple merge of defaults into an empty config.
// We are now simply pulling getConfigDefaults, not using merge...
//func TestLoadConfigDefaults(t *testing.T) {
//	config := new(Config)
//	defaults := getConfigDefaults()
//
//	// Merge defaults into config.
//	config.mergeConfig(getConfigDefaults())
//
//	if !reflect.DeepEqual(config, &defaults) {
//		t.Fatalf("Loaded config not as expected.\nFound: %v\nExpected: %v", config, &defaults)
//	}
//}

// Tests the default config path behavior.
func TestDefaultConfigPath(t *testing.T) {
	// Test no conf file using the default value.
	setOSFlagsForTesting([]string{
		"splendid",
	})
	parseConfigFlags()
	pathByDefault := configFlagsGetConfigPath()
	defaults := getConfigDefaults()

	if pathByDefault != defaults.ConfigFile {
		t.Fatalf("Error. Expected: %v Found: %v", defaults.ConfigFile, pathByDefault)
	}

	// Test passing a flag for the conf file.
	expectedConf := "pathByFlag.conf"
	setOSFlagsForTesting([]string{
		"splendid",
		"-c=" + expectedConf,
	})
	parseConfigFlags()
	pathByFlag := configFlagsGetConfigPath()

	if pathByFlag != expectedConf {
		t.Fatalf("Error. Expected: %v Found: %v", expectedConf, pathByFlag)
	}
}

// Test simple INI Loading.
func TestLoadConfigFile(t *testing.T) {
	//config := getFileConfig("../test.conf")

	// We expect other values to be the Zero value.
	expect := new(Config)
	expect.DefaultUser = "splendid"
	expect.DefaultPass = "splendid"
	expect.Timeout = 1337 * time.Second
	expect.Concurrency = 4
	expect.Devices = []DeviceConfig{
		{"localhost", "pfsense", "pfuser", "pfpass", 22, 0, 0},
	}

	config := new(Config)
	mergeConfigFile(config, "../test.conf")

	if !reflect.DeepEqual(config, expect) {
		t.Fatalf("Loaded config not as expected.\nFound: %v\nExpected: %v", config, expect)
	}
}

// Test merging flag values into an empty config.
func TestLoadConfigFlags(t *testing.T) {
	setOSFlagsForTesting([]string{
		"splendid",
		"--debug",
		"-c=testLoadConfigFlags.conf",
	})
	parseConfigFlags()

	// Merge flags into config.
	config := new(Config)
	mergeConfigFlags(config)

	if config.ConfigFile != "testLoadConfigFlags.conf" {
		t.Fatal("Config did not pick up `testLoadConfigFlags.conf` flag value.")
	}
	if config.Debug != true {
		t.Fatal("Config did not pick up `true` debug flag value.")
	}
}
