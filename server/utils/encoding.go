package utils

import (
	"encoding/json"
	"strconv"
)

func Serialize(data any) []byte {
	json_data, err := json.Marshal(data)
	Panic(err)
	return json_data
}

func Unserialize(data string, v any) {
	err := json.Unmarshal([]byte(data), v)
	Panic(err)
}

func ToString(data any) (ret string) {
	switch data.(type) {
	case int64:
		ret = strconv.Itoa(int(data.(int64)))
	case int32:
		ret = strconv.Itoa(int(data.(int32)))
	case int:
		ret = strconv.Itoa(data.(int))
	case float64:
		ret = strconv.FormatFloat(data.(float64), 'e', 10, 64)
	case float32:
		ret = strconv.FormatFloat(data.(float64), 'e', 10, 32)
	}
	return ret
}

func ToBytes(data any) []byte {
	return []byte(ToString(data))
}
