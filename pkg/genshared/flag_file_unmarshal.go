package genshared

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/golang/protobuf/proto"
	strcase "github.com/stoewer/go-strcase"
	yaml "gopkg.in/yaml.v2"
)

const (
	InputFilesTag = "file"
)

var DefaultFileUnmarshalers = FileUnmarshalers{
	"json": {Func: json.Unmarshal},
	"xml":  {Func: xml.Unmarshal},
	"toml": {Func: toml.Unmarshal},
	"yaml": {Func: yaml.Unmarshal},
	"proto": {
		Test: func(ptr interface{}) error {
			if _, ok := ptr.(proto.Message); !ok {
				return fmt.Errorf("`file` tag with `proto` only support proto.Message, but got: %T", ptr)
			}
			return nil
		},
		Func: func(data []byte, v interface{}) error {
			return proto.Unmarshal(data, v.(proto.Message))
		},
	},
}

type PreUnmarshaler interface {
	PreUnmarshal()
}

type PostUnmarshaler interface {
	PostUnmarshal() error
}

func NewInputParser(v interface{}) (*InputFilesParser, error) {
	return DefaultFileUnmarshalers.NewInputParser(v)
}

type FileUnmarshalers map[string]FileUnmarshaler

// NewInputFiles check `file` tags on every field:
// arg,unmarshaler,omitempty
func (fus FileUnmarshalers) NewInputParser(v interface{}) (*InputFilesParser, error) {
	typ := reflect.TypeOf(v)
	val := reflect.ValueOf(v)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		val = val.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return nil, fmt.Errorf("InputFilesParser only support struct, but got %T", v)
	}

	fileArgs := make([]*inputFileArg, 0, typ.NumField())
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)

		tag, ok := field.Tag.Lookup(InputFilesTag)
		if !ok {
			continue
		}

		var isPtr bool
		var isNil bool
		ftyp := field.Type
		fval := val.Field(i)
		if ftyp.Kind() == reflect.Ptr {
			isPtr = true
			isNil = fval.IsNil()
			ftyp = ftyp.Elem()
		}
		if ftyp.Kind() != reflect.Struct {
			return nil, fmt.Errorf("`%s` tag only support struct field, but got %s", InputFilesTag, ftyp.Name())
		}

		var setBack func()
		fvalPtr := fval
		if !isPtr {
			fvalPtr = fval.Addr()
		} else if isNil {
			fvalPtr = reflect.New(ftyp)
			setBack = func() { fval.Set(fvalPtr) }
		}
		ptr := fvalPtr.Interface()

		// tag:
		// arg,unmarshaler
		tags := strings.Split(tag, ",")
		if len(tags) != 2 {
			return nil, fmt.Errorf("invalid `%s` tag segments: %s", InputFilesTag, tag)
		}

		unmarshaler, ok := fus[tags[1]]
		if !ok {
			return nil, fmt.Errorf("`%s` tag unmarshaler not supported: %s", InputFilesTag, tags[1])
		}

		if unmarshaler.Test != nil {
			err := unmarshaler.Test(ptr)
			if err != nil {
				return nil, err
			}
		}

		// arg name
		name := tags[0]
		if name == "" {
			name = strcase.SnakeCase(field.Name)
		}
		fileArgs = append(fileArgs, &inputFileArg{
			name:      name,
			ptr:       ptr,
			unmarshal: unmarshaler.Func,
			isPtr:     isPtr,
			isNil:     isNil,
			setBack:   setBack,
		})
	}

	for _, arg := range fileArgs {
		required := "the required "
		if arg.isPtr {
			required = "the optional "
		}
		flag.StringVar(&arg.file, arg.name, "", required+arg.name+" input file")
	}
	return &InputFilesParser{args: fileArgs}, nil
}

func (fs *InputFilesParser) Parse() (err error) {
	if fs.parsed {
		return fs.err
	}
	fs.parsed = true

	defer func() { fs.err = err }()

	for _, arg := range fs.args {
		if !arg.isPtr && arg.file == "" {
			return fmt.Errorf("missing arg `%s`", arg.name)
		}
	}

	var b []byte
	for _, arg := range fs.args {
		if arg.isPtr && arg.file == "" {
			continue
		}
		b, err = ioutil.ReadFile(arg.file)
		if err != nil {
			return err
		}

		if pu, ok := arg.ptr.(PreUnmarshaler); ok {
			pu.PreUnmarshal()
		}

		err = arg.unmarshal(b, arg.ptr)
		if err != nil {
			return err
		}

		if pu, ok := arg.ptr.(PostUnmarshaler); ok {
			err = pu.PostUnmarshal()
			if err != nil {
				return err
			}
		}

		if arg.isNil {
			arg.setBack()
		}
	}
	return nil
}

type FileUnmarshaler struct {
	Test func(v interface{}) error
	Func func(data []byte, v interface{}) error
}

type InputFilesParser struct {
	parsed bool
	err    error
	args   []*inputFileArg
}

type inputFileArg struct {
	name      string
	file      string
	ptr       interface{}
	unmarshal func(data []byte, v interface{}) error
	isPtr     bool

	isNil   bool
	setBack func()
}
