package dartpb

import (
	"bytes"
	"fmt"
	"sort"

	"github.com/empirefox/protoc-gen-dart-ext/pkg/dart"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/pgde/format"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/pgde/zero"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/util"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/golang/protobuf/ptypes/wrappers"
	pgs "github.com/lyft/protoc-gen-star"
)

type Zeros struct {
	File          *File
	Plan          string
	ImportManager *dart.ImportManager
	pbFile        *dart.ImportFile

	MemberPrefix     string
	MemberElemSuffix string
}

func (s *Zeros) PbFile() *dart.ImportFile {
	if s.pbFile == nil {
		s.pbFile = s.ImportManager.ImportPbFile(s.File.Pgs)
	}
	return s.pbFile
}

func (s *Zeros) FullPbClass(pgsNty pgs.Message) (dart.Qualifier, error) {
	return s.ImportManager.FullPbClassOrEnum(pgsNty)
}
func (s *Zeros) FullPbEnum(pgsNty pgs.Enum) (dart.Qualifier, error) {
	return s.ImportManager.FullPbClassOrEnum(pgsNty)
}

func ZeroDisabled(msg pgs.Message) (disabled bool, err error) {
	_, err = msg.Extension(zero.E_Disabled, &disabled)
	return
}
func (v *Zeros) Disabled(msg pgs.Message) (disabled bool, err error) {
	return ZeroDisabled(msg)
}

type zeroAction struct {
	Action string
}

func (z zeroAction) IsOnSave() bool { return z.Action == "onZeroSave" }
func (z zeroAction) Actions() []string {
	return []string{"onZeroCreate", "onZeroLoad", "onZeroSave"}
}

type ZeroMessage struct {
	zeroAction
	*Message
	dart.ImportManagerCommonFiles
	Disabled bool
}

func (zm *ZeroMessage) ForAction(action string) *ZeroMessage {
	n := *zm
	n.Action = action
	return &n
}

func (nty *ZeroMessage) FullPbClass() (dart.Qualifier, error) {
	return nty.Zeros().FullPbClass(nty.Pgs)
}

type ZeroOneOf struct {
	zeroAction
	*OneOf
	Default *ZeroField

	defaultNumber int32
}

func (zo *ZeroOneOf) ForAction(action string) *ZeroOneOf {
	n := *zo
	n.Action = action
	return &n
}

func (nty *ZeroOneOf) FullPbWhichEnum() (dart.Qualifier, error) {
	return nty.Zeros().ImportManager.FullPbWhichEnum(nty.Pgs)
}

type ZeroField struct {
	zeroAction
	*Field
	dart.ImportManagerCommonFiles

	Extension zero.Zero

	// Context fields:

	Tpl string

	owner  dart.Qualifier
	getter dart.Qualifier
	setter dart.Qualifier

	Rules *zero.Zero

	SetTo interface{}
}

func (im *ZeroField) PbFile() *dart.ImportFile { return im.Zeros().PbFile() }

func (f *ZeroField) Member() string {
	name := f.Zeros().MemberPrefix + f.DartName.String()
	if f.isInElem() {
		name += f.Zeros().MemberElemSuffix
	}
	return name
}

func (f *ZeroField) isInElem() bool { return f.Rules != &f.Extension }

func (f *ZeroField) AddBuild() (bool, error) {
	if f == nil {
		return false, nil
	}
	plan := f.Zeros().Plan
	switch f.Action {
	case "all":
		return f.Rules.GetOnCreate().BuildOn("") ||
			f.Rules.GetOnLoad().BuildOn("") ||
			f.Rules.GetOnSave().BuildOn(""), nil
	case "onZeroCreate":
		return f.Rules.GetOnCreate().BuildOn(plan), nil
	case "onZeroLoad":
		return f.Rules.GetOnLoad().BuildOn(plan), nil
	case "onZeroSave":
		return f.Rules.GetOnSave().BuildOn(plan), nil
	}
	return false, fmt.Errorf("add build check for empty or unknown action: %s", f.Action)
}

func (f *ZeroField) resolveJson(js []byte) (value interface{}, err error) {
	msg, err := f.File.ReflectMessage(f.Message.Pgs)
	if err != nil {
		return nil, err
	}

	if len(js) == 0 {
		return nil, nil
	}

	var b bytes.Buffer
	b.WriteString(`{"`)
	b.WriteString(f.Pgs.Descriptor().GetJsonName())
	b.WriteString(`":`)
	b.Write(js)
	b.WriteByte('}')
	err = msg.UnmarshalJSON(b.Bytes())
	if err != nil {
		return nil, err
	}

	return msg.GetFieldByNumber(int(f.Pgs.Descriptor().GetNumber())), nil
}

type ZeroMapEntry struct {
	Key   interface{}
	Value interface{}
}

type ZeroEvalNow struct {
	Add    int64
	Format dart.Qualifier
}

func (n *ZeroField) withRules(rules *zero.Zero) error {
	n.Tpl = "none"
	switch r := rules.GetType().(type) {
	case *zero.Zero_Float:
		n.SetTo, n.Tpl = r.Float, "num"
	case *zero.Zero_Double:
		n.SetTo, n.Tpl = r.Double, "num"
	case *zero.Zero_Int32:
		n.SetTo, n.Tpl = r.Int32, "num"
	case *zero.Zero_Int64:
		n.SetTo, n.Tpl = r.Int64, "int64"
	case *zero.Zero_Uint32:
		n.SetTo, n.Tpl = r.Uint32, "num"
	case *zero.Zero_Uint64:
		n.SetTo, n.Tpl = r.Uint64, "int64"
	case *zero.Zero_Sint32:
		n.SetTo, n.Tpl = r.Sint32, "num"
	case *zero.Zero_Sint64:
		n.SetTo, n.Tpl = r.Sint64, "int64"
	case *zero.Zero_Fixed32:
		n.SetTo, n.Tpl = r.Fixed32, "num"
	case *zero.Zero_Fixed64:
		n.SetTo, n.Tpl = r.Fixed64, "int64"
	case *zero.Zero_Sfixed32:
		n.SetTo, n.Tpl = r.Sfixed32, "num"
	case *zero.Zero_Sfixed64:
		n.SetTo, n.Tpl = r.Sfixed64, "int64"
	case *zero.Zero_Bool:
		n.SetTo, n.Tpl = r.Bool, "bool"
	case *zero.Zero_String_:
		n.SetTo, n.Tpl = r.String_, "string"
	case *zero.Zero_Bytes:
		n.SetTo, n.Tpl = r.Bytes, "bytes"
	case *zero.Zero_Enum:
		if r.Enum != 0 {
			value, err := n.parseEnumValue(r.Enum)
			if err != nil {
				return err
			}
			n.SetTo, n.Tpl = value, "enum"
		}
	case *zero.Zero_Message:
		result, err := n.resolveJson(r.Message.GetJson())
		if err != nil {
			return err
		}
		n.SetTo, n.Tpl = result, "message"
	case *zero.Zero_Repeated:
		if n.isInElem() {
			return fmt.Errorf("range rules cannot have range rules")
		}
		result, err := n.resolveJson(r.Repeated.GetJson())
		if err != nil {
			return err
		}
		n.SetTo, n.Tpl = result, "repeated"
	case *zero.Zero_Map:
		if n.isInElem() {
			return fmt.Errorf("range rules cannot have range rules")
		}
		result, err := n.resolveJson(r.Map.GetJson())
		if err != nil {
			return err
		}
		if result != nil {
			raw := result.(map[interface{}]interface{})
			items := make([]ZeroMapEntry, len(raw))
			i := 0
			for k, v := range raw {
				items[i] = ZeroMapEntry{k, v}
				i++
			}
			sort.Slice(items, func(i, j int) bool {
				return fmt.Sprintf("%v", items[i].Key) < fmt.Sprintf("%v", items[j].Key)
			})
			n.SetTo, n.Tpl = items, "map"
		}
	case *zero.Zero_Any:
		if len(r.Any.GetValue()) != 0 {
			n.SetTo, n.Tpl = r.Any, "wkt"
		}
	case *zero.Zero_Duration:
		result, err := r.Duration.Eval()
		if err != nil {
			return err
		}
		// check field type to set Tpl to num or Int64
		tpl, u, err := n.elementOrIntTimeTpl()
		if err != nil {
			return err
		}
		if tpl != "" {
			t, err := ptypes.Duration(result)
			if err != nil {
				return err
			}
			setTo := u.Duration(t)
			if setTo == 0 {
				return fmt.Errorf("duration => int: result is 0: %s",
					n.Pgs.FullyQualifiedName())
			}
			n.SetTo, n.Tpl = setTo, tpl
		} else {
			n.SetTo, n.Tpl = result, "wkt"
		}
	case *zero.Zero_Time:
		result, err := r.Time.Eval()
		if err != nil {
			return err
		}
		if result != nil {
			// check field type to set Tpl to num or Int64
			tpl, u, err := n.elementOrIntTimeTpl()
			if err != nil {
				return err
			}
			// seperate EvalTime
			if result.Wkt != nil {
				if tpl != "" {
					t, err := ptypes.Timestamp(result.Wkt)
					if err != nil {
						return err
					}

					setTo := u.Time(t)
					if setTo == 0 {
						return fmt.Errorf("time => int: result is 0: %s",
							n.Pgs.FullyQualifiedName())
					}
					n.SetTo, n.Tpl = setTo, tpl
				} else {
					n.SetTo, n.Tpl = result.Wkt, "wkt"
				}
			} else {
				if tpl != "" {
					var add int64
					if result.NowAdd != 0 {
						add = u.Duration(result.NowAdd)
						if add == 0 {
							return fmt.Errorf("duration => int: result is 0: %s",
								n.Pgs.FullyQualifiedName())
						}
					}
					setTo := &ZeroEvalNow{
						Add:    add,
						Format: n.PgdeFile().AsDotString(u.String() + "Time"),
					}
					n.SetTo, n.Tpl = setTo, tpl
				} else {
					n.SetTo, n.Tpl = result.NowAdd, "evalTime"
				}
			}
		}
	case *zero.Zero_NoChange:
		if n.IsOnSave() {
			topType := n.Pgs.Type()
			elem := topType.Element()
			if topType.IsRepeated() {
				if !n.isInElem() && elem.IsEmbed() && !elem.Embed().IsWellKnown() {
					n.Tpl = "repeatedNc"
				}
			} else if topType.IsMap() {
				if !n.isInElem() && elem.IsEmbed() && !elem.Embed().IsWellKnown() {
					n.Tpl = "mapNc"
				}
			} else if topType.IsEmbed() {
				if !topType.Embed().IsWellKnown() {
					n.Tpl = "messageNc"
				} else {
					n.Tpl = "wktNc"
				}
			} else {
				switch topType.ProtoType() {
				case pgs.DoubleT, pgs.FloatT,
					pgs.Int32T, pgs.UInt32T, pgs.SFixed32, pgs.SInt32, pgs.Fixed32T:
					n.Tpl = "eq0Nc"
				case pgs.Int64T, pgs.UInt64T, pgs.SFixed64, pgs.SInt64, pgs.Fixed64T:
					n.Tpl = "isZeroNc"
				case pgs.BoolT:
					n.Tpl = "falseNc"
				case pgs.StringT, pgs.BytesT:
					n.Tpl = "isEmptyNc"
				default:
					panic(fmt.Errorf("wkt field type is not supported: %v",
						topType.ProtoType()))
				}
			}
		}
	case nil:
		// if not do this, repeated, map and message will not go deeply
		topType := n.Pgs.Type()
		elem := topType.Element()
		if topType.IsRepeated() {
			if !n.isInElem() && elem.IsEmbed() && !elem.Embed().IsWellKnown() {
				n.Tpl = "repeated"
				rules = &zero.Zero{Type: new(zero.Zero_Repeated)}
			}
		} else if topType.IsMap() {
			if !n.isInElem() && elem.IsEmbed() && !elem.Embed().IsWellKnown() {
				n.Tpl = "map"
				rules = &zero.Zero{Type: new(zero.Zero_Map)}
			}
		} else if topType.IsEmbed() && !topType.Embed().IsWellKnown() {
			n.Tpl = "message"
			rules = &zero.Zero{Type: new(zero.Zero_Message)}
		}
	default:
		return fmt.Errorf("zero type assert failed: %#T", rules.GetType())
	}

	n.Rules = rules
	return nil
}

func (f *ZeroField) elementOrIntTimeTpl() (tpl string, unit format.TimeFormat_IntUnit, err error) {
	switch f.ElementOrFieldProtoType() {
	case pgs.DoubleT, pgs.FloatT:
		// TODO support floating?
		return
	case pgs.Int32T, pgs.UInt32T, pgs.SFixed32, pgs.SInt32, pgs.Fixed32T:
		tpl = "num"
	case pgs.Int64T, pgs.UInt64T, pgs.SFixed64, pgs.SInt64, pgs.Fixed64T:
		tpl = "int64"
	default:
		return
	}

	// get time unit
	var pf format.Format
	_, err = f.Pgs.Extension(format.E_To, &pf)
	if err != nil {
		return
	}

	unit = pf.GetTime().GetIntUnit()
	if unit == format.TimeFormat_und {
		unit = format.TimeFormat_second
	}
	return
}

// Literal map or repeated will not go here, every range item goes here, so it must
// cover all range item types and must not use Member or render template.
func (zf *ZeroField) Literal(value interface{}) (string, error) {
	switch v := value.(type) {
	case bool, float32, float64, int32, uint32:
		return fmt.Sprintf("%v", v), nil
	case int64, uint64:
		return fmt.Sprintf(`%s(%d)`, zf.Int64Type(), v), nil
	case string:
		return dart.RawString(v), nil
	case []byte:
		return "const <int>[" + util.BytesLiteral(v) + "]", nil
	case pgs.EnumValue:
		targetFullEnumClass, err := zf.Zeros().FullPbEnum(v.Enum())
		if err != nil {
			return "", err
		}
		return targetFullEnumClass.Dot(zf.File.Dart.NameOf(v)).String(), nil
	case *any.Any:
		return zf.AnyType().Init().
			DotString(".typeUrl=" + v.TypeUrl).
			DotString(".value=const <int>[" + util.BytesLiteral(v.Value) + "]").
			String(), nil
	case *duration.Duration:
		return zf.DurationType().Init().
			DotString(fmt.Sprintf(".seconds=%d", v.Seconds)).
			DotString(fmt.Sprintf(".nanos=%d", v.Nanos)).
			String(), nil
	case *timestamp.Timestamp:
		return zf.TimestampType().Init().
			DotString(fmt.Sprintf(".seconds=%d", v.Seconds)).
			DotString(fmt.Sprintf(".nanos=%d", v.Nanos)).
			String(), nil
	case *ZeroEvalNow:
		t := v.Format.Dot("ofTime(DateTime.now())").String()
		if v.Add != 0 {
			t += fmt.Sprintf("%+d", v.Add)
		}
		return t, nil
	case *wrappers.DoubleValue:
		return fmt.Sprintf("%s.DoubleValue()..value=%v",
			zf.WrappersFile(), v.Value), nil
	case *wrappers.FloatValue:
		return fmt.Sprintf("%s.FloatValue()..value=%v",
			zf.WrappersFile(), v.Value), nil
	case *wrappers.Int64Value:
		return fmt.Sprintf("%s.Int64Value()..value=%s(%d)",
			zf.WrappersFile(), zf.Int64Type(), v.Value), nil
	case *wrappers.UInt64Value:
		return fmt.Sprintf("%s.UInt64Value()..value=%s(%d)",
			zf.WrappersFile(), zf.Int64Type(), v.Value), nil
	case *wrappers.Int32Value:
		return fmt.Sprintf("%s.Int32Value()..value=%v",
			zf.WrappersFile(), v.Value), nil
	case *wrappers.UInt32Value:
		return fmt.Sprintf("%s.UInt32Value()..value=%v",
			zf.WrappersFile(), v.Value), nil
	case *wrappers.BoolValue:
		return fmt.Sprintf("%s.BoolValue()..value=%v",
			zf.WrappersFile(), v.Value), nil
	case *wrappers.StringValue:
		return fmt.Sprintf("%s.StringValue()..value=%s",
			zf.WrappersFile(), dart.RawString(v.Value)), nil
	case *wrappers.BytesValue:
		return fmt.Sprintf("%s.BytesValue()..value=const <int>[%s]",
			zf.WrappersFile(), util.BytesLiteral(v.Value)), nil
	case proto.Message:
		b, err := proto.Marshal(v)
		if err != nil {
			return "", err
		}
		embed, err := zf.Zeros().FullPbClass(zf.ElementOrEmbed())
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%s.fromBuffer(const <int>[%s])",
			embed, util.BytesLiteral(b)), nil
	default:
		return "", fmt.Errorf("bug: unsupported literal type: %#T", value)
	}
}

func (zf *ZeroField) Owner() dart.Qualifier {
	if zf.owner != "" {
		return zf.owner
	}
	return "proto"
}

func (zf *ZeroField) Getter() dart.Qualifier {
	if zf.getter != "" {
		return zf.getter
	}
	return zf.Owner().Dot(zf.DartName)
}

func (zf *ZeroField) Setter() dart.Qualifier {
	if zf.setter != "" {
		return zf.setter
	}
	return zf.Getter()
}

func (f *ZeroField) TemplateName() string      { return f.Tpl }
func (f *ZeroField) ConstTemplateName() string { return f.Tpl + "Const" }

func (zf *ZeroField) Elem(owner, getter, setter dart.Qualifier) (n *ZeroField, err error) {
	nn := *zf
	n = &nn
	n.owner = owner
	n.getter = getter
	n.setter = setter
	rules := n.Extension.GetMap().GetValues()
	if rules == nil {
		rules = n.Extension.GetRepeated().GetItems()
	}
	err = n.withRules(rules)
	return
}

func (zf *ZeroField) ForAction(action string) (n *ZeroField, err error) {
	nn := *zf
	n = &nn
	n.Action = action
	err = n.withRules(&n.Extension)
	return
}

func (f *ZeroField) EnumFullPbDefault() (dart.Qualifier, error) {
	nty := f.ElementOrEnum()
	typ, err := f.Zeros().FullPbEnum(nty)
	if err != nil {
		return "", err
	}
	return typ.Dot(f.File.Dart.NameOf(nty.Values()[0])), nil
}

func (f *ZeroField) MessageFullPbClass() (dart.Qualifier, error) {
	return f.Zeros().FullPbClass(f.ElementOrEmbed())
}

func (f *ZeroField) MessageFullZeroAccessor() (dart.Qualifier, error) {
	target := f.ElementOrEmbed()
	yes, err := ZeroDisabled(target)
	if err != nil {
		return "", err
	}

	if yes {
		return "", nil
	}

	name := f.File.Dart.MessageNames(target).ZeroName()
	return f.Zeros().ImportManager.ZeroFileDot(target, name)
}

func (f *ZeroField) TypeForField() (dart.Qualifier, error) {
	return f.ImportManager().TypeForFieldType(f.Pgs.Type())
}

func (f *ZeroField) TypeForRangeKey() (dart.Qualifier, error) {
	return f.ImportManager().TypeForNonRangeFieldType(f.Pgs.Type().Key())
}
func (f *ZeroField) TypeForRangeElem() (dart.Qualifier, error) {
	return f.ImportManager().TypeForNonRangeFieldType(f.Pgs.Type().Element())
}

func (f *ZeroField) RangeElemIsMsg() bool {
	return f.Pgs.Type().Element().ProtoType() == pgs.MessageT
}

func (f *ZeroField) MessageHasMethodOr() dart.Qualifier {
	if f.Pgs.Type().IsEmbed() {
		return "!" + f.Owner().Dot(f.Names.HasMethodName()) + "()||"
	}
	return ""
}
func (f *ZeroField) WktFieldsIsZero() string {
	getter := []byte(f.Getter())
	var b bytes.Buffer
	for i, nty := range f.ElementOrEmbed().Fields() {
		if i != 0 {
			b.WriteString("||")
		}
		switch nty.Type().ProtoType() {
		case pgs.DoubleT, pgs.FloatT,
			pgs.Int32T, pgs.UInt32T, pgs.SFixed32, pgs.SInt32, pgs.Fixed32T:
			b.Write(getter)
			b.WriteByte('.')
			b.WriteString(string(f.File.Dart.NameOf(nty)))
			b.WriteString("==0")
		case pgs.Int64T, pgs.UInt64T, pgs.SFixed64, pgs.SInt64, pgs.Fixed64T:
			b.Write(getter)
			b.WriteByte('.')
			b.WriteString(string(f.File.Dart.NameOf(nty)))
			b.WriteString(".isZero")
		case pgs.BoolT:
			b.WriteByte('!')
			b.Write(getter)
		case pgs.StringT, pgs.BytesT:
			b.Write(getter)
			b.WriteByte('.')
			b.WriteString(string(f.File.Dart.NameOf(nty)))
			b.WriteString(".isEmpty")
		default:
			panic(fmt.Errorf("wkt field type is not supported: %v",
				nty.Type().ProtoType()))
		}
	}
	return b.String()
}
func (f *ZeroField) WktClone(to, from string) string {
	var b bytes.Buffer
	b.WriteString(to)
	for _, nty := range f.ElementOrEmbed().Fields() {
		name := []byte(f.File.Dart.NameOf(nty))
		b.WriteString("..")
		b.Write(name)
		b.WriteByte('=')
		b.WriteString(from)
		b.WriteByte('.')
		b.Write(name)
	}
	return b.String()
}
