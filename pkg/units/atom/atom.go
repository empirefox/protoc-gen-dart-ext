package atom

type Atom struct {
	Name   string
	Symbol string
	Plural string
}

func (a *Atom) PrintSymbol(plural bool) string {
	if a.Symbol != "" {
		return a.Symbol
	}
	if plural {
		return a.Plural
	}
	return a.Name
}

var (
	// length
	meter = &Atom{"meter", "m", "meters"}
	foot  = &Atom{"foot", "ft", "feet"}
	inch  = &Atom{"inch", "in", "inches"}
	yard  = &Atom{"yard", "yd", "yards"}
	mile  = &Atom{"mile", "mi", "miles"}
	// 1 nautical mile = 1852 m
	nauticalMile = &Atom{"nautical mile", "NM", "nautical miles"}
	lightYear    = &Atom{"light year", "LY", "light years"}

	// area
	hectare = &Atom{"hectare", "ha", "hectares"}
	are     = &Atom{"are", "a", "ares"}

	// volume
	liter  = &Atom{"liter", "L", "liters"}
	gallon = &Atom{"gallon", "gal", "gallons"}
	barrel = &Atom{"barrel", "bbl", "barrels"}

	// mass
	gram      = &Atom{"gram", "g", "grams"}
	metricTon = &Atom{"ton", "t", "tons"}
	pound     = &Atom{"pound", "lbs", "pounds"}
	ounce     = &Atom{"ounce", "oz", "ounces"}

	// time
	second  = &Atom{"second", "s", "seconds"}
	minute  = &Atom{"minute", "min", "minutes"}
	hour    = &Atom{"hour", "h", "hours"}
	day     = &Atom{"day", "d", "days"}
	week    = &Atom{"week", "", "weeks"}
	month   = &Atom{"month", "", "months"}
	year    = &Atom{"year", "yr", "years"}
	century = &Atom{"century", "", "centuries"}

	// angle
	secondAngle = &Atom{"second", "''", "seconds"}
	minuteAngle = &Atom{"minute", "'", "minutes"}
	degree      = &Atom{"degree", "°", "degrees"}

	// electric current, sound
	ampere       = &Atom{"ampere", "A", "amperes"}
	electronvolt = &Atom{"electronvolt", "eV", "electronvolts"}
	bel          = &Atom{"bel", "B", "bels"}

	// thermodynamic temperature
	kelvin = &Atom{"kelvin", "K", "kelvins"}

	// amount of substance
	mole = &Atom{"mole", "mol", "moles"}

	// luminous intensity
	candela = &Atom{"candela", "cd", "candelas"}

	// ratio
	percent     = &Atom{"percent", "%", ""}
	perthousand = &Atom{"per thousand", "‰", ""}

	bit       = &Atom{"bit", "b", "bits"}
	byte      = &Atom{"byte", "B", "bytes"}
	character = &Atom{"character", "", "characters"}
	word      = &Atom{"word", "", "words"}

	// Table 3.  SI derived units with special names and symbols
	radian        = &Atom{"radian", "rad", "radians"}
	steradian     = &Atom{"steradian", "sr", "steradians"}
	hertz         = &Atom{"hertz", "Hz", ""}
	newton        = &Atom{"newton", "N", "newtons"}
	pascal        = &Atom{"pascal", "Pa", "pascals"}
	joule         = &Atom{"joule", "J", "joules"}
	watt          = &Atom{"watt", "W", "watts"}
	coulomb       = &Atom{"coulomb", "C", "coulombs"}
	volt          = &Atom{"volt", "V", "volts"}
	farad         = &Atom{"farad", "F", "farads"}
	ohm           = &Atom{"ohm", "Ω", "ohms"}
	siemens       = &Atom{"siemens", "S", ""}
	weber         = &Atom{"weber", "Wb", "webers"}
	tesla         = &Atom{"tesla", "T", "teslas"}
	henry         = &Atom{"henry", "H", "henries"}
	degreeCelsius = &Atom{"degree Celsius", "°C", "degrees Celsius"}
	lumen         = &Atom{"lumen", "lm", "lumens"}
	lux           = &Atom{"lux", "lx", "luxes"}
	becquerel     = &Atom{"becquerel", "Bq", "becquerels"}
	gray          = &Atom{"gray", "Gy", "grays"}
	sievert       = &Atom{"sievert", "Sv", "sieverts"}
	katal         = &Atom{"katal", "kat", "katals"}
)

var Atoms = []*Atom{
	// length
	meter, //= &Atom{"meter", "m", "meters"}
	foot,  //= &Atom{"foot", "ft", "feet"}
	inch,  //= &Atom{"inch", "in", "inches"}
	yard,  //= &Atom{"yard", "yd", "yards"}
	mile,  //= &Atom{"mile", "mi", "miles"}
	// 1 nautical mile = 1852 m
	nauticalMile, //= &Atom{"nautical mile", "NM", "nautical miles"}
	lightYear,    //= &Atom{"light year", "LY", "light years"}

	// area
	hectare, //= &Atom{"hectare", "ha", "hectares"}
	are,     //= &Atom{"are", "a", "ares"}

	// volume
	liter,  //= &Atom{"liter", "L", "liters"}
	gallon, //= &Atom{"gallon", "gal", "gallons"}
	barrel, //= &Atom{"barrel", "bbl", "barrels"}

	// mass
	gram,      //= &Atom{"gram", "g", "grams"}
	metricTon, //= &Atom{"ton", "t", "tons"}
	pound,     //= &Atom{"pound", "lbs", "pounds"}
	ounce,     //= &Atom{"ounce", "oz", "ounces"}

	// time, angle
	second,  //= &Atom{"second", "s", "seconds"}
	minute,  //= &Atom{"minute", "min", "minutes"}
	hour,    //= &Atom{"hour", "h", "hours"}
	day,     //= &Atom{"day", "d", "days"}
	week,    //= &Atom{"week", "", "weeks"}
	month,   //= &Atom{"month", "", "months"}
	year,    //= &Atom{"year", "yr", "years"}
	century, //= &Atom{"century", "", "centuries"}

	// angle
	secondAngle, //= &Atom{"second", "''", "seconds"}
	minuteAngle, //= &Atom{"minute", "'", "minutes"}
	degree,      //= &Atom{"degree", "°", "degrees"}

	// electric current, sound
	ampere,       //= &Atom{"ampere", "A", "amperes"}
	electronvolt, //= &Atom{"electronvolt", "eV", "electronvolts"}
	bel,          //= &Atom{"bel", "B", "bels"}

	// thermodynamic temperature
	kelvin, //= &Atom{"kelvin", "K", "kelvins"}

	// amount of substance
	mole, //= &Atom{"mole", "mol", "moles"}

	// luminous intensity
	candela, //= &Atom{"candela", "cd", "candelas"}

	// ratio
	percent,     //= &Atom{"percent", "%", ""}
	perthousand, //= &Atom{"per thousand", "‰", ""}

	bit,       //= &Atom{"bit", "b", "bits"}
	byte,      //= &Atom{"byte", "B", "bytes"}
	character, //= &Atom{"character", "", "characters"}
	word,      //= &Atom{"word", "", "words"}

	// Table 3.  SI derived units with special names and symbols
	radian,        //= &Atom{"radian", "rad", "radians"}
	steradian,     //= &Atom{"steradian", "sr", "steradians"}
	hertz,         //= &Atom{"hertz", "Hz", ""}
	newton,        //= &Atom{"newton", "N", "newtons"}
	pascal,        //= &Atom{"pascal", "Pa", "pascals"}
	joule,         //= &Atom{"joule", "J", "joules"}
	watt,          //= &Atom{"watt", "W", "watts"}
	coulomb,       //= &Atom{"coulomb", "C", "coulombs"}
	volt,          //= &Atom{"volt", "V", "volts"}
	farad,         //= &Atom{"farad", "F", "farads"}
	ohm,           //= &Atom{"ohm", "Ω", "ohms"}
	siemens,       //= &Atom{"siemens", "S", ""}
	weber,         //= &Atom{"weber", "Wb", "webers"}
	tesla,         //= &Atom{"tesla", "T", "teslas"}
	henry,         //= &Atom{"henry", "H", "henries"}
	degreeCelsius, //= &Atom{"degree Celsius", "°C", "degrees Celsius"}
	lumen,         //= &Atom{"lumen", "lm", "lumens"}
	lux,           //= &Atom{"lux", "lx", "luxes"}
	becquerel,     //= &Atom{"becquerel", "Bq", "becquerels"}
	gray,          //= &Atom{"gray", "Gy", "grays"}
	sievert,       //= &Atom{"sievert", "Sv", "sieverts"}
	katal,         //= &Atom{"katal", "kat", "katals"}
}
