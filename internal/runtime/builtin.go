package runtime

import "math"

var (
	// Angle
	CONST_RAD    = withDocs(NewNumber(1), "Radian, a unit of angle equal to the angle subtended at the center of a circle by an arc equal in length to the radius of the circle.")
	CONST_DEG    = withDocs(NewNumber(180/math.Pi), "Degree, a unit of angle equal to 1/360 of a circle.")
	CONST_GRAD   = withDocs(NewNumber(63.6619), "Gradian, a unit of angle equal to 1/400 of a circle.")
	CONST_ARCMIN = withDocs(NewNumber(3437.75), "Arcminute, a unit of angle equal to 1/60 of a degree.")
	CONST_ARCSEC = withDocs(NewNumber(206265), "Arcsecond, a unit of angle equal to 1/60 of an arcminute.")

	// Area
	CONST_SQMM    = withDocs(NewNumber(0.000001), "Square millimeter, a unit of area equal to 1/1000000 square meters.")
	CONST_SQCM    = withDocs(NewNumber(0.0001), "Square centimeter, a unit of area equal to 1/10000 square meters.")
	CONST_SQM     = withDocs(NewNumber(1), "Square meter, a unit of area equal to 1 square meters.")
	CONST_SQKM    = withDocs(NewNumber(1000000), "Square kilometer, a unit of area equal to 1000000 square meters.")
	CONST_SQIN    = withDocs(NewNumber(0.00064516), "Square inch, a unit of area equal to 1/144 square feet.")
	CONST_SQFT    = withDocs(NewNumber(0.09290304), "Square foot, a unit of area equal to 1/9 square yards.")
	CONST_SQYD    = withDocs(NewNumber(0.83612736), "Square yard, a unit of area equal to 9 square feet.")
	CONST_SQMI    = withDocs(NewNumber(2589988.110336), "Square mile, a unit of area equal to 640 acres.")
	CONST_ACRE    = withDocs(NewNumber(4046.8564224), "Acre, a unit of area equal to 43560 square feet.")
	CONST_HECTARE = withDocs(NewNumber(10000), "Hectare, a unit of area equal to 10000 square meters.")

	// Bytes
	CONST_BIT   = withDocs(NewNumber(0.125), "Bit, a unit of information equal to 1/8 bytes.")
	CONST_KBIT  = withDocs(NewNumber(128), "Kilobit, a unit of information equal to 1024 bits.")
	CONST_MBIT  = withDocs(NewNumber(131072), "Megabit, a unit of information equal to 1024 kilobits.")
	CONST_GBIT  = withDocs(NewNumber(134217728), "Gigabit, a unit of information equal to 1024 megabits.")
	CONST_TBIT  = withDocs(NewNumber(137438953472), "Terabit, a unit of information equal to 1024 gigabits.")
	CONST_PBIT  = withDocs(NewNumber(140737488355328), "Petabit, a unit of information equal to 1024 terabits.")
	CONST_BYTE  = withDocs(NewNumber(1), "Byte, a unit of information equal to 8 bits.")
	CONST_KBYTE = withDocs(NewNumber(1024), "Kilobyte, a unit of information equal to 1024 bytes.")
	CONST_MBYTE = withDocs(NewNumber(1048576), "Megabyte, a unit of information equal to 1024 kilobytes.")
	CONST_GBYTE = withDocs(NewNumber(1073741824), "Gigabyte, a unit of information equal to 1024 megabytes.")
	CONST_TBYTE = withDocs(NewNumber(1099511627776), "Terabyte, a unit of information equal to 1024 gigabytes.")
	CONST_PBYTE = withDocs(NewNumber(1125899906842624), "Petabyte, a unit of information equal to 1024 terabytes.")

	// Length
	CONST_MM = withDocs(NewNumber(0.001), "Millimeter, a unit of length equal to 1/1000 meters.")
	CONST_CM = withDocs(NewNumber(0.01), "Centimeter, a unit of length equal to 1/100 meters.")
	CONST_DM = withDocs(NewNumber(0.1), "Decimeter, a unit of length equal to 1/10 meters.")
	CONST_M  = withDocs(NewNumber(1), "Meter, a unit of length equal to 1 meters.")
	CONST_KM = withDocs(NewNumber(1000), "Kilometer, a unit of length equal to 1000 meters.")
	CONST_IN = withDocs(NewNumber(0.0254), "Inch, a unit of length equal to 1/12 foot.")
	CONST_FT = withDocs(NewNumber(0.3048), "Foot, a unit of length equal to 1/3 yard.")
	CONST_YD = withDocs(NewNumber(0.9144), "Yard, a unit of length equal to 3 feet.")
	CONST_MI = withDocs(NewNumber(1609.344), "Mile, a unit of length equal to 1760 yards.")
	CONST_NM = withDocs(NewNumber(1852), "Nautical mile, a unit of length equal to 1852 meters.")
	CONST_LY = withDocs(NewNumber(9460730472580800), "Light year, a unit of length equal to the distance that light travels in one year.")
	CONST_AU = withDocs(NewNumber(149597870700), "Astronomical unit, a unit of length equal to the mean distance between the Earth and the Sun.")

	// Time
	CONST_MS         = withDocs(NewNumber(0.001), "Millisecond, a unit of time equal to 1/1000 seconds.")
	CONST_MICROS     = withDocs(NewNumber(0.000001), "Microsecond, a unit of time equal to 1/1000000 seconds.")
	CONST_NS         = withDocs(NewNumber(0.000000001), "Nanosecond, a unit of time equal to 1/1000000000 seconds.")
	CONST_PS         = withDocs(NewNumber(0.000000000001), "Picosecond, a unit of time equal to 1/1000000000000 seconds.")
	CONST_SEC        = withDocs(NewNumber(1), "Second, a unit of time equal to 1 seconds.")
	CONST_MIN        = withDocs(NewNumber(60), "Minute, a unit of time equal to 60 seconds.")
	CONST_HOUR       = withDocs(NewNumber(3600), "Hour, a unit of time equal to 3600 seconds.")
	CONST_DAY        = withDocs(NewNumber(86400), "Day, a unit of time equal to 86400 seconds.")
	CONST_WEEK       = withDocs(NewNumber(604800), "Week, a unit of time equal to 604800 seconds.")
	CONST_MONTH      = withDocs(NewNumber(2629800), "Month, a unit of time equal to 2629800 seconds.")
	CONST_YEAR       = withDocs(NewNumber(31557600), "Year, a unit of time equal to 31557600 seconds.")
	CONST_DECADE     = withDocs(NewNumber(315576000), "Decade, a unit of time equal to 315576000 seconds.")
	CONST_CENTURY    = withDocs(NewNumber(3155760000), "Century, a unit of time equal to 3155760000 seconds.")
	CONST_MILLENNIUM = withDocs(NewNumber(31557600000), "Millennium, a unit of time equal to 31557600000 seconds.")

	// Volume
	CONST_ML  = withDocs(NewNumber(0.001), "Milliliter, a unit of volume equal to 1/1000 liters.")
	CONST_L   = withDocs(NewNumber(1), "Liter, a unit of volume equal to 1 liters.")
	CONST_CU  = withDocs(NewNumber(1000), "Cubic meter, a unit of volume equal to 1000 liters.")
	CONST_GAL = withDocs(NewNumber(3.785411784), "Gallon, a unit of volume equal to 231 cubic inches.")

	// Weight
	CONST_MG = withDocs(NewNumber(0.000001), "Milligram, a unit of mass equal to 1/1000 grams.")
	CONST_CG = withDocs(NewNumber(0.00001), "Centigram, a unit of mass equal to 1/100 grams.")
	CONST_DG = withDocs(NewNumber(0.0001), "Decigram, a unit of mass equal to 1/10 grams.")
	CONST_G  = withDocs(NewNumber(0.001), "Gram, a unit of mass equal to 1 grams.")
	CONST_KG = withDocs(NewNumber(1), "Kilogram, a unit of mass equal to 1000 grams.")
	CONST_T  = withDocs(NewNumber(1000), "Ton, a unit of mass equal to 1000 kilograms.")
	CONST_OZ = withDocs(NewNumber(0.028349523125), "Ounce, a unit of mass equal to 1/16 pounds.")
	CONST_LB = withDocs(NewNumber(0.45359237), "Pound, a unit of mass equal to 16 ounces.")

	// Math
	CONST_PI   = withDocs(NewNumber(math.Pi), "Pi, a mathematical constant equal to the ratio of a circle's circumference to its diameter.")
	CONST_E    = withDocs(NewNumber(math.E), "Euler's number, a mathematical constant equal to the base of the natural logarithm.")
	CONST_PHI  = withDocs(NewNumber(1.618033988749895), "Golden ratio, a mathematical constant equal to the ratio of two quantities such that the ratio of the sum of the quantities to the larger quantity is equal to the ratio of the larger quantity to the smaller one.")
	CONST_INF  = withDocs(NewNumber(math.Inf(1)), "Infinity, a mathematical concept representing a quantity that is greater than any real number.")
	CONST_NINF = withDocs(NewNumber(math.Inf(-1)), "Negative infinity, a mathematical concept representing a quantity that is less than any real number.")
)

func registerConstants(scope *Scope) {
	// Mathematical constants
	scope.Set("pi", CONST_PI)
	scope.Set("e", CONST_E)
	scope.Set("phi", CONST_PHI)
	scope.Set("inf", CONST_INF)
	scope.Set("ninf", CONST_NINF)

	// Weight
	scope.Set("mg", CONST_MG)
	scope.Set("cg", CONST_CG)
	scope.Set("dg", CONST_DG)
	scope.Set("g", CONST_G)
	scope.Set("kg", CONST_KG)
	scope.Set("ton", CONST_T)
	scope.Set("oz", CONST_OZ)
	scope.Set("lb", CONST_LB)

	// Volume
	scope.Set("ml", CONST_ML)
	scope.Set("l", CONST_L)
	scope.Set("cu", CONST_CU)
	scope.Set("gal", CONST_GAL)

	// Time
	scope.Set("ms", CONST_MS)
	scope.Set("micros", CONST_MICROS)
	scope.Set("ns", CONST_NS)
	scope.Set("ps", CONST_PS)
	scope.Set("sec", CONST_SEC)
	scope.Set("min", CONST_MIN)
	scope.Set("hour", CONST_HOUR)
	scope.Set("day", CONST_DAY)
	scope.Set("week", CONST_WEEK)
	scope.Set("month", CONST_MONTH)
	scope.Set("year", CONST_YEAR)
	scope.Set("decade", CONST_DECADE)
	scope.Set("century", CONST_CENTURY)
	scope.Set("millennium", CONST_MILLENNIUM)

	// Length
	scope.Set("mm", CONST_MM)
	scope.Set("cm", CONST_CM)
	scope.Set("dm", CONST_DM)
	scope.Set("m", CONST_M)
	scope.Set("km", CONST_KM)
	scope.Set("in", CONST_IN)
	scope.Set("ft", CONST_FT)
	scope.Set("yd", CONST_YD)
	scope.Set("mi", CONST_MI)
	scope.Set("nm", CONST_NM)
	scope.Set("ly", CONST_LY)
	scope.Set("au", CONST_AU)

	// Bytes
	scope.Set("b", CONST_BIT)
	scope.Set("Kb", CONST_KBIT)
	scope.Set("Mb", CONST_MBIT)
	scope.Set("Gb", CONST_GBIT)
	scope.Set("Tb", CONST_TBIT)
	scope.Set("Pb", CONST_PBIT)
	scope.Set("B", CONST_BYTE)
	scope.Set("KB", CONST_KBYTE)
	scope.Set("MB", CONST_MBYTE)
	scope.Set("GB", CONST_GBYTE)
	scope.Set("TB", CONST_TBYTE)
	scope.Set("PB", CONST_PBYTE)

}
