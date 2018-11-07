package json_mapper

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/tidwall/gjson"
)

func Mapper(src string, template string) string {

	return "a"
}

func keys(str string, src string) ([]string, error) {
	var settings map[string]interface{}
	if err := json.NewDecoder(bytes.NewReader([]byte(str))).Decode(&settings); err != nil {
		panic(err)
	}

	for k, v := range settings {
		grap := gjson.Get(src, v.(string))
		fmt.Println("\n k:", k)
		fmt.Println("\n v:", v)
	}
}
