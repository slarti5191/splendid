package configuration

import (
	"github.com/go-ini/ini"
	"fmt"
)

// LoadConfig loads the saved config file
func loadConfig(configFile string) (*SplendidConfig, error) {
	// Initialize Splendid Config
	conf := new(SplendidConfig)
	conf.ConfigFile = configFile

	// Load the INI file.
	cfg, err := ini.Load(conf.ConfigFile)
	if err != nil {
		return nil, err
	}

	// map config file to SplendidConfig struct
	if err = cfg.Section("main").MapTo(conf); err != nil {
		return nil, fmt.Errorf("error mapping main config: %v", err)
	}

	// TODO: Refactor this out into separate methods.

	/* I think this can go now?
	// Get a reflection Value for the SplendidConfig conf variable.
	sConf := reflect.ValueOf(conf).Elem()

	// Loop through all of the main values, and set them on the conf variable.
	main := cfg.Section("main")
	for _, key := range main.KeyStrings() {
		value := main.Key(key)

		// Get the specific field from the conf that matches the config file.
		field := sConf.FieldByName(key)

		// Use reflection to determine the type, and set the value.
		switch field.Kind() {
		case reflect.Bool:
			field.SetBool(value.MustBool())
		case reflect.String:
			field.SetString(value.String())
		default:
			return nil, errors.New("Unrecognized field type: " + field.Kind().String())
		}
	}
	*/

	// Iterate through all other sections and create DeviceConfigs
	//for a, b := range cfg.Sections() {
	//	fmt.Println("====Sections====")
	//	fmt.Println(a)
	//	fmt.Println(b)
	//	fmt.Println("-----")
	//}

	return conf, err
}
