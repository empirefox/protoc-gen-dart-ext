package util

import (
	"strings"

	"github.com/iancoleman/strcase"
)

func JoinEntityId(entity, id string) string {
	if entity == "" {
		return id
	}
	return PowerLowerCamel(entity) + PowerCamel(id)
}

func PowerCamel(s string) string {
	return strcase.ToCamel(strings.ReplaceAll(s, ".", "_"))
}

func PowerLowerCamel(s string) string {
	return strcase.ToLowerCamel(strings.ReplaceAll(s, ".", "_"))
}
