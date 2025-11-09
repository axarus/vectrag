package domain

import "fmt"

type FieldType string

const (
	FieldString   FieldType = "string"
	FieldText     FieldType = "text"
	FieldNumber   FieldType = "number"
	FieldBoolean  FieldType = "boolean"
	FieldDate     FieldType = "date"
	FieldDateTime FieldType = "datetime"
	FieldRelation FieldType = "relation"
)

var fieldTypeRegistry = map[FieldType]struct{}{
	FieldString:   {},
	FieldText:     {},
	FieldNumber:   {},
	FieldBoolean:  {},
	FieldDate:     {},
	FieldDateTime: {},
	FieldRelation: {},
}

func IsValidType(t string) bool {
	_, ok := fieldTypeRegistry[FieldType(t)]
	return ok
}

func ListFieldTypes() []string {
	types := make([]string, 0, len(fieldTypeRegistry))
	for ft := range fieldTypeRegistry {
		types = append(types, string(ft))
	}

	return types
}
func GetFieldSQLType(ft FieldType) (string, error) {
	switch ft {
	case FieldString:
		return "VARCHAR(255)", nil
	case FieldText:
		return "TEXT", nil
	case FieldNumber:
		return "DOUBLE PRECISION", nil
	case FieldBoolean:
		return "BOOLEAN", nil
	case FieldDate:
		return "DATE", nil
	case FieldDateTime:
		return "TIMESTAMP", nil
	case FieldRelation:
		return "VARCHAR(255)", nil
	default:
		return "", fmt.Errorf("unsupported field type: %s", ft)
	}
}
