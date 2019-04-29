package module

const defaultsTplStr = `// This library provides Type to default instance mapping.
{{- $name := .InputPath.BaseName }}
import 'package:protobuf/protobuf.dart' show GeneratedMessage;
import './{{ $name }}.pb.dart';

T defaultOf_{{ $name }}<T extends GeneratedMessage>(Type type) {
  switch (type) {
    {{- range .Messages }}
    case {{ .Name }}: return {{ .Name }}.getDefault() as T;
    {{- end }}
    // default: throw ArgumentError.value(T);
    default: return null;
  }
}
`
