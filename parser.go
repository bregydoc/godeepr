package godeepr

import "encoding/json"

func parseJsonToQueries(jsonQuery string) (map[string]interface{}, error) {
	queries := map[string]interface{}{}
	if err := json.Unmarshal([]byte(jsonQuery), &queries); err != nil {
		return nil, err
	}

	return queries, nil
}

func parseQueries(queries map[string]interface{}) error {
	for queryName, description := range queries {
		attributes, isObject := description.(map[string]interface{})
		if isObject {
			for name, attr	:= range attributes {
				if !attributeHasValidName(name) {
					return ErrInvalidAttributeName
				}
				if attr == contextOperator {

				}


			}
		}

		elem, isArray := description.([]interface{})
		if isArray {

		}
	}
}
