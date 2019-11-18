//go:generate enumer -type=Action -text $GOFILE
package main

type Action int

const (
	// Add non-exist lang
	Add Action = iota

	// Merge like Add, but also include exist langs
	Merge

	// Replace the entire LangInfos
	Replace
)
