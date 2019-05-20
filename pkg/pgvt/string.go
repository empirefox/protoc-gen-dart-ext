package pgvt

const stringConstTpl = stringTplDef + `
static final RegExp {{ $kPattern }} = RegExp({{ lit $r.GetPattern }});

{{ template "inConst" . }}
`

const stringTplDef = `{{ $r := .Rules }}
{{- $kPattern := constRef . "kPattern" }}
`

const stringTpl = stringTplDef + `
{{ template "const" . }}
{{ template "in" . }}

{{ if or $r.Len $r.MinLen $r.MaxLen }}
	final _vrl = _v.runes.length;
{{ end }}

{{ if or $r.Len (and $r.MinLen $r.MaxLen (eq $r.GetMinLen $r.GetMaxLen)) }}
	{{ if $r.Len }}
		if (_vrl != {{ $r.GetLen }})
			throw $pgde.LenConstError.character(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateEq, {{ $r.GetLen }});
	{{ else }}
		if (_vrl != {{ $r.GetMinLen }})
			throw $pgde.LenConstError.character(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateEq, {{ $r.GetMinLen }});
	{{ end }}
{{ else if $r.MinLen }}
	{{ if $r.MaxLen }}
		if (_vrl < {{ $r.GetMinLen }} || _vrl > {{ $r.GetMaxLen }})
			throw $pgde.RangeLenError.character(const ErrorRange.inEE(info, {{ $.Number }}, {{ $.L10nField }}, {{ $r.GetMinLen }}, {{ $r.GetMaxLen }}));
	{{ else }}
		if (_vrl < {{ $r.GetMinLen }})
			throw $pgde.LenConstError.character(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateGte, {{ $r.GetMinLen }});
	{{ end }}
{{ else if $r.MaxLen }}
	if (_vrl > {{ $r.GetMaxLen }})
		throw $pgde.LenConstError.character(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateLte, {{ $r.GetMaxLen }});
{{ end }}

{{ if or $r.Len $r.MaxBytes $r.MinBytes }}
	final _vrl = $pgde.Lists.len(_v);
{{ end }}

{{ if or $r.LenBytes (and $r.MinBytes $r.MaxBytes (eq $r.GetMinBytes $r.GetMaxBytes)) }}
	{{ if $r.LenBytes }}
		if (_vbl != {{ $r.GetLenBytes }})
			throw $pgde.LenConstError.byte(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateEq, {{ $r.GetLenBytes }});
	{{ else }}
		if (_vbl != {{ $r.GetMinBytes }})
			throw $pgde.LenConstError.byte(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateEq, {{ $r.GetMinBytes }});
	{{ end }}
{{ else if $r.MinBytes }}
	{{ if $r.MaxBytes }}
		if (_vbl < {{ $r.GetMinBytes }} || _vbl > {{ $r.GetMaxBytes }})
			throw $pgde.RangeLenError.byte(const ErrorRange.inEE(info, {{ $.Number }}, {{ $.L10nField }}, {{ $r.GetMinBytes }}, {{ $r.GetMaxBytes }}));
	{{ else }}
		if (_vbl < {{ $r.GetMinBytes }})
			throw $pgde.LenConstError.byte(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateGte, {{ $r.GetMinBytes }});
	{{ end }}
{{ else if $r.MaxBytes }}
	if (_vbl > {{ $r.GetMaxBytes }})
		throw $pgde.LenConstError.byte(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateLte, {{ $r.GetMaxBytes }});
{{ end }}

{{ if $r.Prefix }}
	if !_v.startsWith({{ lit $r.GetPrefix }})
		throw $pgde.ConstError(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validatePrefix, {{ lit $r.GetPrefix }});
{{ end }}

{{ if $r.Suffix }}
	if !_v.endsWith({{ lit $r.GetSuffix }})
		throw $pgde.ConstError(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateSuffix, {{ lit $r.GetSuffix }});
{{ end }}

{{ if $r.Contains }}
	if !_v.contains({{ lit $r.GetContains }})
		throw $pgde.ConstError(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateContains, {{ lit $r.GetContains }});
{{ end }}

{{ if $r.GetIp }}
	try {
		Uri.parseIPv4Address(_v);
	} on FormatException {
		try {
			Uri.parseIPv6Address(_v);
		} on FormatException {
			throw $pgde.BeSomethingError(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateIp);
		}
	}
{{ else if $r.GetIpv4 }}
	try {
		Uri.parseIPv4Address(_v);
	} on FormatException {
		throw $pgde.BeSomethingError(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateIpv4);
	}
{{ else if $r.GetIpv6 }}
	try {
		Uri.parseIPv6Address(_v);
	} on FormatException {
		throw $pgde.BeSomethingError(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateIpv6);
	}
{{ else if $r.GetEmail }}
	if (!$email_validator.EmailValidator.validate(_v))
		throw $pgde.BeSomethingError(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateEmail);
{{ else if $r.GetHostname }}
	try {
		Uri(host: _v);
	} on FormatException {
		throw $pgde.BeSomethingError(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateHostname);
	}
{{ else if $r.GetUri }}
	try {
		if (!Uri.parse(_v).isAbsolute) throw $pgde.BeSomethingError(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateUri);
	} on FormatException {
		throw $pgde.BeSomethingError(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateUri);
	}
{{ else if $r.GetUriRef }}
	try {
		Uri.parse(_v);
	} on FormatException {
		throw $pgde.BeSomethingError(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateUri);
	}
{{ end }}

{{ if $r.Pattern }}
	if (!{{ $kPattern }}.hasMatch(_v))
		throw $pgde.ConstError(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validatePattern, {{ $kPattern }}.pattern);
{{ end }}
`
