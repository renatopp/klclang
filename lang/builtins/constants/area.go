package constants

import "klc/lang/builtins"

var SQMM = builtins.WithDoc(builtins.NewNumber(0.000001), "Square millimeter, a unit of area equal to 1/1000000 square meters.")
var SQCM = builtins.WithDoc(builtins.NewNumber(0.0001), "Square centimeter, a unit of area equal to 1/10000 square meters.")
var SQM = builtins.WithDoc(builtins.NewNumber(1), "Square meter, a unit of area equal to 1 square meters.")
var SQKM = builtins.WithDoc(builtins.NewNumber(1000000), "Square kilometer, a unit of area equal to 1000000 square meters.")
var SQIN = builtins.WithDoc(builtins.NewNumber(0.00064516), "Square inch, a unit of area equal to 1/144 square feet.")
var SQFT = builtins.WithDoc(builtins.NewNumber(0.09290304), "Square foot, a unit of area equal to 1/9 square yards.")
var SQYD = builtins.WithDoc(builtins.NewNumber(0.83612736), "Square yard, a unit of area equal to 9 square feet.")
var SQMI = builtins.WithDoc(builtins.NewNumber(2589988.110336), "Square mile, a unit of area equal to 640 acres.")
var ACRE = builtins.WithDoc(builtins.NewNumber(4046.8564224), "Acre, a unit of area equal to 43560 square feet.")
var HECTARE = builtins.WithDoc(builtins.NewNumber(10000), "Hectare, a unit of area equal to 10000 square meters.")
