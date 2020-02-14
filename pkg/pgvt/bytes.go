package pgvt

const bytesConstTpl = bytesTplDef + `
{{- if $r.Const }}
	static const List<int> {{ $kConst }} = {{ lit . $r.GetConst }};
	static const String {{ $kConstPrint }} = {{ bytesStr $r.GetConst }};
{{- end -}}
{{- if $r.Pattern }}
	static final RegExp {{ $kPattern }} = RegExp({{ lit . $r.GetPattern }});
{{- end -}}
{{- if $r.Prefix }}
	static const List<int> {{ $kPrefix }} = {{ lit . $r.GetPrefix }};
	static const String {{ $kPrefixPrint }} = {{ bytesStr $r.GetPrefix }};
{{- end -}}
{{- if $r.Suffix }}
	static const List<int> {{ $kSuffix }} = {{ lit . $r.GetSuffix }};
	static const String {{ $kSuffixPrint }} = {{ bytesStr $r.GetSuffix }};
{{- end -}}
{{- if $r.Contains }}
	static const List<int> {{ $kContains }} = {{ lit . $r.GetContains }};
	static const String {{ $kContainsPrint }} = {{ bytesStr $r.GetContains }};
{{- end -}}
{{- if $r.In }}
	static const List<List<int>> {{ $kIn }} = [
		{{- range $r.In }}
		{{ lit $ . }},
		{{- end }}
	];
	static final List<String> {{ $kInPrint }} = [
		{{- range $r.In }}
		{{ bytesStr . }},
		{{- end }}
	];
{{- end -}}
{{- if $r.NotIn }}
	static const List<List<int>> {{ $kNotIn }} = [
		{{- range $r.NotIn }}
		{{ lit $ . }},
		{{- end }}
	];
	static const List<String> {{ $kNotInPrint }} = [
		{{- range $r.NotIn }}
		{{ bytesStr . }},
		{{- end }}
	];
{{- end -}}`

const bytesTplDef = `{{ $r := .Pgv.Rules }}
{{- $kConst := constRef . "kConst" }}
{{- $kConstPrint := constRef . "kConstPrint" }}
{{- $kIn := constRef . "kIn" }}
{{- $kInPrint := constRef . "kInPrint" }}
{{- $kNotIn := constRef . "kNotIn" }}
{{- $kNotInPrint := constRef . "kNotInPrint" }}
{{- $kPattern := constRef . "kPattern" }}
{{- $kPrefix := constRef . "kPrefix" }}
{{- $kPrefixPrint := constRef . "kPrefixPrint" }}
{{- $kContains := constRef . "kContains" }}
{{- $kContainsPrint := constRef . "kContainsPrint" }}
{{- $kSuffix := constRef . "kSuffix" }}
{{- $kSuffixPrint := constRef . "kSuffixPrint" }}`

const bytesTpl = bytesTplDef + `
{{ if or $r.Len $r.MinLen $r.MaxLen }}
	final _bl = {{ .PgdeFile.AsDot "Lists" }}.len({{ .Accessor }});
{{ end }}

{{ if or $r.Len (and $r.MinLen $r.MaxLen (eq $r.GetMinLen $r.GetMaxLen)) }}
	{{ if $r.Len }}
		if (_bl != {{ $r.GetLen }})
			throw {{ .PgdeFile.AsDot "LenConstError" }}.byte ({{ .Err4Args }}, {{ .InfoAccessor }}.l10n.validateEq, {{ $r.GetLen }});
	{{ else }}
		if (_bl != {{ $r.GetMinLen }})
			throw {{ .PgdeFile.AsDot "LenConstError" }}.byte ({{ .Err4Args }}, {{ .InfoAccessor }}.l10n.validateEq, {{ $r.GetMinLen }});
	{{ end }}
{{ else if $r.MinLen }}
	{{ if $r.MaxLen }}
		if (_bl < {{ $r.GetMinLen }} || _bl > {{ $r.GetMaxLen }})
			throw {{ .PgdeFile.AsDot "RangeLenError" }}.byte(const ErrorRange.inEE ({{ .Err3Args }}, {{ $r.GetMinLen }}, {{ $r.GetMaxLen }}));
	{{ else }}
		if (_bl < {{ $r.GetMinLen }})
			throw {{ .PgdeFile.AsDot "LenConstError" }}.byte ({{ .Err4Args }}, {{ .InfoAccessor }}.l10n.validateGte, {{ $r.GetMinLen }});
	{{ end }}
{{ else if $r.MaxLen }}
	if (_bl > {{ $r.GetMaxLen }})
		throw {{ .PgdeFile.AsDot "LenConstError" }}.byte ({{ .Err4Args }}, {{ .InfoAccessor }}.l10n.validateLte, {{ $r.GetMaxLen }});
{{ end }}

{{ if $r.Prefix }}
	if (!{{ .PgdeFile.AsDot "Bytes" }}.hasPrefix({{ .Accessor }}, {{ $kPrefix }}))
		throw {{ .PgdeFile.AsDot "ConstError" }}({{ .Err4Args }}, {{ .InfoAccessor }}.l10n.validatePrefix, {{ $kPrefixPrint }});
{{ end }}

{{ if $r.Suffix }}
	if (!{{ .PgdeFile.AsDot "Bytes" }}.hasSuffix({{ .Accessor }}, {{ $kSuffix }}))
		throw {{ .PgdeFile.AsDot "ConstError" }}({{ .Err4Args }}, {{ .InfoAccessor }}.l10n.validateSuffix, {{ $kSuffixPrint }});
{{ end }}

{{ if $r.Contains }}
	if ({{ .PgdeFile.AsDot "Bytes" }}.indexOf({{ .Accessor }}, {{ $kContains }}) == -1)
		throw {{ .PgdeFile.AsDot "ConstError" }}({{ .Err4Args }}, {{ .InfoAccessor }}.l10n.validateContains, {{ $kContainsPrint }});
{{ end }}

{{ if $r.In }}
	if ({{ .PgdeFile.AsDot "Bytes" }}.indexOfList({{ $kIn }}, {{ .Accessor }}) == -1)
		throw {{ .PgdeFile.AsDot "InError" }}({{ .Err4Args }}, {{ $kInPrint }});
{{ else if $r.NotIn }}
	if ({{ .PgdeFile.AsDot "Bytes" }}.indexOfList({{ $kNotIn }}, {{ .Accessor }}) != -1)
		throw {{ .PgdeFile.AsDot "InError" }}.not ({{ .Err3Args }}, {{ $kNotInPrint }});
{{ end }}

{{ if $r.Const }}
	if (!{{ .PgdeFile.AsDot "Bytes" }}.equal({{ .Accessor }}, {{ $kConst }}))
		throw {{ .PgdeFile.AsDot "ConstError" }}({{ .Err4Args }}, {{ .InfoAccessor }}.l10n.validateEq, {{ $kConstPrint }});
{{ end }}

{{ if $r.GetIp }}
	if (!{{ .PgdeFile.AsDot "Bytes" }}.isIP({{ .Accessor }}))
		throw {{ .PgdeFile.AsDot "BeSomethingError" }}({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validateIp);
{{ else if $r.GetIpv4 }}
	if (!{{ .PgdeFile.AsDot "Bytes" }}.isIpv4({{ .Accessor }}))
		throw {{ .PgdeFile.AsDot "BeSomethingError" }}({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validateIpv4);
{{ else if $r.GetIpv6 }}
	if (!{{ .PgdeFile.AsDot "Bytes" }}.isIpv6({{ .Accessor }}))
		throw {{ .PgdeFile.AsDot "BeSomethingError" }}({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validateIpv6);
{{ end }}

{{ if $r.Pattern }}
	try {
		if (!{{ $kPattern }}.hasMatch({{ .ConvertLib.AsDot "utf8" }}.decode({{ .Accessor }})))
			throw {{ .PgdeFile.AsDot "ConstError" }}({{ .Err4Args }}, {{ .InfoAccessor }}.l10n.validatePattern, {{ $kPattern }}.pattern);
	} on FormatException {
		throw {{ .PgdeFile.AsDot "ConstError" }}({{ .Err4Args }}, {{ .InfoAccessor }}.l10n.validatePattern, {{ $kPattern }}.pattern);
	}
{{ end }}
`
