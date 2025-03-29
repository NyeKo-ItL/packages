package utils

import (
	"encoding/json"
	"reflect"
	"strings"
)

// UnmarshalWithUnmatched takes any struct as input, unmarshals matched JSON fields,
// and stores unmatched fields in a map.
func UnmarshalWithUnmatched(data []byte, targetStruct any, unmatchedFields *map[string]any) error {
	// Step 1: Unmarshal JSON into a temporary map to capture all fields
	tempMap := make(map[string]any)
	if err := json.Unmarshal(data, &tempMap); err != nil {
		return err
	}

	// Step 2: Unmarshal JSON directly into the target struct to populate known fields
	if err := json.Unmarshal(data, &targetStruct); err != nil {
		return err
	}

	// Step 3: Reflect on the struct's type to identify its field tags
	targetValue := reflect.ValueOf(targetStruct).Elem()
	targetType := targetValue.Type()

	// Step 4: Remove all known fields from the temp map
	for i, field := range make([]reflect.StructField, targetType.NumField()) {
		field = targetType.Field(i)
		jsonTag := field.Tag.Get("json")

		// Handle cases where struct tags may contain options like `omitempty`
		tagKey := jsonTag
		if commaIdx := strings.Index(jsonTag, ","); commaIdx != -1 {
			tagKey = jsonTag[:commaIdx]
		}

		// Remove matched fields from the temporary map
		if tagKey != "" {
			delete(tempMap, tagKey)
		}
	}

	// Step 5: Store any remaining (unmatched) fields in unmatchedFields
	*unmatchedFields = tempMap

	return nil
}
