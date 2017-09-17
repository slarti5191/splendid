package configuration

import (
	"reflect"
	"time"
	"log"
	"os"
	"fmt"
)

type DeviceConfig struct {
	Host           string
	Type           string
	Target         string
	User           string
	Pass           string
	Port           int
	Timeout        time.Duration
	CommandTimeout time.Duration
}

type SplendidConfig struct {
	Debug         bool
	Interval      int //time.Duration
	Timeout       int //time.Duration
	GitPush       bool
	Insecure      bool
	Concurrency   int
	HttpListen    string
	HttpEnabled   bool
	SmtpString    string
	Workspace     string
	ExecutableDir string
	ToEmail       string
	FromEmail     string
	UseSyslog     bool
	DefaultUser   string
	DefaultPass   string
	DefaultMethod string
	CmwPass       string
	//Devices       []DeviceConfig
	ConfigFile    string
}

// GetConfigs loads the config file, then parses flags
func GetConfigs(configFile string) (*SplendidConfig, error) {
	// Goal
	// 1) Need to load flags first to determine which config file to use.
	// 2) Create base "defaults" config.
	// 3) Load in config from file, fill missing values with defaults.
	// 4) If flag value is different than default value, override config.
	defaults := SplendidConfig{
		false,
		30,
		120,
		false,
		false,
		30,
		"localhost:5001",
		true,
		"server:port",
		"/workspace",
		"/",
		"",
		"",
		false,
		"user",
		"pass",
		"none",
		"none",
	//	nil,
		"sample.conf",
	}


	flags := new (SplendidConfig)
	parseConfigFlags(flags, defaults)

	config, err := loadConfig(flags.ConfigFile, defaults)
	if err != nil {
		return nil, err
	}

	config.flagUpdate(*flags, defaults)
	os.Exit(0)
	return config, nil
}

func (c *SplendidConfig) flagUpdate(flags SplendidConfig, defaults SplendidConfig) {

	//e, d, err := Compare(flags, defaults)
	//if err != nil {
	//	panic(err)
	//}
	//log.Println(e)
	//log.Println(d)
	//os.Exit(0)

	defaultValue := reflect.ValueOf(defaults)
	//loadedValue := reflect.ValueOf(*c)
	flagValue := reflect.ValueOf(flags)

	//final := reflect.ValueOf(c).Elem()

	for i := 0; i < flagValue.NumField(); i++ {
		fmt.Println("=--------------=")
		v1 := defaultValue.Field(i)
		v2 := flagValue.Field(i)
		v1 = reflect.Indirect(v1)
		v2 = reflect.Indirect(v2)

		//vo1 := flagValue.Field(6)
		//switch vo1.Interface().(type) {
		//case int, bool, string, float64:
		//	log.Println("int, bool, string, float64")
		//default:
		//	log.Println("none of the above")
		//}
		//os.Exit(1)


		if v1.Type() != v2.Type() {
			log.Println("TYPE FAIL")
		}
		if !v1.IsValid() || !v2.IsValid() {
			log.Println("Valid FAIL")
		}
		if v1.Interface() != v2.Interface() {
			log.Printf("v1(%v): %v\n", v1.Kind(), v1.Interface())
			log.Printf("v2(%v): %v\n", v2.Kind(), v2.Interface())

		//if !reflect.DeepEqual(defaultValue.Field(i), flagValue.Field(i)) {
		//if reflect.ValueOf(defaultValue.Field(i)) != reflect.ValueOf(flagValue.Field(i)) {
		//if reflect.ValueOf(defaultValue.Field(i)) != reflect.ValueOf(flagValue.Field(i)) {
			//log.Printf("Overwrite: %v\n", loadedValue.Type().Field(i).Name)
			switch v1.Interface().(type) {
			case int, bool, string, float64:
				log.Println("int, bool, string, float64")
			default:
				log.Println("none of the above")
			}

			if v2.Interface() == reflect.Zero(v2.Type()).Interface() {
				log.Println("EmptyValue")
				continue
			} else {
				log.Printf("NotEmpty[[%v]][[%v]]", v2, reflect.Zero(v2.Type()))
			}
			if !defaultValue.Field(i).IsValid() {
				log.Println("InvalidInvalid")
			}

			log.Printf("Overwrite: %v\n", defaultValue.Type().Field(i).Name)
			log.Printf("Overwrote: %v from %v\n", v1.Interface(), v2.Interface())

			// FINALLY got this working...

			// TODO: Clean up, and merge over flags that are not empty.

		} else {
			log.Printf("As it was: %v\n", defaultValue.Type().Field(i).Name)
		}

		//if !flagValue.Field(i).IsValid() {
		//	final.Field(i).Set(flagValue.Field(i))
		//} else {
		//	final.Field(i).Set(loadedValue.Field(i))
		//}
	}
	return
}



type CompareResult struct {
	FieldName string
	Value1    interface{}
	Value2    interface{}
}

func Compare(struct1 interface{}, struct2 interface{}) (areEqual bool, differences []*CompareResult, err error) {

	if struct1 == nil || struct2 == nil {
		return false, nil, fmt.Errorf("One of the inputs cannot be nil. struct1: %v, struct2 : %v ", struct1, struct2)
	}

	//Get values of the structs
	v1, v2 := reflect.ValueOf(struct1), reflect.ValueOf(struct2)

	//Handle pointers, if a non-pointer struct is passed in, Indirect will just return the element
	v1, v2 = reflect.Indirect(v1), reflect.Indirect(v2)
	if !v1.IsValid() || !v2.IsValid() {
			return false, nil, fmt.Errorf("Types cannot be nil. v1 %v - v2 %v", v1.IsValid(), v2.IsValid())
	}

	//Cache v1 struct type
	structType := v1.Type()

	//Verify both v1 and v2 are the same type
	if structType != v2.Type() {
		return false, nil, fmt.Errorf("Structs must be the same type. Struct1 %v - Stuct2 -%v", structType, v2.Type())
	}

	//Verify v1 is a struct, if v1 is a struct then v2 is also a struct because we have already verified the types are equal
	if v1.Kind() != reflect.Struct {
		return false, nil, fmt.Errorf("Types must both be structs.  Kind1: %v, Kind2 :v", v1.Kind(), v2.Kind())
	}

	//Initialize differences to ensure length of 0 on return
	differences = make([]*CompareResult, 0)

	for i, numFields := 0, v1.NumField(); i < numFields; i++ {
		//Get values of the structure's fields
		field1, field2 := v1.Field(i), v2.Field(i)

		//Get a reference to the field type
		fieldType := structType.Field(i)

		//If the field name is unexported, skip
		if fieldType.PkgPath != "" {
			continue
		}

		//Handle nil pointers, if a non-pointer field is passed in, Indirect will just return the element
		field1, field2 = reflect.Indirect(field1), reflect.Indirect(field2)

		switch valid1, valid2 := field1.IsValid(), field2.IsValid(); {
		//If both are valid, do nothing
		case valid1 && valid2:
			//If only field1 is valid, set field2 to reflect.Zero
		case valid1:
			field2 = reflect.Zero(field1.Type())
			//If only field1 is valid, set field2 to reflect.Zero
		case valid2:
			field1 = reflect.Zero(field2.Type())
			//Both are invalid so skip loop body
		default:
			continue
		}

		if field1.Kind() == reflect.Interface {
			return false, nil, fmt.Errorf("Type of field cannot be interface. field1: %v, field2: %v", field1, field2)
		}

		switch val1, val2 := field1.Interface(), field2.Interface(); val1.(type) {
		case int, bool, string, float64, time.Time:
			if val1 != val2 {
				log.Println(val1, val2)
				result := &CompareResult{FieldName: fieldType.Name, Value1: val1, Value2: val2}
				differences = append(differences, result)
			}
		default:
			return false, nil, fmt.Errorf("Unsupported type: %v", val1)
		}

	}

	return len(differences) == 0, differences, nil
}