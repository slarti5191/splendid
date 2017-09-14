package configuration

import (
	"errors"
	"github.com/go-ini/ini"
	"reflect"
)

// LoadConfig loads the saved config file
func loadConfig(configFile string) (*SplendidConfig, error) {
	conf := new(SplendidConfig)
	// Load config file
	conf.ConfigFile = configFile
	cfg, err := ini.Load(conf.ConfigFile)
	if err != nil {
		// How are we improving on simply passing the error up the stack?
		//return nil, err
		return nil, errors.New("Load failed: " + err.Error())
	}

	// TODO: DARN... after all of this, I saw MapTo(p)

	// TODO: Refactor this out into separate methods.

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

	//conf.Debug = cfg.Section("main").Key("Debug").MustBool()
	//conf.DefaultUser = cfg.Section("main").Key("DefaultUser").Value()

	// TODO: Iterate through all other sections and create DeviceConfigs
	//for a, b := range cfg.Sections() {
	//	fmt.Println("====Sections====")
	//	fmt.Println(a)
	//	fmt.Println(b)
	//	fmt.Println("-----")
	//}

	return conf, nil
}
