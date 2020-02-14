package util

import (
	"reflect"

	"github.com/golang/protobuf/proto"
)

func ProtoMessageOneofFieldName(is interface{}) string {
	typ := reflect.TypeOf(is)
	for typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	var p proto.Properties
	p.Parse(typ.Field(0).Tag.Get("protobuf"))
	return p.OrigName
}
