package pgvt

import (
	"fmt"
	"reflect"
	"strings"
	"text/template"
	"unicode"

	"github.com/empirefox/protoc-gen-dart-ext/pkg/dart"

	"github.com/envoyproxy/protoc-gen-validate/templates/shared"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/lyft/protoc-gen-star"
	"github.com/lyft/protoc-gen-star/lang/go"
)

func Register(tpl *template.Template, params pgs.Parameters) {
	fns := dartFuncs{pgsgo.InitContext(params)}

	tpl.Funcs(map[string]interface{}{
		"byteStr":       fns.byteStr,
		"cmt":           pgs.C80,
		"durGt":         fns.durGt,
		"durLit":        fns.durLit,
		"durStr":        fns.durStr,
		"err":           fns.err,
		"errCause":      fns.errCause,
		"errIdx":        fns.errIdx,
		"errIdxCause":   fns.errIdxCause,
		"errname":       fns.errName,
		"inKey":         fns.inKey,
		"inType":        fns.inType,
		"isBytes":       fns.isBytes,
		"lit":           fns.lit,
		"lookup":        fns.lookup,
		"msgTyp":        fns.msgTyp,
		"name":          fns.Name,
		"oneof":         fns.oneofTypeName,
		"pkg":           fns.PackageName,
		"tsGt":          fns.tsGt,
		"tsLit":         fns.tsLit,
		"tsStr":         fns.tsStr,
		"typ":           fns.Type,
		"unwrap":        fns.unwrap,
		"externalEnums": fns.externalEnums,
		"enumPackages":  fns.enumPackages,
	})

	template.Must(tpl.New("msg").Parse(msgTpl))
	template.Must(tpl.New("const").Parse(constTpl))
	template.Must(tpl.New("ltgt").Parse(ltgtTpl))
	template.Must(tpl.New("in").Parse(inTpl))

	template.Must(tpl.New("none").Parse(noneTpl))
	template.Must(tpl.New("float").Parse(numTpl))
	template.Must(tpl.New("double").Parse(numTpl))
	template.Must(tpl.New("int32").Parse(numTpl))
	template.Must(tpl.New("int64").Parse(numTpl))
	template.Must(tpl.New("uint32").Parse(numTpl))
	template.Must(tpl.New("uint64").Parse(numTpl))
	template.Must(tpl.New("sint32").Parse(numTpl))
	template.Must(tpl.New("sint64").Parse(numTpl))
	template.Must(tpl.New("fixed32").Parse(numTpl))
	template.Must(tpl.New("fixed64").Parse(numTpl))
	template.Must(tpl.New("sfixed32").Parse(numTpl))
	template.Must(tpl.New("sfixed64").Parse(numTpl))

	template.Must(tpl.New("bool").Parse(constTpl))
	template.Must(tpl.New("string").Parse(strTpl))
	template.Must(tpl.New("bytes").Parse(bytesTpl))

	template.Must(tpl.New("email").Parse(emailTpl))
	template.Must(tpl.New("hostname").Parse(hostTpl))

	template.Must(tpl.New("enum").Parse(enumTpl))
	template.Must(tpl.New("repeated").Parse(repTpl))
	template.Must(tpl.New("map").Parse(mapTpl))

	template.Must(tpl.New("any").Parse(anyTpl))
	template.Must(tpl.New("timestampcmp").Parse(timestampcmpTpl))
	template.Must(tpl.New("durationcmp").Parse(durationcmpTpl))

	template.Must(tpl.New("wrapper").Parse(wrapperTpl))
}

type dartFuncs struct{ pgsgo.Context }

func (fns dartFuncs) errName(m pgs.Message) pgs.Name {
	return fns.Name(m) + "ValidationError"
}

func (fns dartFuncs) errIdxCause(ctx shared.RuleContext, idx, cause string, reason ...interface{}) string {
	f := ctx.Field
	n := fns.Name(f)

	var fld string
	if idx != "" {
		fld = fmt.Sprintf(`fmt.Sprintf("%s[%%v]", %s)`, n, idx)
	} else if ctx.Index != "" {
		fld = fmt.Sprintf(`fmt.Sprintf("%s[%%v]", %s)`, n, ctx.Index)
	} else {
		fld = fmt.Sprintf("%q", n)
	}

	causeFld := ""
	if cause != "nil" && cause != "" {
		causeFld = fmt.Sprintf("cause: %s,", cause)
	}

	keyFld := ""
	if ctx.OnKey {
		keyFld = "key: true,"
	}

	return fmt.Sprintf(`%s{
		field: %s,
		reason: %q,
		%s%s
	}`,
		fns.errName(f.Message()),
		fld,
		fmt.Sprint(reason...),
		causeFld,
		keyFld)
}

func (fns dartFuncs) err(ctx shared.RuleContext, reason ...interface{}) string {
	return fns.errIdxCause(ctx, "", "nil", reason...)
}

func (fns dartFuncs) errCause(ctx shared.RuleContext, cause string, reason ...interface{}) string {
	return fns.errIdxCause(ctx, "", cause, reason...)
}

func (fns dartFuncs) errIdx(ctx shared.RuleContext, idx string, reason ...interface{}) string {
	return fns.errIdxCause(ctx, idx, "nil", reason...)
}

func (fns dartFuncs) lookup(f pgs.Field, name string) string {
	return fmt.Sprintf(
		"_%s_%s_%s",
		fns.Name(f.Message()),
		fns.Name(f),
		name,
	)
}

func (fns dartFuncs) lit(x interface{}) string {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Interface {
		val = val.Elem()
	}

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	switch val.Kind() {
	case reflect.String:
		return fmt.Sprintf("%q", x)
	case reflect.Uint8:
		return fmt.Sprintf("0x%X", x)
	case reflect.Slice:
		els := make([]string, val.Len())
		for i, l := 0, val.Len(); i < l; i++ {
			els[i] = fns.lit(val.Index(i).Interface())
		}
		return fmt.Sprintf("%T{%s}", val.Interface(), strings.Join(els, ", "))
	default:
		return fmt.Sprint(x)
	}
}

func (fns dartFuncs) isBytes(f interface {
	ProtoType() pgs.ProtoType
}) bool {
	return f.ProtoType() == pgs.BytesT
}

func (fns dartFuncs) byteStr(x []byte) string {
	printable := true
	elms := make([]string, len(x))
	for i, b := range x {
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

func (fns dartFuncs) oneofTypeName(f pgs.Field) pgsgo.TypeName {
	return pgsgo.TypeName(fns.OneofOption(f)).Pointer()
}

func (fns dartFuncs) inType(f pgs.Field, x interface{}) string {
	switch f.Type().ProtoType() {
	case pgs.BytesT:
		return "string"
	case pgs.MessageT:
		switch x.(type) {
		case []*duration.Duration:
			return "time.Duration"
		default:
			return pgsgo.TypeName(fmt.Sprintf("%T", x)).Element().String()
		}
	default:
		return fns.Type(f).String()
	}
}

func (fns dartFuncs) inKey(f pgs.Field, x interface{}) string {
	switch f.Type().ProtoType() {
	case pgs.BytesT:
		return fns.byteStr(x.([]byte))
	case pgs.MessageT:
		switch x := x.(type) {
		case *duration.Duration:
			dur, _ := ptypes.Duration(x)
			return fns.lit(int64(dur))
		default:
			return fns.lit(x)
		}
	default:
		return fns.lit(x)
	}
}

func (fns dartFuncs) coreDur(dur *duration.Duration) (string, error) {
	if dur.GetNanos()%1000 != 0 {
		return "", fmt.Errorf("Duration.nanos must be microsecond based")
	}
	return fmt.Sprintf("Duration(seconds: %d, microseconds: %d)",
		dur.GetSeconds(), dur.GetNanos()/1000), nil
}

func (fns dartFuncs) durLit(dur *duration.Duration) (string, error) {
	if dur.GetNanos()%1000 != 0 {
		return "", fmt.Errorf("Duration.nanos must be microsecond based")
	}
	return fmt.Sprintf("$wkt.Duration()..seconds=%d..nanos=%d",
		dur.GetSeconds(), dur.GetNanos()), nil
}

func (fns dartFuncs) durGt(a, b *duration.Duration) bool {
	ad, _ := ptypes.Duration(a)
	bd, _ := ptypes.Duration(b)

	return ad > bd
}

func (fns dartFuncs) coreTs(ts *timestamp.Timestamp) (string, error) {
	if dur.GetNanos()%1000 != 0 {
		return "", fmt.Errorf("Timestamp.nanos must be microsecond based")
	}
	return fmt.Sprintf("DateTime.fromMicrosecondsSinceEpoch(%d, isUtc: true)",
		ts.GetSeconds()*1000000+ts.GetNanos()/1000), nil
}

func (fns dartFuncs) tsLit(ts *timestamp.Timestamp) (string, error) {
	if ts.GetNanos()%1000 != 0 {
		return "", fmt.Errorf("Timestamp.nanos must be microsecond based")
	}
	return fmt.Sprintf("$wkt.Timestamp()..seconds=%d..nanos=%d",
		ts.GetSeconds(), ts.GetNanos()), nil
}

func (fns dartFuncs) tsGt(a, b *timestamp.Timestamp) bool {
	at, _ := ptypes.Timestamp(a)
	bt, _ := ptypes.Timestamp(b)

	return bt.Before(at)
}

func (fns dartFuncs) unwrap(ctx shared.RuleContext, name string) (shared.RuleContext, error) {
	ctx, err := ctx.Unwrap("wrapper")
	if err != nil {
		return ctx, err
	}

	ctx.AccessorOverride = fmt.Sprintf("%s.Get%s()", name,
		pgsgo.PGGUpperCamelCase(ctx.Field.Type().Embed().Fields()[0].Name()))

	return ctx, nil
}

func (fns dartFuncs) msgTyp(message pgs.Message) pgsgo.TypeName {
	return pgsgo.TypeName(fns.Name(message))
}

func (fns dartFuncs) externalEnums(file pgs.File) []pgs.Enum {
	var out []pgs.Enum

	for _, msg := range file.AllMessages() {
		for _, fld := range msg.Fields() {
			if en := fld.Type().Enum(); fld.Type().IsEnum() && en.Package().ProtoName() != fld.Package().ProtoName() && fns.PackageName(en) != fns.PackageName(fld) {
				out = append(out, en)
			}
		}
	}

	return out
}

func (fns dartFuncs) enumPackages(enums []pgs.Enum) map[pgs.FilePath]pgs.Name {
	out := make(map[pgs.FilePath]pgs.Name, len(enums))

	for _, en := range enums {
		out[fns.ImportPath(en)] = fns.PackageName(en)
	}

	return out
}
