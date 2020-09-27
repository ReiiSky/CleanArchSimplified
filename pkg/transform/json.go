package transform

import jsoniter "github.com/json-iterator/go"

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// JSONUnmarshal prefer use jsoniter module than standard mod
// due to performance issues
func JSONUnmarshal(d []byte, v interface{}) error {
	return json.Unmarshal(v)
}

// JSONMarshal function to serialize struct to array of byte
func JSONMarshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}