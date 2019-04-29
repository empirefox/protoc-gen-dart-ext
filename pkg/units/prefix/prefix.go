package prefix

type Type int

const (
	SI Type = iota
	Binary
)

type Prefix struct {
	Type     Type
	Name     string
	Symbol   string
	Base     int32
	Exponent int32
}

var (
	// si
	yotta = &Prefix{SI, "yotta", "Y", 10, 24}
	zetta = &Prefix{SI, "zetta", "Z", 10, 21}
	exa   = &Prefix{SI, "exa", "E", 10, 18}
	peta  = &Prefix{SI, "peta", "P", 10, 15}
	tera  = &Prefix{SI, "tera", "T", 10, 12}
	giga  = &Prefix{SI, "giga", "G", 10, 9}
	mega  = &Prefix{SI, "mega", "M", 10, 6}
	kilo  = &Prefix{SI, "kilo", "k", 10, 3}
	hecto = &Prefix{SI, "hecto", "h", 10, 2}
	deka  = &Prefix{SI, "deka", "da", 10, 1}
	deci  = &Prefix{SI, "deci", "d", 10, -1}
	centi = &Prefix{SI, "centi", "c", 10, -2}
	milli = &Prefix{SI, "milli", "m", 10, -3}
	micro = &Prefix{SI, "micro", "µ", 10, -6}
	nano  = &Prefix{SI, "nano", "n", 10, -9}
	pico  = &Prefix{SI, "pico", "p", 10, -12}
	femto = &Prefix{SI, "femto", "f", 10, -15}
	atto  = &Prefix{SI, "atto", "a", 10, -18}
	zepto = &Prefix{SI, "zepto", "z", 10, -21}
	yocto = &Prefix{SI, "yocto", "y", 10, -24}

	// binary
	kibi = &Prefix{Binary, "kibi", "Ki", 2, 10}
	mebi = &Prefix{Binary, "mebi", "Mi", 2, 20}
	gibi = &Prefix{Binary, "gibi", "Gi", 2, 30}
	tebi = &Prefix{Binary, "tebi", "Ti", 2, 40}
	pebi = &Prefix{Binary, "pebi", "Pi", 2, 50}
	exbi = &Prefix{Binary, "exbi", "Ei", 2, 60}
)

var SIPrefixes = []*Prefix{
	yotta,
	zetta,
	exa,
	peta,
	tera,
	giga,
	mega,
	kilo,
	hecto,
	deka,
	deci,
	centi,
	milli,
	micro,
	nano,
	pico,
	femto,
	atto,
	zepto,
	yocto,
}

var BinaryPrefixes = []*Prefix{
	kibi,
	mebi,
	gibi,
	tebi,
	pebi,
	exbi,
}

func AllPrefix() []*Prefix { return merge(SIPrefixes, BinaryPrefixes) }
func merge(a, b []*Prefix) (c []*Prefix) {
	c = make([]*Prefix, len(a)+len(b))
	copy(c[copy(c, a):], b)
	return
}
