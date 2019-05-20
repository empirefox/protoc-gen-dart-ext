package pgvt

const bytesConstTpl = bytesTplDef + `
{{- if $r.Const }}
	static const List<int> {{ $kConst }} = const {{ bytesLit $r.GetConst }};
	static const String {{ $kConstPrint }} = {{ byteStr $r.GetConst }};
{{- end -}}
{{- if $r.Pattern }}
	static final RegExp {{ $kPattern }} = RegExp({{ lit $r.GetPattern }});
{{- end -}}
{{- if $r.Prefix }}
	static const List<int> {{ $kPrefix }} = const {{ bytesLit $r.GetPrefix }};
	static const String {{ $kPrefixPrint }} = {{ byteStr $r.GetPrefix }};
{{- end -}}
{{- if $r.Suffix }}
	static const List<int> {{ $kSuffix }} = const {{ bytesLit $r.GetSuffix }};
	static const String {{ $kSuffixPrint }} = {{ byteStr $r.GetSuffix }};
{{- end -}}
{{- if $r.Contains }}
	static const List<int> {{ $kContains }} = const {{ bytesLit $r.GetContains }};
	static const String {{ $kContainsPrint }} = {{ byteStr $r.GetContains }};
{{- end -}}
{{- if $r.In }}
	static const List<List<int>> {{ $kIn }} = [
		{{- range $r.In }}
		{{ bytesLit . }},
		{{- end }}
	];
	static final List<String> {{ $kInPrint }} = [
		{{- range $r.In }}
		{{ byteStr . }},
		{{- end }}
	];
{{- end -}}
{{- if $r.NotIn }}
	static const List<List<int>> {{ $kNotIn }} = [
		{{- range $r.NotIn }}
		{{ bytesLit . }},
		{{- end }}
	];
	static const List<String> {{ $kNotInPrint }} = [
		{{- range $r.NotIn }}
		{{ byteStr . }},
		{{- end }}
	];
{{- end -}}`

const bytesTplDef = `{{ $r := .Rules }}
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
	final _vl = $pgde.Lists.len(_v);
{{ end }}

{{ if or $r.Len (and $r.MinLen $r.MaxLen (eq $r.GetMinLen $r.GetMaxLen)) }}
	{{ if $r.Len }}
		if (_vl != {{ $r.GetLen }})
			throw $pgde.LenConstError.byte(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateEq, {{ $r.GetLen }});
	{{ else }}
		if (_vl != {{ $r.GetMinLen }})
			throw $pgde.LenConstError.byte(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateEq, {{ $r.GetMinLen }});
	{{ end }}
{{ else if $r.MinLen }}
	{{ if $r.MaxLen }}
		if (_vl < {{ $r.GetMinLen }} || _vl > {{ $r.GetMaxLen }})
			throw $pgde.RangeLenError.byte(const ErrorRange.inEE(info, {{ $.Number }}, {{ $.L10nField }}, {{ $r.GetMinLen }}, {{ $r.GetMaxLen }}));
	{{ else }}
		if (_vl < {{ $r.GetMinLen }})
			throw $pgde.LenConstError.byte(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateGte, {{ $r.GetMinLen }});
	{{ end }}
{{ else if $r.MaxLen }}
	if (_vl > {{ $r.GetMaxLen }})
		throw $pgde.LenConstError.byte(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateLte, {{ $r.GetMaxLen }});
{{ end }}

{{ if $r.Prefix }}
	if (!$pgde.Bytes.hasPrefix(_v, {{ $kPrefix }}))
		throw $pgde.ConstError(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validatePrefix, {{ $kPrefixPrint }});
{{ end }}

{{ if $r.Suffix }}
	if (!$pgde.Bytes.hasSuffix(_v, {{ $kSuffix }}))
		throw $pgde.ConstError(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateSuffix, {{ $kSuffixPrint }});
{{ end }}

{{ if $r.Contains }}
	if ($pgde.Bytes.indexOf(_v, {{ $kContains }}) == -1)
		throw $pgde.ConstError(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateContains, {{ $kContainsPrint }});
{{ end }}

{{ if $r.In }}
	if ($pgde.Bytes.indexOfList({{ $kIn }}, _v) == -1)
		throw $pgde.InError(info, {{ $.Number }}, {{ $.L10nField }}, {{ $kInPrint }});
{{ else if $r.NotIn }}
	if ($pgde.Bytes.indexOfList({{ $kNotIn }}, _v) != -1)
		throw $pgde.InError.not(info, {{ $.Number }}, {{ $.L10nField }}, {{ $kNotInPrint }});
{{ end }}

{{ if $r.Const }}
	if (!$pgde.Bytes.equal(_v, {{ $kConst }}))
		throw $pgde.ConstError(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateEq, {{ $kConstPrint }});
{{ end }}

{{ if $r.GetIp }}
	if (!$pgde.Bytes.isIP(_v))
		throw $pgde.BeSomethingError(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateIp);
{{ else if $r.GetIpv4 }}
	if (!$pgde.Bytes.isIpv4(_v))
		throw $pgde.BeSomethingError(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateIpv4);
{{ else if $r.GetIpv6 }}
	if (!$pgde.Bytes.isIpv6(_v))
		throw $pgde.BeSomethingError(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateIpv6);
{{ end }}

{{ if $r.Pattern }}
	try {
		if (!{{ $kPattern }}.hasMatch(utf8.decode(_v)))
			throw $pgde.ConstError(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validatePattern, {{ $kPattern }}.pattern);
	} on FormatException {
		throw $pgde.ConstError(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validatePattern, {{ $kPattern }}.pattern);
	}
{{ end }}
`
