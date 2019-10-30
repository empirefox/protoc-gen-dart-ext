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
	final _bl = {{ .PgdeFile.As }}.Lists.len({{ .Accessor }});
{{ end }}

{{ if or $r.Len (and $r.MinLen $r.MaxLen (eq $r.GetMinLen $r.GetMaxLen)) }}
	{{ if $r.Len }}
		if (_bl != {{ $r.GetLen }})
			throw {{ .PgdeFile.As }}.LenConstError.byte ({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validateEq, {{ $r.GetLen }});
	{{ else }}
		if (_bl != {{ $r.GetMinLen }})
			throw {{ .PgdeFile.As }}.LenConstError.byte ({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validateEq, {{ $r.GetMinLen }});
	{{ end }}
{{ else if $r.MinLen }}
	{{ if $r.MaxLen }}
		if (_bl < {{ $r.GetMinLen }} || _bl > {{ $r.GetMaxLen }})
			throw {{ .PgdeFile.As }}.RangeLenError.byte(const ErrorRange.inEE ({{ .Err3Args }}, {{ $r.GetMinLen }}, {{ $r.GetMaxLen }}));
	{{ else }}
		if (_bl < {{ $r.GetMinLen }})
			throw {{ .PgdeFile.As }}.LenConstError.byte ({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validateGte, {{ $r.GetMinLen }});
	{{ end }}
{{ else if $r.MaxLen }}
	if (_bl > {{ $r.GetMaxLen }})
		throw {{ .PgdeFile.As }}.LenConstError.byte ({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validateLte, {{ $r.GetMaxLen }});
{{ end }}

{{ if $r.Prefix }}
	if (!{{ .PgdeFile.As }}.Bytes.hasPrefix({{ .Accessor }}, {{ $kPrefix }}))
		throw {{ .PgdeFile.As }}.ConstError({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validatePrefix, {{ $kPrefixPrint }});
{{ end }}

{{ if $r.Suffix }}
	if (!{{ .PgdeFile.As }}.Bytes.hasSuffix({{ .Accessor }}, {{ $kSuffix }}))
		throw {{ .PgdeFile.As }}.ConstError({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validateSuffix, {{ $kSuffixPrint }});
{{ end }}

{{ if $r.Contains }}
	if ({{ .PgdeFile.As }}.Bytes.indexOf({{ .Accessor }}, {{ $kContains }}) == -1)
		throw {{ .PgdeFile.As }}.ConstError({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validateContains, {{ $kContainsPrint }});
{{ end }}

{{ if $r.In }}
	if ({{ .PgdeFile.As }}.Bytes.indexOfList({{ $kIn }}, {{ .Accessor }}) == -1)
		throw {{ .PgdeFile.As }}.InError({{ .Err3Args }}, {{ $kInPrint }});
{{ else if $r.NotIn }}
	if ({{ .PgdeFile.As }}.Bytes.indexOfList({{ $kNotIn }}, {{ .Accessor }}) != -1)
		throw {{ .PgdeFile.As }}.InError.not ({{ .Err3Args }}, {{ $kNotInPrint }});
{{ end }}

{{ if $r.Const }}
	if (!{{ .PgdeFile.As }}.Bytes.equal({{ .Accessor }}, {{ $kConst }}))
		throw {{ .PgdeFile.As }}.ConstError({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validateEq, {{ $kConstPrint }});
{{ end }}

{{ if $r.GetIp }}
	if (!{{ .PgdeFile.As }}.Bytes.isIP({{ .Accessor }}))
		throw {{ .PgdeFile.As }}.BeSomethingError({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validateIp);
{{ else if $r.GetIpv4 }}
	if (!{{ .PgdeFile.As }}.Bytes.isIpv4({{ .Accessor }}))
		throw {{ .PgdeFile.As }}.BeSomethingError({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validateIpv4);
{{ else if $r.GetIpv6 }}
	if (!{{ .PgdeFile.As }}.Bytes.isIpv6({{ .Accessor }}))
		throw {{ .PgdeFile.As }}.BeSomethingError({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validateIpv6);
{{ end }}

{{ if $r.Pattern }}
	try {
		if (!{{ $kPattern }}.hasMatch({{ .DartConvertLib.As }}.utf8.decode({{ .Accessor }})))
			throw {{ .PgdeFile.As }}.ConstError({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validatePattern, {{ $kPattern }}.pattern);
	} on FormatException {
		throw {{ .PgdeFile.As }}.ConstError({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validatePattern, {{ $kPattern }}.pattern);
	}
{{ end }}
`
