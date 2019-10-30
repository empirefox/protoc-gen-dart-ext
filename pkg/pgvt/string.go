package pgvt

const stringConstTpl = stringTplDef + `
{{- if $r.Pattern }}
	static final RegExp {{ $kPattern }} = RegExp({{ lit . $r.GetPattern }});
{{- end -}}

{{ template "inConst" . }}
`

const stringTplDef = `{{ $r := .Pgv.Rules }}
{{- $kPattern := constRef . "kPattern" }}
`

const stringTpl = stringTplDef + `
{{ template "const" . }}
{{ template "in" . }}

{{ if or $r.Len $r.MinLen $r.MaxLen }}
	final _vrl = {{ .Accessor }}.runes.length;
{{ end }}

{{ if or $r.Len (and $r.MinLen $r.MaxLen (eq $r.GetMinLen $r.GetMaxLen)) }}
	{{ if $r.Len }}
		if (_vrl != {{ $r.GetLen }})
			throw {{ .PgdeFile.As }}.LenConstError.character ({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validateEq, {{ $r.GetLen }});
	{{ else }}
		if (_vrl != {{ $r.GetMinLen }})
			throw {{ .PgdeFile.As }}.LenConstError.character ({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validateEq, {{ $r.GetMinLen }});
	{{ end }}
{{ else if $r.MinLen }}
	{{ if $r.MaxLen }}
		if (_vrl < {{ $r.GetMinLen }} || _vrl > {{ $r.GetMaxLen }})
			throw {{ .PgdeFile.As }}.RangeLenError.character(const ErrorRange.inEE ({{ .Err3Args }}, {{ $r.GetMinLen }}, {{ $r.GetMaxLen }}));
	{{ else }}
		if (_vrl < {{ $r.GetMinLen }})
			throw {{ .PgdeFile.As }}.LenConstError.character ({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validateGte, {{ $r.GetMinLen }});
	{{ end }}
{{ else if $r.MaxLen }}
	if (_vrl > {{ $r.GetMaxLen }})
		throw {{ .PgdeFile.As }}.LenConstError.character ({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validateLte, {{ $r.GetMaxLen }});
{{ end }}

{{ if or $r.Len $r.MaxBytes $r.MinBytes }}
	final _vrl = {{ .PgdeFile.As }}.Lists.len({{ .Accessor }});
{{ end }}

{{ if or $r.LenBytes (and $r.MinBytes $r.MaxBytes (eq $r.GetMinBytes $r.GetMaxBytes)) }}
	{{ if $r.LenBytes }}
		if (_vbl != {{ $r.GetLenBytes }})
			throw {{ .PgdeFile.As }}.LenConstError.byte ({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validateEq, {{ $r.GetLenBytes }});
	{{ else }}
		if (_vbl != {{ $r.GetMinBytes }})
			throw {{ .PgdeFile.As }}.LenConstError.byte ({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validateEq, {{ $r.GetMinBytes }});
	{{ end }}
{{ else if $r.MinBytes }}
	{{ if $r.MaxBytes }}
		if (_vbl < {{ $r.GetMinBytes }} || _vbl > {{ $r.GetMaxBytes }})
			throw {{ .PgdeFile.As }}.RangeLenError.byte(const ErrorRange.inEE ({{ .Err3Args }}, {{ $r.GetMinBytes }}, {{ $r.GetMaxBytes }}));
	{{ else }}
		if (_vbl < {{ $r.GetMinBytes }})
			throw {{ .PgdeFile.As }}.LenConstError.byte ({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validateGte, {{ $r.GetMinBytes }});
	{{ end }}
{{ else if $r.MaxBytes }}
	if (_vbl > {{ $r.GetMaxBytes }})
		throw {{ .PgdeFile.As }}.LenConstError.byte ({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validateLte, {{ $r.GetMaxBytes }});
{{ end }}

{{ if $r.Prefix }}
	if !{{ .Accessor }}.startsWith({{ lit . $r.GetPrefix }})
		throw {{ .PgdeFile.As }}.ConstError({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validatePrefix, {{ lit . $r.GetPrefix }});
{{ end }}

{{ if $r.Suffix }}
	if !{{ .Accessor }}.endsWith({{ lit . $r.GetSuffix }})
		throw {{ .PgdeFile.As }}.ConstError({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validateSuffix, {{ lit . $r.GetSuffix }});
{{ end }}

{{ if $r.Contains }}
	if !{{ .Accessor }}.contains({{ lit . $r.GetContains }})
		throw {{ .PgdeFile.As }}.ConstError({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validateContains, {{ lit . $r.GetContains }});
{{ end }}

{{ if $r.GetIp }}
	try {
		Uri.parseIPv4Address({{ .Accessor }});
	} on FormatException {
		try {
			Uri.parseIPv6Address({{ .Accessor }});
		} on FormatException {
			throw {{ .PgdeFile.As }}.BeSomethingError({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validateIp);
		}
	}
{{ else if $r.GetIpv4 }}
	try {
		Uri.parseIPv4Address({{ .Accessor }});
	} on FormatException {
		throw {{ .PgdeFile.As }}.BeSomethingError({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validateIpv4);
	}
{{ else if $r.GetIpv6 }}
	try {
		Uri.parseIPv6Address({{ .Accessor }});
	} on FormatException {
		throw {{ .PgdeFile.As }}.BeSomethingError({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validateIpv6);
	}
{{ else if $r.GetEmail }}
	if (!{{ .EmailValidatorFile.As }}.EmailValidator.validate({{ .Accessor }}))
		throw {{ .PgdeFile.As }}.BeSomethingError({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validateEmail);
{{ else if $r.GetHostname }}
	try {
		Uri(host: {{ .Accessor }});
	} on FormatException {
		throw {{ .PgdeFile.As }}.BeSomethingError({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validateHostname);
	}
{{ else if $r.GetUri }}
	try {
		if (!Uri.parse({{ .Accessor }}).isAbsolute) throw {{ .PgdeFile.As }}.BeSomethingError({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validateUri);
	} on FormatException {
		throw {{ .PgdeFile.As }}.BeSomethingError({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validateUri);
	}
{{ else if $r.GetUriRef }}
	try {
		Uri.parse({{ .Accessor }});
	} on FormatException {
		throw {{ .PgdeFile.As }}.BeSomethingError({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validateUri);
	}
{{ end }}

{{ if $r.Pattern }}
	if (!{{ $kPattern }}.hasMatch({{ .Accessor }}))
		throw {{ .PgdeFile.As }}.ConstError({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validatePattern, {{ $kPattern }}.pattern);
{{ end }}
`
