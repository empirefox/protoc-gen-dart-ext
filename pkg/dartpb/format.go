package dartpb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/empirefox/protoc-gen-dart-ext/pkg/dart"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/pgde/format"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/util"
	"github.com/iancoleman/strcase"
	pgs "github.com/lyft/protoc-gen-star"
)

type FormatField struct {
	*Field
	Extension format.Format
}

func (f *FormatField) Render(im *dart.ImportManager) (dart.Qualifier, error) {
	switch o := f.Extension.Type.(type) {
	case nil:
		return "null", nil
	case *format.Format_Time:
		b := o.Time.Type
		bt, err := f.beTime(im, f.timeOrOfDayOrDur(b))
		if err != nil {
			return "", err
		}

		return "const " + im.PgdeFile().
			AsDotString(strcase.ToCamel(b.String())+"Formatter").
			InitWith("const "+bt), nil
	case *format.Format_Currency:
		typ := o.Currency.Type
		if typ == nil {
			return "", fmt.Errorf("currency format type is not set")
		}

		show := util.ProtoMessageOneofFieldName(typ)
		var code format.Currency
		switch t := typ.(type) {
		case *format.CurrencyFormat_Code:
			code = t.Code
		case *format.CurrencyFormat_Symbol:
			code = t.Symbol
		case *format.CurrencyFormat_Name:
			code = t.Name
		case nil:
		default:
			log.Fatalf("currency format type not found: %s", show)
		}
		return "const " + im.PgdeFile().AsDot("CurrencyFormatter").
			DotString(show).
			InitWithString(code.String(), strconv.FormatBool(o.Currency.CentMode)), nil
	case *format.Format_Number:
		number := o.Number
		getter := im.PgdeFile().AsDot("NumberFormatGetter").
			DotStringer(number.Type)
		return "const " + im.PgdeFile().AsDot("NumberFormatter").
			InitWithString(
				string(getter),
				dart.RawStringOrNull(number.Locale),
				strconv.FormatBool(number.Ordinal),
				string(im.RenderUnit(number.Unit)),
			), nil
	case *format.Format_Bytes:
		return im.PgdeFile().AsDot("BytesFormatter").DotStringer(o.Bytes.Type), nil
	default:
		log.Fatalf("format type not found: %s", util.ProtoMessageOneofFieldName(o))
	}
	return "error", nil
}

func (f *FormatField) timeOrOfDayOrDur(b format.TimeFormat_Builtin) string {
	switch b {
	case format.TimeFormat_mediumDate,
		format.TimeFormat_fullDate,
		format.TimeFormat_mediumDatetime,
		format.TimeFormat_mediumDatetime24H,
		format.TimeFormat_fullDatetime,
		format.TimeFormat_fullDatetime24H:
		return "Time"
	case format.TimeFormat_ofDay,
		format.TimeFormat_ofDay24H:
		return "OfDay"
	case format.TimeFormat_duration:
		return "Dur"
	default:
		log.Fatalf("time format type not found: %s", b)
	}
	return ""
}

func (f *FormatField) beTime(im *dart.ImportManager, timeOrOfDayOrDur string) (dart.Qualifier, error) {
	unit := f.Extension.GetTime().GetIntUnit()
	switch f.ElementOrFieldProtoType() {
	case pgs.Int32T, pgs.UInt32T, pgs.SFixed32, pgs.SInt32, pgs.Fixed32T:
		return im.PgdeFile().AsDot("BeTime").DotString(unit.String() + timeOrOfDayOrDur), nil
	case pgs.Int64T, pgs.UInt64T, pgs.SFixed64, pgs.SInt64, pgs.Fixed64T:
		return im.PgdeFile().AsDot("BeTime").DotString(unit.String() + timeOrOfDayOrDur + "64"), nil
	case pgs.MessageT:
		switch f.Pgs.Type().Embed().WellKnownType() {
		case pgs.DurationWKT:
			return im.DurationBeTime(), nil
		case pgs.TimestampWKT:
			return im.TimestampBeTime(), nil
		}
	}
	// TODO support floating?
	return "", fmt.Errorf("only int type can be format to time: %s",
		f.Pgs.FullyQualifiedName())
}
