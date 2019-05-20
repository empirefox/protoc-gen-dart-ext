package protodesc

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/jhump/protoreflect/desc"
)

func GetMessageExtension(message string, extension *proto.ExtensionDesc) (interface{}, error) {
	md, err := desc.LoadMessageDescriptor(message)
	if err != nil {
		return nil, err
	}

	return proto.GetExtension(md.GetMessageOptions(), extension)
}

func GetFieldExtension(message string, tagNumber int32, extension *proto.ExtensionDesc) (interface{}, error) {
	md, err := desc.LoadMessageDescriptor(message)
	if err != nil {
		return nil, err
	}

	fd := md.FindFieldByNumber(tagNumber)
	if fd == nil {
		return nil, fmt.Errorf("tagNumber not found: %s[%d]", message, tagNumber)
	}

	return proto.GetExtension(fd.GetFieldOptions(), extension)
}
