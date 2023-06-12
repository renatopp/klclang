package constants

import "klc/lang/builtins"

var MM = builtins.WithDoc(builtins.NewNumber(0.001), "Millimeter, a unit of length equal to 1/1000 meters.")
var CM = builtins.WithDoc(builtins.NewNumber(0.01), "Centimeter, a unit of length equal to 1/100 meters.")
var DM = builtins.WithDoc(builtins.NewNumber(0.1), "Decimeter, a unit of length equal to 1/10 meters.")
var M = builtins.WithDoc(builtins.NewNumber(1), "Meter, a unit of length equal to 1 meters.")
var KM = builtins.WithDoc(builtins.NewNumber(1000), "Kilometer, a unit of length equal to 1000 meters.")
var IN = builtins.WithDoc(builtins.NewNumber(0.0254), "Inch, a unit of length equal to 1/12 foot.")
var FT = builtins.WithDoc(builtins.NewNumber(0.3048), "Foot, a unit of length equal to 1/3 yard.")
var YD = builtins.WithDoc(builtins.NewNumber(0.9144), "Yard, a unit of length equal to 3 feet.")
var MI = builtins.WithDoc(builtins.NewNumber(1609.344), "Mile, a unit of length equal to 1760 yards.")
var NM = builtins.WithDoc(builtins.NewNumber(1852), "Nautical mile, a unit of length equal to 1852 meters.")
var LY = builtins.WithDoc(builtins.NewNumber(9460730472580800), "Light year, a unit of length equal to the distance that light travels in one year.")
var AU = builtins.WithDoc(builtins.NewNumber(149597870700), "Astronomical unit, a unit of length equal to the mean distance between the Earth and the Sun.")
