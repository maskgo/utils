package utils

import "reflect"

func ReflectMethodArgsConvert(l int, m reflect.Value, data []interface{}) []reflect.Value {
	args := make([]reflect.Value, l)

	for i := 0; i < l; i++ {
		t := m.Type().In(i)
		switch t.Kind() {
		case reflect.Slice:
			args[i] = reflect.ValueOf(SliceInterfaceConvert(data[i].([]interface{}), t))
		case reflect.Map:
			args[i] = reflect.ValueOf(StringMapConvert(data[i].(map[string]interface{}), t))
		default:
			args[i] = reflect.ValueOf(data[i]).Convert(t)
		}
	}

	return args
}
