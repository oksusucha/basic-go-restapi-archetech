package utils

import "encoding/json"

func PrettyPrint(input interface{}) string {
	pretty, _ := json.MarshalIndent(input, "", "  ")
	return string(pretty)
}
