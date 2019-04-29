package util

import strcase "github.com/stoewer/go-strcase"

func JoinEntityId(entity, id string) string {
	if entity == "" {
		return id
	}
	return PowerLowerCamel(entity) + PowerCamel(id)
}

func PowerCamel(s string) string {
	return strcase.UpperCamelCase(strcase.SnakeCase(s))
}

func PowerLowerCamel(s string) string {
	return strcase.LowerCamelCase(strcase.SnakeCase(s))
}
