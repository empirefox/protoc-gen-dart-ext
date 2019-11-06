package form

import (
	"errors"
	"fmt"
)

var (
	ErrNoColor = errors.New("no color")
	ErrNoIcon  = errors.New("no icon")
)

func (color *MaterialColor) ToDartSource() (string, error) {
	if non := color.GetNon(); non != nil {
		return "Colors." + non.String(), nil
	} else if primary := color.GetPrimary(); primary != nil {
		return fmt.Sprintf("Colors.%s[%s]", primary.Color, primary.Shade.String()[1:]), nil
	} else if accent := color.GetAccent(); accent != nil {
		return fmt.Sprintf("Colors.%s[%s]", accent.Color, accent.Shade.String()[1:]), nil
	}
	return "", ErrNoColor
}

func (i MaterialIcon) ToDartSource() (string, error) {
	if i == MaterialIcon_noMaterialIcon {
		return "", ErrNoIcon
	}
	return fmt.Sprintf("const Icon(Icons.%s)", i), nil
}
