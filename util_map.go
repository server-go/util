package util

import (
	"encoding/json"
	"fmt"
)

// Json converts to map.
func JsonToMap(jsonStr string) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		fmt.Printf("Unmarshal with error: %+v\n", err)
		return nil, err
	}
	return m, nil
}

// Map converts to json.
func MapToJson(m map[string]interface{}) (string, error) {
	jsonByte, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("Marshal with error: %+v\n", err)
		return "", nil
	}
	return string(jsonByte), nil
}

// String converts to map.
func StringToMap(content string) (map[string]interface{}, error) {
	var resMap map[string]interface{}
	err := json.Unmarshal([]byte(content), &resMap)
	if err != nil {
		fmt.Printf("String to map error:%+v", err)
	}
	return resMap, err
}

// Merging multiple maps.
func MergeMaps(maps ...map[string]interface{}) map[string]interface{} {
	newObj := map[string]interface{}{}
	for _, m := range maps {
		for k, v := range m {
			newObj[k] = v
		}
	}
	return newObj
}
