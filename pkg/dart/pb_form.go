package dart

import (
	"github.com/empirefox/protoc-gen-dart-ext/pkg/pgde/form"
)

func (im *ImportManager) RenderMaterialColor(color *form.MaterialColor) Qualifier {
	material := im.Material()
	switch c := color.GetColor().(type) {
	case *form.MaterialColor_Non:
		return material.Colors().DotStringer(c.Non.Color)
	case *form.MaterialColor_Primary:
		return material.Colors().DotStringer(c.Primary.Color).
			IndexString(c.Primary.Shade.String()[1:])
	case *form.MaterialColor_Accent:
		return material.Colors().DotStringer(c.Accent.Color).
			IndexString(c.Accent.Shade.String()[1:])
	}
	return "null"
}

func (im *ImportManager) RenderMaterialIcon(i form.MaterialIcon) Qualifier {
	if i == form.MaterialIcon_noMaterialIcon {
		return "null"
	}
	material := im.Material()
	return "const " + material.Icon().InitWith(material.Icons().DotStringer(i))
}
