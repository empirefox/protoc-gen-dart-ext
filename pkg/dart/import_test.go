package dart

import (
	"testing"
)

func TestRenderImports(t *testing.T) {
	d := NewDart()
	im := NewDefaultImportManager(d, "pkg/a.dart")

	collectionLib := im.Import("dart:collection")
	materialLib := im.Import("package:flutter/material.dart")

	collectionLib.AsDot(Qualifier("HashSet"))
	materialLib.AsDot(Qualifier("LocalizationssDelegate"))
	materialLib.AsDot(Qualifier("Locale"))

	const expected = `import 'dart:collection' as $0 show HashSet;
import 'package:flutter/material.dart' as $1 show Locale, LocalizationssDelegate;`

	result := im.RenderImports()
	if result != expected {
		t.Errorf("Got wrong imports: %s", result)
	}
}
