package constants

import "klc/lang/builtins"

var RAD = builtins.WithDoc(builtins.NewNumber(1), "Radian, a unit of angle equal to the angle subtended at the center of a circle by an arc equal in length to the radius of the circle.")
var DEG = builtins.WithDoc(builtins.NewNumber(57.2958), "Degree, a unit of angle equal to 1/360 of a circle.")
var GRAD = builtins.WithDoc(builtins.NewNumber(63.6619), "Gradian, a unit of angle equal to 1/400 of a circle.")
var ARCMIN = builtins.WithDoc(builtins.NewNumber(3437.75), "Arcminute, a unit of angle equal to 1/60 of a degree.")
var ARCSEC = builtins.WithDoc(builtins.NewNumber(206265), "Arcsecond, a unit of angle equal to 1/60 of an arcminute.")
