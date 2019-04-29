package module

const l10n_base = `
// ignore_for_file: non_constant_identifier_names
import 'dart:core' show int, String, Type;
import 'package:protobuf/protobuf.dart' show ProtobufEnum;
import 'package:flutter/material.dart' show IconData, Icons;
{{ range .Files -}}
import './{{ .InputPath.BaseName }}.pb.dart' as pb;
{{ end }}

{{ $File := pkgDartName . }}

{{ if .AllEnums }} // editor line 14
class {{ $File }}Icons {
  IconData fromEnum(ProtobufEnum v) {
    switch (v.runtimeType) {
      {{- range .AllEnums }}
      case pb.{{ nameOf . }}: return from{{ nameOf . }}(v);
      {{- end }}
    }
    assert(false, 'fromEnum() called for unsupported enum "$v"');
    return null;
  }

  {{ range .AllEnums }}{{ $Enum := nameOf . }}
  IconData from{{ $Enum }}(pb.{{ $Enum }} v) {
    switch (v) {
      {{- range .Values }}
      case pb.{{ $Enum }}.{{ nameOf . }}: return Icons.star;
      {{- end }}
    }
    assert(false, 'from{{ $Enum }}() called for unsupported enum "$v"');
    return null;
  }
  {{ end }}
}
{{ end }}

abstract class {{ $File }}Localizations {
  const {{ $File }}Localizations();

  {{ range .AllMessages }}{{ $Message := nameOf . }}
    // {{ $Message }}
    String get {{ $Message }};

    {{- range .OneOfs -}}
    String get {{ $Message }}_{{ nameOf . }};
    {{- end }}

    {{- range .Fields -}}
    String get {{ nameOf . }}_{{ $Message }};
    {{- end }}
  {{ end }}

	{{ if .AllMessages }}
	  String fromMessage(Type type) {
	    switch (type) {
	      {{- range .AllMessages }}{{ $Message := nameOf . }}
	      case pb.{{ $Message }}: return {{ $Message }};
	      {{- end }}
	      default:
	        // type not found
	        return '@${type}';
	    }
	  }
	
	  String fromField(Type type, int tag) {
	    switch (type) {
	      {{- range .AllMessages }}{{ $Message := nameOf . }}
	      case pb.{{ $Message }}:
	        switch (tag) {
	          {{- range .Fields }}
	          case {{ .Descriptor.Number }}: return {{ nameOf . }}_{{ $Message }};
	          {{- end }}
	        }
	        break;
	      {{- end }}
	      default:
	        // type not found
	        return '$tag.@$type';
	    }
	    // tag not found
	    return '@$tag.$type';
	  }
	{{ end }}

	{{ if hasOneofPkg . }}
		String fromOneof(Type type, int oneof) {
			switch (type) {
				{{- range .AllMessages }}{{ $Message := nameOf . }}
					{{- with $oneofs := .OneOfs }}
					case pb.{{ $Message }}:
						switch (oneof) {
							{{- range $idx, $Oneof := $oneofs }}
							case {{ $idx }}: return {{ $Message }}_{{ nameOf $Oneof }};
							{{- end }}
						}
						break;
					{{- end }}
				{{- end }}
				default:
					// type not found
					return '($oneof).@$type';
			}
			// oneof not found
			return '@($oneof).$type';
		}
	{{ end }}

	{{ range .AllEnums }}{{ $Enum := nameOf . }}{{ $l10n := l10n_enum . }}{{ $values := .Values }}
	  // {{ $Enum }}
	  String get {{ $Enum }};
		{{ if not $l10n.GetIgnore }}
			{{- range $values }}
			String get {{ $Enum }}_{{ nameOf . }};
			{{- end }}

			{{ range $variant := $l10n.GetVariants }}
				{{- range $values }}{{ $value := nameOf . }}
				String get {{ $Enum }}_{{ variant_enum_value $variant $value }};
				{{- end }}
			{{ end -}}
		{{ end }}
	{{ end }}

	{{ if hasEnumPkg . }}
	  String fromEnumType(Type type) {
	    switch (type) {
				{{- range .AllEnums }}{{ $Enum := nameOf . }}
	      case pb.{{ $Enum }}: return {{ $Enum }};
	      {{- end }}
		}
		// type not found
	    return '@$type';
	  }
	
	  String fromEnum(ProtobufEnum v) {
	    switch (v.runtimeType) {
	      {{- range .AllEnums }}{{ $Enum := nameOf . }}{{ $l10n := l10n_enum . }}
				{{- if not $l10n.GetIgnore }}
	      case pb.{{ $Enum }}: return from{{ $Enum }}(v);
				{{- end }}
	      {{- end }}
	    }
		// enum not found
	    return '$v.@${v.runtimeType}';
	  }
	{{ end }}


	{{ range .AllEnums }}{{ $Enum := nameOf . }}{{ $l10n := l10n_enum . }}{{ $values := .Values }}
		{{ if not $l10n.GetIgnore }}
			String from{{ $Enum }}(pb.{{ $Enum }} v) {
				switch (v) {
					{{- range $values }}
					case pb.{{ $Enum }}.{{ nameOf . }}: return {{ $Enum }}_{{ nameOf . }};
					{{- end }}
				}
			// enum not found
				return '@$v';
			}
		
			{{ range $variant := $l10n.GetVariants }}
			String from{{ variant_enum $variant $Enum }}(pb.{{ $Enum }} v) {
				switch (v) {
					{{- range $values }}{{ $value := nameOf . }}
					case pb.{{ $Enum }}.{{ $value }}: return {{ $Enum }}_{{ variant_enum_value $variant $value }};
					{{- end }}
				}
			// enum not found
				return '@$v';
			}
			{{ end -}}
		{{ end }}
  {{ end }}
}
`
