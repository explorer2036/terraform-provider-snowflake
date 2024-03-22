package testint

import (
	"reflect"
	"testing"
)

func TestInt_ShowByID(t *testing.T) {
	client := testClient(t)

	// val := reflect.ValueOf(&t)("Geeks").Call([]reflect.Value{})
	value := reflect.ValueOf(*client)
	for i := 0; i < value.NumField(); i++ {
		fieldName := value.Type().Field(i).Name
		switch fieldName {
		case "config", "db", "sessionID", "accountLocator", "dryRun", "traceLogs":
			continue
		case "ContextFunctions", "ConversionFunctions", "SystemFunctions", "ReplicationFunctions":
			continue
		}
		// if fieldName == "" {
		// 	continue
		// }

		// in := make([]reflect.Value, 1)
		t.Logf("name: %s", fieldName)
		// method := value.Field(i).MethodByName("ShowByID")
		// for j := 0; j < method.NumField(); j++ {
		// 	t.Logf("%v", method.Field(j).Interface())
		// }
		// in[0] = reflect.ValueOf(value.Field(i).Interface())
		// defV.Field(i).MethodByName("Conversion").Call(in)
	}
}
