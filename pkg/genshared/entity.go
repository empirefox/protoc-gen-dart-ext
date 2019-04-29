package genshared

import (
	"fmt"
	"unicode"
)

type VersionEntity struct {
	LowerCamelCase string
	Version        uint32
}

func (ve *VersionEntity) LowerCamel() string {
	return ve.LowerCamelCase
}

func (ve *VersionEntity) UpperCamel() string {
	return string(unicode.ToTitle(rune(ve.LowerCamelCase[0]))) + ve.LowerCamelCase[1:]
}

func (ve *VersionEntity) LowerCamelV() string {
	if ve.Version == 0 {
		return ve.LowerCamelCase
	}
	return fmt.Sprintf("%sV%d", ve.LowerCamelCase, ve.Version)
}

func (ve *VersionEntity) UpperCamelV() string {
	if ve.Version == 0 {
		return ve.UpperCamel()
	}
	return fmt.Sprintf("%sV%d", ve.UpperCamel(), ve.Version)
}
