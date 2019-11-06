package pgvt

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
	"unicode"

	"github.com/empirefox/protoc-gen-dart-ext/pkg/dart"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/dartpb"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/genshared"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/iancoleman/strcase"
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

func Register(tpl *template.Template, params pgs.Parameters) {
	fns := dartFuncs{pgsgo.InitContext(params)}

	tpl.Funcs(map[string]interface{}{
		"bytesStr":        fns.bytesStr,
		"constRef":        fns.constRef,
		"durGt":           fns.durGt,
		"isOfMessageType": fns.isOfMessageType,
		"lit":             fns.lit,
		"render":          Render(tpl),
		"renderConstants": fns.renderConstants(tpl),
		"tsGt":            fns.tsGt,
		"unwrap":          fns.unwrap,
	})

	template.Must(tpl.Parse(fileTpl))
	template.Must(tpl.New("msg").Parse(msgTpl))
	template.Must(tpl.New("oneOf").Parse(oneOfTpl))
	template.Must(tpl.New("oneOfConst").Parse(oneOfConstTpl))

	template.Must(tpl.New("none").Parse(noneTpl))
	template.Must(tpl.New("required").Parse(requiredTpl))
	template.Must(tpl.New("const").Parse(constTpl))
	template.Must(tpl.New("ltgt").Parse(ltgtTpl))
	template.Must(tpl.New("in").Parse(inTpl))
	template.Must(tpl.New("inConst").Parse(inConstTpl))

	template.Must(tpl.New("float").Parse(numTpl))
	template.Must(tpl.New("floatConst").Parse(numConstTpl))
	template.Must(tpl.New("double").Parse(numTpl))
	template.Must(tpl.New("doubleConst").Parse(numConstTpl))
	template.Must(tpl.New("int32").Parse(numTpl))
	template.Must(tpl.New("int32Const").Parse(numConstTpl))
	template.Must(tpl.New("int64").Parse(numTpl))
	template.Must(tpl.New("int64Const").Parse(numConstTpl))
	template.Must(tpl.New("uint32").Parse(numTpl))
	template.Must(tpl.New("uint32Const").Parse(numConstTpl))
	template.Must(tpl.New("uint64").Parse(numTpl))
	template.Must(tpl.New("uint64Const").Parse(numConstTpl))
	template.Must(tpl.New("sint32").Parse(numTpl))
	template.Must(tpl.New("sint32Const").Parse(numConstTpl))
	template.Must(tpl.New("sint64").Parse(numTpl))
	template.Must(tpl.New("sint64Const").Parse(numConstTpl))
	template.Must(tpl.New("fixed32").Parse(numTpl))
	template.Must(tpl.New("fixed32Const").Parse(numConstTpl))
	template.Must(tpl.New("fixed64").Parse(numTpl))
	template.Must(tpl.New("fixed64Const").Parse(numConstTpl))
	template.Must(tpl.New("sfixed32").Parse(numTpl))
	template.Must(tpl.New("sfixed32Const").Parse(numConstTpl))
	template.Must(tpl.New("sfixed64").Parse(numTpl))
	template.Must(tpl.New("sfixed64Const").Parse(numConstTpl))

	template.Must(tpl.New("bool").Parse(constTpl))
	template.Must(tpl.New("string").Parse(stringTpl))
	template.Must(tpl.New("stringConst").Parse(stringConstTpl))
	template.Must(tpl.New("bytes").Parse(bytesTpl))
	template.Must(tpl.New("bytesConst").Parse(bytesConstTpl))

	template.Must(tpl.New("message").Parse(messageTpl))
	template.Must(tpl.New("enum").Parse(enumTpl))
	template.Must(tpl.New("enumConst").Parse(enumConstTpl))
	template.Must(tpl.New("repeated").Parse(repeatedTpl))
	template.Must(tpl.New("repeatedConst").Parse(repeatedConstTpl))
	template.Must(tpl.New("map").Parse(mapTpl))
	template.Must(tpl.New("mapConst").Parse(mapConstTpl))

	template.Must(tpl.New("any").Parse(anyTpl))
	template.Must(tpl.New("anyConst").Parse(anyConstTpl))
	template.Must(tpl.New("timestamp").Parse(timestampTpl))
	template.Must(tpl.New("timestampConst").Parse(timestampConstTpl))
	template.Must(tpl.New("duration").Parse(durationTpl))
	template.Must(tpl.New("durationConst").Parse(durationConstTpl))
	template.Must(tpl.New("wrapper").Parse(wrapperTpl))
	template.Must(tpl.New("wrapperConst").Parse(wrapperConstTpl))
}

type dartFuncs struct{ pgsgo.Context }

func (fns dartFuncs) lit(f *dartpb.ValidateField, x interface{}) (string, error) {
	switch v := x.(type) {
	case string:
		return dart.RawString(v), nil
	case byte:
		return fmt.Sprintf("0x%X", x), nil
	case float32, float64, int32, uint32, bool:
		return fmt.Sprintf("%v", x), nil
	case int64, uint64:
		return fmt.Sprintf("%s(%d)", f.Int64Type(), x), nil
	case []byte:
		return `const <int>[` + genshared.ToBytes(v) + `]`, nil
	case *duration.Duration:
		return fns.coreDur(v)
	case *timestamp.Timestamp:
		return fns.coreTs(v)
	}

	return "", fmt.Errorf(`lit only accept types of:
	[string, number, bytes, Duration or Timestamp],
	but got: %#v`, x)
}

func (fns dartFuncs) coreDur(dur *duration.Duration) (string, error) {
	if dur.GetNanos()%1000 != 0 {
		return "", fmt.Errorf("Duration.nanos must be microsecond based")
	}
	return fmt.Sprintf("Duration(seconds: %d, microseconds: %d)",
		dur.GetSeconds(), dur.GetNanos()/1000), nil
}

func (fns dartFuncs) coreTs(ts *timestamp.Timestamp) (string, error) {
	if ts.GetNanos()%1000 != 0 {
		return "", fmt.Errorf("Timestamp.nanos must be microsecond based")
	}
	return fmt.Sprintf("DateTime.fromMicrosecondsSinceEpoch(%d, isUtc: true)",
		ts.GetSeconds()*1000000+int64(ts.GetNanos()/1000)), nil
}

func (fns dartFuncs) bytesStr(x []byte) string {
	printable := true
	elms := make([]string, len(x))
	for i, b := range bytes.Runes(x) {
		elms[i] = fmt.Sprintf(`\x%X`, b)
		if !unicode.IsPrint(b) {
			printable = false
		}
	}

	if printable {
		return dart.RawString(string(x))
	}
	return fmt.Sprintf(`'%s'`, strings.Join(elms, ""))
}

func (fns dartFuncs) durGt(a, b *duration.Duration) bool {
	ad, _ := ptypes.Duration(a)
	bd, _ := ptypes.Duration(b)

	return ad > bd
}

func (fns dartFuncs) tsGt(a, b *timestamp.Timestamp) bool {
	at, _ := ptypes.Timestamp(a)
	bt, _ := ptypes.Timestamp(b)

	return bt.Before(at)
}

func (fns dartFuncs) unwrap(f *dartpb.ValidateField) (*dartpb.ValidateField, error) {
	f, err := f.Unwrap("")
	if err != nil {
		return nil, err
	}
	wrapperField0 := f.File.Dart.NameOf(f.Pgs.Type().Embed().Fields()[0])
	f.Pgv.AccessorOverride = f.Accessor().Dot(wrapperField0).String()
	return f, nil
}

func (fns dartFuncs) isOfMessageType(f *dartpb.ValidateField) bool {
	return f.Pgs.Type().ProtoType() == pgs.MessageT
}

func Render(tpl *template.Template) func(v *dartpb.ValidateField) (string, error) {
	return func(v *dartpb.ValidateField) (string, error) {
		var b bytes.Buffer
		err := tpl.ExecuteTemplate(&b, v.Pgv.Typ, v)
		return b.String(), err
	}
}

func (fns dartFuncs) renderConstants(tpl *template.Template) func(v *dartpb.ValidateField) (string, error) {
	return func(v *dartpb.ValidateField) (string, error) {
		var b bytes.Buffer
		var err error

		hasConstTemplate := false
		for _, t := range tpl.Templates() {
			if t.Name() == v.Pgv.Typ+"Const" {
				hasConstTemplate = true
			}
		}

		if hasConstTemplate {
			err = tpl.ExecuteTemplate(&b, v.Pgv.Typ+"Const", v)
		}

		return b.String(), err
	}
}

func (fns dartFuncs) constRef(f *dartpb.ValidateField, rule string) string {
	return "_" + strcase.ToLowerCamel(f.DartName.String()+"_"+f.Pgv.Index+"_"+rule[1:])
}
