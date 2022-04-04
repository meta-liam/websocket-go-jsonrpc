package conv

import "fmt"

func Interface2String(inter interface{}) string {
	switch inter.(type) {
	case string:
		fmt.Println("string", inter.(string))
		return inter.(string)
	}
	return ""
}

func Interface2Int64(inter interface{}) int64 {
	switch inter.(type) {
	case int64:
		fmt.Println("int", inter.(int64))
		return inter.(int64)
	}
	return 0
}

func Interface2Int(inter interface{}) int {
	switch inter.(type) {
	case int:
		fmt.Println("int", inter.(int))
		return inter.(int)
	}
	return 0
}

func Interface2Float64(inter interface{}) float64 {
	switch inter.(type) {
	case float64:
		fmt.Println("int", inter.(float64))
		return inter.(float64)
	}
	return 0
}

func Interface2Array(inter interface{}) []interface{} {
	switch value := inter.(type) {
	case []interface{}:
		//fmt.Println("int", inter.(float64))
		return value
	}
	return nil
}

func Interface2map(inter interface{}) map[string]interface{} {
	switch value := inter.(type) {
	case map[string]interface{}:
		//fmt.Println("int", inter.(float64))
		return value
	}
	return nil
}

func Print_map(m map[string]interface{}) {
	for k, v := range m {
		switch value := v.(type) {
		case nil:
			fmt.Println(k, "is nil", "null")
		case string:
			fmt.Println(k, "is string", value)
		case int:
			fmt.Println(k, "is int", value)
		case float64:
			fmt.Println(k, "is float64", value)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range value {
				fmt.Println(i, u)
			}
		case map[string]interface{}:
			fmt.Println(k, "is an map:")
			Print_map(value)
		default:
			fmt.Println(k, "is unknown type", fmt.Sprintf("%T", v))
		}
	}
}

func Print_interfaces(v interface{}) {
	//for k, v := range m {
	k := "this "
	switch value := v.(type) {
	case nil:
		fmt.Println(k, "is nil", "null")
	case string:
		fmt.Println(k, "is string", value)
	case int:
		fmt.Println(k, "is int", value)
	case float64:
		fmt.Println(k, "is float64", value)
	case []interface{}:
		fmt.Println(k, "is an array:")
		for i, u := range value {
			fmt.Println(i, u)
		}
	case map[string]interface{}:
		fmt.Println("is an map:")
		Print_map(value)
	default:
		fmt.Println("is unknown type", fmt.Sprintf("%T", v))
	}
	//}
}
