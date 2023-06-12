package constants

import "klc/lang/builtins"

var ML = builtins.WithDoc(builtins.NewNumber(0.001), "Milliliter, a unit of volume equal to 1/1000 liters.")
var L = builtins.WithDoc(builtins.NewNumber(1), "Liter, a unit of volume equal to 1 liters.")
var CU = builtins.WithDoc(builtins.NewNumber(1000), "Cubic meter, a unit of volume equal to 1000 liters.")
var GAL = builtins.WithDoc(builtins.NewNumber(3.785411784), "Gallon, a unit of volume equal to 231 cubic inches.")
