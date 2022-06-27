package jsonEncode

import (
	"encoding/json"
	"fmt"
	"strings"
)

type MyEncode struct {
	result []string
}

func (s *MyEncode) Encode(dat_ string) {
	// s.result = ""
	byt := []byte(dat_)

	// we are parsing it in a generic fashion
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	s.RecurseJsonMap(dat, "@")
}

// func (s *MyEncode) makeResult(str string) {
// 	s.result += str
// }

func (s *MyEncode) RecurseJsonMap(dat interface{}, path string) {
	arr, isArray := dat.([]interface{})
	obj, isMap := dat.(map[string]interface{})
	if isArray {
		for ind, val := range arr {
			newPath := fmt.Sprintf("%s[%d]", path, ind)
			s.RecurseJsonMap(val, newPath)
		}
		return
	}
	if isMap {
		for key, val := range obj {
			// placeholder := (path == "@")? "s%s%":"%s/%s"
			var placeholder string
			if path == "@" {
				path = ""
				placeholder = "%s%s"
			} else {
				placeholder = "%s/%s"
			}
			newPath := fmt.Sprintf(placeholder, path, key)
			s.RecurseJsonMap(val, newPath)
		}
		return
	}

	var item string

	switch dat.(type) {
	default:
		item = fmt.Sprintf("%v", dat)

	case string:
		item = fmt.Sprintf(`"%s"`, dat)
	}

	s.result = append(s.result, path+"="+item)

}

func (s *MyEncode) GetResult() string {
	var result string = "[\n\t"
	result = result + strings.Join(s.result, ",\n\t") + "\n]"
	return result
}
