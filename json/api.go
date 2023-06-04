package json

import "encoding/json"

func Marshal[T any](v T) ([]byte, error) {
	return json.Marshal(v)
}

func Unmarshal[T any](data []byte) (out T, err error) {
	return out, json.Unmarshal(data, &out)
}
