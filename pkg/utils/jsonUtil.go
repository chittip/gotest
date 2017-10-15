package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/tidwall/gjson"
)

func sameStringSlice(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	// create a map of string -> int
	diff := make(map[string]int, len(x))
	for _, _x := range x {
		// 0 value for int is 0, so just increment a counter for the string
		diff[_x]++
	}
	for _, _y := range y {
		// If the string _y is not in diff bail out early
		if _, ok := diff[_y]; !ok {
			return false
		}
		diff[_y] -= 1
		if diff[_y] == 0 {
			delete(diff, _y)
		}
	}
	if len(diff) == 0 {
		return true
	}
	return false
}

func toAStringrray(result gjson.Result) []string {
	r := make([]string, 0, len(result.Array()))
	for _, value := range result.Array() {
		r = append(r, value.String())
	}
	return r
}

// JSONEquals ...
func JSONEquals(json1, json2 string) bool {
	result := false
	gjson.ParseBytes([]byte(json1)).ForEach(func(key, value gjson.Result) bool {
		log.Println("Key:", key, "Value:", value)
		value2 := gjson.Get(json2, key.Str)
		log.Println(value2.Exists())
		if !value2.Exists() {
			result = false
			return false // stop loop
		}
		if value.Type != value2.Type {
			result = false
			return false // stop loop
		}
		log.Println("same key and same type")

		if value2.IsArray() {
			as := toAStringrray(value)
			as2 := toAStringrray(value2)

			if !sameStringSlice(as, as2) {
				log.Println("not sameStringSlice")

				result = false
				return false
			}
			log.Println(" sameStringSlice")

		} else if value2.String() != "" {
			if value2.String() != value.String() {
				result = false
				return false
			}
		}

		result = true
		return true
	})
	return result
}

// AreEqualJSON simple
func AreEqualJSON(s1, s2 string) (bool, error) {
	var o1 interface{}
	var o2 interface{}

	var err error
	err = json.Unmarshal([]byte(s1), &o1)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 1 :: %s", err.Error())
	}
	err = json.Unmarshal([]byte(s2), &o2)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 2 :: %s", err.Error())
	}

	return reflect.DeepEqual(o1, o2), nil
}
